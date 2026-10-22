<!-- Memória da conta do Claude · caminho original: /areas/eras-do-brasil.md · atualizado em 2026-10-19 · export de 21/10/2026 -->

---
name: eras-do-brasil
description: MMORPG do Sidinei, Eras do Brasil — repo antigo (versão idle) a ser arquivado e recomeço com camada ativa
sources: [chat]
aliases: [Eras do Brasil, eras-do-brasil, Eras, eras-do-brasil-idle]
---

- [stated] Eras do Brasil é seu MMORPG, projeto pessoal de longo prazo; sente ansiedade/imediatismo para tê-lo jogável [s:e581350e249b7c38]
- [stated] Repo da versão idle renomeado para github.com/sidinei-silva/eras-do-brasil-idle (código reaproveitável); repo novo é github.com/sidinei-silva/eras-do-brasil; histórico do projeto em github.com/sidinei-silva/ErasDoBrasil-Historico-Legado [s:e581350e249b7c38]
- [stated] Dúvida recorrente entre Go e JS/TS (Node, Deno, Bun) para o servidor; acha o código Go existente embrionário e fácil de refazer [s:e581350e249b7c38]
- [stated] Novo rumo (out/2026): combate por turno instanciado com modo ativo e caçada automática no mesmo motor, 3D low poly com câmera levemente inclinada [s:e581350e249b7c38]
- [stated] Vai escrever à mão o documento de design do novo Eras e depois rever o plano de desenvolvimento e os Projects do Claude ligados ao jogo [s:e581350e249b7c38]
- [stated] O novo repo do Eras será um monorepo (backend, cliente, dados, docs); a doc de design será enxuta, com um ponto de entrada em docs/ linkando um arquivo por mecânica e um de arquitetura [s:e581350e249b7c38]
- [stated] Não quer engine potente; planeja playgrounds locais para testar uma mecânica por vez, começando pelo combate [s:e581350e249b7c38]
- [stated] Indo com TypeScript em tudo (Node) e cliente web empacotado para a Steam via Electron (decisão de out/2026, registrada no docs/arquitetura.md do repo) [s:e581350e249b7c38]
- [stated] IA foi usada no Eras do começo ao fim (ideia, lore, mundo); teme o ciclo recorrente de tentar sem IA, voltar a usar, sentir que o jogo não é seu e recomeçar do zero [s:e581350e249b7c38]
- [stated] Começou a usar IA (agente) para escrever parte do código do core e do playground do repo novo, em out/2026 [s:e581350e249b7c38]
- [stated] Já teve uma versão jogável de que gostou, com NPCs que tinham "vida", no repo github.com/sidinei-silva/eras-do-brasil-game-arquivado [s:e581350e249b7c38]
- [stated] Em 14/10/2026 declarou o jogo totalmente indefinido (sem definição vigente do que o projeto é); a versão com NPCs vivos também já tinha passado por um reset do zero em abr/2026, e depois dele voltou a usar IA [s:e581350e249b7c38]
- [stated] Núcleo original (WhatsApp, jul/2023, sem IA): mistura de jogo de turno com Albion/New World, "sua classe é o que você usa", subir conhecimento da arma e aprender armas novas (hoje pensa em NPCs ou facções por estilo de arma, sem grupo de companheiros) [s:e581350e249b7c38]
- [stated] Sensações que quer no jogo (out/2026): você é o que veste / troca de classe livre sem alt; o trabalho de life skill importa via mercado vivo; poder parar sem perder nada; NPC com vida é discutível no online [s:e581350e249b7c38]
- [stated] Quer tirar do Eras o papel de "jogo perfeito para mim"; ainda dividido entre usá-lo para recuperar o ofício à mão e tê-lo pronto logo; nunca fez arte, lore ou game design, e teme que fazer tudo do zero leve 6+ anos [s:e581350e249b7c38]
- [stated] Ideia em out/2026: armas e skills baseadas no Albion; semi-idle com modo ativo (escolhe skills), modo com ordem de prioridade (como RavenIdle) e automação mais rica (como Dominus Automa); mundo aberto, PvP, coleta e craft, com combate em janela instanciada; movimento no combate não será manual em tempo real [s:ad51c2300a7a4a44]
- [stated] Considerando PvP assíncrono com "eco" (cópia da build/config do jogador), em que o desafiado só recebe uma mensagem com o resultado, parecido com o Lootborne [s:ad51c2300a7a4a44]
- [stated] Em 15/10/2026 disse que o Eras está parado; quer estudar mecânicas de game design de forma solta, sem vincular ao projeto nem tomar decisões [s:ad51c2300a7a4a44]
- [stated] Lore (17/10/2026): formato persistente com eventos; "o mundo não espera por herói" continua o norte; jogador é pessoa comum, então tira o tema "Dom"; quer a Travessia (tutorial) como estado espiritual/abstrato, não ilha física; após o tutorial vem a Era 1, e eras novas devem somar opções, não poder [s:ad51c2300a7a4a44]
- [stated] Lore (17/10/2026): Ruptura como processo contínuo; encantados só como inspiração indireta, evitando polêmica (jogo inspirado no Brasil e folclore, não reprodução histórica); cogita "Eternos" (pessoas que não morrem e desaparecem) como tema central e nome, ex. "Eternos: As Eras do Brasil"; Perdidos = quem não conseguiu se montar ou escolheu ficar; origem do jogador (isekai ou não) ainda em aberto [s:ad51c2300a7a4a44]
- [stated] Lore (18/10/2026): aceita Eternos (puxados sozinhos, não morrem; morrer = se desfazer) e Enraizados (puxados junto com a terra; as facções de cada era); origem moderna do jogador como hipótese a reavaliar; Brasil e eras só como inspiração, sem nomes fortes; acha o enredo misterioso demais e quer história com começo, meio e fim, motivação clara e uma introdução [s:ad51c2300a7a4a44]
- [stated] Plano de estudo (out/2026): pedir à IA projetos de referência separados, por fatia ou mecânica, para ler e estudar, e depois implementar à mão no próprio repo [s:ad51c2300a7a4a44]
- [stated] Estrutura (18/10/2026): eras como expansões, Era 1 = Brasil de 1500–1654; cada era tem tiers de 1 a 6, então equipamento de uma era pode ser mais ou menos forte que o de outra; procura um modelo de isekai com chegadas contínuas de novos jogadores e motivo coerente para ficar ou sair/voltar [s:4842a924735c6910]
- [stated] Escrevendo (18/10/2026) o documento de apresentação do Eras: MMORPG sandbox para PC, público com pouco tempo e sujeito a interrupções, projeto "for fun" sem intuito comercial agressivo, motivado pelo sonho de ter um MMORPG feito por ele; a lista de mecânicas vira índice com uma página por mecânica [s:ad51c2300a7a4a44]
- [stated] Ideia (19/10/2026): NPCs com vida em duas categorias — da vila (guardam fofoca, dão tarefas, comerciantes e tutores das estações) e viajantes por arquétipo (coletor, caçador, comerciante, crafter) para sustentar e equilibrar a economia com poucos jogadores e carregar lore; decidiu, por enquanto, fazer só NPCs de vila e viajantes, com os viajantes apenas levando e trazendo fofocas [s:ad51c2300a7a4a44]
