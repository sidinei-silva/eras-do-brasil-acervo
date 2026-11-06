# Memória do Project do Claude — Eras Do Brasil

Arquivos de memória que o Claude mantinha para este Project (export de 05/11/2026). É texto gerado pelo Claude a partir das conversas, não documento escrito pelo Sidinei. Foi exportado antes de a direção de 05/11 ser confirmada, então a virada para mundo simulado ainda aparece como decisão pendente.

## Arquivo de memória `/areas/eras-do-brasil.md` (atualizado 2026-11-06)

---
name: eras-do-brasil
description: Eras do Brasil, MMORPG sandbox por turno do Sidinei — estado atual (fase 0), forma de trabalho, stack, mecânicas decididas, Travessia e lore
sources: [chat, cowork]
aliases: [Eras do Brasil, eras-do-brasil, Eras]
---

## Repositórios e estado [s:ad51c2300a7a4a44]
- [stated] Repo atual: github.com/sidinei-silva/eras-do-brasil; docs/ é a fonte da verdade, mantida por ele no Obsidian [s:ad51c2300a7a4a44]
- [stated] Histórico de todas as fases: github.com/sidinei-silva/eras-do-brasil-acervo; é evidência, não verdade [s:ad51c2300a7a4a44]
- [stated] Fase atual: fase 1 (luta no terminal em Go) em andamento desde out/2026, plano em docs/gestao/fase-atual.md; núcleo de regras no pacote core/ (não rules/); em 29/10/2026 trabalhava no estado mínimo do combate + ResolveCombatAction(UseSkill) [s:88d8efe5abc3c870]

## Forma de trabalho [s:ad51c2300a7a4a44]
- [stated] Regras da doc: decidido primeiro, depois Propostas e Em aberto; proposta aceita sobe para decidido; arquivo sem pendência vira `validado: true` [s:ad51c2300a7a4a44]
- [stated] Ordem: luta no terminal → cliente Unity → local com save em arquivo até um amigo jogar a Travessia → rede, login e banco; contrato de mensagens desde a noite 1 [s:ad51c2300a7a4a44]
- [stated] Orçamento de noites por fase (noite = 30 min ou mais de trabalho); estourou 1,5×, corta escopo, nunca troca de stack; 1 linha por noite em noites.md; stack só reabre com a fase em curso jogável [s:ad51c2300a7a4a44]
- [stated] Também trabalha no Eras durante o dia, enquanto espera respostas do Claude no trabalho, não só à noite [s:ad51c2300a7a4a44]
- [stated] Quer a gestão do projeto (tarefas, próximo passo, visão macro) o mais automatizada possível no Claude Code, como um PM/Scrum Master, sem nunca commitar por ele [s:ad51c2300a7a4a44]
- [stated] Ele escreve a regra, os casos de teste e o resultado esperado; a IA só escreve o laço que roda a tabela; IA pode em simulador, scripts e testes de rede, nunca na regra nem no que o jogador vê [s:ad51c2300a7a4a44]
- [stated] Teme o ciclo de usar IA, sentir que o jogo não é dele e recomeçar do zero [s:ad51c2300a7a4a44]
- [stated] Acha que refinar em documento foi mais barato que reescrever código; o problema foi não parar de refinar [s:ad51c2300a7a4a44]
- [stated] Acha que parte do loop vem de ter mais tempo longe do PC (trabalho, cansaço) do que para codar: nesse tempo conversa com a IA e adianta regras e docs, gerando regra demais [s:ad51c2300a7a4a44]
- [stated] Em nov/2026 está pensando em transformar o Eras num RPG single player que simula um MMORPG (mundo vivo e persistente, economia e mercado movidos por bots/NPCs com perfis, sem foco social); dúvida entre C# no Unity e Go, e entre 2D e 3D; decisão pendente [s:ad51c2300a7a4a44]
- [stated] Sente que esquece as próprias decisões e regras depois de um tempo (fica só com fragmentos) e acha que isso, mais que a IA em si, fez o jogo não parecer dele; escrever doc para fixar não funciona para ele com o pouco tempo; prefere revisar por um "dojo" diário (ler, ouvir, exercitar o porquê, o caminho e as opções descartadas de cada regra) [s:ad51c2300a7a4a44]
- [stated] Planeja um devlog no Medium Dev .init com posts pequenos: origem (Hello Eras do Brasil), receios com a IA, como o jogo mudou várias vezes, e como o desejo de jogar e de fazer o jogo se misturou [s:ad51c2300a7a4a44]
- [stated] Em 03/11/2026 cogitou uma POC descartável feita com IA (offline, em outra linguagem) e um vídeo de pitch básico, por curiosidade de ver o jogo e para validar a ideia com outras pessoas; decisão pendente [s:1c4e8b72b432c608]

