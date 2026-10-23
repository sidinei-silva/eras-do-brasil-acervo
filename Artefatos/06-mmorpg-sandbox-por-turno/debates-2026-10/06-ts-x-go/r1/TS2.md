# TS2 — abertura (TypeScript)

## 1. Posição e aceitação do Go

TypeScript para `regras/`, `sim/` e `server/`. Aceitação do Go: **3 de 10**. Go é ótimo para escrever proxy HTTP. Para escrever regra de jogo com 15 tipos de evento, ele obriga você a fingir que tem união discriminada com `interface{ event() }` e `switch v := e.(type)` sem checagem de exaustividade. O próprio código Go do dono faz exatamente essa gambiarra: `type Event interface { event() }` (`acervo/Artefatos/04-idle-escoria-e-fusao/eras-do-brasil-idle/backend/internal/game/event.go`). Soma de tipos está em proposta desde 2017 (https://github.com/golang/go/issues/19412) e não saiu.

**Onde o ódio não se aplica:**
- `if err != nil`: o time do Go encerrou as propostas de sintaxe nova para erro (https://go.dev/blog/error-syntax, jun/2025). Mas `aplicar` é puro, sem I/O, e quase não tem `err`. O ruído aparece no servidor e no loader de `data/`, não na regra. Nesse ponto o meu ataque vale pouco.
- Distribuir junto do Unity: `go build` gera um binário estático. Isso é melhor que tudo o que o Node oferece.

## 2. Cinco problemas em TypeScript

**`aplicar` puro.** `Readonly` profundo e `as const` fazem o compilador reclamar de mutação. O Go não tem `const` para struct, e slice dentro de struct copiada continua compartilhando memória.
```ts
type Estado = DeepReadonly<{ rng: number; mobs: Mob[]; turno: number }>;
export function aplicar(e: Estado, a: Acao): { estado: Estado; eventos: Evento[] } {
  const ev: Evento[] = [];
  const mobs = e.mobs.map(m => m.id === a.alvo ? { ...m, pv: m.pv - 5 } : m);
  return { estado: { ...e, mobs }, eventos: ev };
}
```

**Sorteio com semente no estado.** Usar mulberry32 em 32 bits com `Math.imul` e `>>> 0` dá um resultado exato e igual em qualquer runtime JS.
```ts
function d20(s: number): [number, number] {
  let t = (s + 0x6D2B79F5) >>> 0;
  t = Math.imul(t ^ (t >>> 15), t | 1);
  t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
  const r = ((t ^ (t >>> 14)) >>> 0) % 20 + 1;
  return [r, (s + 0x6D2B79F5) >>> 0]; // [rolagem, nova semente]
}
```

**Tipos de evento.** Aqui o Go perde de lavada:
```ts
type Evento =
  | { tipo: "ataque"; de: Id; alvo: Id; rolagem: number }
  | { tipo: "dano"; alvo: Id; valor: number }
  | { tipo: "reacao"; de: Id; gatilho: Evento["tipo"] };
function nunca(x: never): never { throw x; } // no default: esquecer um caso NÃO compila
```
Em Go, se um tipo novo ficar sem `case`, o código compila calado.

**Teste por tabela.** O Vitest tem `test.each`, e o dono já usa Vitest (`acervo/.../05-semi-idle-typescript/eras-do-brasil-semi-idle/package.json`). Empate técnico: a tabela de structs do Go é igualmente boa, e isso eu admito.

**Servidor local junto do Unity.** `bun build --compile server.ts --outfile eras-server` gera um executável único (https://bun.sh/docs/bundler/executables). O Node também tem SEA (https://nodejs.org/api/single-executable-applications.html). Funciona, mas o binário tem ~50–100 MB, contra ~10 MB do Go. Para um jogo Steam que já pesa centenas de MB, a diferença não importa.

## 3. Ponto mais fraco do TS para este dono

1. **`number` é float64.** O briefing exige "só inteiros". Em TS, isso é disciplina (`>>> 0`, `Math.trunc`, nada de `/` solto), não algo que o tipo garante. Um `pv / 2` escapa e vira 7.5 sem nenhum aviso. O Go dá erro de compilação em `int / float`.
2. **Histórico.** Foi justamente no repo TS que a IA escreveu o `position.ts` que ele queria aprender (LINHA-DO-TEMPO, linha 59). O TS facilita colar código de IA, porque é a linguagem em que a IA mais tem exemplo. Para quem quer recuperar autoria, isso é risco, não vantagem.

Nota sobre autoria: `types.ts` do semi-idle eu cito só como estilo de modelagem, sem afirmar que é dele. O `event.go` é da fase idle, quando o Claude assumia o código.

## 4. Perguntas

- **GO1:** com `switch e := ev.(type)` sobre `interface{ event() }`, como o dono descobre, em tempo de compilação, que esqueceu de tratar o evento `reacao` no simulador? Vale linter externo (`exhaustive`)? Se vale, admita que isso é tooling compensando a linguagem.
- **GO2:** o estado tem `[]Mob`. Se `aplicar` recebe `Estado` por valor e faz `e.Mobs[0].PV -= 5`, a entrada foi mutada. Como o Go impede isso sem clonar à mão cada slice e map?
- **TS1:** qual regra concreta de lint ou de tipo você propõe para impedir float no estado (branded `Int`? ESLint?), sem virar burocracia para quem tem 2h por noite?
