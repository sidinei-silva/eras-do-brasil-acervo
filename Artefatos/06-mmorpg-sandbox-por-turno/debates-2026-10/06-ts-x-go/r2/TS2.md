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
