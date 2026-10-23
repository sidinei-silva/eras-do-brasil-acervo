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
