# Rodada 2: o que aconteceu

**Como foi.** Cada agente leu o resumo da rodada 1, abriu os textos de quem ia atacar e respondeu:
- posição nas 10 disputas;
- 2 ataques a outros grupos;
- resposta a quem o atacou;
- o melhor e o pior custo-benefício pelo critério do M1.

Os mediadores fizeram o placar (M1) e a checagem de fatos (M2). Os textos estão em `r2/<ID>.md`.

## O fato da rodada: todo mundo convergiu
- **Melhor custo-benefício:** 17 dos 20 debatedores apontaram a **luta no terminal** do O2: um personagem contra 2 mobs, escrito por ele, cerca de 20 noites, algo na tela por volta da noite 5. Os outros 3 apontaram o orçamento de noites do M1, o descanso e os cortes do M1.
- **Pior custo-benefício**, empatados:
  - **Babylon + Electron** (T3): seria a 6ª troca de stack e perde o Sidekick;
  - **single-player agora** (O1): troca de gênero, a causa nº 1 das mortes no acervo;
  - **offline de 8h** (J3).

## Quem mudou de ideia, e por quê
| Agente | Antes | Depois | O que convenceu |
|---|---|---|---|
| O1 | single-player agora | MMO como destino, mas roda local com save em arquivo até a fase 6 | D1 (motivação) e M2 (o online nunca matou nada por defeito técnico) |
| D1 | servidor desde já | código local primeiro; servidor só antes da Travessia | O1 |
| T1 | Postgres e servidor real cedo | servidor TS embutido no jogo, save em arquivo até a fase 4; contrato e autoridade desde o dia 1 | ordem do M1 |
| J3 | offline de 8h | retirou; fica o **descanso** (teto de 3 noites, nada produzido) | G2 (inflação) e acervo (o offline puxou para idle 2 vezes) |
| J1 | carroça paga | retirou; nenhum nó a mais de ~4 min de uma cidade | G2 (resolve pelo mapa, não por sistema) |
| T5 | OO em tudo | na luta o T4 vence (função pura); OO fica no resto | a própria prova dele: o código citado era da IA |
| T4 | pureza total | puro só na borda (pode mexer na cópia por dentro); nomes em português | T5 (custo de cópia) |
| T3 | Babylon agora | playground no navegador nas fases 0–1; teste Babylon × Unity só na fase 2 | ataques de A1, T2, J1 e O2 |
| T2 | Sidekick "só Unity" | errou: roda na Unreal também (só a API de runtime é Unity); Unity entra depois da luta no terminal | M2 |
| J2 | movimento grátis | retirou; vira gatilho do simulador (se >60% dos turnos forem "anda + bate + bate") | G1 (movimento grátis apaga a intenção do inimigo) |
| G1 | penalidade -5/-10 | no máximo 2 golpes por turno; a 3ª ação é tática | J2 |
| A2 + A3 | silhueta modelada × kit próprio | tier pela cor, forma pela troca de peça do pack; kit da Era 1 depois da Travessia, com teto de 15 noites | trocaram argumentos entre si |
| G3 | cortar a lore | não corta, porque é o primeiro texto feito por ele; vira glossário de 10 termos | acervo, marco 65 |
| I2 | IA escreve teste | ele escreve a tabela e o teste do núcleo; IA só no simulador, pipeline e testes de rede | I1, I3 e a recusa de 18/04 no acervo |
| I3 | ele escreve tudo | aceita a IA escrever só o laço que roda a tabela, sem casos nem asserção | — |
| G2 | bot de 15h como critério | planilha de 1 noite agora; bot depois do núcleo | T5, O2, I2 |
| O2 | 6.200 linhas jogadas fora | corrigiu para ~7.700; exige algo visível até a noite 5 | recontagem |

## Placar do M1
- **Maioria clara:**
  - MMO fica, rede depois do mundo;
  - nada offline + descanso;
  - automático limitado por recurso;
  - Unity;
  - luta como função pura com log;
  - luta no terminal primeiro;
  - Travessia com 10 passos;
  - morte a pé;
  - troca de peça e cor.
- **Empatadas:**
  - quem escreve o código do teste;
  - quantas ações o inimigo faz por rodada;
  - luta parada passa ao automático?;
  - rendimento do automático e teto do descanso;
  - página da Steam.
- **Estimativa com a maioria:** 120–195 noites (6 a 10 meses), contra 205–355 antes.

## Checagem do M2
Cerca de 40 afirmações conferidas; a maioria se confirma.

- **Corrigidos:**
  - linhas do semi-idle (T5);
  - não existe pilar "sem religião" escrito no repo (A3);
  - o Sidekick roda na Unreal (T2);
  - a Hetzner reajustou preços em jun/2026 (T1);
  - o logout de 30 min do OSRS é de cliente, não do jogo (J3).
- **Ressalva:** "o Unity morreu" e "o single-player falhou" morreram no papel, sem código.
- **Contradições reais no repo:**
  - Era 1 até T6 × 8 marcos de maestria;
  - craft com graus de falha × "nada se perde por azar";
  - o pilar "pouco tempo";
  - o `03-tecnico.md`.
