# TS1 — rodada 3

## Respostas
**@N4:** "Com que prova o TS é a linguagem da autoria? Sim ou não?" Prova de commit eu não tenho. Os ~55 commits em Go escritos à mão contra 1 em TS pendem para o Go. O que sobra para o TS é fluência, e fluência ainda não foi medida. Por isso o teste existe. Quanto ao compromisso: **sim** para `tsc --noEmit` strict no script de teste e **sim** para núcleo sem monorepo desde a noite 1.

**@GO1:** "mostre o `DeepReadonly` sem biblioteca." Ele cabe em uma linha:
```ts
type DeepReadonly<T> = { readonly [K in keyof T]: DeepReadonly<T[K]> };
```
Mesmo assim, já tirei esse argumento na r2: quem garante é o teste (`deepFreeze`), como o seu `reflect.DeepEqual`. Nesse ponto empatamos.

**@N3:** "playground como cliente WebSocket do servidor local?" **Sim.** Assim ele testa o contrato real, o mesmo que o Unity vai usar, e tira do caminho a vantagem de "núcleo no navegador" que eu tinha citado.

**@MK2 / @MK3 / @N1:** respondi na r2: VPS comum com `systemd`, ferramental travado no `package.json` e `number` com `>>> 0` mais semente fixa no teste.

## Protocolo do teste de 1 noite
- **O que ele escreve:** `aplicar(estado, {tipo:"ataque"})` com d20 + modificador contra CA, 4 graus no estilo PF2e (±10), dano, RNG no estado e eventos `ataque` e `dano`.
  - Os ~8 casos de tabela ele escreve antes, num JSON único usado pelas duas versões.
  - O laço de teste (DeepEqual ou deepFreeze) vem pronto nas duas, feito pela IA.
- **Tempo:** 45 min por linguagem, **TS primeiro**. Quem vai em segundo já pensou o problema, então o Go leva essa vantagem. É uma compensação justa, porque TS é a linguagem foco dele.
- **O que mede:**
  - minutos até a tabela passar;
  - consultas (doc ou pergunta à IA);
  - se mutou a entrada;
  - no dia seguinte, se explica cada linha sem olhar.
- **Empate (diferença abaixo de 25%):** **Go ganha.** Ele tem 7 semanas de Go escrito à mão, Go é a linguagem do trabalho e Go distribui melhor.
- **Validade:** definitiva. Depois do teste, a linguagem não volta à mesa. Se a fase estourar 1,5× o orçamento, corta-se escopo, conforme o roteiro.

## Inclinação
**Empate.** Sem a medição, nenhum lado tem margem. A união discriminada, a meu ver, não compensa a distribuição e o histórico em Go.

## Condição
Aceito Go sem reclamar se a versão Go fechar a tabela em até 1,25× o tempo da TS, sem mutar a entrada.
