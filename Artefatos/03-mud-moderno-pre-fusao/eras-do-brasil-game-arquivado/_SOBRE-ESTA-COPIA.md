# Sobre esta cópia — eras-do-brasil-game-arquivado

Cópia dos arquivos da branch `main` do repositório
`github.com/sidinei-silva/eras-do-brasil-game-arquivado`, trazida em 21/10/2026 para que o
repositório pudesse ser apagado. Os arquivos ao lado deste são exatamente os do repo; este
arquivo é o único acrescentado.

- **O que era:** o código da fase MUD Moderno (servidor Go + cliente web), a versão em
  que os NPCs tinham "vida" (rotina, conhecimento, fofoca).
- **Período:** 13/03 → 24/05/2026, 90 commits.
- **Reset de 03/04/2026:** todo o código do servidor foi apagado para ser reescrito sem
  IA. Ver `historico/sessions/2026-04-03-reset-reimplementacao-do-zero.md`. O código de
  antes do reset está no commit `4d64406`; o reset em si é o commit `c758ac3`.

## Histórico completo

O histórico inteiro está em `../eras-do-brasil-game-arquivado.bundle`, gerado de um clone
espelho do GitHub. Ele traz os 90 commits e todas as branches:

| Branch | Último commit | Conteúdo |
|---|---|---|
| `main` | `c0a2e5a` (24/05) | igual a esta cópia |
| `feat/1.3-npc-utility-ai-intermediario` | `4dd1201` | Refinamento 1.3 (Utility AI dos NPCs) |
| `feat/1.4-movimento-de-npcs` | `7b8427c` | Refinamento 1.4 (movimento entre blocos) |
| `feat/1.5-npc-fofoca-knowledge-base` | `c0a2e5a` | Refinamento 1.5 (fofoca), já na `main` |
| `incluindo-ideias-futuras` | `0a1f454` | ideias guardadas para depois |

Para recuperar o repositório inteiro:

```bash
git clone eras-do-brasil-game-arquivado.bundle eras-do-brasil-game-arquivado
cd eras-do-brasil-game-arquivado
git show 4d64406          # o código de antes do reset
```
