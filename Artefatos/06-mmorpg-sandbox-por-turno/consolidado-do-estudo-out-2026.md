# Eras do Brasil — Consolidado do estudo (out/2026)

Oct 18, 2026 · @Sidinei

> Documento do Claude Docs escrito pelo Claude em 18–19/10/2026 na conversa "RavenIdle melee vs ranged mechanics", exportado pelo Sidinei em 19/10/2026.

## Escopo e legenda

Este documento consolida o estudo de game design de outubro de 2026: combate, build, progressão, economia, stack e lore. É a base para os próximos passos, e substitui o consolidado F8 e o consolidado idle onde os dois divergem dele (ver "Divergências com as fases anteriores").

**Referência central:** Albion Online para build e progressão, Pathfinder 2e para as regras de resolução.

**MVP:** o equivalente à ilha tutorial do Albion (a Travessia), com equipamentos e skills até o tier dela. O livro de regras antigo será reescrito do zero, enxuto.

| Status | Significado |
| --- | --- |
| **Decidido** | Você confirmou na conversa |
| **Proposto** | Recomendação feita e não recusada; pode virar decisão |
| **Em aberto** | Pergunta sem resposta, ou deixada para depois de propósito |
| **Adiado** | Fora do MVP por escolha |

## Visão e pilares

MMORPG por turno, para PC na Steam, com história e ambientação inspiradas no Brasil e no misticismo do folclore brasileiro. O jogador não é escolhido: ele se destaca pelo que faz, não pelo que é.

| Pilar | Status |
| --- | --- |
| O jogador é uma pessoa comum, nunca o escolhido ("o mundo não espera por herói") | Decidido |
| Você é o que veste: sem classe, a identidade vem do equipamento | Decidido |
| O combate manual é o pilar; o automático é o fallback. Manual rende mais por luta, automático rende mais por hora | Decidido |
| Mundo persistente, com eventos periódicos que deixam marcas permanentes | Decidido |
| Brasil e folclore como inspiração indireta, sem reprodução histórica e sem nomes fortes | Decidido |
| Eras novas somam opções, não poder | Decidido |
| Só PC (Steam), sem mobile | Decidido |
| Equipamento visível no personagem e skills com animação | Decidido |
| Lutas curtas, contra poucos oponentes, sem pressão de tempo | Decidido |

**Frase do pilar (vinda do MUD):** "Você não foi escolhido. Ninguém aqui foi. Você se destaca pelo que faz, não pelo que é."

## Motor de regras

Toda ação com chance (ataque, coleta, craft, conversa) é resolvida por **1d20 + modificador de atributo + bônus de proficiência + bônus de item contra uma Defesa ou CD** (Decidido). O d20 vem do livro antigo; a build vem do Albion. O Albion não tem dado porque é tempo real; por turno, sem dado, as lutas ficariam determinísticas.

### Economia de ações (Decidido)

- **3 ações por turno + 1 reação**, estilo Pathfinder 2e. Sem ação de rodada completa separada.
- **Penalidade de ataque múltiplo:** -5 no 2º ataque e -10 no 3º.
- Andar custa ação para todos. Skills podem custar 1, 2 ou 3 ações (Em aberto: como definir as de 2 e 3).

### Graus de sucesso (Decidido)

| Resultado | Condição |
| --- | --- |
| Sucesso crítico | Passar da CD por 10 ou mais |
| Sucesso | Igualar ou passar da CD |
| Falha | Ficar abaixo da CD |
| Falha crítica | Ficar 10 ou mais abaixo da CD |

O 20 natural sobe um grau e o 1 natural desce um grau. O crítico é uma faixa, não um número fixo. Os mesmos graus valem para craft (qualidade) e coleta.

### Bônus e penalidades (Decidido)

- **Bônus fixos**, como no Pathfinder. Sem vantagem/desvantagem por 2d20.
- **Tipos de bônus** (item, circunstância, status): bônus do mesmo tipo não se somam, só o maior vale. Você disse que vai usar e precisa estudar mais.

### Faixa de números (Proposto)

Os números do livro são só referência; você liberou trocar.

| Fonte de bônus de acerto | Faixa |
| --- | --- |
| Atributo | 0 a +4 |
| Maestria da linha | +1 a +4 |
| Item (tier) | 0 a +3 |
| **Total no topo** | **cerca de +10** |

Defesas e CDs entre 10 e cerca de 25. Motivo: com crítico por ±10, se a diferença entre os lados passar de 10, um lado crita sempre e o outro nunca acerta.

### Determinismo (Decidido)

