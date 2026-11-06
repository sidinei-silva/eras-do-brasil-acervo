# Era 06 — MMORPG sandbox por turno (15/10 a 05/11/2026)

Depois da declaração de 14/10 (*"o jogo esta totalmente indeciso"*), uma semana de estudo
de game design transformou o Eras num MMORPG sandbox por turno para PC: regras d20 no
estilo Pathfinder 2e, combate em grid com 3 ações, progressão do Albion ("você é o que
veste"), sem full loot, e um enredo escrito pelo próprio Sidinei (a Raiz, os Enraizados,
os Puxados e os Perdidos). A stack proposta em 17/10 (Unity + servidor .NET + núcleo em
C#) nunca foi confirmada. Em 22/10, depois dos debates, ficou **Go no núcleo, simulador e
servidor, com cliente em Unity**.

A era termina em 05/11/2026, quando o formato deixa de ser um MMORPG 100% online e passa a
ser um RPG de mundo vivo, com NPCs simulando jogadores, em C# e Unity. A
[linha do tempo](../../LINHA-DO-TEMPO.md#novembro-de-mmorpg-a-mundo-simulado-05112026)
conta o porquê. O que vale das regras desta era continua sendo usado no formato novo.

O repositório [`eras-do-brasil`](https://github.com/sidinei-silva/eras-do-brasil) segue
vivo e não foi alterado. A cópia aqui é o retrato dele no fim da era:

| Arquivo | O que é |
|---|---|
| `eras-do-brasil/` | Cópia dos arquivos do repositório (21/10 → 05/11/2026, commit `c48ba10`): docs, gestão automática e o começo da fase 1 em Go |
| `eras-do-brasil.bundle` | Histórico completo do repositório, 39 commits (`git clone` nele recupera tudo) |
| `claude-projects/eras-do-brasil/memoria-do-project.md` | Memória do Project "Eras Do Brasil" no Claude (export de 05/11/2026), anterior à confirmação da virada |
| `debates-2026-10/` | Debates entre agentes de 21 a 22/10/2026 sobre stack, arte e o jogo inteiro, incluindo a decisão por Go |
| `consolidado-do-estudo-out-2026.md` | "Eras do Brasil — Consolidado do estudo (out/2026)", escrito pelo Claude em 18–19/10 a partir da sessão de estudo: pilares, motor d20, combate, build, progressão, durabilidade, PvP com eco, mundo, stack, lore, divergências com as fases anteriores e riscos |
| `consolidado-atual_2026-10-21.md` | Retrato do jogo no repositório `eras-do-brasil` (commit `f34acec`, 21/10) mais as decisões das conversas: pilares, enredo, Travessia em 20 passos, motor d20, combate, ofícios, durabilidade, fases até o MVP, checklists para nova Era/zona/mecânica. Anexo A compara com o F8; Anexo B guarda o rascunho de zonas herdado do F7 |
| `brainstorm-2026-10-19/dossie-fases.md` | Dossiê com todas as fases do projeto, preparado para o debate entre agentes |
| `brainstorm-2026-10-19/rodada1-resumo.md` e `rodada2-resumo.md` | Resumos das duas rodadas de debate entre agentes com estilos diferentes sobre enredo, lore e mundo — o resultado não convenceu e levou o Sidinei a escrever o enredo ele mesmo |

O consolidado do estudo é o documento exportado do Claude Docs; o dossiê e os resumos foram
escritos pelo Claude dentro da conversa "RavenIdle melee vs ranged mechanics" e
reconstruídos a partir das chamadas de ferramenta registradas no export. O consolidado atual
foi gerado em 21/10 e retrata o repositório naquele dia, antes dos debates e do código em Go.
