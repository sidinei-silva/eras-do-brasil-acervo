# Eras do Brasil — Acervo

> Acervo do projeto Eras do Brasil: tudo o que foi conversado, decidido, escrito e
> programado desde a primeira conversa, em julho de 2023, até o recomeço atual —
> conversas com IAs, documentos de cada fase, cópias completas dos repositórios de código
> antigos e a linha do tempo que liga tudo isso.

## Linha do tempo

Para entender de onde o projeto começou, por quantas formas ele passou e por quê, veja
[`LINHA-DO-TEMPO.md`](LINHA-DO-TEMPO.md). A seção "Marcos" resume as 13 roupagens do jogo
numa tabela.

## Conteúdo

Organizado em dois eixos: **conversas** (por fonte) e **artefatos** (por era do projeto
— cada era foi substituída pela seguinte, então "era anterior" = "histórico").

| Pasta | Descrição |
|-------|-----------|
| `Conversas/GPT/` | Conversas com o ChatGPT: a origem e a fase de RPG de mesa (2023–2025) e o trilho paralelo ao Claude em ago–set/2026 (servidor da Escória, arte do MVP, a fusão) |
| `Conversas/Gemini/` | Conversas com o Gemini: fases e ferramentas (nov/2025–fev/2026) e o enredo do mundo (out/2026) |
| `Conversas/Claude/` | Conversas com o Claude (mar–out/2026); em `memorias/`, a memória da conta sobre o jogo |
| `Conversas/WhatsApp/` | Prints e transcrição da conversa em que a ideia, nascida no ChatGPT em 01/07/2023, foi apresentada pela primeira vez (02–04/07/2023) |
| `Conversas/Copilot/` | Logs de sessões com o Copilot (fev–mar/2026): a última fase de documentação e o pivô para MUD Moderno |
| `Artefatos/01-rpg-de-mesa-e-narrativa/` | GDD inicial (`Livros Antigos/`) e as análises do NotebookLM (2023–jan/2026), quando o jogo ainda era um RPG de mesa/D20 sem código |
| `Artefatos/02-unity-tatico/` | ADRs e docs técnicos da era em que o jogo era um RPG tático em Unity |
| `Artefatos/03-mud-moderno-pre-fusao/` | MUD Moderno (Go + web, mar–mai/2026): o GDD (`gdd-batalha-turno/`), o código (`eras-do-brasil-game-arquivado/` + histórico completo em `.bundle`) e o Project do Claude da fase |
| `Artefatos/04-idle-escoria-e-fusao/` | Fase idle (mai–set/2026): o código do Eras fundido (`eras-do-brasil-idle/` + `.bundle`), os Projects do Claude (Albion Idle/Outland, A Escória, Eras) e o consolidado de decisões |
| `Artefatos/05-semi-idle-typescript/` | O recomeço em TypeScript (03–14/10/2026): o código (`eras-do-brasil-semi-idle/` + `.bundle`) e o consolidado F8 |
| `Artefatos/06-mmorpg-sandbox-por-turno/` | A fase atual (desde 15/10/2026): o consolidado do estudo de outubro e o brainstorm multiagente que deram origem ao repositório atual |
| `Imagens/` | Imagens conceituais e referências visuais (sem data confiável por arquivo para atribuir a uma era) |

Para recuperar um repositório antigo inteiro, com todas as branches e commits:
`git clone Artefatos/<era>/<repositorio>.bundle <pasta>`.

## Projeto atual

- **[eras-do-brasil](https://github.com/sidinei-silva/eras-do-brasil)** — o jogo hoje:
  MMORPG sandbox por turno para PC. Por enquanto só documentação (pilares, mecânicas,
  enredo e roteiro), sem código.
