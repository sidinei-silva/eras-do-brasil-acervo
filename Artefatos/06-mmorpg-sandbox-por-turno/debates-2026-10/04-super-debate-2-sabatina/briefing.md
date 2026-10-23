# Sabatina do Eras do Brasil: briefing

Você está numa sabatina com 22 agentes em 7 grupos. O objetivo é fazer brainstorm, atacar decisões já tomadas, propor coisas novas e defender a sua posição com força, sempre dizendo POR QUÊ.

## O dono e a restrição que vale para tudo
- **Dev solo:** cerca de 2h por noite, com interrupções (bebê). Fullstack web desde 2018, com Node/TypeScript como linguagem principal; Go no trabalho, ainda iniciante; não sabe C#. Iniciante em Blender e em game design.
- **Trabalho:** escreve o código à mão; a IA explica, revisa e ajuda nos testes.
- **Histórico:** em 3 anos, trocou de stack 5 vezes e de roupagem 13 vezes. Toda proposta precisa caber nisso. Proposta que só funciona com equipe ou com tempo integral está errada.
- **Escopo:** projeto for fun, só PC (Steam).
- **Docs:** curtas e fáceis de validar.

## Material (use, cite e confronte)
- **Jogo atual:** `scratchpad/repo` (github.com/sidinei-silva/eras-do-brasil). Comece por `docs/01-o-jogo.md`; mecânicas em `docs/mecanicas/`, MVP em `docs/02-travessia.md`, técnico em `docs/03-tecnico.md`, roteiro em `docs/04-roteiro.md`.
- **Histórico de tudo:** `scratchpad/acervo` (github.com/sidinei-silva/eras-do-brasil-acervo). Use `LINHA-DO-TEMPO.md` (as 13 roupagens e por que cada uma morreu), `MECANICAS-E-CONTEUDOS.md`, `Conversas/` e `Artefatos/` (código antigo em Go, Unity, Godot e TS). Use o acervo para **confrontar** decisões ("isso já foi tentado em X e morreu por Y") e para defender os seus pontos.
- **Debates anteriores:** `scratchpad/sabatina/00a-compilado-stack-e-arte.md` e `00b-compilado-super-debate.md`.
- **Caminhos:** `scratchpad` = `/tmp/claude-0/-home-claude/b2c9da2e-f5a3-5381-b813-0e9760dd9841/scratchpad`.

## Decisões do dono (22/10). Pode questionar, mas o ônus é seu.
1. **Offline:** nada acontece offline, nem coleta nem refino. O jogo precisa estar aberto. O automático existe com o jogo aberto, com limite.
2. **Morte:** você volta na última cidade em que ENTROU, como no Albion, não na mais próxima.
3. **Combate:** é uma instância que mostra o cenário do lugar do mapa. Quem passa perto vê o personagem parado ao lado do grupo de mobs, com um símbolo de combate.
4. **Era 1:** inspirada em 1500–1654 (Descoberta e Colonização). A paisagem segue em aberto.
5. **Já fechado no debate anterior:**
   - mundo em tempo real com clique para andar; turno só na luta;
   - câmera 3/4 fixa (~50°, FOV ~30°, 2 zooms);
   - grupo fora do MVP;
   - mercado único;
   - low poly sombrio + contorno;
   - IA nunca em arte, voz ou placeholder.

## Pauta da sabatina (opine em tudo; aprofunde na sua especialidade)
1. **Funcionamento:** o que as decisões 1–3 mudam. Como o jogo atende quem tem pouco tempo sem offline? Qual é o limite do automático? Como funciona a instância de combate e a reconexão?
2. **Engine do cliente:** Unity × Godot × web (Three.js, Babylon, PlayCanvas). Por quê.
3. **Servidor:** linguagem (TS × Go × outra); framework pronto (Colyseus, Nakama, SpacetimeDB, Heroic Labs) × próprio; banco; hospedagem e custo.
4. **Arquitetura e estrutura técnica:**
   - paradigma: orientado a objetos × funcional × orientado a dados/ECS, por quê, e em qual camada;
   - núcleo puro × misturado;
   - monólito modular × serviços;
   - monorepo e pastas;
   - contrato de mensagens;
   - luta como log de eventos;
   - determinismo;
   - testes e simulação.
5. **Navegador:** o jogo pode rodar no navegador? Por que sim e por que não (Steam, performance, assets, distribuição, sessão aberta).
6. **Direção de arte:** estilo, câmera, pipeline, assets comprados, UI/HUD, leitura do grid, o que evita o infantil.
7. **Mecânicas:** combate d20, as 2 escadas de progressão, coleta por nó, economia com pouca gente, morte e durabilidade, automático.
8. **Mundo, enredo e estrutura:**
   - Raiz, Puxados, Enraizados e Perdidos;
   - a Travessia e a Era 1 (paisagem certa para 1500–1654);
   - facções;
   - não poetizar.
9. **IA no projeto:** onde entra, onde não entra, e a reação dos jogadores.
10. **Viabilidade:** um MMO solo é viável? Alternativas de formato, como servidor pequeno, co-op ou single-player persistente com modo online depois. Ataque ou defenda.
11. **Pontos que ninguém levantou:**
    - contas e login (Steamworks);
    - cheat e autoridade do servidor;
    - backup e migração de dados;
    - LGPD;
    - custo mensal;
    - som e música;
    - UX da primeira hora;
    - acessibilidade;
    - idioma;
    - preço e modelo (grátis × pago);
    - playtest;
    - nome do jogo;
    - como ele não abandona pela 14ª vez.

## Regras
- **Fato × opinião:** diga o que é fato (com fonte e link) e o que é opinião. Pesquise na web sempre que citar versão, preço, licença ou opinião de jogadores.
- **Não poetizar.** Em português.
