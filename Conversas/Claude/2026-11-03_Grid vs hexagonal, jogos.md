## Sidinei:

Qual a diferença de um jogo de grid para um jogo de mapa exagonal estilo aqueles roguelike antigos?

---

## Claude:

A diferença está nos **vizinhos** e em **como se mede a distância**.

| | Quadrado | Hexágono |
|---|---|---|
| Vizinhos | 4 pelos lados + 4 pelas diagonais | 6, todos iguais |
| Distância | A diagonal é mais longa (≈1,41), então é preciso uma regra para ela | Todo passo vale 1, sem caso especial |
| Área circular | Fica com cara de quadrado ou de losango | Fica quase redonda |
| Paredes, salas, construções | Retas e naturais | Ficam serrilhadas |
| Código | Coordenada (x, y) simples | Coordenada axial ou cúbica, um pouco mais chata |
| Arte | Tilemap comum, muito pack pronto | Menos pack, encaixe mais trabalhoso |

**Correção sobre "roguelike antigo":** Rogue e NetHack usam quadrado com 8 direções. Hexágono é mais comum em wargame e tático (Heroes, Battle Brothers).

**O problema clássico do quadrado é a diagonal.** As três saídas comuns:
1. **Diagonal custa 1:** simples, mas andar na diagonal fica vantajoso.
2. **Diagonal alterna 1, 2, 1, 2…:** é a regra do Pathfinder 2e. Fica justa e continua só com inteiros.
3. **Sem diagonal (4 direções):** o mais simples, mas o movimento fica travado.

**Onde isso toca o Eras:** o `combate.md` já fala em "quadrado" e já tem a régua de "casas por ação". Então hexágono seria reabrir uma decisão. O que **não** está na doc é a regra da diagonal. Ela cai na tarefa "Turno por lado: …, andar" da fase 1, junto com a forma do Cleave, que já está no Em aberto. Quando chegar nessa tarefa, eu te dou as opções para você decidir.

---
