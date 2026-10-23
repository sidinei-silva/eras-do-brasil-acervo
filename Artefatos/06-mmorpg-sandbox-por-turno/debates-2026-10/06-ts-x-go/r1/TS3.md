# TS3 (TypeScript) – Abertura

## 1. Posição
Defendo TypeScript para `regras/`, `sim/` e `server/`. Tenho Go em alta conta: o binário único, o `go test` sem configuração e a falta de escolhas a fazer são vantagens reais. Aceito Go como segunda opção boa. Se a decisão final for Go, não considero erro. Defendo TS pelo encaixe com o dono: é a linguagem foco dele desde 2018, chega mais rápido a algo na tela, roda o mesmo `aplicar` no playground do navegador sem ponte e expressa os eventos de luta como união discriminada checada pelo compilador.

## 2. Cinco problemas

**`aplicar` puro.** O `switch` sobre `tipo` com `never` no fim faz o compilador acusar ação sem tratamento.
```ts
type Estado = Readonly<{ turno: number; rng: number; atores: readonly Ator[] }>;
export function aplicar(e: Estado, a: Acao): { estado: Estado; eventos: Evento[] } {
  switch (a.tipo) {
    case "atacar": return atacar(e, a);
    case "mover":  return mover(e, a);
    default: { const x: never = a; throw new Error(`ação sem regra: ${x}`); }
  }
}
```
O `Readonly` só vale na compilação. Nos testes, um `Object.freeze` profundo no estado de entrada pega mutação em tempo de execução.

**Sorteio com semente no estado.** A função devolve o valor e a semente nova, e quem chama grava a semente no estado. `Math.imul` e `>>> 0` mantêm tudo em inteiro de 32 bits.
```ts
export function rolar(s: number, lados: number): [valor: number, semente: number] {
  const s2 = (s + 0x6d2b79f5) >>> 0;              // mulberry32
  let t = Math.imul(s2 ^ (s2 >>> 15), s2 | 1);
  t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
  return [(((t ^ (t >>> 14)) >>> 0) % lados) + 1, s2];
}
```

**Tipos de evento.** É aqui que TS está mais à frente:
```ts
type Evento =
  | { tipo: "ataque"; seq: number; de: Id; alvo: Id; d20: number; total: number }
  | { tipo: "dano";   seq: number; alvo: Id; valor: number }
  | { tipo: "movimento"; seq: number; quem: Id; para: Casa }
  | { tipo: "reacao"; seq: number; quem: Id; gatilho: Evento["tipo"] };
```
Dentro de `if (ev.tipo === "dano")`, o compilador já sabe que existe `ev.valor`. Em Go seria preciso uma interface com um método marcador mais `type switch`, e esquecer um caso só é acusado com linter (`exhaustive`).

**Teste por tabela.** O dono escreve a tabela e a IA escreve só o laço, que já vem pronto no Vitest:
```ts
it.each([
  { d20: 20, mod: 5, cd: 15, grau: "critico" },
  { d20: 11, mod: 5, cd: 15, grau: "sucesso" },
  { d20: 2,  mod: 5, cd: 15, grau: "falha"   },
])("d20=$d20 +$mod contra CD $cd → $grau", (c) => {
  expect(grau(c.d20, c.mod, c.cd)).toBe(c.grau);
});
```
Ele já escreveu testes nesse estilo Dado/Quando/Então e conferiu que "o estado anterior não muda": `packages/core/combat/fight.test.ts` do semi-idle (commits dele de 11 a 14/10; segundo a LINHA-DO-TEMPO #59, foi a IA que escreveu `position.ts` e o playground).

**Servidor local junto do Unity.** Faço o bundle com esbuild num único `server.js` e distribuo junto um `node.exe` em `StreamingAssets`. O Unity sobe esse servidor com `Process.Start`. A alternativa é o Single Executable Application do Node, que tem `--build-sea` desde a v25.5, mas está com estabilidade 1.1, "Active development" ([docs](https://nodejs.org/api/single-executable-applications.html)). Nesse ponto **Go é melhor**: `GOOS=windows go build` gera um binário de poucos MB. Em TS dá para fazer, mas são dezenas de MB e um passo a mais no build.

## 3. O ponto mais fraco do TS para este dono
Em TS, o determinismo depende de disciplina, e o tipo não ajuda. `number` é float64: um `/` sem `Math.trunc` gera fração sem aviso, e passar de 2^53 perde precisão em silêncio. `Math.random()` e `Date.now()` estão sempre a um autocomplete de distância. Isso já aconteceu no acervo: `packages/core/teste-utils.ts` do semi-idle usa `Math.random()` para gerar a arma do teste. Quando um teste usa sorteio sem semente, ele não reproduz o mesmo resultado a cada execução. Some a isso o ferramental: no semi-idle havia workspaces, `vitest.config`, `vite` e `mise`, ou seja, noites de configuração que em Go não existem. A minha mitigação: uma regra de lint que proíbe `Math.random`, `Date` e `/` dentro de `regras/`, mais um tsconfig fixo escrito uma vez só. Mesmo assim, é um custo que Go não tem.

## 4. Perguntas
- **GO1:** em Go, como você garante, **sem linter externo**, que adicionar o evento `reacao` quebra a compilação em todo `switch` que não trata esse caso (no sim, no log e no servidor)? Se a resposta for o `exhaustive`, quantas noites o dono gasta configurando o golangci-lint?
- **GO3:** para o playground no navegador rodar o mesmo `aplicar` sem chamar o servidor, você usa `GOOS=js GOARCH=wasm` ou TinyGo? Qual o tamanho do `.wasm` e quantas noites isso custa até a primeira luta na tela?
