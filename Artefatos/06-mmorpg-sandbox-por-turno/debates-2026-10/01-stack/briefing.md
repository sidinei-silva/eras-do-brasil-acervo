# Briefing: debate de stack do Eras do Brasil

## O jogo (decidido)
- MMORPG sandbox por turno, só PC (Steam). Dev solo, ~2h por noite com interrupções, projeto "for fun".
- Combate por turno em grid visível: d20 + modificadores, graus de sucesso, 3 ações + 1 reação. Modo manual e modo automático (lista de prioridade); o servidor roda lutas sem cliente aberto (caçada automática/offline).
- Você é o que veste: equipamento visível no personagem, skills com animação. Proposta atual: 3D low poly, câmera fixa inclinada.
- Mundo aberto em zonas; ao encostar num grupo de inimigos, a luta abre no próprio lugar (grid recortado).
- Servidor autoritativo, núcleo de regras separado da engine, determinismo (semente), conteúdo em JSON.
- Coleta, refino, craft, durabilidade, mercado entre jogadores. PvP assíncrono contra "eco" (cópia da build) depois do MVP.
- MVP: a Travessia (área tutorial estilo ilha do Albion, 5 zonas, 20 passos).
- Roteiro: 0 Base, 1 Núcleo (simulador de console), 2 Luta na tela, 3 Online, 4 Mundo e ofícios, 5 Travessia, 6 MVP. Cliente com placeholder até a fase 6.

## O dev
- Fullstack web desde 2018: Node.js/TypeScript é a linguagem foco; PHP; Go no trabalho (se considera iniciante); já passou por Java, Kotlin, Python, Ruby. Sem experiência real em C#.
- Escreve o código à mão; a IA escreve testes e explica. Iniciante em Blender, nunca fez arte.
- Tendência declarada: Unity no cliente.
- Proposta atual (não confirmada): Unity (cliente) + servidor .NET + núcleo C# (.NET Standard 2.1) compartilhado.

## Histórico de stack (risco: trocar de novo)
- fev/2026: Unity tático cooperativo (P2P), "GDD 100%, código 0%".
- mar–abr/2026: MUD com servidor Go + cliente web; reset para reescrever à mão.
- mai–ago/2026: idle "A Escória" em Go.
- set/2026: fusão idle Go + Godot (3D low poly, web primeiro).
- 03–14/10/2026: TypeScript + Electron + Three.js com core determinístico.
- Material de cada fase: /tmp/claude-0/-home-claude/b2c9da2e-f5a3-5381-b813-0e9760dd9841/scratchpad/acervo (pastas Artefatos/ e Conversas/). Use como evidência; não é verdade absoluta.