- Gerador aleatório **com semente** por luta (algoritmos pequenos como mulberry32 ou xorshift), nunca o aleatório padrão da linguagem.
- As regras só leem estado + decisões + semente. Nada de relógio do sistema.
- Para reproduzir uma luta: guardar semente, estado inicial e lista de decisões.
- **Contas em inteiros.** Ordem fixa para efeitos simultâneos (Em aberto: qual).

## Combate, movimento e grid

O combate é por turno num **grid visível**, com o mesmo motor para o jogador, os inimigos, o automático e os ecos de PvP (Decidido).

### Modos (Decidido)

- **Manual:** o pilar. Chefes pedem manual; farm aceita automático.
- **Automático:** lista de prioridade de skills + intenção de movimento. Regras condicionais (gambits, estilo FF12) ficam para depois.
- **Contínuo:** só apresentação. Os mesmos turnos rodando sem pausa, como o tempo real do Pathfinder: WotR.
- A IA dos inimigos é o embrião do automático: os dois usam o mesmo decisor.

### Movimento (Decidido)

- Movimento básico + skills que movem. Andar custa ação para todos.
- A armadura define quantas casas se anda por ação (pode ser mais de uma).
- **Ataque de oportunidade** como reação. **Passo** (1 casa) desengaja sem provocar.
- Ranged pode atacar colado, com penalidade, e **provoca ataque de oportunidade** ao atirar colado.
- Sair da borda do mapa fica a cargo de skills.
- Movimento por intenção é a forma de configurar o automático; não substitui o clique manual.
- **Ocupação de espaço** (regra do Pathfinder): ninguém termina no mesmo quadrado; atravessar aliado pode; atravessar inimigo exige skill ou manobra.

### Onde a luta acontece (Proposto)

No **próprio lugar do mundo**: o servidor recorta o grid lógico ao redor do encontro e usa como arena, e o cliente mostra o grid tático sobre a paisagem (como Pathfinder e Wartales). Ver "Mundo, zonas e montagem de mapas".

### Em aberto

- Se o ataque de oportunidade é de todos ou só de certos itens (escudo, haste, placa).
- Teto de diferença de velocidade entre armaduras.
- Skills de 2 e 3 ações.
- Ordem fixa de resolução para efeitos simultâneos.
- Variância do d20 em lutas pequenas (3 ações e graus de sucesso ajudam; acompanhar nos testes).

## Build, equipamento e atributos

A identidade do personagem vem dos itens: skills e passivas moram na arma e na armadura, e os atributos também vêm do equipamento (Decidido).

### Build (Decidido)

- Estilo Albion: skills na arma e na armadura, passivas escolhíveis no item.
- Classes saem. As classes do livro antigo viram traços de itens (ex.: Tiro Preparado do Mosqueteiro vira traço do mosquete; Postura Inabalável vira passiva da armadura pesada).
- **Material da armadura** (pano, couro, placa) define defesa, deslocamento e o grupo de passivas disponíveis.
- Sem gemas. Sem encantamento no início.

### Atributos

**Origem: só do equipamento** (Decidido). Se sentir necessidade, entra depois uma base pequena, como tendência da arma.

Os seis atributos mantêm os nomes do livro (Decidido):

| Atributo | Equivalente no Pathfinder | Papel |
| --- | --- | --- |
| Força Bruta | Força | Acerto e dano corpo a corpo pesado |
| Astúcia | Destreza | Acerto à distância e com armas leves, defesa, iniciativa |
| Vigor | Constituição | Vida, resistir a veneno e doença |
| Conhecimento | Inteligência | Magia elemental e técnica, craft |
| Sabedoria Ancestral | Sabedoria | Magia espiritual, percepção, coleta |
| Presença | Carisma | Resistência mental, provocação, magia mental, testes sociais |

**Distribuição nas peças (Proposto):** cada peça carrega o atributo do seu papel.

| Peça | Atributo principal |
| --- | --- |
| Arma | O ofensivo da linha (espada pesada: Força Bruta; arco: Astúcia; cajado elemental: Conhecimento) |
| Peitoral | O do material (placa: Vigor; couro: Astúcia; pano: Sabedoria Ancestral ou Conhecimento) |
| Capacete | Mente e espírito (Presença, Sabedoria Ancestral) |
| Botas | Astúcia, junto com o deslocamento |

**Escala (Proposto):** sem equipamento, tudo em 3 (modificador 0); os itens somam a partir daí, com teto em 6 (+3). A Defesa é 10 + Astúcia + armadura; com placa, a defesa vem da armadura (como o limite de Destreza do Pathfinder).

