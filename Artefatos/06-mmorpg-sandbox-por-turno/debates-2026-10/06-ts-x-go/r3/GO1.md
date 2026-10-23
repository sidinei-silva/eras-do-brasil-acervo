# GO1, rodada 3

## Respostas

**@N4:** "Aceitam o teste? Se o Go levar o dobro, aceitam o TS? Número de noites."
Sim, sim. Dei o número na rodada 2: chega na noite 6 ou 7 em Go e na noite 5 em TS (estimativa). O teste existe justamente para trocar essa estimativa por medida.

**@N2:** "Se fechar TS e não fechar Go na mesma noite, você cede ou culpa a curva de iniciante?"
Cedo. A curva de iniciante é justamente o custo que está sendo medido. Ele tem 2h por noite e não tem margem para pagar curva.

**@N3:** a parte dos eventos eu respondi na rodada 2 (struct achatada, classe espelho em C#). Sobre o playground: sim, como cliente WebSocket do servidor local. Isso tira do Go a necessidade de WASM.

**@TS1, @TS2, @TS4:** recebi o `DeepReadonly` de 5 a 6 linhas mais o `deepFreeze`. Fico com o `deepFreeze` como trava comum. Em Go, o equivalente é o laço comparar uma cópia JSON da entrada antes e depois.

## Protocolo que aceito

- **O que ele escreve:** `aplicar(estado, ataque)` com d20 sorteado do RNG que está no estado, os graus de sucesso (±10) e os eventos `ataque` e `dano`. São 6 casos que ele escreve antes, num `casos.json` usado pelas duas versões.
- **O que vem pronto:** nas duas linguagens, a IA prepara antes o esqueleto, o laço da tabela, o teste de mutação (`deepFreeze` em TS, comparação JSON em Go) e o mesmo gerador (splitmix32 ou mulberry32). Assim o teste mede a autoria do `aplicar`, não a montagem do projeto.
- **Tempo:** 45 min por linguagem, na mesma noite. A ordem é tirada no cara ou coroa, porque quem vai em segundo já pensou na lógica.
- **O que mede:**
  - minutos até os 6 casos passarem;
  - quantas vezes travou mais de 5 min;
  - quantas consultas fez (doc e pergunta conceitual à IA contam; pedir código desclassifica a versão);
  - se mutou a entrada;
  - uma nota de 1 a 5 para "parece meu".
- **Desempate:** se a diferença de tempo for de até 25%, ou se nenhuma das duas fechar, **ganha TS**, por ser a linguagem foco dele. Go só ganha se fechar e TS não, ou se for mais de 25% mais rápido.
- **Validade:** definitiva até a fase jogável, pela regra anti-troca. Anota numa linha no registro de noites e não repete o teste.

## Inclinação atual
Go, por margem **leve**. O que ainda pesa para o Go: 55 commits à mão em Go contra 1 em TS (N2), inteiro por especificação, uma porta só de ferramental e a distribuição. O que pesa contra: fluência e eventos, que ficam com o TS. O teste vale mais que essa margem.

## Condição
Aceito TS sem reclamar se o teste der empate ou vitória do TS. Também aceito se ele disser que Go à noite parece extensão do expediente. A fadiga de abril é um dado do acervo, não detalhe.
