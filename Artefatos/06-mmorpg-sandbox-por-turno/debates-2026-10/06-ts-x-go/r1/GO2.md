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
