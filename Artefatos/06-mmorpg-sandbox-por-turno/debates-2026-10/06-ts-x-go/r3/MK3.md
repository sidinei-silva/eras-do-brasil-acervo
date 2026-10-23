# MK3, rodada 3

## 1. Respostas

Ninguém me fez pergunta na rodada 2. As minhas perguntas da rodada 1 foram respondidas:
- **TS1:** o ferramental congelado é Node LTS + Vitest, com `tsc --noEmit && vitest run` e versões travadas. Aceito.
- **GO2:** os eventos ficam numa struct achatada com `Tipo` constante e checagem pelo `exhaustive` via `go tool`. Aceito.

A pergunta que fiz ao GO4 na rodada 2 (quem pega o `case` esquecido) já foi respondida pelo GO2: `go get -tool` e uma linha no script. Dou por encerrada.

## 2. Protocolo do teste de 1 noite

Aceito o teste do TS4 com os critérios do N2. Acrescento cinco pontos para evitar vício de medição.

1. **Preparo antes da noite, contado à parte.** Ele monta os dois esqueletos à mão:
   - Go: `go mod init` e um `_test.go` vazio;
   - TS: 1 `package.json`, 1 `tsconfig` e Vitest.

   Anota os minutos de cada um. Esse é o dado de "portas de ferramental" (N4, ponto 3).
2. **O que ele escreve:**
   - os 5 casos da tabela, no papel, **antes** de abrir qualquer editor, e iguais para as duas linguagens;
   - depois, `aplicar(estado, {tipo:"ataque"}, semente)` com d20 vindo da semente e evento `ataque` + `dano`.
3. **Ordem: TS primeiro, Go depois.** Quem vai em segundo herda o raciocínio do primeiro. Assim, o eventual bônus fica com a linguagem em que ele é iniciante, e uma vitória do TS não fica sob suspeita.
4. **Tempo:** até 50 minutos ativos por linguagem. O cronômetro para nas interrupções do bebê.
5. **O que medir:**
   - minutos até a tabela ficar verde;
   - travadas, contando cada vez que ficar 5 minutos sem avançar;
   - consultas à documentação;
   - perguntas de conceito à IA (código pedido à IA desclassifica);
   - se mutou a entrada (`deepFreeze` / `reflect.DeepEqual`, escrito pela IA antes).

**Empate:** diferença de até 25% no tempo, com as duas tabelas verdes. No empate ganha Go. Se a fluência empata, sobram o binário, os inteiros e menos portas de ferramental, e o próprio TS4 admitiu isso.

**Validade:** é definitivo até a fase jogável, pela regra anti-troca do N4. O resultado vai num ADR de uma linha com os números.

## 3. Inclinação atual

- **Pelo mercado:** empate. Longevidade, IA, carreira e hospedagem já foram zeradas pelo N4.
- **No meu recorte** (estabilidade e ferramental): **leve Go**. A linguagem e o ferramental vêm numa peça só, e o TS 7 mostrou que os padrões do lado TS mudam com as versões.

## 4. Condição

Aceito TS sem reclamar se as duas coisas acontecerem:
1. o teste der TS pelo menos 25% mais rápido, ou com menos travadas;
2. o esqueleto TS couber em uma noite, com `tsc --noEmit` no `npm test` desde o primeiro commit, como o TS1 se comprometeu.
