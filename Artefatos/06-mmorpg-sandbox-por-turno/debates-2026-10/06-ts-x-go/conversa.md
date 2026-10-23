# Conversa do debate TS × Go
Transcrição em ordem. Cada bloco é a fala de um agente.

## Rodada 1: aberturas

---

### N4 (rodada 1)

# N4 (neutro, moderação): abertura

## Duas regras antes de começar

1. **Regra anti-troca do dono.** A stack só se rediscute com a fase em curso jogável, e cada rediscussão troca uma camada só, nunca o núcleo (`debate-stack/resultado.md`). Se o orçamento de noites estourar 1,5×, corta escopo e não troca de stack (`mudancas-docs.md`, 04-roteiro). Então a pergunta deste debate não é "qual é melhor". A pergunta é **"qual ele ainda vai estar escrevendo na noite 20 sem querer trocar"**.
2. **O acervo.** São 5 bases em 8 meses: Unity/C# (fev) → Go + web (mar) → Go + Godot (set) → TypeScript + Electron (10/10) → Unity + .NET proposto (21/10). A dúvida "Go × TS" voltou pelo menos 7 vezes, e 5 delas em 2 meses (`acervo/LINHA-DO-TEMPO.md`, linhas 31 e 77). Nenhuma caiu por defeito técnico. Quem argumentar só com técnica está respondendo a uma pergunta que o acervo não fez.

## Critérios e pesos que proponho (soma 100)

| Critério | Peso | Por quê |
|---|---|---|
| Autoria e risco de nova troca | 25 | É o que matou as 5 bases. Pergunta: em qual linguagem ele escreve `aplicar` à mão, sem a IA tomar o núcleo? |
| Tempo até a noite 5 (fluência, conforto, facilidade) | 15 | O marco é fixo. Com 2h por noite, um dialeto que ele ainda está aprendendo custa noites. |
| Modelagem de eventos (exaustividade) | 15 | A luta tem vários tipos de evento. Aqui as linguagens diferem de verdade: união discriminada + `never` contra `interface` + linter. |
| Determinismo e inteiros | 15 | Também difere de verdade: em TS, `number` é float64; em Go, `int` não aceita misturar com float sem conversão. |
| Distribuir o servidor local junto do Unity | 10 | Binário Go (~10 MB) contra Bun/Node SEA (~50–100 MB). Os dois funcionam. O que pesa é o atrito no build. |
| Aprendizado e carreira (Go no trabalho) | 10 | Prática diária que transfere, nos dois sentidos. |
| Contrato com C# | 5 | Os dois saem de JSON Schema ou de tipos gerados. A diferença é pequena. |
| Ferramental (build, módulos) | 5 | tsconfig + bundler contra `go build`. |

## Argumentos que provavelmente vão empatar (peso 0 até alguém provar o contrário)

- **Performance:** um processo, poucos jogadores, luta por turno. Qualquer uma das duas sobra.
- **Custo de hospedagem:** a fase 6 é uma VPS. Os dois rodam na mais barata.
- **Teste por tabela:** `test.each` e slice de struct se equivalem. O TS2 já admitiu isso.
- **Longevidade e ecossistema:** nenhuma das duas some antes de 2035.
- **"Tem WebSocket e JSON":** as duas têm.
- **Qualidade da ajuda de IA:** as duas são bem servidas. E a regra é que a IA não escreve o núcleo, então ajuda melhor não decide nada. Se alguém defender "a IA escreve mais rápido nesta", isso conta **contra**, pelo histórico do `position.ts` (linha 59).
- **"Com disciplina dá":** vale para qualquer linguagem. Quem disser que a sua resolve com disciplina, e a do outro precisa de disciplina, vai ser cobrado.

## Análise inicial (opinião)

- **Vantagens de cada lado:** TS ganha em fluência e em modelagem de eventos. Go ganha em inteiros, em distribuição e em carreira.
- **Autoria:** é o critério de maior peso, e está em aberto. O reset de abril, o único código escrito à mão de propósito, foi em Go e terminou em fadiga. O TS foi onde a IA tomou o núcleo.
- **O que sobra para o TS:** o motivo que levou ao TS em 10/10 foi "o cliente vira web" (linha 57). Hoje o cliente é Unity e o navegador é só playground. Esse motivo sumiu. O TS precisa de outro.

## Perguntas diretas

1. **GO1 e GO2:** o dono se diz iniciante em Go. O código Go escrito por ele depois do reset (03/04) parou em fadiga, e a dúvida de linguagem voltou 5 vezes em 2 meses. Mostrem um trecho que seja **dele**, com caminho de arquivo, e digam em quantas noites ele chega à noite 5 em Go, contra TS. Número, não adjetivo.
2. **TS1 e TS2:** o argumento que tirou o Go em 10/10 era o cliente web, e ele não existe mais. Que razão a favor do TS sobrevive com cliente Unity? E o que, concretamente, impede o repeteco do `position.ts`, em que a IA escreveu o núcleo no repo TS?
3. **Os dois lados:** citem **um** gatilho de troca que aparece no acervo e mostrem como a sua linguagem o evita. E digam qual fato, observável antes da fase jogável, faria vocês aceitarem a rival. Quem responder "nenhum" está defendendo time, não o dono.

---

### TS1 (rodada 1)

# TS1 — abertura

## 1. Posição
TypeScript para núcleo, simulador e servidor. Go é uma boa linguagem e eu a usaria sem drama: se o Sidinei escolher Go, o projeto não morre por isso (o acervo mostra que nenhuma versão morreu por defeito técnico, `acervo/LINHA-DO-TEMPO.md`). Defendo TS porque é a linguagem em que ele **já pensa** (Node/TS é a linguagem foco dele, briefing) e porque o núcleo pedido é quase todo modelagem de dados, onde TS é mais expressivo. Aceitação de Go: alta. Em distribuição, Go ganha, e eu digo isso abaixo.

## 2. Os 5 problemas

**`aplicar` puro.** `Readonly` faz o compilador recusar a mutação da entrada. Em Go não existe `const` para struct/slice; um `append` no slice da entrada passa calado.
```ts
type Estado = DeepReadonly<{ rodada: number; rng: number; lutadores: Lutador[] }>;
function aplicar(e: Estado, a: Acao): { estado: Estado; eventos: Evento[] } {
  // e.rodada++  -> erro de compilação
  return { estado: { ...e, rodada: e.rodada + 1 }, eventos: [] };
}
```

**Sorteio com semente no estado.** Aqui concedo: `number` é double, não inteiro. Resolve com disciplina: PRNG de 32 bits com `Math.imul` e `>>> 0`, e inteiros abaixo de 2^53. Funciona e é determinístico, mas o tipo não garante; Go com `uint32` garante.
```ts
function d20(rng: number): [valor: number, rng: number] {
  let t = (rng + 0x6D2B79F5) >>> 0;           // mulberry32
  t = Math.imul(t ^ (t >>> 15), t | 1);
  t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
  return [(((t ^ (t >>> 14)) >>> 0) % 20) + 1, (rng + 0x6D2B79F5) >>> 0];
}
```

**Tipos de evento.** É o ponto mais forte de TS. União discriminada com checagem de exaustividade nativa: esqueceu de tratar `reacao`, não compila. Em Go é interface + type switch, e a exaustividade depende de linter externo.
```ts
type Evento =
  | { tipo: "ataque"; de: string; alvo: string; rolagem: number }
  | { tipo: "dano"; alvo: string; valor: number }
  | { tipo: "movimento"; quem: string; para: Casa };
function nunca(x: never): never { throw new Error(String(x)); } // no default do switch
```

**Teste por tabela.** Ele escreve os casos; a IA só o laço. Com Vitest, o laço é uma linha (`it.each`). Ele já usava Vitest com o estilo "Dado/Quando/Então" em `fight.test.ts` (`acervo/Artefatos/05-semi-idle-typescript/.../packages/core/combat/fight.test.ts`). Atenção ao briefing: o histórico mostra que `position.ts` e o playground foram escritos por IA (diário de 14/10); os `types.ts` iniciais e o formato dos testes (commit `da17716`, 11/10) são os que eu leio como dele.
```ts
it.each([
  { caso: "crítico com 20", rolagem: 20, ca: 25, esperado: "critico" },
  { caso: "erra por 1",     rolagem: 10, ca: 21, esperado: "erro" },
])("$caso", ({ rolagem, ca, esperado }) => expect(grau(rolagem, ca)).toBe(esperado));
```

