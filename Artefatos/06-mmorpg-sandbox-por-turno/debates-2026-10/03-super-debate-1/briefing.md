# Briefing do debate geral do Eras do Brasil

Fontes: doc atual em /tmp/claude-0/-home-claude/b2c9da2e-f5a3-5381-b813-0e9760dd9841/scratchpad/docs4 (comece por 01-o-jogo.md; mecânicas em mecanicas/; MVP em 02-travessia.md; técnico em 03-tecnico.md; roteiro em 04-roteiro.md). Histórico de todas as fases: /tmp/claude-0/-home-claude/b2c9da2e-f5a3-5381-b813-0e9760dd9841/scratchpad/acervo (evidência, não verdade). Debates anteriores: /tmp/claude-0/-home-claude/b2c9da2e-f5a3-5381-b813-0e9760dd9841/scratchpad/debate-stack e .../debate-arte (resumos rodada1.md, rodada2.md, resultado.md).

## O dev
Solo, ~2h por noite com interrupções (bebê). Fullstack web desde 2018: Node/TypeScript é a linguagem dele; Go no trabalho (iniciante); passou por Java, Kotlin, PHP, Python. Sem C#. Nunca fez arte, iniciante em Blender. Aprende game design do zero. Escreve o código à mão; IA escreve testes e explica. Histórico: 5 trocas de stack e 13 roupagens em 3 anos. Tende a Unity no cliente. Projeto for fun, sem intuito comercial agressivo, só PC (Steam).

## Como o jogo funciona (decidido)
- MMORPG sandbox persistente em zonas ligadas por saídas (como o Albion). Facções controlam zonas; eventos periódicos deixam marcas.
- Você é o que veste: sem classe; skills, passivas e os 6 atributos vêm do equipamento; equipamento visível no personagem; skills com animação.
- Combate por turno em grid visível: d20 + modificadores contra Defesa/CD, graus de sucesso; 3 ações + 1 reação; ataque de oportunidade; Passo. Lutas curtas, poucos oponentes, sem pressão de tempo.
- Grupo de inimigos visível no mapa; encostar abre a luta no próprio lugar (o servidor recorta um grid ao redor).
- Modo manual (pilar; chefes) e automático (lista de prioridade + intenção de movimento). O servidor roda lutas sem cliente aberto (caçada automática; coleta offline limitada pela ferramenta).
- Progressão por uso em duas escadas (maestria = bônus e gate de tier; especialização = quanto rende) em combate, coleta, refino e craft. Regras de tier como o Albion.
- Coleta por nó (um teste por nó), refino, craft, durabilidade em duas camadas, mercado entre jogadores, comprador NPC. Bolsa e montaria.
- NPCs de vila e viajantes que levam fofoca dentro da Era.
- Sem PvP no MVP; depois PvP assíncrono contra "eco" (cópia da build).
- MVP: a Travessia, área tutorial estilo ilha do Albion, 5 zonas, 20 passos.
- Lore: a Raiz (força viva) rachou e fugiu; rasgos alteram seres vivos; Puxados (jogadores, não morrem de vez, sem origem lembrada), Enraizados (povo de cada Era), Perdidos (sem consciência). Brasil e folclore como inspiração indireta, sem religião nem nomes fortes. Não poetizar mecânica nem lore demais. Visual não pode ser infantil.

## Como o jogo funciona (NÃO decidido — o dono reclamou que o debate de stack ignorou isso)
- Movimento fora do combate: clique para andar, WASD, ou ponto a ponto? Tempo real no mundo e turno só na luta?
- Quantos jogadores numa zona ao mesmo tempo? Eles se veem? Luta de outro jogador é visível? Dá para entrar na luta de outro (grupo)?
- Combate em grupo (vários jogadores no mesmo grid) existe? Como funciona turno com jogadores que podem ficar AFK (interrupções)?
- A luta "no lugar" bloqueia aquele pedaço do mapa para os outros?
- Câmera e perspectiva: isométrica, de cima, inclinada (3/4), de lado? Fixa ou com rotação/zoom? Mesma câmera no mundo e na luta?
- Como a interface mostra o modo automático rodando e o que aconteceu offline.

## Resultados anteriores (podem ser derrubados)
- Stack: cliente Unity (só anima, sem regras), núcleo e servidor em TypeScript, servidor próprio Node + WebSocket + Postgres; fallback do cliente: Godot.
- Arte: low poly sombrio + contorno com 2 tons; base Synty Sidekick + Mixamo; 3 faixas de forma por material; cor do Rasgo como única cor viva; teste da silhueta cinza. Câmera não foi debatida.
