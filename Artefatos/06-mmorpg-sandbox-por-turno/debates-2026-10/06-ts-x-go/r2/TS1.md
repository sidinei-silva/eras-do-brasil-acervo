# TS1 — rodada 2

## Respostas

Uma proposta única responde boa parte das perguntas: **duas travas no laço de teste, que a IA escreve uma vez.**
```ts
const entrada = deepFreeze(caso.estado);            // ESM é strict: mutar lança TypeError
const { estado } = aplicar(entrada, caso.acao);
expect(soInteiros(estado)).toBe(true);              // percorre o JSON: Number.isSafeInteger em todo número
expect(aplicar(entrada, caso.acao)).toEqual(aplicar(entrada, caso.acao)); // replay
```

**@N2:** "`initializeCombat` muta a entrada; o que impede isso na noite 3?" Confirmo: `spawnCombatants` faz `combatant.position = ...` na entrada (`.../semi-idle/packages/core/combat/spawn.ts`). O `readonly` raso não pega isso. O `deepFreeze` acima pega no primeiro teste. Setup de uma noite: 1 `package.json`, 1 `tsconfig`, Vitest, sem monorepo.

**@GO1 / @TS2 / @MK3:** "2^53, `Math.random`, branded `Int`?" Nada de branded. `soInteiros` no laço e uma regra de lint só: `no-restricted-properties` para `Math.random` e `Date.now`. Custa ~10 linhas, escritas uma vez.

**@GO2 / @TS4 / @MK3:** "qual tsconfig e quem roda `tsc`?" Aceito (TS4): `strict`, `noUncheckedIndexedAccess`, `module: nodenext`, `noEmit`. O script `test` é `tsc --noEmit && vitest run`. O tipo é checado a cada teste. Ninguém precisa lembrar de rodar. Versão do TS, do Node LTS e do Vitest travadas no `package.json` até a fase 5.

**@GO3:** "runtime, runner e noites de configuração?" Node LTS via `mise.toml`, Vitest e workspaces. A configuração entrou em 10/10 às 23h14 (`af9443e`) e a primeira regra em 11/10 às 20h23 (`da17716`). Foi uma noite de configuração. Antes disso foram 5 noites de documentação, que não são culpa da linguagem.

**@N1:** "`bigint` ou `>>> 0`, e como ele percebe que esqueceu?" `number` com `>>> 0` (código na r1). `bigint` não serializa em JSON. Ele percebe pelo `soInteiros` e por um teste fixo: a semente 42 dá a sequência X. Se esquecer o `>>> 0`, a sequência muda e o teste quebra.

**@N3:** "qual runtime vai no Steam e o contrato garante `integer`?" Node SEA se estiver estável na fase 5. Se não estiver, `node.exe` junto do `.js`. É mais pesado que Go, admito. Para o contrato, o JSON Schema escrito à mão é a fonte, com `"type": "integer"`, e o C# lê `int`.

**@MK1:** "a vantagem de IA do TS vale aqui?" Pouco. Se a IA não escreve o núcleo, só sobra revisão. Retiro esse argumento.

**@MK2:** "hospedagem sem plataforma e RAM?" VPS comum, o processo sob `systemd`, Postgres ao lado. N1 mediu ~45 MB contra ~7 MB (`scratchpad/n1bench`). Os dois cabem em 1 GB.

**@N4:** "que razão sobrevive com Unity e o que impede outro `position.ts`?" Sobrevivem a fluência, a união discriminada e o playground rodando o mesmo núcleo sem WASM. A linguagem não impede outro `position.ts`. Quem impede é a regra do roteiro. **Gatilho do acervo:** a IA escreveu o núcleo. Menos atrito na linguagem foco reduz a vontade de delegar, e isso é opinião minha. **Fato que me faria aceitar Go:** se o setup em TS passar de uma noite, ou se ele explicar o `aplicar` em Go linha por linha mais rápido do que em TS.

## Réplicas
- **GO1:** "Struct passada por valor já é cópia. O cuidado está em slice e map." Concordo. Spread em TS também é raso. É o mesmo cuidado nas duas, e por isso o `deepFreeze` vale mais que o argumento de linguagem.
- **N1:** "iterar `map` tem ordem aleatória [em Go]". Em TS, `Object.keys` com chave numérica também reordena. O teste de replay do laço acima serve para as duas linguagens.

## Concessão
A vantagem de IA não pesa aqui (MK1). Também cedo que meu "`Readonly` impede mutação" era fraco: só o `deepFreeze` em tempo de execução garante.

## Pergunta
**GO2:** você aceita as mesmas travas no lado Go (replay, `soInteiros`, nada de `map` no estado)? Se sim, sobra alguma diferença no núcleo além da distribuição?
