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
