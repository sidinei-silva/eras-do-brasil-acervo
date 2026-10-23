# MK1, rodada 3

## Respostas

Ninguém me fez pergunta na rodada 2 (busquei "MK1" em `conversa.md`). A minha pergunta ao TS4 ("que habilidade **nova** o TS leva ao mercado?") continua aberta. Registro só que GO4 (linha 1272) e MK2 (linha 1466) cederam o argumento de carreira, e o N4 já o pôs em peso 10. Para o mercado, o assunto está fechado.

## Protocolo do teste de 1 noite

Aceito o teste do TS4, com o critério do N2 e as medidas do GO4. Acrescento quatro pontos:

1. **O que ele escreve:**
   - antes de abrir o editor, em papel ou `.md`, e sem contar tempo: 3 casos (acerto, erro, crítico) e o formato do estado;
   - nas duas linguagens: `aplicar(ataque)`, a tabela com esses 3 casos e um teste que prova que a entrada não mudou.
2. **Tempo:** 50 minutos por linguagem, **com setup incluído** (`go mod init` ou `tsconfig` + Vitest). Assino o GO4: setup pedido à IA conta contra o TS, e o `Clonar()` pedido à IA conta contra o Go.
3. **Ordem:** quem vai em segundo herda o desenho já pensado. Proponho **sortear por moeda e registrar**. Opção mais rigorosa: duas noites, cada uma começando por uma linguagem diferente.
4. **O que medir (anotado na hora, num `noites.md`):**
   - minutos até a tabela ficar verde;
   - quantas vezes consultou documentação (vale) ou a IA (conta);
   - se mutou a entrada;
   - **nota de 1 a 5 para "quero continuar amanhã nesta"**. Para quem trocou de stack 5 vezes, essa é a medida que prevê abandono melhor que o cronômetro. Do ponto de vista de mercado, é ela que decide se o portfólio vai existir.

**Desclassificação:** se pedir código do núcleo à IA, aquela linguagem perde.

**Empate:** diferença de até 20% no tempo, e notas iguais. Nesse caso vai **Go**, pelas vantagens que não dependem do teste: distribuição (5,8 MB contra 86 MB, N3) e o peso 10 de prática diária.

**Validade:** definitiva até a fase jogável, como manda a regra anti-troca do N4. Se o resultado não valer, o teste vira só mais uma reabertura, a 8ª.

## Inclinação atual

**Empate, com leve viés Go.** O mercado empata: as duas são mainstream e têm backend de jogo (Nakama, Colyseus). A carreira dá a Go um ganho pequeno de prática, não de currículo. Quem decide é o teste, não eu.

## Condição

Aceito TS sem reclamar se acontecer uma destas três coisas:
- a versão TS ficar verde e a Go não;
- o TS ficar pelo menos 1,5× mais rápido;
- a nota de "quero continuar" do TS ser maior.

Para o mercado, um jogo terminado em TS vale mais que um jogo abandonado em Go.
