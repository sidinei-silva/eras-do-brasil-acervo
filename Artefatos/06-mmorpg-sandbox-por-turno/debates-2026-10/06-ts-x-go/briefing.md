# Debate TypeScript × Go: briefing

Você está num debate com 15 agentes sobre UMA decisão: em que linguagem escrever o **núcleo de regras, o simulador e o servidor** do Eras do Brasil. O cliente é Unity (C#) e isso não está em debate.

## Quem é o dono (leia com atenção)
- **Rotina:** Sidinei, dev solo, cerca de 2h por noite (21h–0h), com interrupções (bebê). Casado, única renda da casa.
- **Carreira:** fullstack web desde 2018. Começou em PHP e foi para Node.js/TypeScript, a linguagem foco dele. Passou por Java, Kotlin, Ruby e Python.
- **Go:** usa no trabalho (Gamersclub, projetos novos em Go, core em PHP legado) e se considera iniciante.
- **C#:** não sabe.
- **Trabalho no projeto:** escreve o código do jogo à mão. A IA explica, revisa e escreve só o laço dos testes, a partir dos casos que ele escreve. Perdeu prática programando quase só com IA num emprego anterior e quer recuperar autoria.
- **Histórico:** em 3 anos, 13 roupagens e 5 trocas de stack. Nenhuma morreu por defeito técnico: todas caíram por trocar de gênero antes de entregar, pela IA escrevendo o núcleo, ou por documentação no lugar de código.

## O projeto
- **Formato:** MMORPG sandbox por turno, só PC (Steam). "Você é o que veste", d20 estilo Pathfinder 2e, progressão do Albion.
- **Repo atual:** `scratchpad/repo` (github.com/sidinei-silva/eras-do-brasil). Leia `docs/01-o-jogo.md`, `docs/03-tecnico.md`, `docs/04-roteiro.md` e `docs/mecanicas/`. O repo ainda não tem código.
- **Mudanças já decididas** (o repo ainda não reflete): `scratchpad/mudancas-docs.md`.
- **Sabatina anterior:** `scratchpad/sabatina/resultado.md` e `rodada-1/2/3.md`.
- **Histórico:** `scratchpad/acervo/LINHA-DO-TEMPO.md`.
- **Código antigo dele, para ver como ele resolve problemas:**
  - Go, MUD (mar–abr/2026): `scratchpad/acervo/Artefatos/03-mud-moderno-pre-fusao/eras-do-brasil-game-arquivado/server`
  - Go, idle (set/2026): `scratchpad/acervo/Artefatos/04-idle-escoria-e-fusao/eras-do-brasil-idle/backend`
  - TS, semi-idle (out/2026): `scratchpad/acervo/Artefatos/05-semi-idle-typescript/eras-do-brasil-semi-idle`
  - Atenção: parte desse código foi escrita por IA (veja a LINHA-DO-TEMPO). Diga qual parte é dele quando usar como prova.

**Caminho base:** `scratchpad` = `/tmp/claude-0/-home-claude/b2c9da2e-f5a3-5381-b813-0e9760dd9841/scratchpad`

## O que a linguagem precisa fazer (já decidido)
1. **Núcleo como função pura:** `aplicar(estado, ação, semente) → {estado, eventos}`.
   - Nada que muda fica guardado em objeto; o estado é dado puro.
   - Não muta a entrada; só inteiros; sorteio com semente dentro do estado.
   - Replay, reconexão e teste por tabela dependem disso.
2. **Primeiro entregável:** luta no terminal (1 personagem × 2 mobs), com marcos nas noites 5, 10 e 20.
3. **Simulador:** roda 1.000 lutas e mede o ritmo.
4. **Servidor:** um processo, sem framework, WebSocket + JSON, contrato com id por comando e sequência por evento.
   - Até um amigo jogar, roda LOCAL junto do cliente Unity, com save em arquivo.
   - Depois: VPS, Postgres e login Steam.
5. **Contrato de mensagens** compartilhado com o cliente Unity em C#.
6. **Playground no navegador:** só para testar.
7. **Eventos de luta** com vários tipos (ataque, dano, movimento, reação...).

## Critérios obrigatórios
O dev · o projeto · performance · o que é ideal para o projeto · conforto · facilidade.

Mais estes:
- determinismo (inteiros, overflow);
- modelagem de dados (tipos de evento, união discriminada × interface);
- teste por tabela;
- distribuir o servidor local junto do jogo Unity;
- integração do contrato com C#;
- tempo até algo na tela;
- aprendizado e carreira (Go no trabalho);
- qualidade da ajuda de IA em cada linguagem;
- ferramental (build, módulos, tsconfig × go toolchain);
- risco de abandono e de troca de novo, o que o acervo mostra;
- longevidade e ecossistema;
- custo de hospedagem.

## Regras do debate
- Converse de verdade: responda pelo nome, cite a frase que está rebatendo e responda TODA pergunta feita a você.
- Fato tem fonte (link ou caminho de arquivo). Opinião é dita como opinião.
- Pode usar trechos de código de até 10 linhas para comparar. O código é do debate, não do jogo.
- Em português, sem poetizar.
