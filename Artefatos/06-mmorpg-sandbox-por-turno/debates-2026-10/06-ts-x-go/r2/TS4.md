# TS4, rodada 2

## Respostas

Ninguém me fez pergunta pelo nome. Respondo as que foram feitas ao "lado TS" ou a "quem defender TS".

**@N4:** "Citem um gatilho de troca do acervo, como a sua linguagem o evita, e que fato faria aceitar a rival."
- **Gatilho:** "IA no código → recomeço" (LINHA-DO-TEMPO, linha 81).
- **TS não evita isso.** O `position.ts` foi escrito por IA justamente em TS. O que TS faz é diminuir a *necessidade* de pedir ajuda, porque ali ele não é iniciante. Isso reduz o risco, mas não impede.
- **Fato que me faz aceitar Go:** o teste de uma noite que propus na rodada 1. Ele escreve à mão o mesmo `aplicar(ataque)` nas duas linguagens, contra casos de teste dele, sem pedir código à IA. Se terminar em Go no mesmo tempo, eu passo para o lado Go.

**@GO4:** "Como garantir semente de 64 bits sem `BigInt`? E quanto o `Readonly<>` protege, sendo raso?"
- **64 bits:** o jogo não precisa. O mulberry32 tem período de 2^32, cerca de 4 bilhões de rolagens. 1.000 lutas × ~100 rolagens dão 10^5. Se quiser mais folga, o sfc32 usa 4 inteiros de 32 bits, que cabem no JSON sem truque.
- **`Readonly<>`:** é raso, sim. O `DeepReadonly` é um tipo de 5 linhas, mas só vale na compilação. A proteção real é um `deepFreeze` da entrada no teste: em ESM, mutar objeto congelado lança `TypeError`. Isso pega o erro na execução, e não só no editor.

**@N1:** "Tamanho real do servidor TS empacotado e RAM ociosa. Número medido."
Medi agora um servidor WebSocket + JSON mínimo (scratchpad `r2tmp/m-ts` e `r2tmp/m-go`):

| | `bun build --compile` 1.4.2 | Go 1.24 (`net/http`, sem lib de WebSocket) |
|---|---|---|
| `.exe` Windows | **86 MB** | **8,5 MB** |
| RSS ocioso (Linux) | **18 MB** | **7,6 MB** |

Go ganha. Os 86 MB pesam pouco ao lado de um build Unity, mas são 10× mais.

## Réplicas

**GO4:** "Para ele, carreira e projeto apontam para o mesmo lado."
- Concordo em parte. O retorno duplo existe, mas só se a noite em Go for escrita *por ele*.
- O critério de maior peso do N4 é autoria. Se ser iniciante empurrar para "IA, monta aí", o retorno duplo vira perda dupla: nem treina Go, nem o jogo é dele.
- Isso é opinião. Por isso proponho um teste em vez de um adjetivo.

**N2:** "Nos dois casos, ele nunca escreveu um teste por tabela nem uma função pura."
- Concordo, e isso corta o meu argumento pela metade. Fluência em TS prova que ele escreve *tipos* bem. Não prova que escreve funções puras.
- O músculo que falta independe da linguagem. A vantagem de TS fica reduzida a sintaxe e eventos, e não a disciplina.

## Concessão

Conferi o commit à mão `da17716` (11/10), no bundle do semi-idle. A `initializeCombat` faz `combatant.position = { lane: 0, step: 2 }` no objeto que recebeu.

Ou seja, na língua foco dele, sem IA, a primeira regra do núcleo ("não mexer na entrada") quebrou. O `readonly` não estava lá, e o `tsc` nunca rodou no core. Retiro a ideia de que "conforto em TS" protege o `aplicar`. O que protege é o teste de mutação, e ele custa o mesmo nas duas linguagens.

Continuo em TS por dois motivos: o tempo até a noite 5 e a união discriminada nos eventos. A minha margem caiu.

## Nova pergunta

**GO4:** você aceita o teste de uma noite (o mesmo `aplicar(ataque)` à mão nas duas linguagens, sem código da IA) como desempate do critério de autoria? Se aceita, qual medida conta: tempo, testes passando ou número de perguntas à IA?
