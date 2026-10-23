# Rodada 1 · IA · crítico do uso de IA (autoria e aprendizado)

## 1. Como o jogo funciona (opinião)
- **Movimento:** clique para andar, em tempo real no mundo; turno só dentro da luta.
- **Zona:** até ~30 jogadores visíveis por zona. A luta de outro aparece como um grid fantasma no chão, sem bloquear o mapa: quem não está nela atravessa por fora.
- **Grupo:** até 4 jogadores no mesmo grid. Quem chega entra pela borda no início da rodada seguinte.
- **AFK:** solo não tem relógio. Em grupo há um limite por turno (~60 s); quando estoura, o **automático joga aquele turno** com a lista de prioridade do jogador. O sistema já existe, então não custa uma mecânica nova.

## 2. Câmera (opinião)
3/4 inclinada, ângulo fixo, zoom curto e sem rotação. A mesma câmera no mundo e na luta: o grid só acende no chão. Isso mantém a "luta no lugar" coerente, e sem rotação some metade dos problemas de oclusão e de UI.

## 3. Stack (opinião)
O núcleo em TS serve, e é a linguagem onde ele **menos precisa de IA**. Isso é critério de autoria, não só técnico. O item 1 acrescenta posição em tempo real e área de interesse por zona, o que Node + WebSocket aguenta nessa escala. O risco está no cliente: Unity sem C# é onde ele mais vai pedir código pronto. Cliente "burro" e pequeno, de propósito.

## 4. Arte (opinião)
Low poly Synty + contorno funciona em 3/4 fixa. Basta testar a silhueta no zoom mais afastado. E tem um bônus: são assets feitos por gente, o que o protege no item 5.

## 5. IA no projeto

**Fatos:**
- A Steam exige declarar conteúdo gerado por IA (o que vem pronto e o que é gerado ao vivo); ferramentas de código ficam isentas (regra reescrita em jan/2026, [VGC](https://videogameschronicle.com/news/valve-has-significantly-rewritten-steams-rules-for-how-developers-much-disclose-ai-use)).
- ~8 mil jogos (7% do catálogo e ~20% dos lançamentos de 2025) declaram IA, 60% deles em arte ([VGC/Totally Human](https://www.videogameschronicle.com/news/steam-games-disclosing-generative-ai-use-are-up-800-this-year/)).
- GDC 2026: **52%** dos devs acham a IA generativa prejudicial (eram 30% e 18% nos dois anos anteriores); 64% em arte, 63% em design/narrativa e 59% em programação ([BusinessWire](https://www.businesswire.com/news/home/20260129438528/en/2026-State-of-the-Game-Industry-Report-Reveals-Widening-Effect-of-Layoffs-Broader-Perspectives-on-Generative-AI-Unionization-Tariffs-and-More)).
- *Clair Obscur* perdeu os prêmios do Indie Game Awards (dez/2025) por um **placeholder** de IA que escapou para o jogo final ([AV Club](https://www.avclub.com/clair-obscur-genai-iga-awards-rescinded)). A Larian desistiu de IA em concept art depois da reação do público ([Hitmarker](https://hitmarker.net/news/larian-confirms-that-it-will-no-longer-use-generative-ai-for-divinity-concept-art-and-writing-following-backlash-1615314)). Arc Raiders regravou vozes de IA com atores ([PCGamesN](https://www.pcgamesn.com/arc-raiders/ai-voices-replaced)).
- Os NPCs com LLM de *Where Winds Meet* foram manipulados: jogadores subiam a reputação com bajulação e mentiras ([AllThings.how](https://allthings.how/where-winds-meets-ai-npcs-are-already-being-pushed-to-their-breaking-point/)).
- Direitos autorais: nos EUA, prompt sozinho não gera autoria; o que se protege é a seleção ou modificação humana ([Copyright Office, jan/2025](https://www.jenner.com/en/news-insights/client-alerts/us-copyright-office-issues-report-on-copyrightability-of-works-incorporating-ai-generated-material)). No Brasil, a Lei 9.610 (art. 11) só reconhece como autor pessoa física, e o PL 2338 está parado na Câmara justamente por causa dos direitos autorais ([Teletime](https://teletime.com.br/25/03/2026/votacao-do-pl-de-ia-antes-das-eleicoes-e-improvavel/)).
- Aprendizado: num estudo, quem delegou código à IA tirou 50% na prova de compreensão contra 67% de quem fez à mão. Quem só fazia perguntas conceituais manteve o aprendizado ([Anthropic, 2026](https://www.anthropic.com/research/AI-assistance-coding-skills)).

**O caso dele (acervo):** em abril apagou o código porque "todo código que existia era feito por agente". Em 14/10 a IA escreveu as 460 linhas do `position.ts`, justamente o núcleo que ele queria aprender, e veio o desabafo: "vou começar de novo do zero, e isso vai acontecer sempre". Em 08/09 ele viu que lores da IA convergiam para o caminho estatístico mais genérico. Em 20–21/10 escreveu o enredo à mão e rejeitou o que tinha "cara de IA".

**Opinião, por uso:**

| Uso | Veredito |
|---|---|
| Código do núcleo e do servidor | **Não.** Só perguntas conceituais |
| Testes | **Contra o roteiro atual.** Num motor de regras d20, o teste *é* a regra. Se a IA escreve o teste, ela decide o design. Ele escreve os casos (entrada → saída esperada); a IA só sugere casos de borda, que ele aceita ou não |
| Cola do Unity, API e editor | A IA explica. A referência fica em repo separado e fechada antes de escrever (já está no roteiro) |
| Ferramentas descartáveis (visualizador, scripts) | Pode, numa pasta marcada fora do jogo |
| Arte, inclusive placeholder | **Não.** Cubo cinza é placeholder honesto |
| Lore, nomes, textos | **Não.** IA só caça furo |
| Voz | **Não.** Sem voz ou sons sem fala |
| NPC com LLM | **Não.** Custo por token, declaração "ao vivo" na Steam, jailbreak e quebra da regra de não poetizar. A fofoca determinística a partir de eventos do servidor é melhor e é dele |
| Página da loja | Sem IA |

**Anti-ciclo:** a pureza não resolve. Ele mesmo disse que fazer sem IA também frustra e o devolve à IA. O que resolve é um **contrato fixo** (AGENTS.md com a tabela acima). Usar a IA dentro do contrato não é recaída, então não dispara reset. Duas regras:
- **Teste da explicação:** tudo que entra ele explica sem olhar.
- **Nunca apagar o repo:** se a IA escreveu algo fora do contrato, ele reescreve só aquele arquivo.

## 6. Ponto crítico: este próprio debate
Quinze agentes decidindo movimento, câmera e AFK é a IA gerando o design de novo, a mesma fonte do "não é meu" no nível do design. Proposta: a saída do debate é um **cardápio** com 2 opções por ponto. Ele escolhe testando no papel, e cada doc ganha o campo `decidido_por: sidinei`. O que a IA decidir sozinha não fica validado.
