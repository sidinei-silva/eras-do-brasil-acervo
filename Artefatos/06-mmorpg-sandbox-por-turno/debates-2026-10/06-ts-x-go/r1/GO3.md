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