### Em aberto

- Base neutra (3) ou zero sem equipamento. Pesa na punição da morte: com zero, o personagem nu é inútil.
- Um atributo por peça ou principal + secundário.
- O papel da Presença num jogo sem combate social.
- Quantas skills por arma; se armas do mesmo grupo compartilham skills.
- Trocar de roupa antes de uma conversa muda o bônus de Presença: decidir se isso é desejado.

## Progressão: maestria, especialização e craft

A progressão segue o Destiny Board do Albion em duas escadas, traduzidas para degraus de d20: a **maestria da linha** dá acerto e o gate de tier; a **especialização do item** dá dano (Proposto; você quer ver como funciona na prática).

### Como o Albion faz (referência)

- **Maestria da linha** (ex.: "Lutador com Botas de Placas"): cada nível dá +0,2 de poder para todas as peças da linha, e certos níveis liberam quais itens da linha se pode usar.
- **Especialização do item** (ex.: "Botas de Soldado"): cada nível dá +2 de poder só para aquele item e +0,2 para a linha.
- **Nós iniciais** (Aventureiro Novato, Lutador Aprendiz, Fabricante Aprendiz) são o gate de uso e de craft por tier.
- No craft: maestria e especialização dão qualidade e economia de foco, não acerto.

### A tradução para d20 (Proposto)

Num d20, cada +1 vale 5% de chance, então o acerto sobe em degraus; o dano pode subir aos poucos.

| Escada | Sobe usando | Dá |
| --- | --- | --- |
| Maestria da linha (ex.: Espadas) | Qualquer arma da linha | Gate de tier + acerto em degraus |
| Especialização do item (ex.: Espada Larga) | Só aquele item | Dano e força das skills daquele item; uma fração também enche a maestria |

**Maestria (níveis 0 a 100):**

| Nível | Grau | Acerto | Libera |
| --- | --- | --- | --- |
| 0–19 | Aprendiz | +1 | T1–T2 |
| 20–49 | Competente | +2 | T3–T4 |
| 50–79 | Especialista | +3 | T5–T6 |
| 80–100 | Mestre | +4 | T7–T8 |

**Especialização (níveis 0 a 100):** +1 de dano a cada 20 níveis (máximo +5).

- **O gate de tier é a maestria** (Proposto), como no Albion: sem nível de personagem. Você é "Mestre em espadas e Aprendiz em arcos".
- **O tier continua existindo** como valor do item: bônus de item de 0 a +3, atributos e passivas melhores.
- Os limiares não precisam ser lineares: primeiros degraus rápidos, últimos lentos (o Albion pede 184 mil de 204 mil de fama no nível 72).
- Sempre inteiros e degraus, nunca +0,2.

### Craft (Proposto)

- **Maestria de fabricação** (ex.: Livros): gate de receitas + bônus no teste de craft em degraus.
- **Especialização** (ex.: Tomo de Feitiços): menos material e mais chance de sucesso crítico naquele item.
- Os graus de sucesso decidem a qualidade: falha crítica perde material, falha não produz, sucesso produz, sucesso crítico sobe a qualidade.
- **Qualidade em 5 níveis** (Muito Baixa a Excelente no livro; Normal a Obra-prima no Albion). A raridade do livro se sobrepõe ao tier e pode sair.

## Durabilidade, morte e risco

Nada é destruído por azar: a durabilidade tem duas camadas, mostradas como uma barra, e o teto é o ralo de itens que substitui o full loot (Decidido).

### Duas camadas (Decidido)

| Camada | Cai com | Em zero | Volta com |
| --- | --- | --- | --- |
| **Atual** | Uso (leve) e morte (porcentagem) | Inutilizável até reparar | Reparo em campo ou no NPC |
| **Teto** | Morte no ativo (um pouco) e cada reparo em campo | Inutilizável até restaurar | Materiais ou um item raro descartado ("descarte de luxo") |

Na barra, a parte que o reparo não alcança aparece escurecida.

### Reparo (Decidido)

- **Em campo:** kit por tipo de item; gasta teto.
- **No NPC da cidade:** cobra prata; não reduz o teto.

### Morte (Decidido)

- **No ativo:** perde uma porcentagem de durabilidade, um pouco de teto e parte do inventário carregado, com limite.
- **No automático:** só gasta durabilidade atual.
- **Sem full loot.**
- **Na lore:** um Eterno não morre, ele **se desfaz** e se refaz no último abrigo. O que carregava fica para trás, e os objetos perdem um pouco da memória que guardavam (perda de inventário e de teto). Rumor de sabor, não regra: "quem se desfaz vezes demais começa a esquecer de si".

