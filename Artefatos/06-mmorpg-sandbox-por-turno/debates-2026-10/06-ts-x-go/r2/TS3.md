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