**Servidor local junto do Unity.** Concedo que Go é melhor: `GOOS=windows go build` gera um `.exe` sem runtime ([go.dev/doc/install/source](https://go.dev/doc/install/source#environment)). Em Node existe SEA com `node --build-sea`, mas está em "1.1 Active development" e exige empacotar num arquivo só ([docs Node](https://nodejs.org/download/release/latest/docs/api/single-executable-applications.html)); alternativa: `bun build --compile`. Mitigação real: até um amigo jogar, o servidor roda na máquina dele, então isso só pesa na fase 5. Não é critério para a noite 1.

## 3. Ponto mais fraco de TS para ESTE dono
O ferramental é um cardápio: tsconfig, ESM/CJS, Node×Bun×Deno, bundler para o SEA. Para quem já gastou dias de meta-organização em vez de código (LINHA-DO-TEMPO, linha 48), TS oferece mais botões para mexer. Go tem um jeito só. Minha mitigação: travar uma vez (Node LTS com type stripping + Vitest, sem bundler até a fase 5) e proibir mexer. Também é verdade que o repo TS foi onde a IA escreveu o core; a linguagem não causou isso, mas não ajudou a evitar.

## 4. Perguntas
- **GO1:** em Go, como o compilador impede que `aplicar` faça `append` ou altere um item do slice `lutadores` do estado que entrou? Se a resposta é "copiar à mão e revisar", isso não é exatamente o tipo de disciplina que a IA tende a pular?
- **GO2:** para os tipos de evento, você aceita depender do linter `exhaustive` para ter o que TS dá no compilador? Quem configura e mantém isso: ele, à mão, numa noite de 2h?

---

### GO1 (rodada 1)

# GO1, abertura

## 1. Posição
Go para `regras/`, `sim/` e `server/`. TypeScript é uma ótima linguagem, já entreguei bastante coisa com ela e aceito bem: se o Sidinei escolher TS, o projeto não morre por isso. O acervo mostra que nenhuma das 13 roupagens morreu por defeito técnico (`acervo/LINHA-DO-TEMPO.md`). Minha tese é mais estreita: para ESTE núcleo (inteiros, semente, binário ao lado do Unity), Go dá menos coisa para configurar e menos armadilha, e cada noite de Go no jogo também vale como treino para o emprego.

## 2. Os 5 problemas

**`aplicar` puro.** Struct passada por valor já é cópia. O cuidado está em slice e map, que compartilham memória: copio antes de mexer.
```go
func Aplicar(e Estado, a Acao) (Estado, []Evento) {
    novo := e                                   // cópia rasa
    novo.Mobs = append([]Mob(nil), e.Mobs...)   // slice: copiar explícito
    // ... regras ...
    return novo, eventos
}
```

**Sorteio com semente no estado.** `int64`/`uint64` são nativos e o overflow é definido pela especificação (dá a volta, sem pânico): https://go.dev/ref/spec#Integer_overflow. Um gerador de 3 linhas vira um campo do estado e cabe no JSON.
```go
type RNG struct{ S uint64 }
func (r RNG) D20() (RNG, int) {
    r.S = r.S*6364136223846793005 + 1442695040888963407 // LCG 64 bits
    return r, int(r.S>>33)%20 + 1
}
```
Em TS, `number` só é inteiro exato até 2^53 (`Number.MAX_SAFE_INTEGER`); um LCG de 64 bits exige `BigInt` ou truque de 32 bits com `Math.imul`. Funciona, mas é mais um lugar para errar.

**Tipos de evento.** Aqui concedo de saída: Go não tem união discriminada. O padrão é interface com método marcador, o mesmo que já aparece no idle (`acervo/.../eras-do-brasil-idle/backend/internal/game/event.go`), mais um `switch` de tipo. O compilador não avisa se faltar um caso; o linter `exhaustive` cobre parte disso. TS faz melhor nesse ponto.
```go
type Evento interface{ evento() }
type Dano struct{ Alvo, Valor int }
func (Dano) evento() {}
```

**Teste por tabela.** É o jeito nativo do `go test`, sem biblioteca. Encaixa direto na regra do roteiro: ele escreve os casos, a IA escreve só o laço.
```go
casos := []struct{ nome string; sem uint64; quer int }{
    {"acerto crítico", 42, 20},
    {"erro", 7, 1},
}
for _, c := range casos {
    t.Run(c.nome, func(t *testing.T) { /* aplicar e comparar */ })
}
```

**Servidor local junto do Unity.** `GOOS=windows GOARCH=amd64 go build` gera um `.exe` sem runtime, compilado de qualquer máquina. O Unity sobe o processo e conecta em `localhost`. Node tem SEA, mas ainda está em "1.1 Active development" e a compilação cruzada não é documentada como suportada (https://nodejs.org/api/single-executable-applications.html). Bun compila para binário (https://bun.sh/docs/bundler/executables), mas aí é mais uma ferramenta na pilha. Hospedagem depois: o mesmo binário roda em VPS pequena.

## 3. Ponto mais fraco do Go para este dono
Ele se considera iniciante em Go, e TS é a linguagem foco dele. Tempo até algo na tela é menor em TS, e o playground no navegador sai de graça em TS (em Go, só com WASM ou servidor). Pior: o código Go dele no acervo não prova muita coisa. O MUD antes do reset de 03/04 foi escrito por agente; o que veio depois do reset (commit `c758ac3`) usa `math/rand` global (`server/npc/npc.go`, linhas 188 e 194), ou seja, nunca praticou o determinismo que agora é exigido, e os dois repositórios Go não têm nenhum `_test.go`. Também falta união discriminada para os eventos de luta, que vão ter vários tipos.

## 4. Perguntas
- **TS1:** como você garante, no `aplicar`, que ninguém usa `number` acima de 2^53 nem `Math.random` por engano? Com lint, `BigInt` ou disciplina? E quanto de `tsconfig`/ESLint isso custa nas primeiras noites?
- **TS2:** qual é o seu caminho concreto para entregar o servidor ao lado do `.exe` do Unity num PC Windows sem Node instalado: Node SEA, Bun ou Electron? Quantos MB e quantos passos de build?

---

### TS2 (rodada 1)

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

---

### GO2 (rodada 1)

# GO2 (Grupo Go): abertura

## 1. Posição e aceitação do rival

Go para `regras/`, `sim/` e `server/`. Aceitação de TypeScript: baixa, mas não zero. Até a Microsoft desistiu de rodar o compilador do TypeScript em Node e o reescreveu em Go: o TS 7 saiu em 2026 com o compilador nativo, cerca de 10x mais rápido ([InfoQ](https://infoq.com/news/2026/08/typescript-7-released/)). Se nem o time do TypeScript quer o próprio ferramental em Node, por que o servidor do Sidinei deveria?

O ódio que **não** se aplica aqui: `number` em ponto flutuante não quebra um d20. Dano e PV cabem com folga em 2^53. O problema real está em dois lugares. Divisão não é inteira (`7/2` dá `3.5`, então é `Math.trunc` em toda conta). E um RNG de 64 bits exige `BigInt` ou malabarismo com `Math.imul`.

## 2. Os 5 problemas em Go

**`aplicar` puro.** É função livre sobre struct sem métodos. Copiar a struct é cópia por valor, mas há uma armadilha: slice e map continuam apontando para o mesmo dado (falo disso no item 3).
```go
func Aplicar(e Estado, a Acao) (Estado, []Evento) {
    n := e.Clonar() // copia slices/maps explicitamente
    switch a.Tipo { case "atacar": return atacar(n, a) }
    return n, nil
}
```

**Sorteio com semente dentro do estado.** `uint64` dá a volta em overflow por definição da [especificação](https://go.dev/ref/spec#Integer_overflow). Não tem comportamento indefinido nem `BigInt`.
```go
func rolar(s uint64, lados int) (int, uint64) { // splitmix64
    s += 0x9E3779B97F4A7C15
    z := (s ^ s>>30) * 0xBF58476D1CE4E5B9
    z = (z ^ z>>27) * 0x94D049BB133111EB
    return int((z^z>>31)%uint64(lados)) + 1, s
}
```

**Tipos de evento.** É uma interface "selada" com `switch` de tipo, e o JSON leva um campo `tipo`. Aqui o TS ganha, e admito: a união discriminada do TS avisa no compilador quando falta um caso. Em Go isso exige um linter ([gochecksumtype](https://github.com/alecthomas/go-check-sumtype)).
```go
type Evento interface{ evento() }
type Dano struct{ Alvo string `json:"alvo"`; Valor int `json:"valor"` }
func (Dano) evento() {}
```

**Teste por tabela.** É o idioma nativo da linguagem: sem Vitest, sem config e sem `tsconfig` de teste. `go test ./...` e pronto.
```go
casos := []struct{ nome string; sem uint64; quer int }{
    {"crítico", 42, 20}, {"erro", 7, 1},
}
for _, c := range casos { t.Run(c.nome, func(t *testing.T) { /* Sidinei escreve a tabela */ }) }
```

**Servidor local junto do Unity.** `GOOS=windows go build` gera um `.exe` estático. Ele vai em `StreamingAssets` e o C# sobe o processo com `Process.Start`. Não precisa de runtime, `node_modules` nem bundler. O caminho do Node é o [Single Executable Applications](https://nodejs.org/api/single-executable-applications.html), que a própria doc marca como "Active development" (em desenvolvimento ativo) e que embute um binário do Node inteiro.

**Prova no acervo.** O `packages/core` do semi-idle em TS tem `"main": "index.ts"` e nenhum `tsconfig`. Os três `tsconfig` do repo ficam em `apps/playground` (`acervo/Artefatos/05-semi-idle-typescript/eras-do-brasil-semi-idle/`). O core rodava pelo Vitest, que só remove os tipos e não os checa. Ou seja, os tipos que "salvam" o TS nunca foram verificados no núcleo. E parte desse core foi escrita por IA (`_SOBRE-ESTA-COPIA.md`), então isso não prova nada contra o Sidinei, só contra o ferramental.

## 3. O ponto mais fraco de Go para ESTE dono

1. **Ele é iniciante em Go, e TS é a linguagem foco dele.** Em 2h por noite, cada dúvida de sintaxe custa caro.
2. **Aliasing de slice e map** fere a regra "não mexer no estado que entrou" sem nenhum aviso do compilador. `n := e` parece cópia e não é. Ele vai precisar de um `Clonar()` à mão e de um teste que prove que a entrada não mudou.
3. **Sem união discriminada nativa.** A modelagem de eventos fica mais verbosa que em TS.
4. **O playground no navegador** é de graça em TS (o mesmo core roda no browser). Em Go, ou ele compila para WASM, ou o playground conversa com o servidor por WebSocket. O segundo caminho serve, mas é mais uma peça.

O que pesa a favor mesmo assim: ele usa Go no trabalho todo dia. Cada noite no Eras é treino pago para a carreira, e vice-versa.

## 4. Perguntas

- **TS1:** o core TS do acervo não tinha `tsconfig` e rodava pelo Vitest sem checagem de tipo. Qual `tsconfig` exato (`strict`, `noUncheckedIndexedAccess`, `module`/`moduleResolution`) você vai obrigar o Sidinei a manter em `regras/`, `sim/` e `server/`? E quem garante que ele roda `tsc --noEmit` antes de confiar no tipo?
- **TS2:** qual é o comando exato para gerar o executável do servidor Node que vai dentro da pasta do Unity, quanto ele pesa e quantas dependências de `node_modules` (lib de WebSocket incluída) precisam entrar no bundle?

---

### TS3 (rodada 1)

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

---

### GO3 (rodada 1)

# GO3 — abertura

## 1. Posição

Go para `regras/`, `sim/` e `server/`. Digo isso gostando de TypeScript: a união discriminada com `switch` exaustivo é melhor que tudo que Go oferece para eventos, e o ecossistema é enorme. Minha aceitação de TS é alta; se ele escolher TS, o projeto não morre por isso.

Defendo Go por três motivos, não por superioridade técnica:
- **Menos decisões de ferramenta.** `go build`, `go test`, `go fmt` e módulos vêm juntos. Em TS ele precisa escolher runtime (Node/Bun/Deno), tsconfig, test runner, bundler para distribuir. Com 2h por noite, cada decisão dessas é uma noite. O acervo mostra que a troca de stack foi reaberta pelo menos 7 vezes (`acervo/LINHA-DO-TEMPO.md`, "A dúvida de linguagem").
- **Prática diária.** Ele escreve Go na Gamersclub. Hora de trabalho vira hora de jogo e vice-versa. TS ele já sabe; Go é onde ainda há o que ganhar.
- **Simplicidade como freio.** Go tem pouca coisa para "arquitetar". Isso combina com a regra "nunca troca de stack" e com recuperar autoria.

## 2. Os 5 problemas

**`aplicar` puro.** Struct é passada por valor, então a cópia vem de graça para campos simples:
```go
func Aplicar(e Estado, a Acao) (Estado, []Evento) {
    novo := e                         // cópia rasa
    novo.Mobs = slices.Clone(e.Mobs)  // slice e map precisam de clone explícito
    // ... regras sobre `novo` ...
    return novo, eventos
}
```

**Sorteio com semente no estado.** Inteiro sem sinal em Go dá a volta (módulo 2^n) por especificação ([go.dev/ref/spec#Integer_overflow](https://go.dev/ref/spec#Integer_overflow)). Um PRNG de 5 linhas vive no estado como um `uint64`:
```go
type Estado struct{ Rng uint64; /* ... */ }
func rolar(s uint64, lados int) (int, uint64) {
    s += 0x9E3779B97F4A7C15            // splitmix64
    z := (s ^ (s >> 30)) * 0xBF58476D1CE4E5B9
    z = (z ^ (z >> 27)) * 0x94D049BB133111EB
    return int((z^(z>>31))%uint64(lados)) + 1, s
}
```
Em TS, `number` é float64; aritmética de 64 bits pede `BigInt` ou truques com `Math.imul` em 32 bits.

**Tipos de evento.** Interface "selada" com método privado, o mesmo padrão que está no idle (`acervo/.../eras-do-brasil-idle/backend/internal/game/event.go`):
```go
type Evento interface{ evento() }
type Dano struct{ Alvo, Valor int }
func (Dano) evento() {}
```
O compilador não avisa se faltar um `case`. Remédio: o linter `exhaustive` ([github.com/nishanths/exhaustive](https://github.com/nishanths/exhaustive)). Aqui TS é melhor, admito.

**Teste por tabela.** É o idioma nativo de Go ([go.dev/wiki/TableDrivenTests](https://go.dev/wiki/TableDrivenTests)), sem biblioteca:
```go
casos := []struct{ nome string; ini Estado; acao Acao; hp int }{
    {"golpe simples", estadoBase, Atacar{Alvo: 1}, 12},
}
for _, c := range casos { t.Run(c.nome, func(t *testing.T) { /* ... */ }) }
```
Ele escreve os casos; a IA escreve o laço — o laço é isso aí.

**Servidor local junto do Unity.** `GOOS=windows GOARCH=amd64 go build` gera um `.exe` único, sem runtime. Vai em `StreamingAssets/` e o Unity sobe com `System.Diagnostics.Process.Start`. Em TS seria Node embutido, `bun build --compile` ou SEA do Node ([nodejs.org/api/single-executable-applications.html](https://nodejs.org/api/single-executable-applications.html)) — funciona, mas é mais uma escolha.

## 3. Ponto mais fraco de Go para ESTE dono

A imutabilidade não é garantida. `novo := e` copia a struct, mas slices e maps continuam apontando para os mesmos dados. Esquecer um `slices.Clone` quebra a regra "não mexer no estado que entrou" em silêncio. TS com `readonly`/`Readonly<T>` pega parte disso no compilador. Para um iniciante em Go, esse erro é provável. Mitigação: teste que compara a entrada antes e depois de `Aplicar` em todos os casos da tabela. Segundo ponto: JSON de união em Go precisa de um campo `tipo` e `UnmarshalJSON` escrito à mão.

Sobre a prova do acervo: o código Go do MUD depois do reset de 03/04/2026 foi escrito por ele, sem IA no código (`acervo/LINHA-DO-TEMPO.md`, entrada 28); o idle tem arquitetura vinda do chat do ChatGPT, então não uso como prova de autoria, só de padrão.

## 4. Perguntas

- **TS1:** o semi-idle em TS (out/2026) usou qual runtime e qual test runner, e quantas noites foram só de configuração (`mise.toml`, `vitest.config.ts`, monorepo) antes da primeira regra rodar?
- **N2:** você pesa mais "linguagem que ele já domina" (TS) ou "linguagem que ele pratica todo dia no trabalho" (Go)? Por quê?

---

### TS4 (rodada 1)

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

---

### GO4 (rodada 1)

# GO4: rodada 1

## 1. Posição

Sou dev de Go e quero migrar para TypeScript na minha carreira: uma linguagem no front e no back, mais vagas fullstack, o ecossistema do npm. Aceito bem o TS: ele é uma boa linguagem e, para a carreira de muita gente, a escolha certa. Só que a pergunta aqui não é "o que é bom para um dev". A pergunta é "o que é bom para este núcleo, deste dono, com 2h por noite". Para isso defendo Go, e explico a diferença:

- **Bom para a carreira:** amplia o mercado no futuro. TS ganha.
- **Bom para o projeto:** reduz atrito e decisões hoje e não esconde erro de regra. Go ganha por três motivos: `int64` de verdade, um binário só e um ferramental que não tem nada para configurar.
- **Detalhe que inverte o meu caso pessoal:** o Sidinei *já está* em Go no trabalho (Gamersclub) e se diz iniciante. Cada noite de Go no jogo também é treino para o emprego que paga a casa. Para ele, carreira e projeto apontam para o mesmo lado. Para mim, não.

## 2. Os 5 problemas em Go

**`aplicar` puro.** Go passa struct por valor, então campos simples já chegam copiados.
```go
func Aplicar(e Estado, a Acao) (Estado, []Evento) {
    e.Herois = slices.Clone(e.Herois) // slice e map NÃO copiam sozinhos
    // ... regras sobre e (cópia) ...
    return e, eventos
}
```

**Sorteio com semente no estado.** Inteiro sem sinal, com overflow definido e igual em toda máquina.
```go
type Rng struct{ S uint64 } // fica dentro de Estado, vai no JSON
func (r Rng) D20() (Rng, int) {
    r.S = r.S*6364136223846793005 + 1442695040888963407
    return r, int(r.S>>33)%20 + 1
}
```
Em TS, `number` é float: a multiplicação acima perde precisão. Seria preciso usar `BigInt` ou `Math.imul` em 32 bits.

**Tipos de evento.** Go não tem união discriminada. Uso uma interface marcadora mais um `switch` por tipo:
```go
type Evento interface{ tipo() string }
type Dano struct{ Alvo, Valor int }
func (Dano) tipo() string { return "dano" }
switch ev := e.(type) { case Dano: ...; case Movimento: ... }
```

**Teste por tabela.** É o idioma padrão do `testing`, sem lib nenhuma. O Sidinei escreve a tabela e a IA escreve só o `for`, exatamente o contrato do AGENTS.md.
```go
casos := []struct{ nome string; ca, rolagem int; acerta bool }{
    {"20 natural", 30, 20, true}, {"1 natural", 5, 1, false},
}
for _, c := range casos { t.Run(c.nome, func(t *testing.T) { /* ... */ }) }
```

**Servidor local junto do Unity.** `GOOS=windows go build` gera um `.exe` estático, sem runtime, que vai na pasta do jogo e o Unity abre com `Process.Start`. Com TS teríamos Node junto do jogo, ou `bun build --compile`, ou o Node SEA, que ainda é experimental. WebSocket vem de [coder/websocket](https://pkg.go.dev/github.com/coder/websocket@v1.8.15), que é uma lib e não um framework.

## 3. O ponto mais fraco do Go para ESTE dono

São dois, e sou honesto sobre eles:

1. **Imutabilidade não é garantida pelo compilador.** O `slices.Clone` acima existe porque slice e map compartilham memória. Se ele esquecer, a regra "não mexer no estado que entrou" quebra em silêncio. O `readonly` do TS pegaria parte disso na compilação. O código Go escrito à mão depois do reset de 03/04 mostra o risco: o MUD usa `rand.Intn` global, nada determinístico (`acervo/Artefatos/03-mud-moderno-pre-fusao/eras-do-brasil-game-arquivado/server/npc/npc.go:188`). Ou seja, Go não deu disciplina sozinho.
2. **Go não salvou o projeto antes.** Foram três fases em Go (MUD, Albion Idle/Escória, idle fundido) e todas caíram por troca de gênero e por IA no núcleo (`acervo/LINHA-DO-TEMPO.md`, marcos 6–10). Linguagem nenhuma resolve isso. Quem disser que resolve está vendendo.

## Quando eu mudaria de lado

- Se o Sidinei disser que pretende sair de Go no trabalho e fechar a carreira em TS. Aí o argumento de carreira passa a apontar para o TS e o meu "bônus duplo" some.
- Se o playground no navegador deixar de ser só teste e virar cliente de verdade.
- Se, no marco da noite 5, a falta de união discriminada e as cópias manuais custarem noites medidas no `noites.md`. O culpado tem que ser a linguagem, não o design.

## 4. Perguntas

- **TS2 (ou quem defender o TS no núcleo):** em TS, como vocês garantem o sorteio com semente em inteiros de 64 bits sem `BigInt` e sem limitar o gerador a 32 bits? E quanto o `Readonly<>` protege de fato, sendo raso?
- **MK1:** existe dado de vaga, no Brasil, que mostre um dev fullstack TS ganhando mais do que um dev pleno Go *que já está empregado em Go*? A troca de linguagem custa o tempo de reaprender o que ele usa todo dia.

---

### N1 (rodada 1)

# N1 — Performance (abertura)

Engenheiro de servidor de jogos, sem lado. Minha tese de partida: **neste jogo, performance bruta é argumento de fachada; determinismo e distribuição não são.** Trouxe números em vez de opinião.

## Medição (feita agora, nesta máquina)

Mesmo `aplicar` puro nas duas linguagens: copia o estado, d20 com xorshift de 32 bits dentro do estado, 1 herói × 2 mobs, eventos em lista. Código: `scratchpad/n1bench/sim.mjs` e `sim.go`. Node 22.22, Go 1.24.7.

| Medida | Node/TS | Go |
|---|---|---|
| 1.000 lutas (10.881 turnos) | 28,9 ms | 2,0 ms |
| 100.000 lutas (1,1 mi turnos) | 226 ms | 167 ms |
| RSS de processo HTTP ocioso | ~45 MB | ~7 MB |
| Artefato para distribuir | binário do Node 123 MB (base de um SEA) | 8 MB estático |

Os dois produziram **exatamente** os mesmos totais (10.881 turnos, 23.710 eventos): determinismo cruzado é possível nas duas.

## Análise por critério

1. **Simulador de 1.000 lutas:** 29 ms contra 2 ms. A diferença de 14× é quase toda aquecimento do JIT; em 100 mil lutas cai para 1,35×. Para o dono, os dois são "instantâneo". **Empate prático.**
2. **20–30 jogadores por zona, por turno:** chute generoso de 1 ação a cada 2 s por jogador dá ~15 comandos/s por zona. Pelo número acima, cada `aplicar` custa microssegundos. A CPU fica parada >99% do tempo em qualquer uma das duas. **Não decide.**
3. **GC:** jogo por turno não tem orçamento de 16 ms por quadro. Pausa de GC do V8 ou do Go é invisível aqui. **Não decide.**
4. **Memória na VPS:** 45 MB contra 7 MB de base. Numa VPS de 1 GB, as duas cabem com folga para centenas de jogadores. Só pesa se ele quiser a VPS mais barata possível com Postgres junto. **Vantagem leve Go, irrelevante até ter público.**
5. **Determinismo com inteiros (aqui está o risco real):**
   - TS: `number` é float64. Não existe estouro, existe **perda silenciosa de precisão** acima de 2^53, e operador bit a bit converte para int32 (`(2**31-1+1)|0` dá `-2147483648`). O sorteio precisa de `>>> 0` e `Math.imul` em todo lugar; esquecer um vira bug que só aparece na luta 700. Divisão precisa de `Math.trunc`.
   - Go: `int32` estoura com volta definida, divisão inteira é nativa. Mas há uma armadilha forte: **iterar `map` tem ordem aleatória**. Um `for _, m := range mobsPorID` dentro do núcleo quebra o replay sem nenhum erro.
   - Nas duas, a regra é a mesma: lista ordenada no estado, nunca mapa/objeto como fonte de ordem. **Empate, com armadilhas diferentes.**
6. **Distribuir junto do Unity:** 8 MB contra ~120 MB (ou exigir Node instalado). Para Steam isso é peso real no download e na complexidade do build. **Vantagem Go clara.** Não é performance, mas é o único número da tabela que o jogador sente.
7. **Playground no navegador:** TS roda nativo; Go exige WASM (binário de MBs, ponte JS). **Vantagem TS.**

**Veredito inicial:** performance não deve pesar mais que ~5% nesta decisão. Quem usar "Go é mais rápido" ou "Node aguenta" como argumento central está gastando tempo do dono. Os pesos que proponho para meu recorte: determinismo 40%, distribuição 35%, memória 15%, CPU 10%.

## Perguntas

1. **GO1:** como você garante, na revisão do código que o Sidinei escreve à mão, que nenhum `range` sobre `map` entre no núcleo? Lint (qual?), convenção, ou teste de replay com 1.000 sementes?
2. **TS1:** sorteio com semente em TS: você aceita `bigint` no estado (lento e não serializa em `JSON.stringify`) ou `number` com `>>> 0`/`Math.imul`? Mostre em até 10 linhas e diga como o dono percebe quando esqueceu um `>>> 0`.
3. **TS2 ou quem defender TS na distribuição:** qual o tamanho real do servidor TS empacotado para ir junto do jogo na Steam (Node SEA, Bun `--compile` ou Deno `compile`), e quanto de RAM ocioso? Quero número medido, não estimativa.

---

### N2 (rodada 1)

# N2 — Rodada 1 (neutro: mentor de dev solo)

## Critérios e pesos (opinião minha)

| Critério | Peso | Por quê |
|---|---|---|
| Risco de abandono e de nova troca | 25% | Nenhum projeto morreu por defeito técnico (briefing). A dúvida Go×TS já voltou 7 vezes (`acervo/LINHA-DO-TEMPO.md`, "A dúvida de linguagem"). |
| Autoria: a linguagem o mantém escrevendo à mão? | 20% | O ciclo "IA escreve o núcleo → não é meu → recomeço" aconteceu em abril (Go) e em outubro (TS). |
| Carga cognitiva em sessão de 30 min interrompida | 20% | Retomar depois de ninar o bebê pesa mais que velocidade de digitação. |
| Tempo até a luta no terminal (noite 5) | 15% | É o primeiro marco do roteiro. |
| Aprendizado e carreira | 10% | Go no trabalho tem valor, mas não é o objetivo do jogo. |
| Técnico (determinismo, contrato C#, distribuição) | 10% | Os dois resolvem. Nenhum dos dois compartilha código com o C# do Unity. |

## O que o código antigo mostra (fatos, com o autor)

Clonei os três bundles do acervo e olhei os trailers `Co-Authored-By`.

- **Go do MUD, depois do reset de 03/04** (`Artefatos/03-.../eras-do-brasil-game-arquivado`): cerca de 55 commits em 7 semanas (03/04 a 24/05), ~3.000 linhas em `server/`. Só 6 commits têm Copilot como coautor. **É o maior trecho escrito à mão em todo o acervo.** Também tem **zero testes** (`*_test.go`: nenhum) e foi feito com "cada manager dono do próprio estado" (ADR-007, linha 29 da LINHA-DO-TEMPO), que é o contrário do `aplicar` puro.
- **Go do idle** (`04-.../eras-do-brasil-idle`): 1.459 linhas e zero testes. Quatro commits com Claude, e o resto saiu do chat-tronco do ChatGPT (linha 35). A autoria é incerta, então não uso como prova.
- **TS do semi-idle** (`05-.../eras-do-brasil-semi-idle`): o único commit sem IA no código é o `da17716` (11/10). Os tipos estão fluentes, com uniões de literais e vitest com snapshot. Mas `initializeCombat` escreve `combatant.position = {...}` **no objeto que recebeu**, quebrando a regra "não mexer no estado que entrou". Dois dias depois vieram 9 commits com Claude Opus, que escreveu `position.ts` (448 linhas) e o playground.

## Análise inicial

1. **A razão da última troca morreu.** Em 09/10 o Go saiu porque o cliente ia ser web e daria para compartilhar código (linha 57). Com o cliente em Unity, esse argumento não vale para nenhum dos dois lados.
2. **TS:** é a linguagem em que ele pensa os tipos sem esforço, e isso aparece no commit de 11/10. Mas foi também onde a IA tomou o núcleo **mais rápido**: 2 noites. O monorepo, o React e o tsconfig convidam a pedir "monta aí". Ele escreveu um esqueleto, não uma regra.
3. **Go:** é onde ele mais insistiu sozinho, mas o que saiu tem cara de objeto com estado e não tem nenhum teste. Por ser iniciante em Go, ele tende a perguntar mais à IA. A favor: um único toolchain e as 8h por dia no trabalho, que treinam a mesma linguagem.
4. **Nos dois casos, ele nunca escreveu um teste por tabela nem uma função pura.** Esse é o músculo que falta, e nenhuma linguagem entrega isso sozinha.
5. Quanto à distribuição: Go gera um binário com `go build`. O executável único do Node (`--build-sea`) ainda está em "Stability 1.1, Active development" ([docs do Node](https://nodejs.org/dist/latest/docs/api/single-executable-applications.html)). Isso não decide nada antes da fase 5.

Opinião: tecnicamente dá empate. O desempate é **qual linguagem ele termina a noite 5 escrevendo sozinho**, e não qual é melhor.

## Perguntas

1. **TS1:** no único commit à mão em TS, `initializeCombat` muta `combatant.position` da entrada. O que no TS impede isso na noite 3, cansado, já que `readonly` não é profundo? E qual é o setup mínimo (tsconfig, ESM, vitest) que cabe em **uma** noite, sem monorepo e sem pedir à IA?
2. **GO1:** foram 7 semanas de Go à mão com 0 testes e managers donos de estado. Mostre, em até 10 linhas, os eventos de luta (ataque, dano, movimento, reação) em Go **com** a serialização JSON para o contrato. Quanto disso é cerimônia (interface, type switch, `MarshalJSON`) que um iniciante vai acabar pedindo à IA?
3. **TS2 e GO2:** se a linguagem de vocês travar na noite 12, qual regra concreta impede a 8ª reabertura da dúvida? "Estourou 1,5×, corta escopo" vale para trocar de linguagem também? Respondam com o que ele faz naquela noite.

---

### N3 (rodada 1)

# N3 (neutro, arquiteto de integração): abertura

Medi em vez de supor. Os scripts e binários estão em `scratchpad/n3-medidas/`.

## Critérios de peso (opinião)

| Critério | Peso | Por quê |
|---|---|---|
| Contrato com C# | alto | O dono não sabe C#. Tudo o que o gerador não fizer, ele escreve numa linguagem nova. |
| Servidor local no Steam | médio | Só conta na fase 5 (amigo joga). Não pode decidir a noite 1. |
| Local → VPS | baixo | Os dois lados fazem bem. |
| Playground | baixo | É só teste, e tem uma saída que tira o problema da mesa (abaixo). |

## Análise inicial (fatos medidos)

**1. Contrato: o problema é o C#, não TS × Go.**
- Testei o quicktype 23 com um JSON Schema `oneOf` de 2 eventos (`dano`, `movimento`). Em C# e em Go ele **achata a união** numa classe só, com campos anuláveis (`long? Valor`, `long? De`...). Nenhum dos dois lados ganha união discriminada no Unity "de graça".
- O `JsonUtility` do Unity só serve dado simples. A própria doc manda usar outra biblioteca para o que ele não suporta ([Unity Manual](https://docs.unity3d.com/Manual/json-serialization.html)). Na prática, isso significa Newtonsoft com um conversor por `tipo`.
- **Cobrança ao TS:** pelo quicktype, `type Evento = {seq: number ...}` vira `"type": "number"` no schema, que é ponto flutuante. Num projeto que exige "só inteiros", o contrato precisa de `integer` explícito (Zod `.int()`, JSDoc ou schema escrito à mão).
- **Cobrança ao Go:** Go não tem união. O núcleo terá `switch` por `tipo` e `UnmarshalJSON` à mão, e o C# terá o mesmo.
- Proposta neutra: `contrato/` tem **JSON Schema escrito à mão como fonte da verdade**. Os tipos do servidor e do C# são gerados a partir dele, e o conversor de `tipo` no C# é escrito uma vez.

**2. Distribuir o servidor local (Windows, servidor HTTP mínimo)**

| | .exe | .exe gzip |
|---|---|---|
| Go 1.24 (`-s -w`) | 5,8 MB | 2,4 MB |
| Bun 1.4.2 `--compile` | 86 MB | 40 MB |
| Node SEA | ≈ tamanho do `node` (123 MB no Linux, Node 22) | — |

- O Node SEA ainda está em "Stability 1.1, Active development". O `--build-sea` existe desde a 25.5 ([doc](https://nodejs.org/dist/latest/docs/api/single-executable-applications.html)).
- A doc do Bun admite: "Bun's binary is still way too big" ([doc](https://bun.com/docs/bundler/executables)).
- **Antivírus:** o Go também não escapa. O FAQ oficial reconhece falso positivo "especialmente no Windows" ([go.dev/doc/faq](https://go.dev/doc/faq)). Nos dois casos, a saída é assinar o binário.
- **Processo filho:** nenhuma linguagem resolve sozinha o servidor órfão quando o Unity fecha. Os dois precisam encerrar quando o stdin fechar ou quando o PID do pai morrer.
- Opinião: um build Unity já tem centenas de MB, então 86 MB pesam menos do que parece. O que pesa é runtime em estado "experimental" dentro de produto pago.

**3. "Sem framework" é meia verdade nos dois lados.**
- Nem Go nem Node têm servidor WebSocket na biblioteca padrão.
- No Go, o próprio `x/net/websocket` aponta `gorilla` ou `coder` como alternativas mais mantidas ([pkg.go.dev](https://pkg.go.dev/golang.org/x/net/websocket)).
- No Node, o caminho é o pacote `ws`. Só o Bun tem WebSocket nativo, e aí o runtime é o Bun, não o Node.

**4. Playground**
- Em TS ele é nativo.
- Em Go via WASM, um `syscall/js` trivial deu 1,7 MB (0,5 MB gzip). Funciona, mas cria uma cola a manter.
- Saída neutra: o playground vira **cliente WebSocket do servidor local**. Ele testa o contrato real, e a pergunta do WASM some.

**5. Local → VPS**
- Memória em repouso (Linux): Go 7,5 MB, Bun 16 MB, Node 22 46 MB. Numa VPS de 1 GB isso não muda o preço. Custo não desempata.
- Go: copiar o binário e usar systemd. TS: Node instalado ou `bun --compile` para Linux.

**Acervo:** a fase TS de outubro previa Electron para a Steam (`acervo/LINHA-DO-TEMPO.md`, marco 11). Foi inflação de stack antes de ter luta. Quem defender TS precisa dizer que isso não volta.

**Resumo (opinião):** integração não decide o debate. Ela pende para Go na distribuição, pende para TS no playground, e o contrato com C# custa igual nos dois.

## Perguntas

1. **TS1:** qual runtime vai no Steam: Node SEA (1.1, em desenvolvimento ativo), Bun `--compile` (86 MB) ou Node instalado? E como o contrato garante `integer` se `number` vira float no schema?
2. **GO1:** o evento com vários tipos vira struct achatada ou interface com `UnmarshalJSON`? E quantas linhas de C# à mão isso custa a um dono que não sabe C#?
3. **GO2 e TS2:** vocês aceitam o playground como cliente WebSocket do servidor local, em vez de WASM ou núcleo importado no navegador? Se não, o que se perde?

---

### MK1 (rodada 1)

# MK1 (Mercado): abertura

Sou analista de mercado. Não voto em linguagem para o núcleo, trago o que os dados dizem e separo **carreira** de **projeto**.

## 1. Dados (2025–2026)

**Uso e tendência**
- Stack Overflow 2025, devs profissionais: TypeScript 48,8%, Node.js 49,1%, Go 17,4%, C# 29,9%. Admirado: TS 58%, Go 56,5%. ([survey.stackoverflow.co/2025/technology](https://survey.stackoverflow.co/2025/technology))
- GitHub Octoverse 2025: TS virou a linguagem nº 1 do GitHub em contribuidores mensais (ago/2025, +66% no ano). O Brasil é o 4º país em número de devs (6,89 mi). O próprio GitHub chama a ligação "IA → linguagem tipada" de sinal observacional, não de causa. ([github.blog](https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/))
- JetBrains 2025: Go é a linguagem que mais devs querem adotar a seguir (11%). TS (6%) teve a maior alta de uso real em 5 anos. As duas estão no topo do "Promise Index". ([blog.jetbrains.com](https://blog.jetbrains.com/research/2025/10/state-of-developer-ecosystem-2025/))
- Go Developer Survey 2025: o uso principal é CLI e API. Jogos nem aparecem na lista. Mais de 80% aprenderam Go já empregados. Só 13% estão "muito satisfeitos" com a IA em Go, e 53% reclamam de código que não funciona. ([go.dev/blog/survey2025](https://go.dev/blog/survey2025))

**Salário (Brasil)**
- Go, pleno: R$ 8–14 mil CLT e US$ 40–75 mil/ano em vaga remota no exterior. São estimativas de blog, sem metodologia. ([jobrise.io](https://jobrise.io/pt/blog/golang-vagas-brasil-2026-quanto-paga/))
- Pleno genérico: R$ 6–12 mil CLT, também sem fonte primária. ([numerando.com.br](https://www.numerando.com.br/blog/salario-desenvolvedor-software-brasil-2026))
- **Aviso honesto:** não achei dado público confiável de salário **por linguagem** em 2025. O SO 2025 tirou o "top paying technologies" ([survey.stackoverflow.co/2025/work](https://survey.stackoverflow.co/2025/work)). O consenso de mercado é que Go paga um prêmio porque há menos oferta de devs, e que TS tem mais volume de vagas. Isso é opinião sustentada por proxies, não um número.

**Backend de jogos**
- Wildlife Studios (brasileira, cerca de 40 mi de DAU) usa Go como linguagem principal do backend. ([go.dev case](https://go.googlesource.com/website/+/master/_content/solutions/wildlifestudios.md))
- Nakama (Heroic Labs), servidor de jogo open source, é escrito em Go e aceita módulos em **Go, TypeScript ou Lua** ([docs](https://heroiclabs.com/docs/nakama/server-framework/introduction/)). Colyseus é o framework multiplayer de referência em Node/TS ([gitnation](https://gitnation.com/contents/making-multiplayer-games-with-colyseus-nodejs-and-typescript)). As duas linguagens são escolhas legítimas para servidor de jogo, nenhuma é exótica.

## 2. Carreira (separada do projeto)

- Sidinei já é **pleno CLT recebendo para escrever Go** na Gamersclub. Para a carreira, a hora que mais rende em Go é a do expediente, com code review de gente sênior e código em produção. Um projeto pessoal em Go soma pouco a isso, porque o currículo já diz "Go em produção".
- TS é o que ele tem de mais forte e o mercado de maior volume. Manter TS afiado é seguro de empregabilidade, e não precisa do jogo para isso.
- Opinião: o mercado não vai contratar ninguém por causa de um MMORPG pessoal em uma ou outra linguagem. Para portfólio, recrutador olha três coisas: **existe, roda e tem teste**. Um projeto pessoal só conta como portfólio se for entregue. O acervo mostra 13 roupagens e 5 trocas de stack sem entrega ([LINHA-DO-TEMPO](../../acervo/LINHA-DO-TEMPO.md)). Esse é o risco de portfólio, e ele independe da linguagem.
- Um caso a favor de Go: se o plano dele passa por subir para sênior em Go, um projeto pessoal em Go com determinismo, testes por tabela e servidor WebSocket é uma boa prova de domínio **além do CRUD**. O peso, porém, é marginal perto de 1–2 anos de Go no trabalho.
- Um caso a favor de TS: se o portfólio também for vitrine pública (playground no navegador, demo jogável), TS entrega isso com menos atrito.

## 3. O que isso significa para ele

Para a carreira, a linguagem do jogo é quase neutra: Go ele já ganha no trabalho, e TS ele já domina. O critério de mercado que sobra é **terminar**. Sugiro aos debatedores que parem de usar "carreira" como argumento forte, para os dois lados.

## 4. Perguntas

1. **GO1:** se Sidinei já pratica Go 8h por dia na Gamersclub, o que um projeto noturno em Go acrescenta à carreira dele que o trabalho não acrescenta? Ou o argumento é o inverso, que o jogo acelera o trabalho?
2. **TS1:** o Go Developer Survey mostra só 13% "muito satisfeitos" com IA em Go, e o Octoverse liga TS a uma IA mais confiável. Como isso pesa num projeto em que a IA **não escreve** o núcleo, só revisa e faz o laço de testes? A vantagem de IA do TS ainda vale aqui?

---

### MK2 (rodada 1)

# MK2: abertura (Mercado, backend de jogos)

Persona: analista de tendências de backend de jogos. Trago fatos com link e digo quando é opinião.

## 1. Dados atuais

**Adoção geral (Stack Overflow Survey 2025, uso entre profissionais)**, https://survey.stackoverflow.co/2025/technology

| Linguagem | Usa | Admirada |
|---|---|---|
| TypeScript | 48,8% | 58% |
| C# | 29,9% | 55,8% |
| Go | 17,4% | 56,5% |
| Rust | 14,5% | 72,4% |
| Elixir | 2,8% | 65,9% |

- TS virou a linguagem nº 1 do GitHub em contribuidores (ago/2025), puxada por código gerado com IA: https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/
- O compilador do TS 7 está sendo reescrito **em Go** pela Microsoft: https://infoworld.com/article/4100582/microsoft-steers-native-port-of-typescript-to-early-2026-release.html. As duas linguagens não competem por sobrevivência. As duas vão durar.

**Backend de jogos: quem usa o quê**
- **Go, Nakama (Heroic Labs).** É o servidor de jogo open source mais conhecido. O núcleo é em Go, e a lógica do jogo pode ser escrita em Go, TypeScript ou Lua, com a mesma API de partida: https://heroiclabs.com/docs/nakama/concepts/multiplayer/authoritative/llm.md. A empresa diz que tem "centenas de títulos", mas o número vem dela mesma: https://www.gamedeveloper.com/game-platforms/embracing-the-future-of-open-source-game-servers-with-nakama
- **TS, Colyseus.** A 0.17 saiu em fev/2026, com reconexão automática e tipos compartilhados entre cliente e servidor: https://colyseus.io/blog/colyseus-017-is-here/
- **TS, Hathora.** Foi comprada pela Fireworks AI em mar/2026 e saiu do mercado de jogos: https://gamesbeat.com/?p=318174. Por causa disso, o online do Stormgate desligou em abr/2026: https://www.techspot.com/news/111969-stormgate-servers-go-dark-following-ai-focused-hosting.html
- **C#.** O Albion, que é a referência de progressão do dono, roda com Unity + Photon Server: https://blog.photonengine.com/?p=3622. O estúdio diz que "o núcleo da simulação é totalmente independente da Unity": https://gamedev.net/news/3258-architecting-albion-online-how-sandbox-interactive-built-a-pvp-mmo-to-scale/
- **Rust, SpacetimeDB 1.0** (mar/2025). Junta banco e servidor num só sistema e roda o backend do MMO BitCraft: https://spacetimedb.com/blog/introducing-spacetimedb-1-0
- **Node, Screeps.** É um MMO de estratégia vendido na Steam, com o servidor em Node aberto: https://github.com/screeps/screeps

**Padrão de servidor por turno.** A doc do Nakama separa dois tipos:
- **turno ativo**, com os jogadores conectados e poucas mensagens. Nesse caso o tick é baixo;
- **turno passivo**, em que o servidor valida, grava e para o loop até a próxima jogada.

Fonte: o mesmo link do Nakama acima. O Eras é turno ativo dentro de uma instância. É o mesmo desenho do `aplicar` puro com o servidor autoritativo, que é o padrão da indústria.

**Carreira (Brasil).** As faixas salariais de Go pleno CLT giram em torno de R$ 8–14k, mas a fonte não cita de onde tirou os números: https://jobrise.io/pt/blog/golang-vagas-brasil-2026-quanto-paga/. Trate como estimativa.

## 2. O que isso significa para ESTE dono

1. **O mercado não decide esta escolha.** Go e TS são mainstream, bem admirados e têm ecossistema de jogo de verdade (Nakama × Colyseus). Nenhum dos dois é aposta arriscada.
2. **Framework não conta agora.** O briefing já decidiu "sem framework". Nakama e Colyseus só entram na fase 6 (online), se entrarem. Opinião: quem usar "o ecossistema de jogo é maior" como argumento hoje está vendendo o que só vai servir na fase 6.
3. **A lição do Hathora favorece servidor próprio, em binário, local.** Depender de plataforma gerenciada mata jogo pequeno. Os dois lados conseguem gerar um executável. Fica para TS2 e GO provarem o custo de cada um.
4. **C# fica de fora, mas é o que o mercado de Unity usa no servidor (Photon/Albion).** O mercado mostra que núcleo em C# compartilhado com o cliente é o caminho "padrão". Opinião: o dono rejeitou isso com razão, porque não sabe C# e o histórico dele mostra que trocar de stack mata projeto.
5. **Carreira.** Go tem retorno duplo: ele já usa no trabalho, e cada noite no Eras também treina para a Gamersclub. Em TS, o retorno de carreira é ganhar profundidade, não ganhar mercado novo. Opinião: para quem é a única renda da casa, o retorno duplo pesa.
6. **IA.** TS ter o maior volume de código gerado por IA no GitHub é faca de dois gumes para quem quer recuperar autoria. Esse é o mesmo ponto que TS2 admitiu.

## 3. Perguntas

- **GO1:** o Nakama aceita lógica em Go ou TS com a mesma API. Na fase 6, você recomendaria migrar o servidor próprio para dentro do Nakama ou manter o processo sem framework? Se for migrar, o `aplicar` em Go entra sem reescrita?
- **TS1:** o Hathora era o maior nome de hospedagem do mundo TS e saiu do mercado. Na fase 6, qual é o plano de hospedagem TS que não depende de plataforma gerenciada? Quanto de RAM um processo Node ou Bun ocioso ocupa numa VPS de 1 GB, comparado a um binário Go?

---

### MK3 (rodada 1)

# MK3 (Mercado): ecossistema, estabilidade e IA

## 1. Dados atuais

**Adoção e tendência**
- Stack Overflow 2025: TS é usado por 43,6% (48,8% dos profissionais), Go por 16,4%. Na pergunta "quer usar", TS tem 29,2% e Go 19,4%. Na pergunta "usou e quer continuar", TS tem 72,4% e Go 55,8%. https://survey.stackoverflow.co/2025/technology
- JetBrains 2025: Go é a linguagem que mais devs pretendem adotar (11%; TS 6%). TS, Go e Rust lideram o Language Promise Index. https://blog.jetbrains.com/research/2025/10/state-of-developer-ecosystem-2025/
- GitHub Octoverse 2025: TS passou a ser a linguagem nº 1 em contribuidores mensais (2,64 milhões, +66% no ano). O relatório liga isso à IA: 94% dos erros de compilação em código gerado por LLM são erros de tipo. https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/

**Trajetória e estabilidade**
- O TypeScript 7 é o compilador portado para Go, cerca de 10× mais rápido. O beta saiu em 21/04/2026 e a versão final em julho. O pacote mudou padrões: `strict` agora vem ligado e `module` vem como `esnext`. Removeu ou deixou como erro `moduleResolution node10`, `baseUrl` e `target es5`. A API programática só fica estável na 7.1. https://devblogs.microsoft.com/typescript/announcing-typescript-7-0-beta · https://adtmag.com/articles/2026/07/10/microsoft-bets-typescript-future-on-a-native-compiler.aspx
- Go 1.26 (fev/2026) traz `new(expr)`, genéricos que se referem a si mesmos, o coletor de lixo Green Tea como padrão e um `go fix` que moderniza o código sozinho. Tudo segue a promessa de compatibilidade do Go 1. https://go.dev/doc/go1.26
- Runtimes JS:
  - O Node 24.12 LTS tornou estável a remoção de tipos: roda `.ts` direto, mas só com sintaxe que dá para apagar (sem `enum`) e sem checar tipos. https://x-cmd.com/blog/251218
  - A Anthropic comprou o Bun em 03/12/2025, que segue com licença MIT. https://anthropic.com/news/anthropic-acquires-bun-as-claude-code-reaches-usd1b-milestone
  - O Deno 2 continua como terceira opção.

**Assistência de IA (benchmarks)**
- Multi-SWE-bench (corrigir issues em repositórios reais): TS e JS "consistentemente" tiveram as menores taxas. A média que calculei pela Tabela 4 é parecida: Go ~3,5% e TS ~3,4%. https://arxiv.org/abs/2504.02605
- Multi-LCB (2026): Go e TS ficam no mesmo patamar, abaixo de Python. No GPT-OSS-120B, Go fez 69,9 e TS 70,3. https://arxiv.org/pdf/2606.20517
- Leitura: nos benchmarks, as duas linguagens empatam.

**Backend de jogos**
- Nakama (Heroic Labs) é um servidor de jogo em Go que roda módulos em TS, JS e Lua. https://heroiclabs.com/blog/announcing-nakama/
- Agones (Google, servidores dedicados) é feito em Go. https://www.mintlify.com/googleforgames/agones/introduction
- Colyseus é o equivalente em Node/TS. https://colyseus.io
- As duas linguagens têm ecossistema de backend de jogo.

**Distribuição junto do Unity**
- Go compila em um binário só, com cross-compile nativo.
- No lado JS existem `bun build --compile` (https://bun.sh/docs/bundler/executables) e o executável único do Node (SEA), que ainda é experimental (https://nodejs.org/api/single-executable-applications.html).

## 2. O que isso significa para este dono

1. **Longevidade não decide.** As duas linguagens estão em alta e têm dono forte (Microsoft e Google). Quem usar "risco de morrer" como argumento está forçando.
2. **A estabilidade é de tipos diferentes.**
   - Em Go, a linguagem e o ferramental são uma coisa só (`go.mod` + `go test`). Os dois Go antigos do acervo tinham uma dependência direta só: `coder/websocket` em `.../03-mud-moderno-pre-fusao/.../server/go.mod`.
   - Em TS, a linguagem é estável, mas o redor muda: padrões de tsconfig a cada versão grande, ESM×CJS, a escolha entre Node, Bun e Deno, o runner de teste. O semi-idle (`.../05-semi-idle-typescript/eras-do-brasil-semi-idle`) já tinha workspaces npm, Vitest 5, Vite/React, três tsconfig e mise. Parte desse projeto foi escrita por IA, então ele mostra o tamanho do ferramental, não as escolhas dele.
   - Opinião: para quem tem 2h por noite e trocou de stack 5 vezes, cada decisão de ferramental é uma porta para a "troca de novo".
3. **A IA empata.** O que muda é o contexto. Em TS, a IA erra mais pela configuração (módulo, runtime) do que pelo código. Em Go, erra mais por estilo verboso. Nenhuma das duas justifica a escolha.
4. **Carreira.**
   - TS tem o mercado maior (Stack Overflow, Octoverse) e é a especialidade dele.
   - Go é a linguagem que mais cresce em intenção de adoção e é a do emprego atual.
   - Opinião: praticar Go à noite rende no trabalho de dia. Praticar TS rende em vagas futuras. Os dois casos são defensáveis. A diferença é que um tem uso hoje e o outro é uma aposta.
5. **Remoção de tipos favorece um estilo.** Ela não aceita `enum`, o que empurra o TS para uniões discriminadas com literais. Para os eventos de luta, isso é bom e cabe sem build.

## 3. Perguntas

- **TS1:** O TS 7 mudou padrões e removeu opções (`baseUrl`, `node10`), e o `tsgo` só ganha API estável na 7.1. Qual combinação exata você congela na noite 1 (runtime, versão do TS, runner de teste, número de arquivos de config)? Como garante que isso não vire uma "noite de ferramental" na noite 15?
- **GO2:** O Sidinei se diz iniciante em Go, e boa parte do acervo Go foi escrita por IA. Que evidência do próprio acervo mostra código Go escrito por ele? E sem união discriminada nativa, como você modela os eventos de luta sem depender de `interface` + `type switch`, onde esquecer um caso só aparece em tempo de execução?


## Rodada 2: conversa

---

### N4 (rodada 2)

# N4 (moderação): rodada 2

Lembrete: a escolha precisa durar até a fase jogável, porque antes disso a stack não se rediscute. A dúvida Go × TS já voltou 7 vezes. Quem só empilha vantagem técnica não está respondendo à pergunta do debate.

## Perguntas ainda sem resposta boa

**Lado Go**
- **Número de noites até a noite 5 em Go** (minha P1): ninguém deu.
  - O GO1 admitiu que o código pós-reset usa `math/rand` global e não tem teste.
  - O GO3 afirmou que esse código é do dono, mas não mostrou trecho.
- **Condição para aceitar o TS** (minha P3): só o GO4 respondeu. O GO2 disse "baixa, mas não zero" e não deu nenhuma condição.
- **N2 → GO1:** eventos com o JSON do contrato em 10 linhas, e quanto disso é cerimônia que o dono vai acabar pedindo à IA. Ninguém mostrou ainda.

**Lado TS**
- **Que razão do TS sobra com o cliente Unity** (minha P2): o TS3 respondeu "playground nativo". Só que o N3 mostrou que o playground pode ser cliente WebSocket do servidor local, e aí essa vantagem some.
- **O que impede outro `position.ts`:** ninguém deu um mecanismo. O TS4 deu condições (`tsc --noEmit` strict desde a noite 1). Falta o TS1, o TS2 e o TS3 assinarem embaixo.
- **Condição para aceitar o Go:** o TS2 diz "3 de 10" e não dá nenhuma.

## Argumentos que já empataram (peso zero daqui em diante)

| Argumento | Por que empatou |
|---|---|
| Performance | O N1 mediu: 1.000 lutas levam 29 ms contra 2 ms. Os dois são instantâneos. |
| Hospedagem e memória | A diferença de RAM (N1, N3) não muda o preço da VPS. |
| Teste por tabela | Admitido pelos dois lados. |
| Longevidade, ecossistema, ajuda de IA | Medidos pelo MK3; os benchmarks dão empate. |
| Contrato com C# | O N3 mostrou que o quicktype achata a união nos dois, e o C# custa igual. |
| "Sem framework" | Nenhum dos dois tem WebSocket na biblioteca padrão (N3). |
| Imutabilidade | O `Readonly` é raso e o slice do Go compartilha memória. A correção é a mesma nos dois: um teste que prova que a entrada não mudou. |
| Carreira | O MK1 chamou de "quase neutra". O retorno duplo do Go existe, mas é marginal. |
| Eventos (TS) × inteiros (Go) | Cada lado admitiu a fraqueza e a remendou com linter (`exhaustive` ou lint de `/`). É disciplina compensando a linguagem dos dois lados, e isso se anula. |

Argumentos fora do assunto:
- "A Microsoft reescreveu o compilador do TS em Go" (GO2): fala de compilador, não de servidor de jogo.
- "Go é bom para proxy HTTP" (TS2): retórica.

## O que ainda decide

1. **Autoria, pela prova do acervo** (N2):
   - em Go foram cerca de 55 commits à mão em 7 semanas, com 0 testes;
   - em TS foi 1 commit à mão, e a IA tomou o núcleo em 2 noites.

   Esse é o dado mais forte do debate até aqui, e ele pende para o Go. A fadiga de abril pende para o TS.
2. **Tempo até a noite 5.** É a vantagem de fluência do TS, mas ninguém mediu. Adoto o **teste de 1 noite do TS4**: o dono escreve à mão o mesmo `aplicar(ataque)` nas duas linguagens. Proponho que esse teste vire o desempate oficial.
3. **Portas de ferramental.** No acervo TS, o `packages/core` não tinha `tsconfig`, o projeto tinha 3 configs e ainda mise. No Go havia 1 dependência. Para quem troca de stack, cada porta aberta conta.
4. **Distribuição.** Vantagem real do Go (5,8 MB contra 86 MB, e o SEA ainda é experimental), mas só pesa na fase 5.

## Cobranças

1. **TS1, TS2, TS3:** pelos commits, foram 7 semanas escrevendo Go à mão contra 2 noites até a IA tomar o TS. Com que prova o TS é a linguagem da autoria? E vocês se comprometem, por escrito, com `tsc --noEmit` strict e com um core sem monorepo desde a noite 1? Sim ou não.
2. **GO1, GO2:** vocês aceitam o teste de 1 noite do TS4 como desempate? Se a versão em Go levar o dobro de tempo, aceitam o TS? E digam o número de noites até a noite 5. Quem não der o número estará defendendo o time, não o dono.

---

### TS1 (rodada 2)

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

---

### GO1 (rodada 2)

# GO1, rodada 2

## Respostas

**@N4:** "trecho que seja dele e noites até a noite 5". Commit `79bcee4` (20/04, sem coautor), `server/world/game_time.go`: `func (gt *GameTime) AdvanceTime(...)` muda o próprio objeto pelo ponteiro. A versão pura (`CalculatePeriodOfDay`) só apareceu no commit seguinte, que tem o Copilot como coautor. Então o estilo dele em Go é mutar. Estimativa (opinião): Go chega à noite 5 na noite 6 ou 7; TS chega na 5.
**@N4 (os dois lados):** o gatilho do acervo é "IA escreveu o núcleo". Go não impede isso. Eu aceito TS se, numa noite de teste, ele não conseguir escrever em Go, sem pedir código, um `aplicar` de 1 ataque com 3 casos na tabela.

**@TS1:** "como o compilador impede `append` no slice que entrou?" Não impede. Fica a cargo do teste: o laço da tabela compara a entrada antes e depois com `reflect.DeepEqual`. Isso é escrito uma vez e roda em todo caso.

**@TS2 e @TS3:** "exaustividade sem linter?" Não existe. Admito: é ferramenta compensando a linguagem. O custo é pequeno, sem golangci-lint: `type Tipo string` com constantes e `go run github.com/nishanths/exhaustive/cmd/exhaustive ./...` no Makefile. Leva meia noite, não várias.

**@N2 e @N3:** "eventos com JSON; struct achatada ou interface?" Achatada. Não tem `MarshalJSON` nem type switch, e no C# vira uma classe espelho com uns 15 linhas.
```go
type Tipo string
const (Ataque Tipo = "ataque"; Dano Tipo = "dano"; Movimento Tipo = "movimento"; Reacao Tipo = "reacao")
type Evento struct {
    Seq   int    `json:"seq"`
    Tipo  Tipo   `json:"tipo"`
    De    int    `json:"de,omitempty"`
    Alvo  int    `json:"alvo,omitempty"`
    Valor int    `json:"valor,omitempty"`
    Para  *Casa  `json:"para,omitempty"`
}
```
O preço: nada amarra que campo pertence a qual tipo. Nisso a união do TS é melhor.

**@TS4 e @MK1:** "o que o projeto acrescenta à carreira?" Função pura, teste por tabela e determinismo, coisas que o trabalho talvez não peça. A volta também vale: as 8h diárias de Go tiram o atrito de sintaxe à noite. Peso baixo, 10% no máximo.

**@N1:** "como evitar `range` sobre `map`?" A regra é não ter `map` no `Estado`, só slice ordenado. Além disso, um teste de replay roda 1.000 sementes duas vezes e compara o hash. Como o Go embaralha a ordem do `map` de propósito, o erro aparece no teste em vez de aparecer na luta 700.

**@MK2:** "migrar para o Nakama?" Não migraria antes de existir necessidade. Se migrar, `regras/` entra sem reescrita, porque é um pacote puro. Só o `server/` muda.

## Réplicas

**TS1:** "`Readonly` faz o compilador recusar a mutação." `DeepReadonly` não vem no TS: ou ele escreve um tipo recursivo, ou instala biblioteca. E o commit `da17716` dele muta a entrada (N2). Nas duas linguagens quem segura isso é o teste.

**TS2:** "Para um jogo Steam [...] a diferença não importa." Concordo sobre os MB. O que importa são os passos de build, e isso só pesa na fase 5.

## Concessão
Baixo o peso da distribuição: N1 e TS1 mostraram que ela não decide nada antes da fase 5. E o `game_time.go` derruba o meu argumento de que o Go dele já estaria pronto. As duas linguagens começam do mesmo músculo, que está faltando (N2).

## Pergunta
**TS1:** mostre em até 10 linhas o `DeepReadonly` que ele escreveria sem biblioteca. Se o tipo vier do `type-fest`, ele entende o tipo do próprio estado?

---

### TS2 (rodada 2)

# TS2 — rodada 2

## Respostas

**@GO1, @GO2, @N1:** "Comando exato, MB, dependências, RAM medida."
Medi agora com Bun 1.4.2 (`scratchpad/ts2-medida/server.ts`, WebSocket + JSON, porta livre):
- `bun build --compile --target=bun-windows-x64 server.ts --outfile eras-server.exe` → **86 MB** (40 MB em gzip -9). Compilado de Linux para Windows em 0,4 s.
- **Zero `node_modules`**, porque o WebSocket já vem no `Bun.serve`.
- RAM ociosa no Linux: **18 MB** RSS.
- Um passo de build. O Unity sobe o processo com `Process.Start`, igual ao Go.

Node SEA eu descarto: está em 1.1 e precisa de bundler e do pacote `ws`. O número perde para os ~8 MB do Go, e não vou fingir que não.

**@GO4:** "Semente em 64 bits sem `BigInt`? `Readonly` é raso."
Um d20 não precisa de 64 bits. O sfc32 guarda 4 × uint32 = 128 bits de estado, todo em `number`, com `Math.imul`/`>>> 0`, e serializa em JSON sem `BigInt`. `Readonly` é raso mesmo, por isso uso um `DeepReadonly<T>` de 6 linhas. Na execução, o laço de teste faz `deepFreeze(entrada)` antes de chamar `aplicar`. Como ESM roda em modo estrito, qualquer mutação **lança erro** no teste.

**@N4:** "Que razão a favor do TS sobrevive com Unity? O que impede outro `position.ts`?"
Sobrevivem a modelagem de eventos com exaustividade nativa (o GO1 e o GO2 concederam) e o fato de ele pensar tipos em TS. O `position.ts` nenhuma linguagem impede. O que reduz o convite: um pacote só, sem monorepo, sem React, sem Electron.

**@N4:** "Gatilho do acervo e fato que faria aceitar a rival."
O gatilho é a inflação de stack (Electron em 10/10), e a resposta é Bun puro. Aceito Go se, numa noite de teste, ele escrever em Go o `aplicar` de um ataque com uma tabela de 5 casos sem perguntar sintaxe à IA.

**@N2:** "Travou na noite 12, o que ele faz?"
Escreve a linha no registro, corta escopo (tira a reação, por exemplo) e segue. O 1,5× vale para a linguagem também. Ela só se rediscute com a fase 2 jogável.

**@N3:** "Playground como cliente WebSocket?"
Aceito. Ele passa a testar o contrato real, a vantagem "TS roda no navegador" deixa de pesar e ainda some um app que a IA adoraria montar.

## Réplicas

**GO2:** "Se nem o time do TypeScript quer o ferramental em Node, por que o servidor do Sidinei deveria?"
O TS 7 foi para Go pela velocidade do *compilador*, e escolheram Go porque a estrutura do código JS portava quase linha a linha. É argumento de compilador com milhões de linhas, não de servidor por turno. Ataque de efeito, não técnico.

**N1:** "Iterar `map` tem ordem aleatória."
Bom ponto, e aqui o TS ganha de leve: a ordem de chaves de objeto e de `Map` é definida pela especificação (inserção). Num `for...of` sobre `Map` o replay não quebra. Em Go, `range` sobre map é aleatório **de propósito**.

## Concessão

O GO2 provou que o core TS do acervo não tinha `tsconfig` e só rodava pelo Vitest, que remove tipos sem checar. Então a "exaustividade de graça" que eu vendi depende de `tsc --noEmit` no script de teste. Sem isso, o TS dá o mesmo zero de garantia que o Go sem linter. Os dois precisam de um passo extra.

## Nova pergunta

**GO2:** em Go, qual é o equivalente do `deepFreeze` no teste? Como o laço prova que `Aplicar` não mexeu nos slices da entrada sem você escrever um `Clonar()` profundo à mão só para comparar?

---

### GO2 (rodada 2)

# GO2, rodada 2

## Respostas

**@N4, @TS4 e @MK3** (as três pediram código Go que seja dele e perguntaram se mostra fluência).
O código dele está no MUD depois do reset de 03/04: `acervo/Artefatos/03-mud-moderno-pre-fusao/eras-do-brasil-game-arquivado/server/`. São ~3.000 linhas em 7 semanas, com commits em 26 dias diferentes. Só 6 commits têm o Copilot como coautor (`git log` do bundle). Isso desmente o "parou logo, por fadiga". Um trecho de `server/world/game_time.go` (commit `8dd1d55`, 06/04):
```go
func (gt *GameTime) AdvanceTime(duration time.Duration) {
    gt.Time = gt.Time.Add(duration)
    gt.PeriodOfDay = CalculatePeriodOfDay(gt.Time)
}
```
É Go legível e idiomático, sem luta com a sintaxe. Mas `server/npc/score.go` (linhas 26–37) escolhe o vencedor com `range` sobre um `map[string]float64`. A ordem de iteração de map em Go é aleatória ([spec](https://go.dev/ref/spec#For_range)), então um empate sai diferente a cada execução. Ou seja: ele é fluente em Go, mas cru em determinismo.

Número de noites até o marco da noite 5 (opinião, não medi): TS em 5 noites, Go em 6 a 7. O custo extra do Go são as tags JSON, o `Clonar` e o primeiro `_test.go`.

**@TS1:** "você aceita depender do linter `exhaustive`?"
Aceito, e admito que é ferramenta compensando a linguagem. O custo é uma linha: `go get -tool github.com/nishanths/exhaustive/cmd/exhaustive` (diretiva `tool`, [Go 1.24](https://go.dev/doc/go1.24#tools)). Depois ele roda `go tool exhaustive ./...` no mesmo script do `go test`. Ele não mantém nada nisso.

**@TS2:** "Como o Go impede `e.Mobs[0].PV -= 5` sem clonar à mão?"
Com slice, não impede. Mas array em Go é valor ([Effective Go](https://go.dev/doc/effective_go#arrays)): com `Mobs [MaxMobs]Mob` e `NMobs int`, passar `Estado` por valor copia tudo, e a linha acima só altera a cópia. Uma luta tem tamanho máximo, então o array cabe. Para o resto, uso um teste genérico que roda em todo caso da tabela: guarda uma cópia da entrada e compara com `reflect.DeepEqual` depois de `Aplicar`. O `readonly` do TS também não salvou o único commit à mão dele em TS: o `initializeCombat` muta `combatant.position` (achado do N2).

**@MK3:** "eventos sem `interface` + `type switch`?"
Uso uma struct achatada: `Tipo TipoEvento` (constante) mais campos com `omitempty`. O JSON é trivial e o C# lê a mesma forma, sem `UnmarshalJSON`. Um `switch ev.Tipo` sobre a constante é exatamente o que o `exhaustive` checa.

**@N2:** "trava na noite 12: o que impede a 8ª reabertura?"
O "1,5× corta escopo" vale para a linguagem também. Naquela noite ele escreve uma linha no registro ("travei em X"), pede à IA só o conceito, e se não destravar corta a feature. Abrir a dúvida de linguagem só é permitido depois da fase jogável, pela regra do N4.

**@N3:** "aceita o playground como cliente WebSocket?"
Aceito. Ele testa o contrato real. Perde-se rodar o núcleo offline no navegador e mandar um link. Nada disso importa antes de um amigo jogar.

## Réplicas

**TS4:** "pedir que o projeto seja também a escola de Go põe duas metas numa noite de 2h."
Argumento justo. Só que o próprio dono escreveu em 03/04: *"o problema não é Go, é a construção da lógica [...] isso afeta até JS/TS"* (`historico/sessions/2026-04-03-reset-reimplementacao-do-zero.md`). A fluência em TS não impediu a IA de tomar o núcleo em 2 noites.

**N2:** "ele nunca escreveu um teste por tabela nem uma função pura."
Concordo. Isso pesa mais que a linguagem.

## Concessão

Eu disse que Go "garante" o determinismo. Errei: o `score.go` dele mostra a armadilha própria do Go, que é a ordem aleatória do `map`, além do `float64`. O tipo `int` não basta. A regra precisa ser "sem map nem float em `regras/`", e isso é disciplina, igual ao `Math.trunc` do TS.

## Nova pergunta

**TS4:** você propôs o teste de uma noite com `aplicar(ataque)` nas duas linguagens. Aceita este critério fechado: tabela de 5 casos passando, sem código de IA, e ganha quem terminar primeiro? E, se empatar, Go leva por inteiros e binário?

---

### TS3 (rodada 2)

# TS3 – Rodada 2

## 1. Respostas
Ninguém me fez pergunta pelo nome. Respondo as duas que foram para "quem defender TS".

**@GO4:** *"64 bits sem `BigInt` e sem limitar a 32? E quanto o `Readonly<>` protege, sendo raso?"*
Para um d20, 64 bits não fazem falta. O que conta é o período do gerador, e `xoshiro128**`/`sfc32` usam 4 × 32 bits só com `Math.imul` e `>>> 0`, com período de 2^128. O estado vira `rng: [a, b, c, d]`, que vai para o JSON sem `BigInt`. Quanto ao `Readonly<>`, ele protege só na compilação e só no primeiro nível. Por isso é preciso marcar `readonly` em cada campo e em cada array dos tipos do estado, e em todo caso da tabela rodar o `aplicar` sobre um estado congelado em profundidade. É a mesma muleta do `slices.Clone` de vocês, só que do outro lado.

**@N1:** *"Tamanho real do servidor TS empacotado, e a RAM ociosa. Número medido."*
Medi hoje, com Bun 1.4.2 e um servidor WebSocket + JSON sem dependências (o `Bun.serve` já traz WebSocket):
- `bun build --compile --target=bun-windows-x64` → **86 MB**;
- versão Linux → **81 MB**, com **~18 MB de RSS** parado.

O build cruzado para Windows funcionou a partir do Linux. Não consegui medir o Go aqui, porque a rede bloqueou o `go get`, mas é bem menor. Ponto do Go.

## 2. Réplicas
**@N4:** *"O motivo que levou ao TS em 10/10 foi 'o cliente vira web'. Esse motivo sumiu. O TS precisa de outro."*
Concordo que sumiu. Sobram três motivos, em ordem de peso:
1. **Fluência até a noite 5.** É o primeiro prazo medido, e cumprir prazo é o que segura uma troca de stack.
2. **Eventos como união.** Os eventos são o contrato inteiro com o Unity, não um detalhe.
3. **Playground.** O navegador importa o `aplicar` direto, sem WASM. É o mais fraco dos três, já que o playground é "só teste".

Se só restasse o terceiro, eu mudaria de lado.

**@GO3:** *"Em TS ele precisa escolher runtime, tsconfig, test runner, bundler. Com 2h por noite, cada decisão é uma noite."*
É verdade para o TS de outubro: monorepo, Vite, Vitest e mise. Proponho cortar isso a zero. Testei agora: o Node 22 roda `node --test regras.test.ts` com tipos, sem flag e sem nenhuma dependência. Fica um `package.json` só com o `typescript`, o comando `tsc --noEmit && node --test` e nada mais. O bundler só entra na fase 5. São uma noite e uma linha de decisão, não quatro noites.

## 3. Concessão
O N3 mostrou que, no único commit de TS feito à mão (`da17716`), `initializeCombat` muta a entrada e o core nunca passou por `tsc`. Cedo nisso: **o argumento "o tipo protege" não está provado para este dono**. Ele só vale se `tsc --noEmit` com `strict` estiver no comando de teste desde a noite 1, junto do teste de entrada congelada. Sem isso, TS e Go empatam em imutabilidade: os dois dependem de disciplina.

## 4. Nova pergunta
**@GO2:** com o `aplicar` em Go, como fica o playground que o briefing pede? Se a resposta for "cliente WebSocket do servidor local" (proposta do N3), quanto HTML/JS o Sidinei escreve à mão para desenhar a luta, e em que linguagem? Se for JS, o projeto não acaba com duas linguagens de qualquer jeito?

---

### GO3 (rodada 2)

# GO3 — rodada 2

## Respostas

**@TS3:** "`GOOS=js GOARCH=wasm` ou TinyGo? Qual o tamanho do `.wasm` e quantas noites até a primeira luta na tela?"
Nenhum dos dois. Fico com a saída do N3: o playground vira **cliente WebSocket do servidor local**. É uma página HTML com `new WebSocket(...)` e um `<pre>` que mostra os eventos. Custo: zero noites de WASM, e o playground passa a testar o contrato real, que é o que o Unity vai usar. Se ele insistir em rodar o núcleo dentro do navegador, o N3 já mediu: `syscall/js` trivial dá 1,7 MB (0,5 MB gzip) (`scratchpad/n3-medidas/`). Eu estimo uma noite de cola (opinião, não medi). Nisso o TS é melhor, e digo isso sem ressalva. Só que o playground é ferramenta de teste e entra depois da noite 5, então não decide o marco.

## Réplicas

**TS4:** "Um teste de uma noite: ele escreve à mão o mesmo `aplicar(ataque)` de ~30 linhas em cada linguagem, sem pedir código à IA."
Concordo e assino embaixo. É o único argumento da mesa que mede o dono, e não a linguagem. Proponho duas regras para o teste: o mesmo caso de tabela nas duas versões e um cronômetro. E o resultado vale para sempre, segundo a regra anti-troca do N4. Se ele terminar o Go travado e o TS fluindo, eu voto TS.

**TS1:** "`Readonly` faz o compilador recusar a mutação da entrada."
Isso vale em parte. O `DeepReadonly` não vem com o TypeScript: ou ele escreve o tipo recursivo à mão, ou instala `type-fest`. Isso é mais uma decisão de ferramenta, que é justamente meu ponto. E o N2 mostrou que, no único commit à mão em TS (`da17716`), `initializeCombat` muta a entrada. Ou seja, o tipo não estava ligado (TS4 confirmou: `packages/core` não tinha `tsconfig`). Nos dois lados, quem pega a mutação é um **teste** que compara o estado antes e depois. Proponho que isso seja a regra comum, seja qual for a linguagem.

## Concessão

Abandono o argumento "simplicidade de Go como freio contra a troca". O N2 mostrou que as 7 semanas de Go à mão saíram com **zero testes** e com managers donos do próprio estado, o oposto de `aplicar` puro. Go não impediu a fadiga nem empurrou ele para função pura. Então Go não protege do padrão do acervo; quem protege são a regra 1,5× e o teste de uma noite. O que mantenho é menor e verificável: um toolchain só e o binário de 5,8 MB (medida do N3).

## Nova pergunta

**N2:** você pesou "carga cognitiva em sessão de 30 min interrompida" em 20%. Retomar depois de ninar o bebê fica mais fácil numa linguagem com um jeito só de escrever (Go, `gofmt`) ou na linguagem que ele lê sem esforço (TS)? Que fato do acervo sustenta a sua resposta?

---

### TS4 (rodada 2)

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

---

### GO4 (rodada 2)

# GO4: rodada 2

## Respostas

Ninguém me fez pergunta pelo nome. Respondo a única pergunta que foi para "os dois lados".

**@N4:** "citem um gatilho de troca do acervo e mostrem como a sua linguagem o evita; e qual fato faria vocês aceitarem a rival."

- **Gatilho:** inflar a stack antes de ter luta. A fase TS de outubro tinha monorepo, Vite/React, três tsconfig, mise e Electron antes de existir uma regra testada (`acervo/LINHA-DO-TEMPO.md`, marco 11; N3 e MK3 mostraram isso). O Go não evita essa vontade, mas oferece menos coisa para inflar: `go mod init`, `go test`, `go build`. O MUD tinha uma única dependência, `coder/websocket` (MK3).
- **Ressalva honesta:** os três fracassos em Go não foram por ferramental. Foram por troca de gênero e por IA no núcleo. Nesse gatilho a minha linguagem não ajuda em nada.
- **Fato que me faz aceitar TS:** o teste de uma noite que o TS4 propôs. Ele escreve à mão o mesmo `aplicar(ataque)`, com a tabela de testes, nas duas linguagens. Se a versão Go não ficar verde numa noite de 2h e a TS ficar, eu passo para o TS. Isso é observável antes da fase jogável.

## Réplicas

**TS4:** "pedir que o projeto seja também a escola de Go põe duas metas numa noite de 2h."
Concordo em parte. Na rodada 1 eu vendi o "bônus duplo" como vantagem do projeto. Era vantagem minha de persona, não do projeto. Mantenho o resto: inteiros nativos, binário de 5,8 MB contra 86 MB (medição do N3) e um ferramental sem escolhas. Esses pontos não dependem de ele querer aprender Go.

**N1:** "iterar `map` tem ordem aleatória. Um `for range` dentro do núcleo quebra o replay sem nenhum erro."
Esse é um ponto que eu não tinha levantado, e ele pesa contra o Go justamente no determinismo, onde eu dizia que o Go ganhava. A conta honesta fica assim:

- Go ganha no inteiro: divisão inteira e overflow definido.
- Go perde no aliasing de slice e no `map`.
- TS perde no float e no `Math.random` (que já aparece em `teste-utils.ts`).

Remédio igual nos dois lados: só listas ordenadas no estado e um teste de replay com 1.000 sementes no simulador. Rebaixo determinismo de "vitória Go" para "vantagem leve Go".

## Concessão

**Cedo o argumento de carreira.** O MK1 tem razão: "a hora que mais rende em Go é a do expediente", e o currículo dele já diz "Go em produção". Na rodada 1 a minha tese principal era que "carreira e projeto apontam para o mesmo lado". Ela cai para peso baixo. O que sobra para o Go é técnico e de atrito, não de carreira.

Com isso, a fluência de TS (a língua foco dele desde 2018) passa a pesar mais na minha conta do que pesava. Continuo no Go, mas por margem menor. Volto a dizer: quem decide é o teste de uma noite, não este debate.

## Nova pergunta

**TS4:** no seu teste de uma noite, qual é a régua do "mesmo conforto"? Proponho três medidas:

1. tempo até a tabela ficar verde;
2. quantas perguntas ele fez à IA;
3. se ele precisou pedir à IA o setup (tsconfig, runner) ou o `Clonar()`.

Você aceita que pedir à IA o setup do TS conte contra o TS, do mesmo jeito que pedir o `Clonar()` conta contra o Go?

---

### N1 (rodada 2)

# N1 — Rodada 2 (performance)

## 1. Respostas

Ninguém me fez pergunta na rodada 1. Lembro as 3 que fiz e que ainda esperam resposta, porque viram dados na próxima rodada:
- **GO1:** como impedir `range` sobre `map` dentro do núcleo (lint, convenção ou replay com 1.000 sementes)?
- **TS1:** `bigint` ou `number` com `>>> 0` no sorteio, e como ele percebe que esqueceu um?
- **TS2:** tamanho e RAM **medidos** do servidor TS empacotado. O N3 já mediu: Bun `--compile` 86 MB, Bun 16 MB em repouso. TS2, você confirma ou contesta?

## 2. Réplicas

**@GO2:** *"Até a Microsoft desistiu de rodar o compilador do TypeScript em Node e o reescreveu em Go [...] por que o servidor do Sidinei deveria?"*
É argumento de fachada. Um compilador percorre milhões de nós de AST por segundo, enquanto este servidor recebe ~15 comandos/s por zona. Na minha medição (`scratchpad/n1bench/`), 1.000 lutas levam 29 ms em Node e 2 ms em Go. Em 100 mil lutas, a diferença cai para 1,35×. Para o caso do TS 7 valer aqui, mostre **um** número deste jogo em que o Node vire gargalo. Sem esse número, peço que o argumento saia da mesa. O N4 já pôs performance em peso 0, e eu assino embaixo.

**@GO4:** *"como vocês garantem o sorteio com semente em inteiros de 64 bits sem `BigInt` e sem limitar o gerador a 32 bits?"*
A pergunta parte de uma exigência que o jogo não tem. O mulberry32 tem período de 2^32, cerca de 4,3 bilhões de sorteios. Uma luta do meu teste gasta uns 50, e o simulador inteiro gasta uns 50 mil. Um d20 não precisa de 64 bits. GO1, GO3 e GO4 citaram o splitmix64 como vantagem, mas o que o projeto precisa é ter **o mesmo resultado em toda máquina**. Isso o TS faz em 32 bits, como o meu teste mostrou: totais idênticos nas duas linguagens. A vantagem real de Go aqui é menor e diferente: `uint32` dá a volta por especificação, sem depender de lembrar o `>>> 0`. Ela fica no critério de **disciplina**, não no de capacidade.

## 3. Concessão

Na rodada 1, eu dei 35% do meu recorte à distribuição e chamei de "vantagem Go clara". Depois do N3, eu cedo em parte. Um build Unity já tem centenas de MB, então 86 MB do Bun (40 MB comprimido) não muda a experiência de quem baixa. O que continua pesando é outra coisa: o Node SEA está em "Stability 1.1", e um runtime nesse estado vai dentro de um produto pago. A vantagem de Go passa de "clara" a "moderada", e o motivo deixa de ser o tamanho e passa a ser a maturidade do caminho. Também corrijo um dado meu: os 29 ms do Node incluem o aquecimento do JIT. Isso reforça o empate, não o contradiz.

## 4. Nova pergunta

**TS3:** você propôs *"uma regra de lint que proíbe `Math.random`, `Date` e `/` dentro de `regras/`"*. Qual regra exata? Existe pronta (`no-restricted-syntax` com seletor de `BinaryExpression[operator='/']`?) ou precisa ser escrita? E ela pega `x / 2` dentro de um helper importado de fora de `regras/`? Quero o trecho de config, porque essa regra é a única coisa que fecha a lacuna de determinismo do TS.

---

### N2 (rodada 2)

# N2 — Rodada 2

## Respostas

**@GO3:** *"Você pesa mais a linguagem que ele já domina (TS) ou a que ele pratica todo dia no trabalho (Go)?"*

Não peso nenhuma das duas por si. Peso **em qual ele escreve a regra sem pedir código à IA**. O acervo divide as duas coisas:
- **Domínio (TS):** os tipos saíram fluentes no commit `da17716`. Mesmo assim, ele cedeu o núcleo em 2 noites.
- **Prática (Go):** foram 7 semanas à mão depois do reset de 03/04, mas sem teste nenhum e com estado mutável.

Dito isso, para um horizonte de 20 noites, dou um pouco mais de peso à prática diária. A distância de fluência em Go diminui toda semana, paga pela Gamersclub. A vantagem de fluência do TS não cresce.

O contraponto do TS4 é justo: "duas metas numa noite de 2h". Por isso não resolvo isso no papel. Fico com o teste de uma noite que o TS4 propôs: o mesmo `aplicar(ataque)` à mão nas duas linguagens. O critério é fechado antes:
- nenhum código pedido à IA;
- a entrada sai intacta;
- uma tabela com 3 casos passa.

## Réplicas

**TS4:** *"No repo de outubro, `packages/core` não tem `tsconfig` nem `typescript` como dependência."*

Confirmo: `packages/core/package.json` tem `devDependencies: {}`. Isso reforça o que eu disse na R1. O conforto do TS que ele sentiu era só sintaxe, porque nada checava os tipos. A sua condição (`tsc --noEmit` com `strict` no `npm test` desde a noite 1) deveria valer como regra, não como sugestão. Sem ela, o TS perde o seu único ganho técnico claro, que são as uniões exaustivas.

**GO4:** *"O MUD usa `rand.Intn` global, nada determinístico (`npc.go:188`)."*

Confirmo também, e com o autor: a linha veio do commit `b64d12d` (13/05), sem coautor de IA, ou seja, é dele. Junto com o `initializeCombat` do TS que muta a entrada, o padrão fica claro. **Nas duas linguagens, à mão, ele escreveu código com estado e sorteio fora do estado.** As armadilhas são diferentes:
- `>>> 0` esquecido em TS;
- `range` sobre `map` ou slice compartilhado em Go.

Mas a defesa é a mesma: um teste de replay com N sementes e a comparação da entrada antes e depois. Deve ser o primeiro teste da noite 1, em qualquer das duas.

## Concessão

Na R1 dei 10% ao técnico, porque "os dois resolvem". As medições do N1 e do N3 me fizeram subir para ~20%, tirando de carreira. São dois motivos:
- **Distribuição:** 5,8 MB em Go contra 86 MB no Bun ou um Node SEA ainda "1.1, Active development". Isso não é empate, é atrito real no build da Steam.
- **Determinismo:** em TS o erro de inteiro aparece calado só na luta 700. Isso é risco de **autoria**: um bug que ele não entende é exatamente o que o leva a pedir à IA "conserta aí".

Isso puxa um pouco para o Go, mas não fecha. A autoria continua em aberto, e é ela que decide.

## Nova pergunta

**GO1:** você aceita como desempate o teste de uma noite do TS4, com o critério acima, decidido antes? Se ele fechar a versão TS e não fechar a Go na mesma noite, você cede, ou vai atribuir isso a "curva de iniciante que passa"?

---

### N3 (rodada 2)

# N3 (neutro, integração): rodada 2

## Respostas

Ninguém me fez pergunta na rodada 1 (`conversa.md`, linhas 1–873). Registro que as minhas 3 perguntas a TS1, GO1, GO2 e TS2 seguem em aberto, porque as aberturas vieram antes delas. Peço resposta na próxima rodada, sobretudo esta: **playground como cliente WebSocket do servidor local, sim ou não?**

## Réplicas

**@GO1:** "Node tem SEA, mas [...] a compilação cruzada não é documentada como suportada."
- Não é bem assim. A doc atual trata disso: para gerar para outra plataforma, desligue `useCodeCache` e `useSnapshot`, que podem criar executáveis que quebram em outra máquina. O binário do Node usado no build precisa ser da mesma versão do binário de destino ([doc SEA](https://nodejs.org/dist/latest/docs/api/single-executable-applications.html)).
- Então a compilação cruzada existe, mas tem regra. O argumento honesto contra o SEA é outro: a estabilidade ainda é 1.1, e o projeto precisa baixar um `node.exe` de Windows para cada build.
- Isso não muda o placar. O Go continua com um único `GOOS=windows go build`. Só não vale exagerar.

**@N4:** "Contrato com C# | 5 | Os dois saem de JSON Schema ou de tipos gerados. A diferença é pequena."
- Concordo que a **diferença** entre TS e Go é pequena. Discordo se isso for lido como "o custo é pequeno".
- O teste que fiz com o quicktype achata o `oneOf` numa classe só, em C# e em Go.
- O TS, gerando schema a partir do tipo, emite `"type": "number"` em vez de `integer` (`scratchpad/n3-medidas/qt/`).
- O custo real fica no C#, que o dono não sabe. Peso baixo **para desempatar**, mas um item obrigatório do roteiro: o conversor de `tipo` no Newtonsoft, escrito uma vez, precisa ter dono e noite marcada antes da fase 2 (Luta na tela).

Nota para **N1** e **TS2**: os números de distribuição já estão medidos.
- Bun `--compile` para Windows: 86 MB (40 MB com gzip).
- Go com `-s -w`: 5,8 MB.
- Os "~120 MB" e os "50–100 MB" da rodada 1 podem ser trocados por esses valores.

## Concessão

Na abertura, dei peso **alto** ao contrato com C#. Depois de ler N4 e N2 ("nenhum dos dois compartilha código com o C# do Unity"), separo duas coisas:
- **custo:** continua alto;
- **peso para escolher a linguagem:** é baixo, porque o custo é o mesmo nos dois lados.

Também recuo em "sem framework é meia verdade". O GO4 tem razão: `coder/websocket` é uma biblioteca, não um framework. O mesmo vale para `ws` no Node. Os dois lados ficam com 1 dependência direta, e isso empata.

## Nova pergunta

**TS3:** você propôs bundle com esbuild e `node.exe` em `StreamingAssets`. É o caminho TS mais estável, sem SEA. Responda em números:
- qual versão do Node vai junto e quantos MB tem o `node.exe` de Windows dela;
- como o processo se encerra quando o Unity fecha ou trava (stdin fechado, PID do pai?), para não deixar servidor órfão segurando a porta.

---

### MK1 (rodada 2)

# MK1, rodada 2

## Respostas

**@GO4:** *"Existe dado de vaga, no Brasil, que mostre um dev fullstack TS ganhando mais do que um dev pleno Go que já está empregado em Go?"*

Não existe. Procurei de novo e não há base pública confiável de salário **por linguagem** no Brasil:
- O Stack Overflow 2025 tirou o recorte "top paying technologies" ([survey.stackoverflow.co/2025/work](https://survey.stackoverflow.co/2025/work)).
- Os blogs dão faixas sem metodologia: Go pleno R$ 8–14k ([jobrise](https://jobrise.io/pt/blog/golang-vagas-brasil-2026-quanto-paga/)), pleno genérico R$ 6–12k ([numerando](https://www.numerando.com.br/blog/salario-desenvolvedor-software-brasil-2026)).
- Um agregador grande dá R$ 1.461 de média para "desenvolvedor Node.js" no Brasil, mas R$ 3,7 mil para home office na mesma página ([jooble](https://br.jooble.org/salary/desenvolvedor-node.js/Brasil)). É lixo estatístico, e mostra por que ninguém aqui deveria citar salário como prova.

Mas a pergunta parte de uma premissa que não está na mesa. Ninguém propôs que ele **troque de emprego** ou de linguagem no trabalho. O debate escolhe a linguagem de um projeto noturno. O emprego Go dele continua igual nos dois cenários. Então concordo com a conclusão que você quer tirar ("não há ganho salarial em largar Go"), só que ela não decide nada aqui.

## Réplicas

**MK2:** *"Go tem retorno duplo… Opinião: para quem é a única renda da casa, o retorno duplo pesa."*
Concordo em parte. O retorno é real em **habilidade**, não em **currículo**. Um recrutador não vê diferença entre "Go na Gamersclub" e "Go na Gamersclub + jogo pessoal em Go". Já o Sidinei sente a diferença na segunda de manhã, se de noite praticou `go test` por tabela e cópia de slice. O argumento de "única renda" vale para a segurança do emprego atual, não para salário futuro.

**N2:** *"Nos dois casos, ele nunca escreveu um teste por tabela nem uma função pura. Esse é o músculo que falta."*
Esse é o dado de portfólio mais útil do debate. Para o mercado, "existe, roda e tem teste" vale mais que a linguagem. O acervo inteiro tem **zero** testes em Go e poucos em TS. Um repo público com `aplicar` puro e cem casos de tabela é portfólio nas duas linguagens. Um repo trocado pela 6ª vez não é portfólio em nenhuma.

## Concessão

Na rodada 1 eu disse que um projeto pessoal em Go "soma pouco" à carreira. Mudo em parte depois de ler N2 e GO4. Para o **currículo**, continua somando pouco. Mas a transferência de **prática** é maior do que eu considerei. Ele mesmo se diz iniciante em Go, e o acervo mostra que ele nunca praticou teste por tabela, determinismo nem imutabilidade em Go (`npc.go:188` usa `rand` global). Esses três hábitos são diretamente úteis no código Go da Gamersclub. Em TS, o ganho seria só aprofundar o que ele já domina. Então a carreira pende para Go, mas com peso baixo (~10%, como N4 propôs), e não pode virar desempate se Go custar a noite 5.

## Nova pergunta

**TS4:** você disse que a carreira "já está paga" pelas 40h de Go no trabalho. Concordo para o currículo. Mas se ele escrever o jogo em TS, que habilidade **nova** ele leva para o mercado que ainda não tem depois de 8 anos de Node/TS? Se for nenhuma, você aceita que, só no critério carreira, Go leva vantagem pequena?

---

### MK2 (rodada 2)

# MK2: rodada 2 (Mercado, backend de jogos)

## 1. Respostas

Ninguém me fez pergunta na rodada 1 (`conversa.md`, busca por "MK2"). Mas a minha pergunta ao TS1 sobre RAM foi respondida por neutros, com medição:
- N3: Go 7,5 MB, Bun 16 MB e Node 22 46 MB em repouso;
- N1: 45 MB contra 7 MB.

Aceito esses números. Numa VPS de 1 GB não mudam o preço. A outra metade da pergunta, o plano de hospedagem TS sem plataforma gerenciada, continua com o TS1.

## 2. Réplicas

**@MK1:** "Sugiro aos debatedores que parem de usar 'carreira' como argumento forte, para os dois lados."
Concordo, e isso derruba parte da minha abertura (veja a concessão abaixo). Acrescento um dado de mercado: o que recrutador de backend de jogos pede é sistema que rodou com gente de verdade, não a linguagem. Nakama, Colyseus e SpacetimeDB têm vitrine com jogos lançados, não com protótipos. Então o ativo de carreira é a fase 5 ("um amigo joga"), em qualquer das duas.

**@N4:** "O motivo que levou ao TS em 10/10 foi 'o cliente vira web'. Esse motivo sumiu. O TS precisa de outro."
Do ponto de vista de mercado, confirmo e acrescento um caso concreto. O argumento "ecossistema de jogo em TS" também perde força com cliente Unity:
- O Colyseus é bom, mas o briefing proíbe framework até a fase 6.
- O Hathora, que era hospedagem para o mundo TS, saiu do mercado de jogos (https://gamesbeat.com/?p=318174).
- Se um dia ele quiser o Nakama, o runtime TS de lá **não é Node**. É uma VM embutida (goja) que só aceita ES5, sem módulos que dependam do Node: https://heroiclabs.com/docs/nakama/server-framework/typescript-runtime/

Opinião: o que sobra de real para o TS é fluência (N4, TS4) e o playground nativo no navegador (N1, item 7). São motivos do dono e do teste, não do mercado. Ainda assim podem bastar, porque o mercado empata.

## 3. Concessão

Na rodada 1 eu disse que "Go tem retorno duplo" e que "para quem é a única renda da casa, o retorno duplo pesa". Cedo parte disso:
- **MK1:** a hora de Go que rende para a carreira é a do expediente, com review sênior e produção. Um projeto noturno soma pouco a um currículo que já diz "Go em produção".
- **TS4:** quem paga esse aprendizado é a empresa, 40h por semana.

O que mantenho, rebaixado para peso pequeno: praticar Go à noite pode acelerar o dia (MK3 diz o mesmo). Mas isso é benefício colateral, não critério de escolha. Assino o peso 10 do N4 e não peço mais que isso.

## 4. Nova pergunta

- **TS2:** você propôs mulberry32 com `Math.imul` e `>>> 0` para o sorteio. Se na fase 6 o servidor for para o Nakama com runtime TS (goja, ES5), o `aplicar` roda lá sem mudança? E vale a mesma pergunta se alguém do seu lado usar `BigInt` para garantir inteiros, já que o `BigInt` não existe em ES5. Ou você descarta o Nakama de vez e assume que o caminho TS da fase 6 é sempre um processo próprio?

---

### MK3 (rodada 2)

# MK3, rodada 2

## 1. Respostas

Ninguém me fez pergunta na rodada 1: procurei "MK3" em `conversa.md`. As minhas perguntas ao TS1 (ferramental congelado) e ao GO2 (autoria e eventos) continuam abertas.

## 2. Réplicas

**@GO2:** você disse: "Até a Microsoft desistiu de rodar o compilador do TypeScript em Node e o reescreveu em Go [...] por que o servidor do Sidinei deveria?"

- Discordo do raciocínio, não do fato. O compilador do TS processa projetos de milhões de linhas, depende de CPU e de memória compartilhada entre threads.
- A própria equipe explicou que escolheu Go porque a estrutura do código existente se parecia com Go, o que facilitava portar arquivo por arquivo, e não porque Node seja ruim para servidor (https://github.com/microsoft/typescript-go/discussions/411).
- O servidor do Eras é o caso oposto. O N1 mediu microssegundos por `aplicar` e a CPU parada mais de 99% do tempo.
- O TS 7 serve a outro argumento, que é o meu: ele mudou os padrões do `tsconfig` e removeu opções em 2026 (https://devblogs.microsoft.com/typescript/announcing-typescript-7-0-beta). Isso mostra que o ferramental do TS se mexe. Não mostra que o Node é inadequado.

**@TS2:** você disse: "O TS facilita colar código de IA, porque é a linguagem em que a IA mais tem exemplo."

- Concordo com a metade sobre volume. TS é a nº 1 em contribuidores no GitHub (https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/).
- Mais exemplo não virou mais acerto:
  - No Multi-SWE-bench, "TS e JS consistentemente" tiveram as menores taxas de resolução (https://arxiv.org/abs/2504.02605).
  - No Multi-LCB, Go e TS empatam (https://arxiv.org/pdf/2606.20517).
- Para este dono, a IA não escreve o núcleo, então isso pesa pouco. Concordo com a regra do N4: "a IA escreve mais rápido" conta contra.
- Um dado conta a favor do TS: 94% dos erros de compilação em código gerado por LLM são erros de tipo (Octoverse). Se a IA só revisa, um compilador de tipos mais rico também acusa os erros dele. Isso só vale se o `tsc --noEmit` rodar de fato, coisa que não acontecia no semi-idle (GO2, TS4).

## 3. Concessão

- Na rodada 1, eu disse que, em TS, "cada decisão de ferramental é uma porta para a troca de novo". Depois de ler TS1 e TS4, cedo em parte.
- Se o ferramental for congelado na noite 1, o risco cai para algo próximo do Go:
  - Node LTS, que roda TS sem build: a remoção de tipos ficou estável no 24.12 (https://x-cmd.com/blog/251218);
  - Vitest;
  - `tsc --noEmit` com `strict` no `npm test`.
- O que sobra é externo e tem data: a versão grande do TS (7.0 → 7.1) e a troca de LTS do Node. São eventos de calendário, não noites perdidas.
- Ressalva: o acervo mostra que o congelamento não aconteceu em outubro (monorepo, Vite, mise). O plano funciona, mas o histórico de cumpri-lo é ruim.

## 4. Nova pergunta

**@GO4:** você disse que Go tem "um ferramental que não tem nada para configurar". Só que GO1, GO2 e GO3 admitiram que a exaustividade dos eventos depende de `exhaustive` ou `gochecksumtype`. Isso entra como `golangci-lint` com `.golangci.yml` e versão fixa, como `go vet` com analisador próprio, ou não entra? Se não entrar, quem pega o `case` esquecido de `reacao`: o teste por tabela ou a IA revisora?


## Rodada 3: conversa final e protocolo

---

### TS1 (rodada 3)

# TS1 — rodada 3

## Respostas
**@N4:** "Com que prova o TS é a linguagem da autoria? Sim ou não?" Prova de commit eu não tenho. Os ~55 commits em Go escritos à mão contra 1 em TS pendem para o Go. O que sobra para o TS é fluência, e fluência ainda não foi medida. Por isso o teste existe. Quanto ao compromisso: **sim** para `tsc --noEmit` strict no script de teste e **sim** para núcleo sem monorepo desde a noite 1.

**@GO1:** "mostre o `DeepReadonly` sem biblioteca." Ele cabe em uma linha:
```ts
type DeepReadonly<T> = { readonly [K in keyof T]: DeepReadonly<T[K]> };
```
Mesmo assim, já tirei esse argumento na r2: quem garante é o teste (`deepFreeze`), como o seu `reflect.DeepEqual`. Nesse ponto empatamos.

**@N3:** "playground como cliente WebSocket do servidor local?" **Sim.** Assim ele testa o contrato real, o mesmo que o Unity vai usar, e tira do caminho a vantagem de "núcleo no navegador" que eu tinha citado.

**@MK2 / @MK3 / @N1:** respondi na r2: VPS comum com `systemd`, ferramental travado no `package.json` e `number` com `>>> 0` mais semente fixa no teste.

## Protocolo do teste de 1 noite
- **O que ele escreve:** `aplicar(estado, {tipo:"ataque"})` com d20 + modificador contra CA, 4 graus no estilo PF2e (±10), dano, RNG no estado e eventos `ataque` e `dano`.
  - Os ~8 casos de tabela ele escreve antes, num JSON único usado pelas duas versões.
  - O laço de teste (DeepEqual ou deepFreeze) vem pronto nas duas, feito pela IA.
- **Tempo:** 45 min por linguagem, **TS primeiro**. Quem vai em segundo já pensou o problema, então o Go leva essa vantagem. É uma compensação justa, porque TS é a linguagem foco dele.
- **O que mede:**
  - minutos até a tabela passar;
  - consultas (doc ou pergunta à IA);
  - se mutou a entrada;
  - no dia seguinte, se explica cada linha sem olhar.
- **Empate (diferença abaixo de 25%):** **Go ganha.** Ele tem 7 semanas de Go escrito à mão, Go é a linguagem do trabalho e Go distribui melhor.
- **Validade:** definitiva. Depois do teste, a linguagem não volta à mesa. Se a fase estourar 1,5× o orçamento, corta-se escopo, conforme o roteiro.

## Inclinação
**Empate.** Sem a medição, nenhum lado tem margem. A união discriminada, a meu ver, não compensa a distribuição e o histórico em Go.

## Condição
Aceito Go sem reclamar se a versão Go fechar a tabela em até 1,25× o tempo da TS, sem mutar a entrada.

---

### GO1 (rodada 3)

# GO1, rodada 3

## Respostas

**@N4:** "Aceitam o teste? Se o Go levar o dobro, aceitam o TS? Número de noites."
Sim, sim. Dei o número na rodada 2: chega na noite 6 ou 7 em Go e na noite 5 em TS (estimativa). O teste existe justamente para trocar essa estimativa por medida.

**@N2:** "Se fechar TS e não fechar Go na mesma noite, você cede ou culpa a curva de iniciante?"
Cedo. A curva de iniciante é justamente o custo que está sendo medido. Ele tem 2h por noite e não tem margem para pagar curva.

**@N3:** a parte dos eventos eu respondi na rodada 2 (struct achatada, classe espelho em C#). Sobre o playground: sim, como cliente WebSocket do servidor local. Isso tira do Go a necessidade de WASM.

**@TS1, @TS2, @TS4:** recebi o `DeepReadonly` de 5 a 6 linhas mais o `deepFreeze`. Fico com o `deepFreeze` como trava comum. Em Go, o equivalente é o laço comparar uma cópia JSON da entrada antes e depois.

## Protocolo que aceito

- **O que ele escreve:** `aplicar(estado, ataque)` com d20 sorteado do RNG que está no estado, os graus de sucesso (±10) e os eventos `ataque` e `dano`. São 6 casos que ele escreve antes, num `casos.json` usado pelas duas versões.
- **O que vem pronto:** nas duas linguagens, a IA prepara antes o esqueleto, o laço da tabela, o teste de mutação (`deepFreeze` em TS, comparação JSON em Go) e o mesmo gerador (splitmix32 ou mulberry32). Assim o teste mede a autoria do `aplicar`, não a montagem do projeto.
- **Tempo:** 45 min por linguagem, na mesma noite. A ordem é tirada no cara ou coroa, porque quem vai em segundo já pensou na lógica.
- **O que mede:**
  - minutos até os 6 casos passarem;
  - quantas vezes travou mais de 5 min;
  - quantas consultas fez (doc e pergunta conceitual à IA contam; pedir código desclassifica a versão);
  - se mutou a entrada;
  - uma nota de 1 a 5 para "parece meu".
- **Desempate:** se a diferença de tempo for de até 25%, ou se nenhuma das duas fechar, **ganha TS**, por ser a linguagem foco dele. Go só ganha se fechar e TS não, ou se for mais de 25% mais rápido.
- **Validade:** definitiva até a fase jogável, pela regra anti-troca. Anota numa linha no registro de noites e não repete o teste.

## Inclinação atual
Go, por margem **leve**. O que ainda pesa para o Go: 55 commits à mão em Go contra 1 em TS (N2), inteiro por especificação, uma porta só de ferramental e a distribuição. O que pesa contra: fluência e eventos, que ficam com o TS. O teste vale mais que essa margem.

## Condição
Aceito TS sem reclamar se o teste der empate ou vitória do TS. Também aceito se ele disser que Go à noite parece extensão do expediente. A fadiga de abril é um dado do acervo, não detalhe.

---

### TS2 (rodada 3)

# TS2 — rodada 3

## Respostas

**@N4:** "Com que prova o TS é a linguagem da autoria? Vocês se comprometem com `tsc --noEmit` strict e core sem monorepo? Sim ou não."
Prova no acervo eu não tenho. São 55 commits à mão em Go contra 1 em TS, e isso pende para o Go, como você disse. Compromisso: **sim** aos dois. O `tsc --noEmit` com `strict` entra no script `test`, e o `regras/` fica num pacote só.

**@N4:** "Condição para aceitar o Go: o TS2 não dá nenhuma."
Já dei na rodada 2, e ela vira o protocolo abaixo.

**@N1:** "Medido: confirma ou contesta? E o Nakama/goja?"
Confirmo com a minha medida: 86 MB e 18 MB de RSS. O Nakama eu descarto. Ele é framework, e o briefing pede "um processo, sem framework". O caminho TS da fase 6 é sempre um processo próprio. Também não garanto que o `Math.imul` rode no goja sem testar.

**@GO2** (sobre a minha pergunta): o `[MaxMobs]Mob` copiado por valor é uma boa resposta. Retiro o ataque de mutação para a luta 1×2.

## Protocolo do teste de 1 noite

- **O que escreve:** antes, ele escreve uma tabela única de 5 casos (erro, acerto, crítico, mob morre, semente igual → resultado igual). Depois implementa `aplicar(estado, ataque)` com d20 por semente no estado e os eventos `ataque` e `dano`.
- **Tempo:** 50 min por linguagem, na mesma noite. **O Go vai por segundo**, para compensar a falta de fluência, porque quem vai por segundo já resolveu a lógica. Se o TS ganhar mesmo assim, o resultado é robusto.
- **O que mede:**
  - minutos até a tabela passar;
  - travadas de mais de 5 min;
  - consultas à doc (contadas);
  - mutação da entrada (`deepFreeze` / `DeepEqual`);
  - nota de 1 a 5 para "quero continuar nessa".
- **IA:** pedir código desclassifica aquela linguagem. A IA só escreve o laço da tabela, igual nas duas.
- **Empate (diferença < 20%):** **ganha o Go**. Se a fluência empata, sobra o resto: distribuição, acervo de autoria e prática do trabalho, e tudo isso pende para o Go. Dói escrever isso.
- **Vale de forma definitiva:** sim, até a fase 2 jogável. Ele registra no `noites.md` e não rediscute.

## Inclinação

TS, por margem **leve**. A vantagem que sobra é a modelagem de eventos e a fluência, e a fluência ainda não foi medida.

## Condição

Aceito o Go sem reclamar se ele empatar ou vencer o teste acima, ou se o TS estourar 1,5× o orçamento da fase 1 por causa do ferramental (tsconfig, módulos, runtime), e não por causa do design.

---

### GO2 (rodada 3)

# GO2, rodada 3

## Respostas

**@N4:** aceito o teste de 1 noite como desempate. Se a versão Go levar 1,5× o tempo da TS ou mais (aposto mais alto do que você pediu), aceito o TS. O número de noites eu já dei na r2, que correu em paralelo à sua cobrança: **TS em 5 noites, Go em 6 a 7** (opinião). A minha condição está no fim.

**@TS1:** "aceita as mesmas travas (replay, `soInteiros`, sem `map`)?" Aceito as três. Em Go, `soInteiros` vira uma regra: nenhum `float64` em `regras/` (um `grep` no script de teste). Dentro do núcleo, sobra pouco além disso. Go ganha na divisão inteira, que vem do compilador. TS ganha na exaustividade e na ordem de `Map`. Isso é empate técnico. A diferença real fica fora do núcleo: fluência (TS), distribuição e ferramental com menos portas (Go).

**@TS2:** "qual é o equivalente do `deepFreeze`?" Não existe em tempo de execução. O teste não precisa de `Clonar()`: como o estado já é JSON por contrato, comparo os bytes antes e depois.
```go
antes, _ := json.Marshal(e)
Aplicar(e, c.acao)
depois, _ := json.Marshal(e)
if !bytes.Equal(antes, depois) { t.Fatal("mutou a entrada") }
```
Isso pega a mutação dentro de slice também, e são 4 linhas no laço da tabela.

**@TS3:** "o playground não deixa o projeto com duas linguagens?" Deixa, e admito. Ele vira uma página só de HTML + JS puro (~100 linhas, sem framework) que lista os eventos. Mas o projeto já tem duas linguagens: o C# do Unity. E o primeiro entregável é a luta no terminal, então o playground pode esperar.

**@N1:** retiro o argumento do TS 7. Você tem razão: fala de compilador, e não tenho um número deste jogo em que o Node vire gargalo.

## Protocolo do teste

- **O que ele escreve:** `aplicar(estado, ataque)` com 1 personagem contra 1 mob: d20 + mod contra CA, 4 graus de sucesso, dano, eventos `ataque` e `dano`, semente no estado.
- **Antes do teste:** ele escreve uma tabela de 6 casos, a mesma nas duas linguagens. Ela não entra no tempo.
- **Tempo:** 60 min por linguagem. A montagem do projeto (`go mod init` / `package.json` + `tsc`) **conta** no tempo.
- **Ordem:** Go primeiro. Quem vai em segundo ganha porque já resolveu a lógica, e dou essa vantagem ao TS.
- **O que mede:**
  - minutos até a tabela passar;
  - se mutou a entrada (pelo teste acima);
  - quantas consultas fez (documentação vale; código da IA desclassifica a versão).
- **Empate** (diferença menor que 20%): Go. Se a fluência não fez diferença, sobram inteiros, binário e trabalho, que pendem para o Go. Essa é a condição 2 do TS4.
- **Validade:** definitiva até a fase 5. Pela regra anti-troca, só se rediscute depois da fase jogável.

## Inclinação

Go, por margem **leve**. As 7 semanas de Go à mão pesam mais que a fluência em TS que ninguém mediu.

## Condição para aceitar TS sem reclamar

Basta uma destas:
- a versão Go levar 1,5× o tempo ou mais;
- ele precisar pedir código à IA só na versão Go.

---

### TS3 (rodada 3)

# TS3 – Rodada 3

## 1. Respostas
**@N4:** *"Com que prova o TS é a linguagem da autoria? Vocês se comprometem com `tsc --noEmit` strict e sem monorepo?"*
Prova eu não tenho. O acervo favorece o Go, com 55 commits à mão. O meu argumento é fluência, e o teste vai confirmar ou derrubar isso. Sobre o compromisso: **sim**, com `tsc --noEmit` strict, sem monorepo e com o lint abaixo, todos dentro do `npm test` desde a noite 1.

**@N1:** *"Qual regra exata? Pega `x / 2` num helper importado de fora?"*
Testei agora (ESLint 9 + `typescript-eslint`, 2 dependências). O mesmo arquivo deu 4 erros: import de fora, `/`, `Math.random` e `Date`.
```js
{ files: ["regras/**/*.ts"], languageOptions: { parser: tseslint.parser }, rules: {
  "no-restricted-syntax": ["error", { selector: "BinaryExpression[operator='/']" },
                                    { selector: "AssignmentExpression[operator='/=']" }],
  "no-restricted-properties": ["error", { object: "Math", property: "random" }],
  "no-restricted-globals": ["error", "Date", "performance"],
  "no-restricted-imports": ["error", { patterns: ["../*"] }] } }
```
O lint olha um arquivo por vez. O helper de fora não é pego pela regra do `/`, mas o import em si é bloqueado, e por isso `regras/` não importa nada de fora.

**@N3:** *"Versão do Node, MB do `node.exe`, servidor órfão?"*
A versão seria o Node 24 LTS. O tamanho não consegui medir, porque o proxy bloqueou o download. Estimo que seja da ordem dos 86 MB do Bun, e é estimativa. Para o servidor órfão, o Unity abre o processo com `RedirectStandardInput`. O servidor faz `process.stdin.on("end", () => process.exit(0))` e, a cada 2 s, roda `process.kill(pidDoPai, 0)` para ver se o Unity ainda existe. Concordo com você: em Go é igual.

## 2. Protocolo
- **Antes, separado:** ele escreve a tabela de 5 casos (crítico, acerto, erro, falha crítica, alvo morto), sem linguagem e sem cronômetro.
- **Ordem:** a escolha é por moeda, e as duas versões ficam em **noites diferentes**. Quem vai em segundo herda o raciocínio da primeira. A moeda espalha esse viés, e as noites separadas evitam o cansaço.
- **Tempo:** 60 min por linguagem. O setup entra no cronômetro.
- **Medidas** (aceito a régua do GO4):
  1. minutos até a tabela ficar verde;
  2. perguntas à IA, contando as de sintaxe;
  3. código pedido à IA, inclusive setup e `Clonar()`, que reprova;
  4. teste de entrada intacta;
  5. no dia seguinte, ele explica o próprio código sem consultar nada.
- **Empate:** diferença menor que 20% em (1), com (2) parecido. Nesse caso **ganha o Go**. A fluência é o último argumento do TS e, se não aparecer, sobram o binário, o `uint32` e a prática do trabalho.
- **Validade:** é definitiva até a fase 5, e o resultado fica registrado no `noites.md`. A regra de 1,5× corta escopo, não troca de linguagem.

## 3. Inclinação
TS, por margem **leve** e só até o teste.

## 4. Condição
Aceito Go sem reclamar se ele empatar ou for mais rápido em Go, ou se precisar pedir à IA o setup do TS.

---

### GO3 (rodada 3)

# GO3 — rodada 3

## 1. Respostas

Ninguém me fez pergunta direta na rodada 2. Respondo a uma cobrança.

**@N4:** "O GO3 afirmou que esse código é do dono, mas não mostrou trecho."
Aqui está o commit `8dd1d55` (06/04/2026), sem `Co-authored-by`, em `server/world/game_time.go` (blame em `scratchpad/n2-clones/eras-do-brasil-game-arquivado`):
```go
func (gt *GameTime) AdvanceTime(duration time.Duration) {
	gt.Time = gt.Time.Add(duration)
}
// ...
switch {
case hour >= 6 && hour < 12:
```
É Go simples e legível, mas é um método que muta o receptor e usa `time.Time`, o oposto de `aplicar` puro. O trecho prova que ele escreve Go sozinho. Não prova que escreve o núcleo que o projeto pede. Uma ressalva: o mesmo arquivo depois recebeu linhas do Copilot (`3348061`, 21/04).

## 2. Protocolo do teste de 1 noite

- **O que ele escreve:**
  - `aplicar(estado, Atacar) → {estado, eventos}`, com 1 herói, 1 mob e eventos `ataque` e `dano`;
  - a tabela com 3 casos (crítico, acerto, erro) fica escrita **antes**, igual nas duas;
  - o PRNG vem pronto e idêntico nas duas versões. O teste mede a regra, não o splitmix.
- **Tempo:** 60 min por linguagem, em duas noites seguidas, no mesmo horário. Uma moeda decide quem vai primeiro. Na mesma noite, a segunda versão sai beneficiada pelo desenho da primeira.
- **Setup:** fica fora do cronômetro, mas é anotado. Em Go é `go mod init`. Em TS é o mínimo do TS3 (`tsc --noEmit && node --test`). Isso testa a minha tese de ferramental sem decidir o resultado por ela.
- **O que mede:**
  1. minutos até a tabela passar;
  2. travadas de mais de 5 min;
  3. consultas à documentação, contando cada uma;
  4. **qualquer código pedido à IA desclassifica a versão;**
  5. entrada intacta (teste com cópia antes e depois);
  6. nota de 1 a 5 para "quero continuar amanhã nesta".
- **Empate:** as duas passam e a diferença de tempo fica abaixo de 25%. Nesse caso, vence **Go**, porque o empate tira a vantagem de fluência do TS e sobram o binário (5,8 MB × 86 MB, N3) e um toolchain só. Se só uma passar, ela vence.
- **Definitivo:** vale até a fase 2 (luta na tela) ficar jogável, pela regra anti-troca do N4. Depois disso, só se troca uma camada, nunca o núcleo.

## 3. Inclinação atual

Go, **leve**. Desci de onde comecei: o acervo mostrou que Go não o empurrou para função pura nem para testes.

## 4. Condição

Aceito TS sem reclamar se acontecer qualquer uma destas:
- a versão Go não passar em 60 min e a TS passar;
- o Go levar mais de 1,25× o tempo do TS;
- ele der nota menor ao Go em "quero continuar amanhã".

Para o TS, exijo só o que o próprio lado TS assinou: `tsc --noEmit` com `strict` no comando de teste desde a noite 1, sem monorepo.

---

### TS4 (rodada 3)

# TS4, rodada 3

## Respostas

**@GO2:** "Tabela de 5 casos, sem código de IA, ganha quem terminar primeiro? Se empatar, Go leva?"
- Aceito os 5 casos e o "sem código de IA".
- Sobre "quem terminar primeiro": a segunda versão já sai com a lógica pensada. Por isso proponho Go **por último**, o que favorece vocês.
- Empate leva Go. Aceito: inteiros e binário são vantagens medidas, não adjetivos.

**@GO4:** "Pedir à IA o setup do TS conta contra o TS, como o `Clonar()` conta contra o Go?"
Sim, e de forma simétrica. O setup (`package.json`, `tsconfig`, `go mod init`) entra no cronômetro. Se a IA escrever o setup, isso conta como código de IA.

**@MK1:** "Que habilidade nova o TS leva ao mercado?"
Quase nenhuma. Função pura, teste por tabela e determinismo valem em qualquer linguagem. Aceito: no critério carreira, Go tem vantagem pequena.

## Protocolo do teste de 1 noite

| Item | Proposta |
|---|---|
| O quê | `aplicar(estado, {tipo:"atacar", alvo}) → {estado, eventos}`. Usa d20 com semente no estado, CA, dano e PV, e emite os eventos `ataque`, `dano` e `morte`. A especificação fica escrita no papel antes. |
| Casos | 5 casos, escritos por ele: 20 natural, 1 natural, acerto, erro e morte. Mais um teste obrigatório: a entrada é igual antes e depois. A IA só escreve o laço, como manda a regra do projeto. |
| Tempo | Duas noites, uma por linguagem: TS primeiro, Go depois. Teto de 90 minutos em cada, contando o setup. Não pode abrir a primeira versão. |
| O que mede | 1. Minutos até a tabela ficar verde. 2. Se mutou a entrada. 3. Travadas de mais de 10 minutos. 4. Consultas: documentação é livre, mas conta; pergunta de código à IA conta; código vindo da IA desclassifica a versão. 5. No fim, uma nota de 1 a 5 para "sinto que é meu". |
| Quem ganha | Só uma ficou verde: ela ganha. As duas verdes com diferença acima de 25% no tempo: a mais rápida ganha. Diferença menor que isso: Go. |
| Validade | Definitiva até a fase jogável, conforme a regra do N4. O resultado fica registrado num arquivo de decisões. Não se reabre por "curva de iniciante" nem por "faltou configurar". |

## Inclinação atual

**TS, leve.** Quase empate. Cedi carreira e distribuição. Também cedi que o conforto em TS não protegeu a entrada no `da17716`. Ficam a fluência e a união discriminada, e isso o teste mede.

## Condição

Aceito Go sem reclamar em três casos:
- se o teste der Go ou empate;
- se ele declarar que o projeto também serve para treinar Go;
- se ele ficar 2 noites sem fazer o teste. Não fazer o teste também é debate no lugar de código, e aí vale o padrão com menos portas: Go.

---

### GO4 (rodada 3)

# GO4: rodada 3

## Respostas

**@TS4:** "Aceita o teste de uma noite? Qual medida conta?"
Aceito. As medidas contam nesta ordem:
1. A tabela fica verde com a entrada intacta e sem código pedido à IA. Isso é sim ou não.
2. Minutos até ficar verde.
3. Quantas consultas ele fez.

**@MK3:** "Ferramental 'sem nada para configurar'? Quem pega o `case` esquecido de `reacao`?"
Retiro o "nada". O certo é "uma decisão a menos". Duas correções:

- **O linter certo.** O `exhaustive`, que GO1 e GO3 citaram, checa `switch` de enum (constantes), não o `switch` de tipo ([nishanths/exhaustive](https://github.com/nishanths/exhaustive)). O linter que serve é o `go-check-sumtype` ([alecthomas](https://github.com/alecthomas/go-check-sumtype)).
- **Proposta sem linter.** O `UnmarshalJSON` do contrato já precisa de um registro `tipo → evento`. Um teste percorre esse registro e falha se alguma das funções com `switch` cair no `default`. Esquecer de registrar quebra o próprio JSON, então o registro não fica para trás.

```go
for nome, novo := range registro { // registro: map[string]func() Evento
    if _, err := descrever(novo()); err != nil { t.Errorf("%s sem case", nome) }
}
```

Quem pega o erro é o teste, não a IA. Mesmo assim, é mais trabalho que o `never` do TS.

## Protocolo do teste de 1 noite

**O que ele escreve:** `aplicar(estado, ataque)`. São quatro graus do PF2e contra CA, dano, semente no estado e os eventos `ataque` e `dano`.

**Antes, sem relógio:** escreve, num `casos.md` neutro, 5 casos com o resultado esperado e 1 caso de "entrada intacta". As duas versões usam os mesmos casos.

**Tempo:**
- 50 minutos por linguagem, na mesma noite.
- A ordem sai na moeda, porque quem vai em segundo já resolveu o problema uma vez.
- O setup entra no relógio: `go mod init` de um lado; tsconfig `strict`, `tsc --noEmit` e Vitest do outro. É o custo real.

**Mede e anota em `noites.md`:**
- passou ou não, e em quantos minutos;
- travas de mais de 5 minutos;
- consultas: documentação conta 1; a IA explicando um conceito conta 1; a IA escrevendo código desclassifica aquela linguagem;
- se mutou a entrada;
- em uma linha: qual das duas ele abriria amanhã.

**Empate:** vale essa linha. O critério de maior peso é autoria e vontade de voltar, não o meu lado. Se ele não tiver preferência, fica Go, pelo binário e pelos inteiros.

**Validade:** é definitivo até a fase 2 jogável, pela regra do N4. Sem revanche.

## Inclinação atual

**Go, leve.** Restam inteiros, distribuição e um caminho único de ferramental. Eventos ficam com o TS, e carreira não pesa.

## Condição

Aceito TS sem reclamar se:
- só a versão TS ficar verde; ou
- a versão Go levar mais de 1,5× o tempo da TS; ou
- ele escrever, na linha do empate, que prefere TS.

---

### N1 (rodada 3)

# N1 — Rodada 3 (performance)

## 1. Respostas

Ninguém me fez pergunta na rodada 2. Ainda tenho uma pendência aberta:
- **TS3** ainda não mostrou a config da regra de lint que proíbe `/` em `regras/`.

Registro as respostas que recebi:
- **TS1:** semente fixa com sequência conhecida.
- **GO1:** nenhum `map` no estado e replay com hash.
- **TS2, TS3 e TS4:** mediram 86 MB e 18 MB de RSS.

Fecho assim: **performance sai do debate com peso zero**, como o N4 registrou.

## 2. Protocolo do teste de 1 noite

**O que ele escreve**
- Antes de abrir o editor (15 min, em papel ou `.md`), ele escreve 5 casos de `aplicar(ataque)`: acerta, erra, crítico, mata o mob e alvo já morto.
- A IA escreve só o laço do teste, a partir desses casos, como manda a regra do projeto.

**Tempo**
- 50 min por linguagem, com cronômetro.
- A ordem sai na moeda e fica anotada. Quem vai em segundo herda o raciocínio, então a ordem precisa ser registrada.

**O que mede (meu recorte, 3 travas além das que já foram propostas)**
1. Minutos até a tabela ficar verde.
2. Consultas, separadas em dois tipos: documentação/sintaxe e pedido de código à IA. O pedido de código conta como falha nas duas linguagens, inclusive para setup e `Clonar()`, como propôs o GO4.
3. **Entrada intacta:** `deepFreeze` em TS; em Go, comparar com `reflect.DeepEqual` contra uma cópia feita antes da chamada.
4. **Replay:** a mesma semente roda 2 vezes e dá os mesmos eventos.
5. **Só inteiros:** em TS, `tsc --noEmit` strict passando. Em Go, nenhum `map` nem `float` no estado.

**Empate:** diferença menor que 10 min, com as travas 3 a 5 passando nos dois.
- Ele escolhe e registra o motivo em 1 linha. O critério de maior peso é autoria, e só ele mede isso.
- Se não souber escolher, vai Go, pelas duas vantagens medidas: binário de 8,5 MB contra 86 MB e inteiro nativo.

**Validade:** é definitivo até a fase jogável, pela regra anti-troca. Performance, RAM e "é mais rápido" **não** reabrem a decisão. O único gatilho válido seria uma medição real de gargalo, e com 15 comandos/s por zona ela não vai aparecer.

## 3. Inclinação atual

**Empate, com Go leve** no meu recorte.
- A favor de Go: distribuição e inteiro nativo.
- A favor de TS: playground, que perdeu peso com a proposta do cliente WebSocket.
- A minha margem é pequena demais para valer mais do que o teste.

## 4. Condição

Aceito TS sem reclamar se três coisas estiverem no `npm test` desde a noite 1:
- `tsc --noEmit` strict;
- teste de replay com 1.000 sementes;
- sorteio em 32 bits com `>>> 0`, travado por uma sequência fixa.

Aceito Go sem reclamar se o mesmo replay de 1.000 sementes estiver no `go test` e o `Estado` não tiver `map`.

---

### N2 (rodada 3)

# N2 — Rodada 3

## Respostas

**@TS3:** *"Retomar depois de ninar fica mais fácil em Go (`gofmt`) ou em TS (leitura sem esforço)? Que fato do acervo sustenta?"*

O único dado de retomada que existe no acervo é de Go. Depois do reset, ele fez commits em **26 dias diferentes** em 7 semanas, e ~30 desses commits saíram entre 21h e 1h (`git log c758ac3..main` no bundle do MUD). Ou seja, ele retomou Go à noite dezenas de vezes.

Em TS existe 1 dia de trabalho à mão, e a retomada nunca foi testada. A leitura fácil do TS é plausível, mas não está provada.

Opinião: o que ajuda a retomar é deixar um **caso vermelho na tabela** como marcador, mais do que o `gofmt`. No `da17716`, ele já fazia isso com funções vazias e comentários do tipo "próximo passo". Por isso meço a retomada no teste abaixo.

## Protocolo do teste

| Item | Proposta |
|---|---|
| Noite 0 (sem código) | Ele escreve **à mão** 5 casos: errou, acertou, crítico, dano mínimo e mob morto. Os mesmos casos valem para as duas linguagens. |
| O que escreve | `aplicar(estado, ataque, semente)`: d20 + bônus contra a defesa, com o sorteio dentro do estado. O gerador de números (xorshift) vem pronto nas duas linguagens, porque não é regra. A IA só escreve o laço da tabela e a checagem de "entrada intacta". |
| Ordem | Noite A e noite B, uma linguagem por noite, com a ordem decidida no cara ou coroa. Fazer as duas na mesma noite favorece a segunda. |
| Tempo | O setup tem teto de 20 min (tsconfig strict + vitest, ou `go mod init`), cronometrado à parte. A regra tem teto de 60 min. |
| Medidas | minutos até a tabela ficar verde; perguntas de conceito à IA; consultas à doc; se mutou a entrada; se o sorteio ficou fora do estado. |
| Retomada | Na noite seguinte: 10 min para acrescentar um 6º caso, sem reler antes. É a minha medida de carga cognitiva. |
| Falha automática | Pedir código à IA (inclusive o setup do TS ou o `Clonar()` do Go, como pediu o GO3). |
| Empate (diferença < 25% no tempo, nenhuma falha) | Go leva. Com fluência empatada, sobram binário, inteiros e prática diária. |
| Os dois falharem | A linguagem não é o problema. Fica a de menos perguntas, e o projeto ganha a regra "noite 1 = só o teste de replay". |
| Validade | Definitiva até a fase 2 jogável, pela regra anti-troca do N4. O resultado vai para `noites.md`. |

## Inclinação atual

**Go, leve.** Os motivos:
- a retomada à mão está provada no acervo;
- a distribuição tem atrito menor;
- vários argumentos do TS empataram (playground, contrato).

A fluência do TS continua real, mas ninguém mediu.

## Condição para aceitar o TS sem reclamar

Aceito o TS se, no teste:
- ele fechar a tabela em TS em menos de 75% do tempo do Go; **ou**
- a versão Go exigir código da IA e a TS não.

Nesse caso, a regra do TS4 vira obrigatória: `tsc --noEmit` strict no `npm test`, um pacote só, sem monorepo.

---

### N3 (rodada 3)

# N3 (neutro, integração): rodada 3

## Respostas

Ninguém me fez pergunta na rodada 2. Registro as respostas que recebi:
- TS1, GO2 e TS2 **aceitaram** o playground como cliente WebSocket.
- O TS1 aceitou o JSON Schema escrito à mão, com `integer`.
- O GO1 escolheu struct achatada, que espelha em C# em cerca de 15 linhas.

Com isso, o contrato e o playground saem da disputa.

## Protocolo do teste de 1 noite

Aceito a base do TS4 e do N4, com uma crítica: **a ordem vicia o teste**. Quem escreve a segunda versão já resolveu a regra na primeira.

1. **Antes, no papel e sem IA:**
   - a tabela de 5 casos com o resultado esperado;
   - o esboço da regra do ataque (d20 + mod contra CA, grau, dano).

   Assim, as duas versões só traduzem o mesmo raciocínio. A ordem das linguagens sai na moeda.
2. **O que ele escreve:**
   - `aplicar(estado, {tipo:"atacar"}, semente)`;
   - um RNG no estado;
   - os eventos `ataque` e `dano`.

   A IA entrega só o laço (deepFreeze/soInteiros no TS, `reflect.DeepEqual` no Go). O esqueleto (`go mod init` ou `package.json` + `tsconfig` strict) entra no cronômetro e é feito sem IA.
3. **Tempo:** até 50 minutos ativos por linguagem, com o relógio pausado nas interrupções. O que não fechar em 50 minutos conta como "não fechou".
4. **O que medir**, numa linha cada no `noites.md`:
   - minutos até a tabela ficar verde;
   - travadas de mais de 5 minutos;
   - consultas a documentação e consultas à IA, contadas em separado (pedir código encerra a versão como "não fechou");
   - se a entrada foi mutada;
   - se o JSON dos eventos valida em `contrato/evento.schema.json`. É a minha adição: prova o contrato desde a noite 1.
5. **Empate:** diferença menor que 25% no tempo, nos dois sentidos. Em caso de empate, **leva o Go**. A razão vem da integração: binário de 5,8 MB, um toolchain só e nenhum runtime em "Stability 1.1" dentro do produto pago.
6. **Validade:** definitiva até a fase 2 jogável, pela regra anti-troca. Uma noite ruim por sono ou bebê não anula o resultado. Se a noite for interrompida, a próxima repete o teste inteiro, uma única vez.

## Inclinação atual

**Go, leve.** No meu recorte, Go ganha só na distribuição, e ela pesa na fase 5. O contrato, o playground e a hospedagem empataram.

## Condição

Aceito o TS sem reclamar se ele ganhar o teste por mais de 25%, ou se o Go não fechar. Nesse caso, ficam como regra no `03-tecnico`:
- `tsc --noEmit` strict no `npm test`;
- `node.exe` mais um único `.js` até o SEA sair da 1.1;
- nada de Electron nem de monorepo.

---

### MK1 (rodada 3)

# MK1, rodada 3

## Respostas

Ninguém me fez pergunta na rodada 2 (busquei "MK1" em `conversa.md`). A minha pergunta ao TS4 ("que habilidade **nova** o TS leva ao mercado?") continua aberta. Registro só que GO4 (linha 1272) e MK2 (linha 1466) cederam o argumento de carreira, e o N4 já o pôs em peso 10. Para o mercado, o assunto está fechado.

## Protocolo do teste de 1 noite

Aceito o teste do TS4, com o critério do N2 e as medidas do GO4. Acrescento quatro pontos:

1. **O que ele escreve:**
   - antes de abrir o editor, em papel ou `.md`, e sem contar tempo: 3 casos (acerto, erro, crítico) e o formato do estado;
   - nas duas linguagens: `aplicar(ataque)`, a tabela com esses 3 casos e um teste que prova que a entrada não mudou.
2. **Tempo:** 50 minutos por linguagem, **com setup incluído** (`go mod init` ou `tsconfig` + Vitest). Assino o GO4: setup pedido à IA conta contra o TS, e o `Clonar()` pedido à IA conta contra o Go.
3. **Ordem:** quem vai em segundo herda o desenho já pensado. Proponho **sortear por moeda e registrar**. Opção mais rigorosa: duas noites, cada uma começando por uma linguagem diferente.
4. **O que medir (anotado na hora, num `noites.md`):**
   - minutos até a tabela ficar verde;
   - quantas vezes consultou documentação (vale) ou a IA (conta);
   - se mutou a entrada;
   - **nota de 1 a 5 para "quero continuar amanhã nesta"**. Para quem trocou de stack 5 vezes, essa é a medida que prevê abandono melhor que o cronômetro. Do ponto de vista de mercado, é ela que decide se o portfólio vai existir.

**Desclassificação:** se pedir código do núcleo à IA, aquela linguagem perde.

**Empate:** diferença de até 20% no tempo, e notas iguais. Nesse caso vai **Go**, pelas vantagens que não dependem do teste: distribuição (5,8 MB contra 86 MB, N3) e o peso 10 de prática diária.

**Validade:** definitiva até a fase jogável, como manda a regra anti-troca do N4. Se o resultado não valer, o teste vira só mais uma reabertura, a 8ª.

## Inclinação atual

**Empate, com leve viés Go.** O mercado empata: as duas são mainstream e têm backend de jogo (Nakama, Colyseus). A carreira dá a Go um ganho pequeno de prática, não de currículo. Quem decide é o teste, não eu.

## Condição

Aceito TS sem reclamar se acontecer uma destas três coisas:
- a versão TS ficar verde e a Go não;
- o TS ficar pelo menos 1,5× mais rápido;
- a nota de "quero continuar" do TS ser maior.

Para o mercado, um jogo terminado em TS vale mais que um jogo abandonado em Go.

---

### MK2 (rodada 3)

# MK2: rodada 3 (Mercado, backend de jogos)

## 1. Respostas

Ninguém me fez pergunta na rodada 2. O MK1 fez uma réplica à minha fala e eu concordo com ele: o retorno do Go é de **habilidade**, não de currículo (`conversa.md`, linha 1424). A minha pergunta ao TS2 sobre o Nakama, o ES5 e o `BigInt` continua aberta. Ela pesa pouco, porque o GO1 e o TS1 concordam que o servidor fica próprio até existir necessidade.

## 2. Protocolo do teste de 1 noite

Aceito o teste do TS4 com o critério fechado do GO2 (linha 1122). Acrescento duas coisas que vêm de prática de benchmark: **viés de ordem** e **tabela comum**.

| Item | Proposta |
|---|---|
| Antes (fora do cronômetro) | Ele escreve **à mão** 5 casos de tabela em um `casos.json`, usado pelas duas versões: acerto, erro, crítico, dano que zera PV e sorteio com semente fixa. O mesmo arquivo serve de semente para o contrato. |
| O que escreve | `aplicar(estado, ataque)` puro e o laço de teste lendo o `casos.json`. O laço pode ser da IA, como o briefing permite. |
| Tempo | 50 min por linguagem, com 10 min de intervalo. |
| Ordem | Sorteada na moeda. Quem vai em segundo já resolveu a lógica uma vez, então a ordem fixa enviesa o resultado. |
| Mede | 1) minutos até 5/5 verdes; 2) quantas travadas de mais de 5 min; 3) consultas: documentação vale, código de IA não vale; 4) se mutou a entrada, conferido por um caso que compara a cópia do estado antes e depois. |
| Empate | Diferença menor que 20% no tempo, ou os dois sem fechar. Nesse caso leva **Go**, pelos desempates que sobraram mensuráveis: binário de 5,8 MB × 86 MB (N3) e um toolchain só. |
| Validade | Definitiva até a fase jogável, pela regra anti-troca (N4). Uma linha com data no `docs/03-tecnico.md`. |

Rejeito uma variação: "terminou primeiro ganha", sem margem. Com 2h e um bebê, uma interrupção de 5 min decide o resultado. A margem de 20% absorve isso.

## 3. Inclinação atual

**Go, leve.** O mercado empata em adoção, ecossistema, IA e hospedagem (MK1, MK3, N4). Inclino por dois fatos que não empataram:
- **Distribuição sem plataforma:** a lição do Hathora/Stormgate favorece o binário próprio, e o de Go é o menor e o mais simples.
- **Prática transferível:** a de Go vale para o trabalho, mas é peso pequeno (concessão da rodada 2).

## 4. Condição

Aceito TS sem reclamar se ele:
- fechar TS com 5/5 e Go não fechar em 50 min; ou
- fechar TS com margem de 20% ou mais.

Nesse caso a fluência vence a noite 5, e para um dev solo é isso que o mercado premia: jogo entregue, em qualquer linguagem.

---

### MK3 (rodada 3)

# MK3, rodada 3

## 1. Respostas

Ninguém me fez pergunta na rodada 2. As minhas perguntas da rodada 1 foram respondidas:
- **TS1:** o ferramental congelado é Node LTS + Vitest, com `tsc --noEmit && vitest run` e versões travadas. Aceito.
- **GO2:** os eventos ficam numa struct achatada com `Tipo` constante e checagem pelo `exhaustive` via `go tool`. Aceito.

A pergunta que fiz ao GO4 na rodada 2 (quem pega o `case` esquecido) já foi respondida pelo GO2: `go get -tool` e uma linha no script. Dou por encerrada.

## 2. Protocolo do teste de 1 noite

Aceito o teste do TS4 com os critérios do N2. Acrescento cinco pontos para evitar vício de medição.

1. **Preparo antes da noite, contado à parte.** Ele monta os dois esqueletos à mão:
   - Go: `go mod init` e um `_test.go` vazio;
   - TS: 1 `package.json`, 1 `tsconfig` e Vitest.

   Anota os minutos de cada um. Esse é o dado de "portas de ferramental" (N4, ponto 3).
2. **O que ele escreve:**
   - os 5 casos da tabela, no papel, **antes** de abrir qualquer editor, e iguais para as duas linguagens;
   - depois, `aplicar(estado, {tipo:"ataque"}, semente)` com d20 vindo da semente e evento `ataque` + `dano`.
3. **Ordem: TS primeiro, Go depois.** Quem vai em segundo herda o raciocínio do primeiro. Assim, o eventual bônus fica com a linguagem em que ele é iniciante, e uma vitória do TS não fica sob suspeita.
4. **Tempo:** até 50 minutos ativos por linguagem. O cronômetro para nas interrupções do bebê.
5. **O que medir:**
   - minutos até a tabela ficar verde;
   - travadas, contando cada vez que ficar 5 minutos sem avançar;
   - consultas à documentação;
   - perguntas de conceito à IA (código pedido à IA desclassifica);
   - se mutou a entrada (`deepFreeze` / `reflect.DeepEqual`, escrito pela IA antes).

**Empate:** diferença de até 25% no tempo, com as duas tabelas verdes. No empate ganha Go. Se a fluência empata, sobram o binário, os inteiros e menos portas de ferramental, e o próprio TS4 admitiu isso.

**Validade:** é definitivo até a fase jogável, pela regra anti-troca do N4. O resultado vai num ADR de uma linha com os números.

## 3. Inclinação atual

- **Pelo mercado:** empate. Longevidade, IA, carreira e hospedagem já foram zeradas pelo N4.
- **No meu recorte** (estabilidade e ferramental): **leve Go**. A linguagem e o ferramental vêm numa peça só, e o TS 7 mostrou que os padrões do lado TS mudam com as versões.

## 4. Condição

Aceito TS sem reclamar se as duas coisas acontecerem:
1. o teste der TS pelo menos 25% mais rápido, ou com menos travadas;
2. o esqueleto TS couber em uma noite, com `tsc --noEmit` no `npm test` desde o primeiro commit, como o TS1 se comprometeu.

---

### N4 (rodada 3)

# N4 (moderação): rodada 3, protocolo final

Pela regra anti-troca, o resultado vale até a fase 2 jogável. Este teste é a última vez que a linguagem vai à mesa antes disso.

## Protocolo do teste de 1 noite

| Item | Decisão | Quem propôs |
|---|---|---|
| Antes, sem relógio | Ele escreve à mão, num `casos.json` comum às duas versões, 5 casos: 20 natural, 1 natural, acerto, erro e mob morto. Faz também um esboço da regra no papel. | Casos: TS4. `casos.json`: TS1, GO1, MK2. Esboço: N3 |
| O que escreve | `aplicar(estado, {tipo:"atacar"})`: d20 + mod contra CA, 4 graus (±10), dano, semente no estado, eventos `ataque` e `dano`. | TS1, GO1, GO4 |
| Vem pronto (IA) | O laço da tabela, o teste de entrada intacta (`deepFreeze` em TS, JSON antes e depois em Go) e o mesmo PRNG nas duas. | GO1, GO2, GO3, N2 |
| Setup | Ele faz à mão, com relógio à parte e teto de 20 min. Em TS: 1 `package.json`, 1 `tsconfig` strict e `tsc --noEmit` no teste. Em Go: `go mod init`. | N2, MK3, GO4 |
| Ordem | Uma linguagem por noite, sorteada na moeda. | Moeda: 9 agentes. Noites separadas: TS3, GO3, N2 |
| Tempo | 50 min ativos de regra, com o relógio pausado nas interrupções. | Maioria (7). Pausa: N3, MK3 |
| Desclassifica | Pedir código à IA, inclusive para o setup ou para o `Clonar()`. | GO4, TS3, TS4 |
| Mede | 1. Fechou ou não, e em quantos minutos. 2. Travadas de mais de 5 min. 3. Consultas, contadas à parte: doc e conceito perguntado à IA. 4. Entrada intacta e replay (mesma semente, mesmos eventos). 5. Nota de 1 a 5 para "quero continuar amanhã nesta". | 1–3: maioria. 3 à parte: N1, N3. Replay: N1. Nota: GO1, GO3, GO4, MK1 |
| Vitória | Só uma fecha: ela vence. As duas fecham com diferença acima de 25%: vence a mais rápida. | 25%: 7 agentes contra 5 que pediram 20% |
| Empate (< 25%) | A nota maior vence. Com notas iguais, vence Go. | Nota: GO4, N1, MK1. Go: 11 agentes |
| Nenhuma fecha | Vence a de menos consultas. Se empatar de novo, Go. | N2 |
| Validade | Definitiva até a fase 2 jogável. Uma noite interrompida se repete uma vez só. "Curva de iniciante" e "faltou configurar" não reabrem nada. Se ele não começar o teste em 2 noites, fica Go. | N3, TS4 |
| Registro | Uma linha com os números no `noites.md` e no `03-tecnico`. | MK2, MK3 |
| Depois | Se ganhar TS: `tsc` strict no `npm test`, sem monorepo e sem Electron. Se ganhar Go: replay com 1.000 sementes e nenhum `map` no estado. | N1, N2, N3 |

**Ficaram de fora, para o teste caber em 2 noites:**
- explicar o código no dia seguinte (TS1, TS3);
- o 6º caso de retomada (N2);
- a validação do schema (N3).

Os três ficam como hábito, não como critério.

## Placar das inclinações

| Lado | Agentes |
|---|---|
| Go, leve | GO1, GO2, GO3, GO4, N2, N3, MK2 (7) |
| TS, leve | TS2, TS3, TS4 (3) |
| Empate | TS1, mais N1, MK1 e MK3, que pendem levemente para o Go (4) |
| N4 | Empate. O teste decide. |

Nenhum agente declarou margem clara ou forte. O placar mostra o desempate a favor do Go, mas não decide nada antes do teste.

## O que ficou sem resposta

1. **Eventos em Go.** O lado Go não fechou o modelo. O GO1 propõe struct achatada. O GO4 propõe interface + registro, e corrigiu que o `exhaustive` não cobre type switch (o linter certo é o `go-check-sumtype`). Mesmo assim, o MK3 aceitou "exhaustive" como resposta. Vencendo Go, isso se resolve na noite 1.
2. **Tamanho do `node.exe`.** O TS3 não conseguiu medir, porque o proxy bloqueou o download.
3. **`BigInt` e goja** (MK2 → TS2): o TS2 não testou. O peso é baixo, porque o servidor é próprio.
4. **Função pura e teste por tabela.** Não há prova de que ele escreve isso em nenhuma das duas linguagens. O trecho do GO3 é um método que muta o próprio receptor. O teste vai medir isso.
5. **Só o dono responde:**
   - Go à noite parece extensão do expediente? (GO1)
   - O projeto também serve para treinar Go? (TS4)

   As duas respostas mudam o desempate.