### Em aberto

- Porcentagens de durabilidade e tamanho da barra (balanceamento).
- Quanto o reparo em campo tira do teto (recomendação: um pouco, para manter o ralo).

## PvP assíncrono com eco

O PvP é contra o **eco** de outro jogador: uma cópia da build e da configuração de automação dele, como no Lootborne e no Backpack Battles (Decidido).

- **Primeira versão:** os dois lados no automático. **Depois:** você no manual contra o eco (Decidido).
- **Sem perda de itens** no PvP (Decidido).
- **Determinismo obrigatório:** semente no aleatório, para replay e auditoria (Decidido).
- **Grupo com ecos contratados**, estilo pawns do Dragon's Dogma (Adiado).
- **Em aberto:** o que o PvP recompensa (ranking, moeda), a compensação do defensor e o pareamento por poder (para o PvP não virar só teste de equipamento).

## Jornada do jogador

A orientação do jogador mistura o modelo do Albion (metas por atividade) com o do Pathfinder (histórias com escolhas), em cinco camadas com papéis diferentes (Proposto).

| Camada | Inspiração | Responde | Custo |
| --- | --- | --- | --- |
| 1. Progressão | Destiny Board | "Como fico mais forte?" | Já desenhada (maestria e especialização) |
| 2. Diário | Diário de Albion | "O que existe para fazer?" | Baixo: são dados |
| 3. Contratos | Quadro de missões | "O que faço agora?" | Baixo: modelos + estado do mundo |
| 4. Jornadas | Pathfinder | "Por que estou aqui?" | Alto: escritas à mão |
| 5. Desafio periódico | Desafio do Aventureiro | "Por que voltar amanhã?" | Baixo; depois do MVP |

- **Diário:** metas reveladas aos poucos, por categoria (combate, coleta, craft, exploração, NPCs). Ensina os sistemas e empurra a experimentar linhas de arma.
- **Contratos:** gerados de modelos e oferecidos por NPCs. Os rumores lidos do mundo viram contratos, e cumprir contratos sobe a afinidade com o NPC. Contratos de caça funcionam no automático.
- **Jornadas:** poucas e marcantes, com testes de d20 fora de combate (Presença para convencer, Astúcia para se esconder). Pedem o manual, como os chefes.
- **Escolhas num mundo compartilhado** afetam o estado do próprio jogador (afinidade, recompensas, falas), nunca o mundo de todos.
- **Recompensas** preferencialmente fora do mercado (experiência de maestria, afinidade, visuais, itens presos à conta), para não inflar a economia.

**Na Travessia (MVP):** uma jornada curta que ensina cada sistema, a primeira página do diário e um quadro com dois ou três modelos de contrato.

## NPCs e mecânicas reaproveitadas

Os NPCs ganham vida em degraus de custo crescente; para o MVP, os degraus 1 e 2, talvez o 4 (Proposto). Num MMO, cada sistema de NPC precisa responder "vale por jogador ou para todos?".

| Degrau | O que é | Custo |
| --- | --- | --- |
| 1. Agenda simples | Lugares por período do dia (o ferreiro na forja de dia, na taverna à noite) | Baixo |
| 2. Rumores lidos do mundo | O NPC lê o estado real do servidor ("apareceu um bicho raro na mata") | Baixo |
| 3. Afinidade por jogador | Relação com cada NPC ou facção; libera preço, receitas, missões | Médio |
| 4. Reação ao que você veste | Peças de uma tradição abrem ou fecham portas | Médio |
| 5. Testemunhas e reputação | Crime visto, Percepção contra Furtividade | Alto |
| 6. Fofoca simulada | NPCs trocando conhecimento ao se cruzar | Muito alto |

**Trabalhadores (Proposto):** NPCs contratados que coletam ou produzem para o jogador, limitados por algo (diários, comida, salário), como os laborers do Albion. É uma resposta possível para a auto-coleta (Em aberto) sem inflar a economia.

### Mecânicas do livro antigo que cabem (Proposto)

| Do livro | Como encaixa |
| --- | --- |
| Sucesso parcial | Coberto pelos graus de sucesso |
| Qualidade em 5 níveis | Os graus do craft decidem; a raridade sai |
| Recursos com raridade e regeneração | Iguais aos nós de coleta do Albion |
| Eventos aleatórios e roteirizados com gatilhos | Mundo aberto; eventos globais depois |
| Economia regional (escassez, caravanas) | Depois do MVP |
| Condições de status | Quase inteiras, trocando vantagem por +2/-2 |
| Concentração | Igual à ação de sustentar do Pathfinder 2e |
| Fusão de itens | Pode ser a mesma mecânica do "descarte de luxo" |
| Névoa de guerra e obstrução | Mapa (explorar revela) e grid (linha de visão, furtividade) |
| Melhoria gradual de item | Quando o encantamento entrar |