## Stack e arquitetura [s:ad51c2300a7a4a44]
- [stated] Go no núcleo, simulador e servidor (22/10/2026); cliente Unity fino; servidor próprio, um processo, WebSocket + JSON, sem framework; navegador só como playground; só PC (Steam) [s:ad51c2300a7a4a44]
- [stated] Núcleo: função pura aplicar(estado, ação, agora) → (estado, eventos); estado é dado puro, sem map nem ponteiro compartilhado; não mexe na entrada; só inteiros; semente e contador do sorteio dentro do estado [s:ad51c2300a7a4a44]
- [stated] Tempo: combate por turno, sem relógio; mundo e automático por Atividade com início e fim, loop de 1 s só checa o que venceu [s:ad51c2300a7a4a44]
- [stated] Estrutura Go: um go.mod na raiz; core/, sim/, data/ na fase 1; contract/ e server/ na fase 2; sem internal/ e cmd/; evento como interface com uma struct por tipo; testdata/ com as tabelas de casos [s:ad51c2300a7a4a44,s:88d8efe5abc3c870]
- [stated] Idioma (25/10/2026): código todo em inglês (pastas, pacotes, tipos, funções); docs continuam em português (nomes e conteúdo); termos próprios do jogo ficam como nome (Travessia, Enraizado); sem tabela de glossário: cada tipo do domínio tem comentário com o termo do doc, ex. `// Strike é o golpe (combate.md).` [s:ad51c2300a7a4a44]

## Mecânicas decididas [s:ad51c2300a7a4a44]
- [stated] Referências centrais: Albion (build, progressão, economia, tier) e Pathfinder 2e (regras) [s:ad51c2300a7a4a44]
- [stated] Resolução: 1d20 + atributo + proficiência + item contra Defesa/CD, graus de sucesso (±10), em combate, coleta, craft e conversa [s:ad51c2300a7a4a44]
- [stated] CD padrão do tier = 8 + bônus esperado do tier (T1 10, T2 11, T3 13, T4 14, alvo de 65% de acerto); ±2 por tier de diferença; calibrar no simulador [s:ad51c2300a7a4a44]
- [stated] Efeitos simultâneos: reação antes da ação que a disparou; no mesmo lado, ordem da lista; efeitos de turno do mais antigo para o mais novo [s:ad51c2300a7a4a44]
- [stated] Combate em instância com o cenário do lugar; turno por lado; 3 ações + 1 reação, máx. 2 golpes por turno; inimigo até 4 ações; manual pausa sem limite; andar custa ação [s:ad51c2300a7a4a44]
- [stated] Manual é o pilar; automático só com o jogo aberto, limitado por recurso, rende menos que o manual; nada acontece offline [s:ad51c2300a7a4a44]
- [stated] Descanso: após 4h deslogado acumula bônus em pontos para o manual, cumulativo, teto de 3 noites [s:ad51c2300a7a4a44]
- [stated] Você é o que veste: sem classe; skills, passivas e os 6 atributos vêm do equipamento; off-hand desde o T1 [s:ad51c2300a7a4a44]
- [stated] Progressão em duas escadas (maestria = bônus + gate de tier, 8 marcos; especialização = quanto rende) em combate, coleta, refino e craft; sem bônus cruzado [s:ad51c2300a7a4a44]
- [stated] Coleta: um teste por nó; ferramenta coleta até 1 tier acima; T1 sem ferramenta; falha crítica de craft e refino não perde material [s:ad51c2300a7a4a44]
- [stated] Durabilidade em duas camadas (atual e teto); sem full loot; ao morrer volta a pé à última cidade em que entrou [s:ad51c2300a7a4a44]
- [stated] Era 1: 1500–1654, litoral, Mata Atlântica e engenho, até T4 no lançamento; facções só depois da luta jogável, sem povo real nem religião como lado [s:ad51c2300a7a4a44]
- [stated] Arte: peça de pack + tier pela cor; bichos só comprados já animados; visual não pode ser infantil [s:ad51c2300a7a4a44]
- [stated] Sem PvP no MVP [s:ad51c2300a7a4a44]
- [stated] Como jogador, largou jogos de turno pela ação picotada, sem fluxo; dúvida em aberto para o Eras (fluidez ou tempo real com pausa), a avaliar descansado [s:ad51c2300a7a4a44]
- [stated] Não gosta de turno com gerenciamento de grupo (BG3, Divinity); quer controlar só o próprio personagem [s:ad51c2300a7a4a44]

