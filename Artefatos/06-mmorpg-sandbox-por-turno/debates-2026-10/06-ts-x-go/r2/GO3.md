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