## Mundo, zonas e montagem de mapas

Cada zona tem duas camadas, a visual (cena do Unity, só no cliente) e a lógica (JSON, lida pelo servidor e pelo cliente), e o mundo é feito de zonas ligadas por saídas, como no Albion (Proposto).

|  | Camada visual | Camada lógica |
| --- | --- | --- |
| O que é | Terreno, árvores, casas, iluminação | Onde se anda, nós de coleta, surgimento de monstros, NPCs, saídas |
| Onde vive | Cena do Unity | `data/` (JSON) |
| Quem usa | Só o cliente | Servidor e cliente |

### Grid lógico no mundo inteiro (Proposto)

Mesmo com o jogador andando livre na tela, cada zona tem por baixo um grid (ex.: casas de 1 metro) com: se dá para andar, altura e se bloqueia a visão. Isso permite:

1. Validar o movimento no servidor (busca de caminho A\*).
2. Linha de visão e furtividade como dado por casa.
3. Combate no próprio lugar: o servidor recorta o grid ao redor do encontro e usa como arena.

Custo: as zonas precisam ser desenhadas pensando em lutas (espaços abertos, cobertura, desníveis).

### Fluxo de montagem (Proposto)

1. **Esboço** com blocos simples (ProBuilder). Testar andar e lutar antes da arte.
2. **Marcadores lógicos** na cena: nó de minério, surgimento de lobos, NPC com agenda, saída para a zona X.
3. **Exportar** com uma ferramenta de editor própria (`tools/`): o grid é calculado lançando raios de cima para baixo em cada casa; os marcadores viram lista.
4. **Arte** com um kit modular low poly que encaixa no grid. Exportar de novo.
5. **Validar:** casas inalcançáveis, nó dentro de parede, saída sem destino.

**Fixo na cena:** terreno, construções, decoração. **Enviado pelo servidor:** nós, monstros, NPCs, jogadores; o cliente cria o modelo a partir de um catálogo de prefabs indexado por tipo.

**Em aberto:** arena no lugar do encontro ou arenas fixas por bioma (a recomendação é no lugar).

## Stack, arquitetura, rede e repositório

