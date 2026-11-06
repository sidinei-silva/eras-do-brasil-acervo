## Sidinei:

Estava pensando uma coisa, e se a primeira versão ou o jogo todo do eras do brasil não for de mundo aberto total, e ser como são as brumas hoje do albion online zonas instanciadas, onde pode ter mobs e recursos exatamente como é no albion online, e acabar sendo um dungeon crawler praticamente?

---

## Claude:

Vou ler a memória do projeto para achar o repo e depois os arquivos no GitHub.

Vou tentar pelo git.

**Estacionar.** Estrutura do mundo não é da fase 1 (luta no terminal), e `mundo.md` está congelado.

Para ficar tranquilo:

- **Não muda nada da fase 1.** O combate já é decidido em instância ([combate.md](docs/jogo/mecanicas/combate.md)), seja o mundo aberto ou de zonas instanciadas. O núcleo é o mesmo.
- **A hora certa é a fase 3 (A Beira).** É lá que entram zona, grid, coleta e mobs na tela. Os passos 1–6 da Beira já são instanciados e solo (`02-travessia.md`). Lá você testa a ideia na prática antes de decidir o mundo todo.
- **Ao reabrir, vai mexer em decisões já escritas.** Uma é "Zonas ligadas por saídas", em `mundo.md`. A outra é "o Acampamento é o primeiro lugar com outros jogadores", em `02-travessia.md`.

Linha para `docs/gestao/estacionamento.md`, em **Ideias** (ou `/estacionar` no Claude Code):

```
- 2026-11-03 · mundo em zonas instanciadas tipo brumas do Albion (mobs e recursos dentro; quase dungeon crawler), na primeira versão ou no jogo todo · fase 3
```

---
