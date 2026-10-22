# Sobre esta cópia — eras-do-brasil-idle

Cópia dos arquivos da branch `main` do repositório
`github.com/sidinei-silva/eras-do-brasil-idle`, trazida em 21/10/2026 para que o
repositório pudesse ser apagado. Os arquivos ao lado deste são exatamente os do repo; este
arquivo é o único acrescentado.

- **O que era:** o Eras do Brasil idle depois da fusão com A Escória — servidor Go
  autoritativo, MVP "A Travessia", documentação normativa em `docs/`
  (`contexto-do-projeto.md`, `decisoes-de-design.md`, `arquitetura-consolidada.md`,
  `direcao-de-arte.md`...). Até out/2026 se chamava `eras-do-brasil`; foi renomeado quando
  o projeto recomeçou em TypeScript.
- **Período:** 12/09 → 26/09/2026, 74 commits. Fatias 1–3 (conta, login, personagem)
  concluídas; a fatia 4 (entrar no mundo) foi mergeada na `main` em 06/10.
- `WORK_NOTES.md` (perguntas de PvP herdadas da Escória que nunca foram resolvidas) deixou
  de ser versionado no commit `aead883`, por isso não está nesta cópia. A versão do disco
  era idêntica à última commitada: `git show aead883^:WORK_NOTES.md` no bundle.

## Histórico completo

O histórico inteiro está em `../eras-do-brasil-idle.bundle` (branches `main` e
`04-entrar-no-mundo`, ambas no commit `cd7409e`). Ele pesa ~9 MB porque dois binários
compilados do Go (`backend/tmp/server`, gerados pelo hot reload) foram commitados por
engano no meio do caminho; não estão na versão final e foram mantidos no bundle para não
alterar nenhum hash de commit.

```bash
git clone eras-do-brasil-idle.bundle eras-do-brasil-idle
```

## Trabalho que só existia na máquina

O clone local tinha um `git stash` de 17/09/2026 que nunca foi para o GitHub
("criando-world-para-entrar-no-mundo", na branch `04-entrar-no-mundo`, sobre o commit
`d30c70d`): primeiro esboço de mundo/zonas para a fatia 4 (`world.go`, `zone.go`,
`gamedata/zones.go`, `data/mvp/zones.json`). Foi salvo como
`../eras-do-brasil-idle_stash-2026-09-17_criando-world.patch`.

```bash
git clone eras-do-brasil-idle.bundle eras-do-brasil-idle && cd eras-do-brasil-idle
git checkout d30c70d && git apply ../eras-do-brasil-idle_stash-2026-09-17_criando-world.patch
```
