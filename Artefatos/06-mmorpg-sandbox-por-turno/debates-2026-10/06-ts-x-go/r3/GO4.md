# GO4: rodada 3

## Respostas

**@TS4:** "Aceita o teste de uma noite? Qual medida conta?"
Aceito. As medidas contam nesta ordem:
1. A tabela fica verde com a entrada intacta e sem código pedido à IA. Isso é sim ou não.
2. Minutos até ficar verde.
3. Quantas consultas ele fez.

**@MK3:** "Ferramental 'sem nada para configurar'? Quem pega o `case` esquecido de `reacao`?"
Retiro o "nada". O certo é "uma decisão a menos". Duas correções:

- **O linter certo.** O `exhaustive`, que GO1 e GO3 citaram, checa `switch` de enum (constantes), não o `switch` de tipo ([nishanths/exhaustive](https://github.com/nishanths/exhaustive)). O linter que serve é o `go-check-sumtype` ([alecthomas](https://github.com/alecthomas/go-check-sumtype)).
- **Proposta sem linter.** O `UnmarshalJSON` do contrato já precisa de um registro `tipo → evento`. Um teste percorre esse registro e falha se alguma das funções com `switch` cair no `default`. Esquecer de registrar quebra o próprio JSON, então o registro não fica para trás.

```go
for nome, novo := range registro { // registro: map[string]func() Evento
    if _, err := descrever(novo()); err != nil { t.Errorf("%s sem case", nome) }
}
```

Quem pega o erro é o teste, não a IA. Mesmo assim, é mais trabalho que o `never` do TS.

## Protocolo do teste de 1 noite

**O que ele escreve:** `aplicar(estado, ataque)`. São quatro graus do PF2e contra CA, dano, semente no estado e os eventos `ataque` e `dano`.

**Antes, sem relógio:** escreve, num `casos.md` neutro, 5 casos com o resultado esperado e 1 caso de "entrada intacta". As duas versões usam os mesmos casos.

**Tempo:**
- 50 minutos por linguagem, na mesma noite.
- A ordem sai na moeda, porque quem vai em segundo já resolveu o problema uma vez.
- O setup entra no relógio: `go mod init` de um lado; tsconfig `strict`, `tsc --noEmit` e Vitest do outro. É o custo real.

**Mede e anota em `noites.md`:**
- passou ou não, e em quantos minutos;
- travas de mais de 5 minutos;
- consultas: documentação conta 1; a IA explicando um conceito conta 1; a IA escrevendo código desclassifica aquela linguagem;
- se mutou a entrada;
- em uma linha: qual das duas ele abriria amanhã.

**Empate:** vale essa linha. O critério de maior peso é autoria e vontade de voltar, não o meu lado. Se ele não tiver preferência, fica Go, pelo binário e pelos inteiros.

**Validade:** é definitivo até a fase 2 jogável, pela regra do N4. Sem revanche.

## Inclinação atual

**Go, leve.** Restam inteiros, distribuição e um caminho único de ferramental. Eventos ficam com o TS, e carreira não pesa.

## Condição

Aceito TS sem reclamar se:
- só a versão TS ficar verde; ou
- a versão Go levar mais de 1,5× o tempo da TS; ou
- ele escrever, na linha do empate, que prefere TS.
