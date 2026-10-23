# TS4 (Grupo TypeScript), rodada 1

## 1. Posição

Eu sou dev de TypeScript e já tinha planos de migrar para Go na minha carreira. Os motivos são bons: binário único, `int64` de verdade, toolchain sem escolhas, mercado de backend. Aceito Go em **8 de 10**. Mesmo assim, defendo TS para **este** projeto, **este** dono e **agora**.

A distinção que eu quero deixar na mesa:
- **Bom para a carreira** = o que te deixa mais empregável daqui a 2 anos. Para o Sidinei, isso já está pago: ele usa Go na Gamersclub 40h por semana, e quem paga por esse aprendizado é a empresa.
- **Bom para o projeto** = o que faz a luta da noite 5 sair escrita à mão em 2h, com sono e com bebê. Aqui pesa a fluência, não o currículo. TS é a língua foco dele desde 2018 (`briefing.md`). Em Go ele se diz iniciante.

Na minha opinião, pedir que o projeto seja também a escola de Go põe duas metas numa noite de 2h. O acervo mostra o que acontece com metas somadas (`acervo/LINHA-DO-TEMPO.md`, "O 'não' que vira projeto paralelo").

## 2. Os 5 problemas em TS

**`aplicar` puro.** Uso `readonly` no tipo e `Object.freeze` nos testes. Assim, uma mutação quebra o teste, não só o editor.
```ts
type Estado = DeepReadonly<{ rng: number; turno: number; lutadores: Lutador[] }>;
export function aplicar(e: Estado, a: Acao): { estado: Estado; eventos: Evento[] } {
  const [d20, rng] = rolar(e.rng, 20);
  return { estado: { ...e, rng, turno: e.turno + 1 }, eventos: [{ tipo: "rolagem", d20 }] };
}
```

**Sorteio com semente no estado.** Uso um gerador de 32 bits, com `Math.imul` e `>>> 0`, todo em inteiro.
```ts
function rolar(s: number, lados: number): [number, number] {
  s = (s + 0x6d2b79f5) >>> 0;
  let t = Math.imul(s ^ (s >>> 15), s | 1);
  t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
  return [(((t ^ (t >>> 14)) >>> 0) % lados) + 1, s];
}
```
Sendo honesto: aqui Go é melhor. Em TS, `7 / 2` dá `3.5` sem aviso. Isso exige disciplina com `Math.trunc` e um teste que verifique `Number.isInteger` em todo o estado.

**Tipos de evento.** É o ponto mais forte de TS: união discriminada com verificação exaustiva.
```ts
type Evento = { tipo: "ataque"; de: string; alvo: string; d20: number }
            | { tipo: "dano"; alvo: string; valor: number }
            | { tipo: "movimento"; quem: string; para: Casa };
switch (ev.tipo) { /* ... */ default: const _: never = ev; }
```
Se ele esquecer um caso, o compilador acusa. Em Go isso vira interface com type switch, e o esquecimento só aparece em tempo de execução.

**Teste por tabela.** Uso o `it.each` do Vitest. Ele já escreveu testes no estilo "Dado/Quando/Então" em `semi-idle/packages/core/combat/spawn.test.ts`; a autoria dos testes não está clara na LINHA-DO-TEMPO.
```ts
it.each([[{ d20: 20 }, "critico"], [{ d20: 1 }, "falha"]])("%o → %s", (ent, saida) =>
  expect(resultado(ent)).toBe(saida));
```

**Servidor local junto do Unity.** Tanto o Node SEA quanto o `bun build --compile` geram um executável que o Unity abre como processo filho. Sendo honesto: são de 50 a 100 MB, contra ~10 MB em Go. Go ganha aqui, mas esse custo só cai na loja Steam, que vem bem depois da noite 20.

## 3. O ponto mais fraco de TS para este dono

Esse ponto não é teórico; está no acervo dele:
- **Os tipos nunca foram checados.** No repo de outubro, `packages/core` não tem `tsconfig` nem `typescript` como dependência. Só o playground tem, e ele foi escrito por IA (LINHA-DO-TEMPO, linha 59). O Vitest apenas remove os tipos, então o "conforto do TS" estava desligado sem ninguém perceber.
- **`teste-utils.ts` usa `Math.random()`**, o oposto de determinismo.
- **TS oferece decisões demais:** Node, Bun ou Deno, `tsconfig`, ESM, test runner. Cada uma é uma porta para trocar de stack de novo. Go tem uma resposta para cada uma dessas perguntas.

Minha condição para TS valer: `tsc --noEmit` com `strict` no `npm test` desde a noite 1, e Node puro, que já remove tipos nativamente desde a v22.18 ([nodejs.org](https://nodejs.org/en/blog/release/v22.18.0)).

## Quando eu mudo de lado

1. **Se ele disser que o projeto também é para aprender Go.** Aí a meta é outra, e Go ganha.
2. **Um teste de uma noite:** ele escreve à mão o mesmo `aplicar(ataque)` de ~30 linhas em cada linguagem, sem pedir código à IA. Se terminar a versão Go com o mesmo conforto, a vantagem de fluência some, e o resto (inteiros, binário, toolchain) favorece Go.
3. **Se até a noite 10 o tempo for mais para ferramental TS do que para regra.**

## 4. Perguntas

- **GO1:** se o Go do trabalho já cuida da carreira por 40h semanais, o que o *projeto* ganha em ser também uma escola de Go, numa janela de 2h com interrupções?
- **GO2:** em abril ele apagou o MUD em Go para reescrever sem IA e parou por fadiga (LINHA-DO-TEMPO, linha 55). Que parte desse código Go é dele, e ela mostra fluência ou luta com a linguagem?
- **TS1:** você aceita que o grupo TS se comprometa com `tsc --noEmit` `strict` e um teste de "só inteiros" desde a noite 1? Sem isso, nosso argumento de tipos é o mesmo que falhou em outubro.
