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
