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