A recomendação, avaliada só pelo formato do jogo e sem levar em conta o que você já domina, é **3D low poly com câmera fixa inclinada, cliente em Unity (C#), servidor .NET (C#) e um núcleo de regras em C# puro compartilhado** (Proposto, base do estudo). O Godot 4 com C# é a segunda opção, bem próxima.

### Por quê

- **3D, não 2D:** equipamento visível + animações. Em 3D, um esqueleto e um conjunto de animações servem para todos os itens; em 2D por quadros, cada item precisa ser desenhado em cada quadro de cada animação e direção.
- **Fora:** Phaser, PixiJS e Ebitengine (2D); linguagem pura (seria construir editor e ferramentas); web 3D + Electron (sem editor, animação fraca); Unreal (pesado demais para o caso).
- **Unity sobre Godot:** a loja de assets (personagens modulares low poly e animações) reduz o custo de arte, a parte mais difícil para um dev solo. O Unity Personal é gratuito até US$ 200 mil de receita em 12 meses, e a taxa por instalação foi cancelada.
- **C# dos dois lados:** um único núcleo de regras. O cliente calcula a chance de acerto para a interface e reproduz replays com a mesma regra do servidor.
- **Caso de referência:** o Albion usa Unity no cliente, servidor em C# e um núcleo de simulação independente do Unity.

### Arquitetura

1. **Núcleo (`core`, C# puro, .NET Standard 2.1):** combate, graus, gerador com semente, durabilidade, definições lidas de `data/`. Sem dependência do Unity.
2. **Servidor (.NET, autoritativo):** dono do mundo; PostgreSQL; roda lutas sem cliente aberto (caçada automática, PvP com eco); emite um registro de eventos por luta.
3. **Cliente (Unity):** não decide nada; envia comandos e anima eventos; monta o personagem a partir do equipamento (esqueleto + malhas + pontos de fixação).
4. **Login:** Steamworks (ticket de sessão validado pelo servidor).

### Rede (Proposto)

Uma conexão persistente e bidirecional, com três tipos de mensagem:

| Tipo | Direção | Exemplo |
| --- | --- | --- |
| Comando | Cliente → servidor | "Andar até X", "usar skill Y no alvo Z" |
| Evento | Servidor → cliente | "A atacou B, rolou 17, sucesso, 8 de dano" |
| Estado | Servidor → cliente | Estado completo ao entrar ou reconectar; depois só o que mudou |

- **MVP:** WebSocket ou TCP, com mensagens em JSON.
- **Depois:** UDP com canais (ex.: LiteNetLib, C# dos dois lados): canal confiável para combate, canal "só o mais recente" para posições; formato binário (MessagePack ou Protobuf).
- O transporte fica atrás de uma camada própria, para trocar sem mexer no jogo.
- O mundo aberto precisa de área de interesse (o servidor só manda o que está perto do jogador).

### Repositório (Proposto)

```
/
├── core/     regras (C# puro)
├── server/   servidor .NET
├── sim/      simulador de console
├── tests/    testes do núcleo
├── tools/    exportador de mapas etc.
├── data/     itens, skills, zonas (JSON)
└── client/   projeto Unity
```

- O `core` entra no Unity como **pacote local** com arquivo de definição de assembly (sem cópia de DLL).
- **Git LFS** desde o início para modelos, texturas e áudio; `.gitignore` do Unity (`Library/`, `Temp/`, `Obj/`, `Logs/`, `Build/`); `.meta` visíveis e serialização em texto; conferir a cota de LFS do provedor.
- Ficar numa versão LTS do Unity.

## Processo: fatias e método de estudo

O caminho até o MVP tem oito fatias, cada uma terminando com algo jogável ou testável, e a rede só entra na quinta (Proposto).

1. **Núcleo sem tela:** simulador de console com d20, graus, 3 ações, penalidade de ataque múltiplo, grid, ataque de oportunidade, semente e registro de eventos. Mil lutas sozinhas.
2. **Combate visível:** no Unity, uma arena, dois personagens, uma arma; o cliente anima o registro de eventos. Jogável no manual.
3. **Equipamento:** trocar arma e armadura muda o visual e os atributos. Duas ou três linhas de arma.
4. **Automático:** lista de prioridade + intenção de movimento; a IA dos inimigos usa o mesmo sistema.
5. **Servidor e rede:** o combate passa a rodar no servidor. Login, persistência, reconexão.
6. **Mundo aberto:** uma zona, grid lógico exportado, movimento validado, luta no lugar.
7. **Economia básica:** coleta, craft, maestria e especialização, durabilidade, reparo no NPC.
8. **Travessia:** jornada curta, primeira página do diário, contratos, NPCs com agenda e rumores. É o MVP.

### Método de estudo (Decidido, com regras Propostas)

Você vai pedir à IA projetos de referência separados, por fatia ou mecânica, para ler e estudar, e depois implementar à mão no seu repositório. Regras para isso funcionar:

- Repositório de referência sempre separado; nunca copiar trechos.
- **Fechar a referência antes de escrever.** Se travar, olhar, fechar e continuar.
- Pedir a referência pequena, comentada e com um texto explicando cada decisão.
- **No núcleo de regras, pedir testes em vez de implementação:** a IA escreve os testes a partir das regras; você escreve o código que passa neles.
- **Na parte de Unity, a referência vale mais** (APIs e editor).
- Fechar cada fatia com revisão do seu código e comparação com a referência.

Na prática: fatia 1 com testes da IA e código seu; fatias 2, 3 e 6 com projetos de referência; revisão no fim de cada uma.

## Lore, enredo e mundo

A base do mundo está decidida (Eternos, Enraizados, Perdidos, a Ruptura contínua), mas a espinha do enredo ainda não: falta uma história com começo, meio e fim, uma motivação clara e uma introdução. Mistério demais não motiva.

### Decidido

- **Formato:** mundo persistente com eventos periódicos.
- **O jogador é uma pessoa comum.** O tema "Dom" sai (quem não é herói não tem dom).
- **"Você é o que veste" na lore:** os itens carregam a memória de quem os fez e usou; vestir uma tradição dá as skills dela. O craft continua sem justificativa mística: fazer uma peça é reproduzir uma técnica.
- **A Ruptura é um processo contínuo**, não um evento único: a Raiz está rasgando. Rasgos pequenos puxam pessoas; rasgos grandes puxam terras inteiras com seu povo (as eras); rasgos momentâneos são os eventos.
- **Inspiração indireta em tudo**, para evitar polêmica: os encantados (gente que some sem morrer) inspiram o tema, sem usar a religião. O jogo não reproduz a história.
- **Os Eternos:** pessoas puxadas sozinhas, que não morrem. Nome cogitado: "Eternos: As Eras do Brasil". Cuidado: "Eternos" é o título em português do filme da Marvel, o que atrapalha busca na Steam.
- **Os Enraizados:** gente puxada junto com a sua terra, que vive como se ainda estivesse no seu tempo. **As facções de cada era são Enraizados.**
- **Os Perdidos:** quem não conseguiu se montar, ou escolheu ficar.
- **A morte é se desfazer** (ver "Durabilidade, morte e risco").
- **A Travessia (tutorial)** é um estado, não um lugar: um "entre" feito de fragmentos de memória. Sair dela é conseguir se montar; o arco do tutorial do idle (acordar sem nada → vestir-se → atravessar) já conta isso.
- **Depois da Travessia vem a Era 1** (inspirada em 1500–1654). Cada era tem tiers de 1 a 6; eras novas somam opções, não poder.
- **Brasil como inspiração, sem nomes fortes:** nomes descritivos (a era das naus; os de Além-Mar, os da Terra, os do Fundo; a Costa, a Mata Fechada) e criaturas com nomes próprios do jogo.

### Hipótese a reavaliar

- **Origem moderna do jogador** (um isekai brasileiro). Referências de isekai sem escolhido: Grimgar (gente comum sem memória, virando soldado para sobreviver) e Log Horizon (milhares presos construindo uma sociedade).

### Proposto, ainda não aceito

- **Dois relógios:** o mundo existe fora do tempo; as datas dizem de onde as coisas vieram, não quando o jogo acontece. O "agora" é contado pelos grandes rasgos ("três estações depois do Rasgo do Ouro").
- **Uma era nova chega como um grande rasgo:** o lançamento de expansão é um evento do mundo.
- **Eras sem aumento de poder:** cada era tem seu gradiente de tier da borda ao centro; mesmo teto de bônus de item por tier; o gate é a maestria, não a era; cada peça nova é melhor em alguma situação, não em geral.

### Em debate (brainstorm de 18/10)

A última proposta (voltar para casa, companheiro na Travessia, um ato por era) **não convenceu**: empilhou conceitos sem espinha, e "voltar para casa" briga com um MMO persistente. As opções em aberto:

| Pergunta | Opções |
| --- | --- |
| Que "não escolhido" é o jogador? | Colono entre milhares (Asheron's Call, New World); à sombra de heróis canônicos (LOTRO); membro de uma causa (Dark Age of Camelot); construtor (Wurm); ou uma mistura |
| Qual a motivação? | Voltar para casa (termina, e pode virar motivo do antagonista); sobreviver e se estabelecer; pertencer; descobrir; deixar legado |
| Como as eras se conectam? | A: independentes; B: em sequência; C: independentes, ligadas por um fio que avança um passo por era (a separação que você propôs entre narrativa do mundo e narrativa das eras) |
| Por que a crise nunca termina (o fio)? | Ciclo natural da Raiz; ferida que reabre; alguém rasgando de propósito (ex.: Eternos rasgando para voltar para casa); algo do outro lado empurrando |
| Onde a Raiz se encaixa? | Lugar (o centro do mapa), ser (se defende pelos Guardiões) e memória (Ecos e itens); as três se completam |
| Teremos antagonista? | Em camadas: da era (local), do mundo (o fio), a força (o mistério longo), e lutas sem vilão (o Guardião da Fenda) |
| A introdução | Precisa existir e ambientar; uma opção barata é narração sobre ilustrações paradas, com som |

**Aviso:** "imortais que não são escolhidos porque todos são" é a premissa de New World (Aeternum). O que diferencia precisa ficar claro.

### Material reaproveitável das fases anteriores

- **MUD:** a frase do pilar; a causa trágica da Ruptura (um ritual para proteger a terra de um presságio falhou, e a profecia se cumpriu pela tentativa de evitá-la); Restauradores e Fatalistas; o Guardião da Fenda (resposta imunológica da Raiz, não vilão); uma facção antagonista que avança sozinha a cada temporada; o arco Tensão → Instabilidade → Apogeu → Legado decidido pela comunidade; os Ecos para quem chega depois.
- **Livro antigo:** a Raiz do Mundo; a Linha Eco-Histórica (uma lista pronta de rupturas); as 18 mini-campanhas, a reescrever ancoradas em região em vez de classe.
- **Idle:** a Travessia e seu tutorial de 18 passos; o forasteiro sem origem; as eras lado a lado ao redor da Raiz; a Era 1 projetada (22 zonas, 43 mobs, 17 NPCs, 3 facções); "O que fica na borda"; "lore é puxada, não empurrada".

## Divergências com as fases anteriores

O estudo reabre várias decisões do idle e do F8 porque o formato mudou (de idle sem controle de movimento para combate ativo por turno). Seguindo a regra do próprio idle, "reabrir exige dizer o que mudou e por quê", ficam registradas aqui.

| Tema | Antes | Agora | Por quê |
| --- | --- | --- | --- |
| Resolução | Idle: combate D20 "morto de vez" | d20 estilo Pathfinder 2e | Por turno, sem dado, as lutas ficam determinísticas |
| Espaço | Idle: sem posicionamento; F8: três filas | Grid visível | Movimento e distância viraram parte do combate |
| Turno | F8: andar + item + uma ação | 3 ações + 1 reação | Mais escolhas por turno; equilíbrio do ranged |
| Tempo de combate | Idle: dois relógios, sem turno | Turnos; o modo contínuo é só apresentação | Manual como pilar |
| Progressão | Idle: maestria vira IP (+12 por nível) | Maestria em degraus de acerto + especialização em dano | No d20 não existe +0,2 |
| Atributos | Idle: derivados do IP | Seis atributos do livro, vindos das peças | Encaixam no d20 |
| Mini-campanhas | Idle: fora do design | Reaproveitáveis como jornadas regionais | Jornadas é a camada escrita à mão |
| Durabilidade | F8: arma não desgasta com uso; fuga falha desgasta | Uso desgasta (leve); teto cai com morte e reparo em campo | Teto como ralo de itens |
| Stack | Idle: Go + Godot web; F8: TypeScript + Electron + Three.js | Unity + .NET + núcleo C# | Equipamento modular e animação com menos arte; só PC |
| PvP | Idle: consentido, com flag e perda por banda | Assíncrono contra eco, sem perda de itens | Escolha do estudo |
| Lore | Idle: forasteiro sem origem; MUD: Ruptura de 1497 e Dom | Eternos, Enraizados, Ruptura contínua, sem Dom | Pessoa comum, sem escolhido |

## Em aberto e riscos

A maior pendência é a espinha do enredo; nas mecânicas, o que falta é sobretudo número e alguns detalhes de regra.

### Em aberto

- [ ] Espinha do enredo: modelo de jogador, motivação, conexão entre eras, o fio, antagonista, introdução
- [ ] Origem do jogador (moderna ou não) e nome do jogo
- [ ] Ataque de oportunidade universal ou por item
- [ ] Skills de 2 e 3 ações; quantas skills por arma; skills compartilhadas no grupo de armas
- [ ] Base dos atributos sem equipamento; atributos por peça; papel da Presença
- [ ] Fração da especialização que alimenta a maestria
- [ ] Ordem fixa para efeitos simultâneos
- [ ] Recompensas do PvP com eco e compensação do defensor
- [ ] Auto-coleta (ou trabalhadores NPC)
- [ ] Balanceamento: maestria, especialização, CDs, defesas, durabilidade, velocidade das armaduras

### Riscos

| Risco | Mitigação atual |
| --- | --- |
| Escopo: combate tático, automação, PvP, mundo, coleta, craft e economia, para um dev solo | MVP só na Travessia; rede só na fatia 5 |
| Qualidade da IA (inimigos, automático e ecos dependem dela) | Mesmo decisor para todos; simulador de console |
| Manual × automático: se o manual render demais, o auto pune; se render igual, ninguém joga manual | Manual rende por luta, auto por hora; jornadas e chefes pedem manual |
| Variância do d20 em lutas pequenas | 3 ações, graus de sucesso; observar nos testes |
| Faixa de números estourar o crítico por ±10 | Orçamento de bônus até cerca de +10 |
| Inflação (caçada automática gerando loot) | Durabilidade com uso, reparo pago, recompensas fora do mercado |
| PvP virar teste de equipamento | Pareamento por poder; sem perda de itens |
| Volume de conteúdo (skills e passivas em cada peça; duas escadas por linha e item) | Pacotes de assets; dados em JSON; uma tela de progressão como o Destiny Board |
| Eras "horizontais" virarem verticais na prática | Cada peça nova melhor numa situação, não em geral |
| Núcleo de regras depender do Unity | Biblioteca isolada desde o primeiro dia |
| Referências de IA virarem o projeto | Regras do método de estudo |
| Tema sensível (colonização, povos, religiões vivas) | Inspiração indireta, sem nomes fortes; pesquisa e leitura cuidadosa |