## Travessia (MVP) [s:ad51c2300a7a4a44]
- [stated] O MVP vai até o T2 [s:d8863b1fc711a11d]
- [stated] Tutorial no estilo da ilha do Albion, 15 passos, mesmas regras das Eras; separada da Era 1 por bruma; o Mateiro leva e traz, com volta [s:ad51c2300a7a4a44]
- [stated] A morte é ensinada só com aviso, sem matar o jogador de propósito [s:ad51c2300a7a4a44]

## Lore [s:ad51c2300a7a4a44]
- [stated] Escreve a lore à mão; Brasil e folclore só como inspiração indireta, sem religião, sem nomes fortes e sem nada que exista de verdade [s:ad51c2300a7a4a44]
- [stated] A Raiz rachou, abriu rasgos que alteram os seres vivos e fugiu, puxando terras para se firmar [s:ad51c2300a7a4a44]
- [stated] Enraizados vieram com a terra e vivem normalmente dentro da sua era; Puxados (jogadores) vieram sozinhos e não morrem de vez; Perdidos perderam a consciência [s:ad51c2300a7a4a44]
- [stated] Pendente: quem recebeu o presságio; nome do jogo [s:ad51c2300a7a4a44]

## Referências para estudar [s:ad51c2300a7a4a44]
- [stated] asked to remember (29/10/2026): está testando a demo de SoulStonia (https://store.steampowered.com/app/4448060/Soulstonia_Demo/); o jogo se joga manualmente como um jogo normal, mas tem um lado idle; achou interessante o automático, chamado "Objetivos": o jogador monta uma lista de objetivos (coleta, produção, combate, locais) com quantidade, e o jogo segue na ordem; se um objetivo de produção depende de material, ele acrescenta sozinho o objetivo de coletar esse material (ex.: fabricar barra de bronze acrescentou coletar minério de cobre e de estanho, como subobjetivos). Na estação de refino de minério, como no RuneScape, você deposita o que quer para depois fazer. Ainda não sabe como funciona para combate [s:ad51c2300a7a4a44]

## Arquivo de memória `/preferences.md` (atualizado 2026-10-23)

---
name: preferences
description: Como o Sidinei quer as respostas e o trabalho no Project Eras do Brasil — tamanho, formato, autoria do código e como decidir
sources: [cowork]
---

- [stated] Respostas e documentos curtos e fáceis de validar: dev solo, ~2h por noite com interrupções, aprendendo game design; texto denso gerado por IA vira leitura demais e a validação fica falha [s:ad51c2300a7a4a44]
- [stated] A explicação vai no chat; o documento guarda só a regra ("pense simples, faça simples"; "feito é melhor que perfeito, mas não mal feito") [s:ad51c2300a7a4a44]
- [stated] Docs curtos e visuais (tabela, Mermaid, canvas do Obsidian), poucos arquivos, só separações que façam sentido, com fluxo de leitura claro [s:ad51c2300a7a4a44]
- [stated] Escreve o código do projeto à mão: quer explicação de conceito, comparação de abordagens, ajuda a depurar e revisão do que ele escreveu, não código pronto [s:ad51c2300a7a4a44]
- [stated] Em arquitetura, quer entender cada decisão e o porquê dela [s:ad51c2300a7a4a44]
- [stated] Para escolher algo, quer no máximo 3 opções com o porquê de cada uma; ele decide [s:ad51c2300a7a4a44]
- [stated] Quando algo vira decisão, quer a linha exata e o arquivo onde escrever [s:ad51c2300a7a4a44]
- [stated] Ideia fora da fase atual: marcar como "estacionar" e não desenvolver [s:ad51c2300a7a4a44]
- [stated] Não poetizar mecânica nem poetizar demais a lore [s:ad51c2300a7a4a44]
- [stated] Um chat por assunto: validar arquivo, dúvida de conceito ou revisão de código vai em chat novo do Project; trabalho pesado (debate com agentes, análise do acervo, pacote de arquivos) vai numa sessão aberta só para isso e encerrada no fim [s:ad51c2300a7a4a44]
- [stated] Quando o chat atual já cumpriu o objetivo, ou ficou longo/resumido a ponto de misturar versões, Claude deve dizer que ele cumpriu o papel e sugerir abrir um chat novo, já com o primeiro prompt pronto [s:ad51c2300a7a4a44]
- [stated] Antes de encerrar um chat: a decisão vai para o arquivo certo do repo e o porquê importante vai para o acervo; o que não está escrito no repo não existe [s:ad51c2300a7a4a44]

## Arquivo de memória `/topics/jogos.md` (atualizado 2026-11-06)

---
name: jogos
description: Gostos de jogador no contexto do Eras — o que o afasta de jogos de turno e de visual infantil
sources: [cowork]
---

- [stated] Já largou vários jogos de turno porque a luta fica picotada, sem ação conectada; foi um dos motivos de pensar no Eras com o personagem fazendo as ações para assistir [s:ad51c2300a7a4a44]
- [stated] Não gosta de jogos de visual infantil (descartou Dicey Dungeons e Backpack Battles por isso); acha Stoneshard difícil demais [s:ad51c2300a7a4a44]
- [stated] Também não gosta de visual muito sombrio (descartou Tainted Grail: Conquest e Ring of Pain por isso); procura o meio-termo, nem infantil nem sombrio [s:ad51c2300a7a4a44]
- [stated] Não gosta de jogos de gerenciar grupo (tipo Baldur's Gate); quer controlar só o próprio personagem [s:ad51c2300a7a4a44]
- [stated] Jogo que procura para si: coletar, refinar, craftar armas e armaduras, masmorras e luta em mundo aberto (turno ou ação), que dê para largar no meio quando o bebê chora e voltar depois [s:ad51c2300a7a4a44]
- [stated] Tem receio de jogos de sobrevivência que dependam de grupo/coop; entre Valheim e Enshrouded prefere Enshrouded; acha Scapewatch legal, mas não como jogo principal [s:ad51c2300a7a4a44]
- [stated] Não acha bonito o visual low poly [s:ad51c2300a7a4a44]
- [stated] O que mais procura num MMORPG é economia, mercado e sensação de mundo vivo; quase não usa grupo, guilda, raid ou PvP de larga escala [s:ad51c2300a7a4a44]
