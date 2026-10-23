# Resultado: TypeScript × Go

**Veredito:** empate técnico. Quem decide é um teste de 1 noite por linguagem, escrito por você. Se empatar, fica Go.

## Placar final (15 agentes)
| Lado | Quem | Margem |
|---|---|---|
| Go | GO1, GO2, GO3, GO4, N2, N3, MK2 | leve |
| TS | TS2, TS3, TS4 | leve |
| Empate | TS1, N1, MK1, MK3 (os 3 últimos pendem para o Go), N4 | — |

Ninguém declarou margem clara ou forte.

## O que empatou (não decide)
- **Performance:** 1.000 lutas em 29 ms (Node) × 2 ms (Go). As duas sobram para jogo por turno.
- **Hospedagem e RAM:** diferença pequena demais para pesar na VPS.
- **Teste por tabela:** Vitest `it.each` × `go test`.
- **Ajuda de IA:** os benchmarks empatam.
- **Contrato com o C#:** o custo de manter fica no lado C#, igual para as duas.
- **Imutabilidade:** nenhuma das duas impede você de mexer na entrada sozinha. As duas precisam de um teste de "entrada intacta".
- **Carreira:** você já pratica Go no trabalho, e o mercado premia jogo entregue, em qualquer linguagem.
- **Mercado:** TS tem mais uso (48,8% × 17,4% no Stack Overflow 2025). As duas têm backend de jogo.

## O que pesa para cada lado
| Go | TS |
|---|---|
| `int64` de verdade, com overflow definido | É a sua linguagem foco; algo na tela mais rápido (noite 5 em ~5 noites × 6–7 em Go) |
| Binário de 6–8 MB para ir junto do Unity (Bun: 86 MB; Node SEA ainda experimental) | União discriminada: o compilador avisa quando falta tratar um tipo de evento |
| Uma ferramenta só (`go mod init`) | Playground no navegador |
| No acervo: 55 commits seus à mão em Go × 1 em TS | — |

**Contra o Go:**
- a ordem de percorrer um `map` é aleatória, então nada de `map` no estado;
- falta modelo fechado para os eventos;
- slice e map copiados compartilham dados.

**Contra o TS:**
- `number` é float;
- no semi-idle, o `tsconfig` nunca checou os tipos;
- o `Math.random` apareceu nos testes.

**Dos dois lados:** no código feito à mão, você mexeu no estado de entrada e sorteou fora dele, tanto em TS quanto em Go. A disciplina vem do teste, não da linguagem.

## Protocolo do teste (versão do N4)
- **Antes:** você escreve à mão, em `casos.json`, 5 casos: 20 natural, 1 natural, acerto, erro e mob morto.
- **O que escrever:** `aplicar(estado, atacar)` com d20 + mod contra a Defesa, 4 graus, dano, semente no estado e eventos `ataque` e `dano`.
- **Vem pronto da IA:** só o laço da tabela, o teste de entrada intacta e o mesmo gerador de sorteio nas duas.
- **Setup à mão:** teto de 20 min. TS: `package.json` + `tsconfig` strict + `tsc --noEmit` no teste. Go: `go mod init`.
- **Ordem e tempo:** uma linguagem por noite, ordem tirada na moeda, 50 min de regra com o relógio pausado nas interrupções.
- **Desclassifica:** pedir código à IA, inclusive no setup.
- **Mede:**
  - se fechou e em quantos minutos;
  - quantas vezes travou mais de 5 min;
  - quantas consultas fez;
  - se a entrada ficou intacta e se o replay bate;
  - nota de 1 a 5 para "quero continuar amanhã nesta".
- **Vence:**
  - a única que fechar;
  - se as duas fecharem, a mais rápida por mais de 25%;
  - com diferença menor que 25%, a nota maior; com notas iguais, Go;
  - se nenhuma fechar, a de menos consultas; se empatar de novo, Go.
- **Validade:** definitiva até a fase 2 jogável. Se você não fizer o teste em 2 noites, fica Go.
- **Regras depois do teste:**
  - se TS: `tsc` strict no `npm test`, sem monorepo e sem Electron;
  - se Go: replay com 1.000 sementes e nada de `map` no estado.

## Só você responde (mudam o desempate)
1. Go à noite parece extensão do expediente?
2. O projeto também serve para treinar Go?
