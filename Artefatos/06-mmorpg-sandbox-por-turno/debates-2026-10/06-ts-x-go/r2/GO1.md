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
