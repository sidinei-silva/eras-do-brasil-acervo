## Sidinei:

Projeto Eras do Brasil (MMORPG por turno, dev solo, Go). Leia do GitHub, branch main, de sidinei-silva/eras-do-brasil: AGENTS.md, docs/00-como-documentar.md, docs/jogo/mecanicas/motor-de-regras.md, combate.md, personagem.md, docs/jogo/02-travessia.md e docs/gestao/fase-atual.md. A doc é a fonte da verdade. O repo eras-do-brasil-acervo é histórico: serve de inspiração, nunca de regra.

Objetivo: montar os dados T1 do primeiro combate no terminal (fase 1) e uma regra de conversão Albion → Eras que eu consiga refazer à mão para itens e mobs novos.

Como responder:
- Sou dev solo, com pouco tempo por sessão. Curto e fácil de validar; tabela em vez de parágrafo.
- No máximo 3 opções, com o porquê. Eu decido.
- Ideia fora da fase 1: diga "estacionar" e não desenvolva.
- Quando algo virar decisão, me dê a linha exata e o arquivo.
- Não escreva código de core/ nem casos de teste.
- Use o Albion como referência de formato e de proporção (tipo, custo, papel, força relativa). Não copie nomes nem textos.

Já decidido (não reabra):
- Acerto: 1d20 + HitModifier da skill contra a Defense do alvo. Graus a ±10. 20 natural sobe e 1 natural desce um grau.
- Dano: dado por skill (XdY+Z); só rola se acertou; o crítico dobra o total.
- Turno: 3 ações; no máximo 2 golpes, o 2º com -5. Ataque básico é uma skill da arma: 1 ação, 1 golpe.
- Tier T1: bônus esperado +2, Defesa do mob 10, acerto ~65% (tabela do motor-de-regras).
- Mob: bônus de ataque = Defesa típica do jogador no T1 − 8. Mob e jogador usam o mesmo formato:
  Fighter{ID, Side, HP, Defense, Skills}
  Skill{ID, Cost, HitModifier, Damage}
  Damage{Count, Sides, Modifier}
- O jogador chega à luta com números prontos. Uma função converte equipamento em números (atributo + maestria + item; bônus do mesmo tipo não soma).
- Skills de arma (modelo do Albion): Q e W vêm do pool da família (todas as espadas compartilham), liberados por tier; o E é único de cada arma. Na fase 1, a espada T1 usa só o que o Albion libera no T1.

Escopo da fase 1: tudo T1. O T1 é a porta de entrada, igual para todos; não existe T1 por estilo de combate.
- Jogador: espada T1 + escudo T1 + couro T1 (peitoral, capacete, botas).
- Inimigos: os primeiros do tutorial. O 02-travessia.md propõe 2 inimigos T1 (Caranguejo da Beira, Teiú Rachado) e o chefe como T2. Quero também um mini boss T1. Aponte o conflito com o doc e com o "Pronto quando" da fase 1 (2 mobs), e eu decido.
- Fora do escopo (estacionar): quais famílias de arma vão existir e os arquétipos (Warrior/Hunter/Mage ou o que está no acervo).

Faça nesta ordem e PARE ao fim de cada etapa, esperando a minha decisão:

1. HP: não está em nenhum doc. Até 3 opções (de onde vem, quanto vale no T1 para jogador e mob), com o porquê.
2. Skills de armadura: confirme no Albion como funcionam (ativa por peça? passivas por material?) e me diga antes de montar o armor.json. Para mim é só hipótese.
3. Regra de conversão, em passos curtos:
   - usar PROPORÇÃO (o item ou mob em relação à média do tier na mesma peça), nunca os números absolutos do Albion, porque a escala é outra (tempo real × d20 por turno);
   - os números absolutos vêm da tabela de tiers do motor-de-regras;
   - dizer o que do Albion não tem equivalente no Eras e é ignorado;
   - a média do dado (X×(Y+1)/2 + Z) é o que se compara com o dano relativo do Albion.
4. Aplicar a regra aos itens e mobs, mostrando a conta numa tabela por item (Albion → proporção → Eras).
5. Entregar o JSON de data/ (weapons.json, armor.json, enemies.json) e o formato de equipamento que a função equipamento → números vai ler.
   - Cada item tem um campo Family (ex.: "sword", "shield", "leather"), para eu agrupar em árvores depois sem refazer números.
   - As skills de arma ficam num pool por família, com o tier em que cada uma libera; o E fica na arma.
   - Se o formato tiver escolha, até 3 opções.
6. As linhas exatas para os docs (arquivo e seção).

Fontes (comece por estas, mas pode buscar em outros lugares, como repositórios no GitHub com dump dos dados do Albion, tipo ao-bin-dumps):
- https://wiki.albiononline.com/wiki/Beginner%27s_Broadsword
- https://albiononline2d.com/en/item/id/T1_MAIN_SWORD/spells
- https://wiki.albiononline.com/wiki/Category:Sword_Ability
- https://albiononline2d.com/en/item/id/T1_OFF_SHIELD
- https://albiononline2d.com/en/item/id/T1_ARMOR_LEATHER_SET1
- https://albiononline2d.com/en/item/id/T1_HEAD_LEATHER_SET1
- https://albiononline2d.com/en/item/id/T1_SHOES_LEATHER_SET1
- https://wiki.albiononline.com/wiki/Category:Armor
- https://wiki.albiononline.com/wiki/Plate_Helmets
- https://wiki.albiononline.com/wiki/Tutorial
- https://albiononline2d.com/en/mob/tier/1
Diga de onde saiu cada número do Albion.

---

## Claude:

Vou ler os docs direto do GitHub (main) e começar pela etapa 1 (HP).

**Etapa 1 · HP**

Li o `main` (commit `87db79b`). HP só aparece como "Em aberto" em `combate.md`. O `personagem.md` liga Vigor a Vida, mas os atributos ainda estão em aberto.

**O que o Albion dá** (fonte: `ao-bin-dumps`, `characters.xml` e `mobs.json`)

| Quem | HP Albion | Em relação ao jogador |
|---|---|---|
| Jogador (base) | 1200 | 100% |
| T1 Scavenger (tutorial) | 120 | 10% |
| T1 Archer | 168 | 14% |
| T1 Thief | 228 | 19% |
| T1 Tank | 288 | 24% |
| T2 mini boss do tutorial (Brawler) | 659 | 55% (≈2× um T2 comum) |

**Opções**

| # | De onde vem o HP | T1 jogador / mob | Porquê | Contra |
|---|---|---|---|---|
| **A** | **Turnos para matar**: HP = dano médio por turno × turnos-alvo | ~40 / ~16 (o mob morre em ~3 turnos; o jogador vence 2 mobs com folga) | Não depende de atributo em aberto. O simulador ajusta só os turnos-alvo. O Albion entra só como proporção *entre* mobs (Tank > Thief > Archer) | O número final só fecha depois da etapa 4 (dano) |
| B | Vigor: base + Vigor × k + peça | depende de k | Segue o `personagem.md` (Vigor → Vida) | Puxa questões em aberto (escala de atributo, base 3 ou 0) para dentro da fase 1 |
| C | Proporção direta do Albion: mob = 10–24% do jogador | 40 / 4–10 | Uma conta só | O mob morre em 1 ou 2 golpes. A proporção vem do tempo real, não do turno |

Conta de A (dado provisório 1d8, média 4,5): 1º golpe a 65% com 15% de crítico, mais o 2º golpe a 40% → **~5,6 de dano por turno**. Daí sai mob ≈ 16, e o jogador ≈ 40 para sobreviver aos 2 mobs.

**Conflito do mini boss** (decida junto)

| Fonte | Diz |
|---|---|
| `fase-atual.md` · Pronto quando | 1 personagem × **2 mobs** |
| `02-travessia.md` · Propostas | 2 mobs T1; o chefe é T2 |
| Albion | Não tem mini boss T1; o do tutorial é T2 |

| # | Opção | Porquê |
|---|---|---|
| 1 | Mini boss T1 é o 3º inimigo no `data/` e o "Pronto quando" vira "2 mobs + 1 mini boss" | Testa um HP maior no simulador já agora |
| 2 | **Estacionar** o mini boss | Mantém a fase 1 como está; o mini boss é da Travessia |
| 3 | O mini boss substitui um dos 2 mobs | Não muda a contagem, mas testa um perfil diferente |

Me diga: **HP A, B ou C** e **mini boss 1, 2 ou 3**. Depois disso te passo as linhas e sigo para a etapa 2.

---

## Sidinei:

Decisões da etapa 1 (já estão nos docs; leia o personagem.md de novo no main):

HP:
- Atributo é só modificador: sem equipamento tudo vale 0; as peças somam até +4.
- O peso da armadura (o mesmo da velocidade: placa 2, couro 1, pano 0 por peça; total de 0 a 6) soma no HP. O Vigor é um atributo normal, de 0 a +4.
- Jogador: HP = base + (peso + Vigor) × k, sem nível. Tudo pano sem Vigor ≈ 55% do HP de tudo placa. Calibre base e k pela sua opção A (turnos para matar), respeitando essa proporção.
- Mob: não tem atributo. Recebe o HP pronto pela opção A; a proporção do Albion vale só entre os mobs.

Encontros: no lugar de mini boss 1, 2 ou 3, serão 3 encontros, como o Farol do tutorial do Albion (The Lighthouse; no Eras, a Beira):
1. 3× Wild Scavenger (T1): https://albiononline2d.com/en/mob/id/T1_MOB_HERETIC_SCAVENGER_TUTORIAL
2. 1× Strange Thief (T1): https://albiononline2d.com/en/mob/id/T1_MOB_HERETIC_THIEF
3. 1× Big Spencer (T2, mini boss, sozinho): https://albiononline2d.com/en/mob/id/T2_MOB_RD_HERETIC_BRAWLER_MINIBOSS
O "Pronto quando" já mudou. Use os números T2 do motor-de-regras para o mini boss. Nomes do Eras eu dou depois; use IDs provisórios.

Os 3 encontros são só a primeira zona, não o fim do tutorial. O MVP é a Travessia inteira, mais tarde. Aqui é só o combate da fase 1.

Siga para a etapa 2.

---

## Claude:

**Etapa 2 · Skills de armadura no Albion**

Reli o `personagem.md` no main (`9a0f908`). Suas decisões de HP estão lá.

**O que o Albion faz** (fonte: `ao-bin-dumps`, arquivos `items.xml` e `spells.xml`)

| Pergunta | Resposta | Prova no dump |
|---|---|---|
| Ativa por peça? | **Sim.** Cada peça tem 1 ativa: peitoral = R, capacete = D, botas = F | `activespellslots="1"` nas 3 peças |
| Passiva por material? | **Não.** A passiva também é **por peça**, escolhida de um pool | `passivespellslots`: 0 no T1, 1 do T2 em diante |
| De onde vem o pool? | Do **material + set** e do slot (couro set1 ≠ placa set1), igual à família da arma | `shopsubcategory2="armors_leather_set1"`; o T4 aponta para a lista do T3 |
| Cresce com o tier? | **Sim**, igual à arma | Peitoral de couro: 1 skill no T1, 4 no T2, 6 no T3 |
| Existe T1 de pano ou placa? | **Não.** O T1 é só couro set1, espada, escudo e ferramentas | Lista de itens `T1_*` |
| Escudo T1 tem skill? | **Não**, nem ativa nem passiva | `T1_OFF_SHIELD` com 0 e 0 |

Então sua hipótese estava meio certa: **a ativa é por peça**, e a passiva é **por peça, com pool por material**.

**O couro T1 no Albion, traduzido para o Eras**

| Peça | Skill T1 | Albion | Equivalente na fase 1 |
|---|---|---|---|
| Peitoral | Cura fora de combate | `usableincombat="false"`, recarga 30 s | **Nenhum**: ignorar (a regra de fora da luta fica para a fase de mundo) |
| Capacete | Barreira (escudo que absorve dano) | buff em si, recarga 25 s | Skill de defesa: absorve X de dano |
| Botas | Esquiva (rola 5 m) | movimento, recarga 20 s | Andar algumas casas sem ataque de oportunidade (≈ Passo melhorado) |

**Um conflito que aparece na etapa 4 (espada T1)**

| Fonte | Q | W | E |
|---|---|---|---|
| Wiki (Beginner's Broadsword) | Heroic Strike | Interrupt | Mighty Blow |
| Dump do jogo | Heroic Strike **ou** Cleave | — | Mighty Blow |

Vou seguir o **dump**, porque é o arquivo do cliente e a wiki é editada à mão. Se preferir a wiki, me avise.

**Opções para o `armor.json` da fase 1**

| # | Opção | Porquê | Contra |
|---|---|---|---|
| **A** | Mesmo modelo da arma: pool por família (`leather`) e slot, com tier de liberação. No T1 só entram as 2 ativas usáveis em luta (capacete e botas) | Fica igual ao Albion e não refaz o formato depois | Duas skills a mais para o core resolver: escudo que absorve dano e movimento |
| B | Formato com pool, mas **vazio no T1**: a armadura só dá números (Defesa, peso → HP, velocidade) | O "Pronto quando" não pede skill de armadura. Menos core na fase 1 | O tutorial do Albion ensina o D e o F logo cedo |
| C | Só a skill do capacete (barreira). Botas sem skill na fase 1 | A esquiva repete o Passo, que já está no `combate.md` | Foge um pouco do Albion |

Minha leitura é **B** para a fase 1 (o formato já fica pronto para A). Mas a decisão é sua.

**Linha para o doc** (vale para qualquer opção)

`docs/jogo/mecanicas/personagem.md` › **Build**, no lugar do item "Skills de armadura" de **Em aberto**:

> - Armadura: 1 skill ativa por peça (peitoral, capacete, botas) e, do T2 em diante, 1 passiva por peça. As duas vêm do pool do material e do slot, e as opções aumentam com o tier (como o Albion).

Me diga **A, B ou C** e se fica **dump ou wiki** para a espada. Aí sigo para a etapa 3.

---

## Sidinei:

Correções (vi no jogo, vale mais que o dump e a wiki):
- Espada T1: só Q, com 2 opções (Heroic Strike ou Heroic Cleave). Sem W e sem E no T1.
- Armadura T1: cada peça (peitoral, capacete, botas) tem 1 ativa, 1 opção, sem passiva. O peitoral também tem ativa (a cura fora da luta).
- Off-hand: sem skill e sem passiva em nenhum tier.
- T2: cada peça tem 2 opções de ativa e 2 de passiva. Do T3 em diante, 3 e 3. O peitoral de placa tem um 2º slot de passiva (2 opções). A 1ª ativa de cada peça é a mesma em todos os materiais.
- O MVP vai até o T2.

Decisões (já estão nos docs; releia personagem.md e combate.md no main):
- Armadura: opção A, com pool por material e slot e tier de liberação. A fase 1 inclui tudo o que o T1 tem: as 3 ativas da armadura e o Q da espada com as 2 opções.
- O HP passa de um encontro para o outro. Fora da luta, o HP regenera sozinho (como no Albion), e a ativa do peitoral cura entre os encontros.
- A forma da área do Cleave fica para quando houver grid. Na conversão, trate como área provisória.
- O ataque básico vem da arma, à parte do Q (como o auto-ataque do Albion): 1 ação, 1 golpe. A espada T1 tem ataque básico + Q (2 opções).

Siga para a etapa 3. Na conversão, inclua: o custo em ações de cada skill (1, 2 ou 3), com o porquê; quanto a barreira absorve; quantas casas a esquiva anda; e o regen fora da luta como proporção do HP máximo.

---

## Claude:

**Etapa 3 · Regra de conversão Albion → Eras**

Reli o `combate.md` e o `personagem.md` no main (`e317a5e`). Os números do Albion abaixo saem do `ao-bin-dumps`: `spells.xml`, `items.xml`, `mobs.json` e `characters.xml`.

**A regra, em 6 passos (dá para refazer à mão)**

| # | Passo | Conta |
|---|---|---|
| 1 | Achar a **referência** do Albion: o mesmo tipo de coisa, no mesmo tier | Skill de arma → ataque básico da arma · mob → média dos mobs comuns do tier · utilidade → HP do jogador |
| 2 | **Proporção** = valor Albion ÷ referência Albion | Um número só, como 2,4× ou 16% |
| 3 | **Base do Eras** = tabela de tiers do motor-de-regras + HP da opção A | T1: Defesa 10, bônus +2 · T2: Defesa 11, bônus +3 |
| 4 | **Valor Eras** = base × proporção, arredondado para inteiro | Só inteiros (núcleo) |
| 5 | **Dado**: escolher XdY+Z cuja média X×(Y+1)/2+Z fique perto do valor Eras | Ex.: alvo 8 → 2d6+1 (média 8) |
| 6 | **Custo em ações** (só para skill de dano) = proporção do passo 2 arredondada, de 1 a 3 | Skill forte custa mais ações: é assim que o turno "paga" o que no Albion é recarga |

- **Mob**: no passo 2, compare o **dano por segundo** (dano × velocidade de ataque), porque um turno é uma fatia fixa de tempo.
- **Utilidade instantânea** (barreira, esquiva): custa 1 ação. No Albion as duas têm `castingtime="0"`.

**O que não tem equivalente no Eras e é ignorado**

| Albion | Por quê |
|---|---|
| Energia ou mana (`energyusage`) | O Eras não tem energia; o custo é em ações |
| Tempo de conjuração e de animação | Por turno, não há tempo de animação |
| Bônus de velocidade de ataque e de movimento (pilhas do Heroic, +40% depois da esquiva) | Velocidade no Eras é casas por ação, fixa por peso |
| Armadura física e resistência mágica em número | A Defesa vem da tabela de tiers |
| Item Power, Ability Power e qualidade | O tier já dá o bônus de item |
| Bônus de dano por alvo extra (Cleave +8% por alvo) | Área provisória, sem bônus |
| Imunidade a controle e invencibilidade de 0,8 s da esquiva | Viram "não provoca ataque de oportunidade" |
| Fama, prata, raio de aggro | Fora do combate |

**Os 4 pontos que você pediu**

| Ponto | Albion (fonte) | Proporção | Eras T1 |
|---|---|---|---|
| Custo do **Heroic Strike** | 77,81 de dano (`HEROICSTRIKE2_EFFECT`) ÷ 33 do ataque básico (`T1_MAIN_SWORD attackdamage`) | 2,36× | **2 ações**, 1 golpe |
| Custo do **Cleave** | 60,63 por alvo (`CLEAVE_EFFECT`), área de raio 6 | 1,84× por alvo | **2 ações**, 1 golpe, área provisória |
| **Barreira** (capacete) | absorve 192 por 3 s (`ENERGY_BARRIER_SHIELD`) ÷ 1200 de HP do jogador | 16% do HP máx. | **1 ação**; absorve 16% do HP máx. até o seu próximo turno; não acumula (`maxstack="1"`) |
| **Esquiva** (botas) | rolamento de 5 m (`DODGE castrange`) ÷ 5,5 m/s de movimento (`characters.xml`) | ≈ 1 s de andar | **1 ação**: anda as casas de 1 ação **sem provocar** ataque de oportunidade |
| **Regen fora da luta** | 12 HP/s (`hitpointsregeneration`) ÷ 1200 | **1% do HP máx. por segundo** (cheio em 100 s) | Hipótese: no dump, 12/s é o regen base, e não achei um campo separado para fora da luta |
| **Cura do peitoral** | 9% do HP máx. × 8 pulsos (`OUTOFCOMBATHEAL_EFFECT`, `effectcount="8"`); o dano cancela | **72% do HP máx.** | Fora da luta, entre encontros; 0 ação |

Uma ressalva: os números das skills são os valores crus do dump. Se o tooltip no jogo mostrar outro número, troque e refaça só o passo 2.

**3 decisões antes da etapa 4**

**① Recarga.** O Albion usa recarga (barreira 25 s, esquiva 20 s, Heroic e Cleave 3 s), e o Eras não tem recarga no doc. Sem limite, a barreira sai todo turno.

| # | Opção | Porquê |
|---|---|---|
| A | Recarga de 3 s ou menos: ignorar, só vale o custo. 20 s ou mais: **1 vez por luta** | Uma regra simples e fiel à proporção: 20–25 s ≈ uma luta curta no Albion |
| B | Recarga em turnos = recarga ÷ 6 s, arredondada (barreira 4 turnos, esquiva 3) | Mais fiel, mas cria um campo e um contador novos no estado |
| C | Sem limite; a barreira custa 2 ações | Nenhum campo novo, mas custa caro e vira skill ruim |

**② Big Spencer: qual ID?** O link que você mandou é a versão de dungeon. O tutorial usa outra.

| ID | HP | Dano × vel. | HP ÷ média T2 (425) |
|---|---|---|---|
| `T2_MOB_TUT_HERETIC_BRAWLER_MINIBOSS` (tutorial) | 659 | 66 × 0,4 | **1,55×** |
| `T2_MOB_RD_HERETIC_BRAWLER_MINIBOSS` (seu link) | 1829 | 127 × 0,3 | 4,3× |

Para o Farol, a versão do tutorial faz mais sentido.

**③ Regen entre encontros na fase 1**: o terminal não tem relógio de mundo.

| # | Opção | Porquê |
|---|---|---|
| A | O peitoral cura 72% entre encontros, e o regen fica de fora até a fase de mundo | Testa só o que o T1 tem |
| B | Um campo `rest_seconds` por encontro no `data/` (1%/s) + a cura do peitoral | Já calibra o regen, mas é um campo a mais |

**Linhas para os docs** (as que já são decisão)

`docs/jogo/mecanicas/combate.md` › **Turno**, novo item:
> - Skill de dano custa ações em proporção ao dano: 1, 2 ou 3 (2 ≈ o dobro do ataque básico). Utilidade instantânea custa 1 ação.

`docs/jogo/mecanicas/combate.md` › **Entrada na luta**, completando o item do HP:
> - Regen fora da luta: 1% do HP máx. por segundo. A ativa do peitoral T1 cura 72% do HP máx. fora da luta e para se levar dano.

(Ao aceitar, sai "Velocidade do regen fora da luta" de **Em aberto**.)

Me diga **① A, B ou C**, **② TUT ou RD** e **③ A ou B**. Aí aplico a regra na etapa 4.

---

## Sidinei:

A ideia de construir equivalências faz sentido, mas cada uma vira regra nova no core. O critério que mantém a fase 1 viável é o mesmo de antes: **entra agora o que o T1 do Farol usa; o resto ganha uma equivalência decidida quando a primeira coisa que precisar dela aparecer.**

## Como o PF2e trata essas coisas

- **Mana:** o PF2e não tem. Ele usa a economia de ações (cada magia custa 1, 2 ou 3 ações), magias "infinitas" mais fracas e magias fortes com usos limitados, que recarregam no descanso. É o mesmo caminho que o Eras já segue: custo em ações, mais a recarga.
- **Imunidade a controle:** tem. Depois de sofrer certos efeitos (ficar amedrontado, por exemplo), a criatura fica temporariamente imune àquele efeito. Também existe a regra de *incapacitation*, que enfraquece o controle contra inimigos mais fortes. Invencibilidade, no sentido da esquiva do Albion, não existe.
- **Tempo de conjuração:** o PF2e não tem. Existe em Final Fantasy Tactics (a magia cai numa área depois de alguns turnos), e é a ideia que você descreveu. Não sei dizer como é no Wartales.

## Cada coisa do Albion: equivalente e quando

| Albion | Equivalente no Eras | Quando |
|---|---|---|
| **Recarga** | Em turnos: recarga ÷ 6 s, arredondada (o turno do PF2e equivale a 6 s). O automático solta as skills pela prioridade, respeitando a recarga | **Fase 1** |
| **Bônus por alvo extra** | O Cleave ganha dano a mais por alvo extra | **Fase 1** (o Cleave é T1) |
| **Armadura física** | A parte "armadura" da Defesa (`10 + Astúcia + armadura`), convertida por proporção | **Fase 1** (o couro T1 tem armadura) |
| **Velocidade de ataque** | Já está embutida: a conversão do mob usa dano × velocidade | Já resolvido |
| **Bônus de velocidade de movimento** | Condição de "+N casas por X turnos" | Estacionar: nenhuma skill T1 que entra usa |
| **Resistência mágica** | Redução fixa de dano mágico, como a *resistance* do PF2e | Estacionar até a primeira magia (o cajado é T2) |
| **Mana** | Nenhuma, por enquanto: o custo é em ações e a recarga limita o uso | Reabrir com a primeira magia (T2) |
| **Tempo de conjuração** | Área marcada que cai depois de N turnos | Estacionar até a primeira magia com conjuração |
| **Item Power, qualidade, encantamento** | A qualidade soma no bônus de item; o encantamento já está como "sem no início" | Estacionar até a **fase 3**, quando entra o craft T1 |
| **Imunidade a controle e invencibilidade** | Imunidade temporária depois do controle, como no PF2e | Estacionar até a primeira skill de controle (o Tacape do acervo atordoa) |

## As 3 decisões do outro chat

- **Recarga: B** (em turnos), como você quer.
- **Big Spencer: TUT.** O seu link era da versão de dungeon. A versão do tutorial (659 de HP) é a do Farol.
- **Regen: B**, com um campo de tempo de descanso por encontro, calibrando desde já.

## O que vai para os docs

A conversão termina na etapa 6, que já é a das linhas para os docs. Então não precisa de prompt para o Claude Code agora: deixe o outro chat juntar tudo no fim. Os itens marcados como "estacionar", você guarda com `/estacionar` no Claude Code.

## Para colar no outro chat

```
Decisões:
① Recarga: B (em turnos = recarga ÷ 6 s, arredondada; abaixo de 3 s, sem recarga). O automático solta as skills na ordem de prioridade, respeitando a recarga.
② Big Spencer: TUT (T2_MOB_TUT_HERETIC_BRAWLER_MINIBOSS).
③ Regen: B (campo de descanso por encontro em data/, 1%/s, mais a cura do peitoral).

Sobre o que você ignorou, mudo três:
- Armadura física: não ignore. Ela vira a parte "armadura" da Defesa (Defesa = 10 + Astúcia + armadura, nas Propostas do personagem.md), convertida por proporção dentro do tier. Mostre a conta do couro T1.
- Bônus por alvo extra do Cleave: entra. Converta para inteiro (ex.: +1 de dano por alvo extra) e diga a conta.
- Velocidade de ataque: já está no passo do mob; mantenha.

O resto (mana, tempo de conjuração, bônus de movimento, resistência mágica, Item Power/qualidade/encantamento, imunidade a controle) fica estacionado: não entra na fase 1 e não deve aparecer nos JSONs agora.

Ainda nas linhas para os docs: guarde tudo para a etapa 6 e entregue de uma vez, já incluindo recarga e regen.

Siga para a etapa 4.
```

---

## Claude:

**Etapa 4 · A regra aplicada (T1 do Farol)**

Decisões aplicadas: recarga em turnos, Big Spencer = TUT, descanso por encontro. Os números do Albion saem do `ao-bin-dumps` (`items.xml`, `spells.xml`, `mobs.json`, `characters.xml`).

**0 · Âncoras do Eras** (os únicos números absolutos que eu escolhi; o resto é proporção)

| Âncora | Valor | Porquê |
|---|---|---|
| Ataque básico da espada T1 | 1d6 (média 3,5) | O menor dado que ainda deixa 2,36× e 1,84× virarem dados inteiros |
| Mob comum T1 morre em | 2 turnos | Opção A |
| Mob comum mata o jogador de couro em | 13 turnos | Mantém a razão do Albion: o jogador mata o mob médio em 4,6 s (204 ÷ 44,6) e o mob médio mata o jogador em 29,9 s (1200 ÷ 40,1), ou seja, 6,5× |
| Armadura de placa no T1 | +3 | É o único valor que dá 3 degraus distintos (pano 1, couro 2, placa 3) |
| Arredondamento | Para o mais próximo; ,5 sobe | Por isso uma recarga de 3 s vira 1 turno (3 ÷ 6 = 0,5) |

**1 · Espada T1** (referência: o ataque básico, 33 de dano)

| Skill | Albion | Proporção | Eras | Custo | Recarga |
|---|---|---|---|---|---|
| Ataque básico | 33 | 1,0× | **1d6** (3,5) | 1 | — |
| Q · Heroic Strike | 77,81, recarga 3 s | 2,36× → 8,3 | **2d6+1** (8) | 2 | 1 turno |
| Q · Cleave | 60,63 por alvo, área, recarga 3 s | 1,84× → 6,4 | **1d8+2** (6,5), área provisória | 2 | 1 turno |
| Cleave · bônus por alvo extra | +8% por alvo (`targetcountvaluebonusfactor`) | 2 alvos: 0,52 · 3 alvos: 1,04 | **+1 de dano se atingir 2 ou mais alvos** | — | — |

**2 · Escudo T1**

| Albion | Proporção | Eras |
|---|---|---|
| 1,08% de redução contra mobs, 4% de ameaça, sem armadura e sem HP | Arredonda para 0 | **Tudo 0.** Só ocupa o off-hand (a fase 1 não tem ameaça) |

**3 · Couro T1** (o T1 só tem couro; o pano e a placa saem do peitoral T2, porque a armadura do dump não muda com o tier)

| Peça | Albion | Proporção | Eras |
|---|---|---|---|
| Armadura do peitoral | pano 68 · **couro 108** · placa 143 (o capacete e as botas têm 0) | couro ÷ placa = 0,755 | **+2** (pano 0,476 → +1, placa +3) |
| Peso | — | (`combate.md`) | 1 por peça → **3** |
| Capacete · Barreira | absorve 192 por 3 s, recarga 25 s | 192 ÷ 1200 = 16% | **Absorve 16% do HP máx. (6)**, 1 ação, recarga 4 turnos |
| Botas · Esquiva | 5 m, recarga 20 s | ≈ 1 ação de andar | **4 casas sem provocar ataque de oportunidade**, 1 ação, recarga 3 turnos |
| Peitoral · Cura | 9% × 8 pulsos; o dano cancela; recarga 30 s | 72% | **Cura 72% do HP máx. (29)**, fora da luta, 1 vez por descanso |
| Regen | 12/s ÷ 1200 | 1%/s | `rest_seconds` por encontro (provisório: 30 s → 12 HP) |

**4 · Jogador montado** (função equipamento → números)

| Número | Conta | Valor |
|---|---|---|
| HP | base 28 + (peso 3 + Vigor 0) × k 4 | **40** (placa 52 · pano 28 = 54% ✓) |
| Defesa | 10 + Astúcia 1 (botas) + armadura 2 | **13** |
| HitModifier | Força Bruta 1 (espada) + maestria 1 + item 0 | **+2** ✓ igual à tabela T1 |

Como saem o k e a base: dano médio por turno do mob comum = 3,1 → 3,1 × 13 turnos ≈ 40 de HP. A regra dos 55% dá base ≈ 7,3 × k, então k = 4 e base = 28.

⚠️ Os atributos por peça e "Defesa = 10 + Astúcia + armadura" estão nas **Propostas** do `personagem.md`. Usei esses valores, mas eles ainda não são regra.

**5 · Mobs** (Defesa: no Albion todos os mobs do mesmo tier têm a mesma armadura, 80 no T1 e 97 no T2, então vale o valor da tabela de tiers)

| Mob (ID provisório) | HP Albion ÷ média do tier | HP Eras | Dano/s Albion ÷ média | Dano Eras (2 golpes) | Def. | Acerto |
|---|---|---|---|---|---|---|
| `t1_scavenger` ×3 | 120 ÷ 204 = 0,59 | 12 × 0,59 = **7** | 36,5 ÷ 40,1 = 0,91 | **1d4** | 10 | +5 (13 − 8) |
| `t1_thief` | 228 ÷ 204 = 1,12 | **13** | 38,7 ÷ 40,1 = 0,97 | **1d4** | 10 | +5 |
| `t2_spencer` (TUT) | 659 ÷ 425 = 1,55; base T2 = 12 × 2,08 = 25 | **39** | 26,4 ÷ 39,6 = 0,67 | **1d3** | 11 | +6 (Def. típica T2 = 14) |

- Média comum T1 = 12 HP: o jogador tira em média 6,2 por turno (alternando o turno de Q com o de 2 básicos) × 2 turnos.
- Base T2 = T1 × 2,08 (média de HP T2 ÷ T1 no Albion). O dano por segundo médio do T2 é quase igual ao do T1 (39,6 contra 40,1), então o dado base não muda; o T2 sobe só no acerto.

**Teste rápido** (20 mil corridas dos 3 encontros, com política simples; é só uma sanidade, não substitui o simulador)

| Q escolhido | Morte | Turnos (enc. 1 / 2 / 3) | HP no fim |
|---|---|---|---|
| Heroic Strike | 10% | 5,2 / 2,5 / 7,1 | 25 |
| Cleave | 2% | 3,2 / 3,0 / 8,6 | 23 |

O Cleave ganha no encontro com 3 Scavengers, como no Albion. Ajustar isso é tarefa do relatório de 1.000 lutas.

**Ignorado nesta etapa** (entra na lista final): +50 de HP e +5% de regen da espada (o HP vem só do peso + Vigor), ameaça e redução do escudo, bônus de velocidade de ataque e de movimento do Heroic.

**Decisões antes da etapa 5**

**① O tapa do Big Spencer.** No Albion é o que define o chefe: 800 de dano (67% do HP do jogador), com 1,8 s de aviso numa área em anel e empurrão de 3 m. Fonte: `MOB_HERETIC_BRAWLER_DOUBLESLAP_TUTORIAL`. Depende de tempo de conjuração, que está estacionado, e de grid.

| # | Opção | Porquê |
|---|---|---|
| A | Fase 1 sem o tapa: Spencer só com ataque básico (é assim no teste acima). O tapa vai para o estacionamento junto com o tempo de conjuração | Nada novo no core agora |
| B | O tapa usa a **intenção do inimigo** que já está no `combate.md`: avisa num turno e cai no seguinte. 3 ações, 27 de dano (67% de 40), e a Esquiva ou o andar escapam | Não cria regra nova, mas só dá para fazer depois da tarefa do andar (grid) |
| C | Skill comum de 3 ações, sem aviso: 3× o básico = **3d3** | Fiel ao custo, mas perde o propósito do chefe |

**② Âncoras.** Confirme ou troque: básico 1d6 · 2 turnos para matar · placa +3 · os atributos das Propostas (Força Bruta +1 na espada, Astúcia +1 nas botas, maestria +1).

**③ Recarga de 3 s = 1 turno** (o ,5 sobe, coerente com "abaixo de 3 s, sem recarga"). Com isso, o Q fica disponível 1 turno sim, 1 turno não. Confirma?

---

## Sidinei:

① O tapa do Big Spencer: nenhuma das 3. Ele entra na fase 1 e estreia a conjuração:
- Intenção (combate.md) continua para o mob comum: mostra para onde vai e ataca na hora.
- Qualquer skill com conjuração, de qualquer um (mob ou jogador): ao ser usada, marca a área; cai no início do próximo turno de quem conjurou. Quem sair da área escapa (andar ou esquiva). Na fase 1 não pode ser interrompida (não há controle ainda). O tapa do Spencer é a primeira skill a usar essa regra.
- O tapa entra completo: área em anel, dano (proporção do HP do jogador) e empurrão. Converta: custo em ações, dano em dado, forma do anel em casas, empurrão em casas, recarga em turnos. A forma exata no grid eu confirmo na tarefa do andar.

② Âncoras confirmadas: básico 1d6, mob comum morre em 2 turnos, placa +3. Dos atributos, promovo para regra só o que a fase 1 usa: arma = atributo ofensivo da linha, botas = Astúcia, Defesa = 10 + Astúcia + armadura, maestria +1 no T1. Inclua essas linhas na etapa 6.

③ Recarga de 3 s = 1 turno: confirmado.

Na etapa 6, inclua também: a linha da conjuração no combate.md; a tarefa nova em docs/gestao/fase-atual.md, logo depois da tarefa do Passo e do ataque de oportunidade: "- [ ] Conjuração e o tapa do mini boss: área marcada, cai no próximo turno, empurrão · 2 sessões"; e o orçamento da fase 1 indo de 27 para 29 (fase-atual e roteiro).

Siga para a etapa 5.

---

## Claude:

**Etapa 5 · O tapa do Big Spencer e o `data/`**

Uma correção antes: no dump, o tapa **não é um anel completo**. É um **arco de 35° à frente** do Spencer, entre 1,5 m e 4,5 m dele (`arc innerradius 6.5 outerradius 9.5 angle 35`, com centro 5 m atrás). Quem fica colado ou longe escapa.

**Uma régua para converter metros em casas** (proposta, decida em ①): **1 casa ≈ 2,5 m**, que é o alcance corpo a corpo do Albion (espada 3 m, mobs 2–2,5 m). Isso bate com o movimento: 5,5 m/s × 2 s por ação ≈ 4,4 casas, perto das 5 casas base.

**O tapa convertido** (`MOB_HERETIC_BRAWLER_DOUBLESLAP_TUTORIAL`)

| Ponto | Albion | Conta | Eras |
|---|---|---|---|
| Custo | 1,8 s de conjuração; dano 12× o básico do Spencer (800 ÷ 66) | Passa do teto | **3 ações**, conjuração |
| Dano | 800 ÷ 1200 de HP do jogador = 67% | 67% × 40 = 27 | **6d8** (média 27) |
| Área | arco de 35°, 1,5–4,5 m | ÷ 2,5 → 0,6–1,8 | **Arco à frente, casas 1 a 2** (provisório até o grid) |
| Empurrão | 3 m | ÷ 2,5 = 1,2 | **1 casa** |
| Recarga | 5,5 s | ÷ 6 = 0,92 | **1 turno** (avisa num turno, cai no outro, e pode avisar de novo no seguinte) |
| Quando começa | `IfFightDuration ≥ 5` s | 5 ÷ 6 = 0,83 → 1 turno | **A partir do 2º turno dele** |

**Arquivos** (enviados; formato da opção A abaixo)

| Arquivo | O que tem | De onde vem |
|---|---|---|
| `weapons.json` | Pool `sword` (Q: `heroic_strike`, `cleave`; W e passiva vazios), espada T1 (básico 1d6, Força Bruta +1, E nulo), escudo T1 (tudo 0) | Etapa 4 |
| `armor.json` | Pool `common` com as 3 ativas T1 (todo material começa por elas), pool `leather` vazio, 3 peças (armadura 2 no peitoral, Astúcia +1 nas botas) | Etapa 4 + regra "1ª ativa igual em todos os materiais" |
| `enemies.json` | 3 mobs e 3 encontros com `RestSecondsBefore` | Etapa 4 + tapa |
| `player.json` | **O formato de equipamento que a função lê**: 5 slots, a ativa escolhida por slot, o tier da maestria por linha | — |
| `rules.json` | HP base 28, k 4, Defesa base 10, peso por material, maestria T1 +1, regen 1%/s | Constantes; o `AGENTS.md` pede constante em `data/` |

Notas rápidas:
- Os campos novos (`Cooldown`, `Area`, `Cast`, `Push`, `Range`, `MultiTargetBonus`) só entram no core quando a tarefa que os usa chegar. O núcleo mínimo lê só `Cost`, `HitModifier` e `Damage`.
- `Range`: o Scavenger é de distância no Albion (7 m → 3 casas). Os outros são corpo a corpo (1).
- O peso fica em `rules.json` por material, não na peça: é regra do `combate.md`.
- `HitModifier` não aparece nas skills do jogador, porque é a função que monta (atributo + maestria + item). Nos mobs, ele já vem pronto.

**Decisões antes da etapa 6**

**① Régua de casas.** Isso mexe na Esquiva da etapa 4.

| # | Opção | Porquê |
|---|---|---|
| A | 1 casa ≈ 2,5 m. A Esquiva passa de 4 para **2 casas** (5 m) sem ataque de oportunidade. Tapa nas casas 1–2, empurrão de 1 | Uma régua só para tudo; a Esquiva vira um "Passo de 2" (já está assim nos arquivos) |
| B | Régua só para áreas; a Esquiva fica em "casas de 1 ação" (4) | Mantém a etapa 4, mas são duas réguas |

**② A conjuração rola acerto?**

| # | Opção | Porquê |
|---|---|---|
| A | **Não rola.** Quem está na área no início do turno do conjurador leva o dano; sem crítico | Sair da área já é a defesa. Com crítico, 6d8 dobrado (54) mata qualquer um de uma vez |
| B | Rola 1d20 + acerto contra a Defesa, como um golpe | Usa só o fluxo que já existe, mas acerto + sair da área = defesa dupla, e o crítico mata |

**③ Formato dos pools**

| # | Opção | Porquê |
|---|---|---|
| A | Pools dentro de `weapons.json` e `armor.json` (como enviei) | Só os arquivos que você nomeou, mais o `rules.json` |
| B | Um `skills.json` com todas as skills (arma, armadura, mob); os itens só apontam IDs | Todo número de skill num lugar só; um arquivo a mais |

Me diga **① A ou B**, **② A ou B** e **③ A ou B**. Aí fecho a etapa 6 com todas as linhas dos docs de uma vez.

[Arquivo anexado: weapons.json]

[Arquivo anexado: armor.json]

[Arquivo anexado: enemies.json]

[Arquivo anexado: player.json]

[Arquivo anexado: rules.json]

---

## Sidinei:

① A: 1 casa ≈ 2,5 m para tudo. Esquiva = 2 casas sem ataque de oportunidade, recarga 3 turnos.
② A: conjuração não rola acerto e não tem crítico; quem está na área quando cai leva o dano. (Salvamento básico do PF2e fica estacionado.)
③ B: um skills.json com todas as skills (arma, armadura, mob); itens e mobs só apontam IDs. Refaça os arquivos nesse formato.

Regra extra: a recarga de skill com conjuração conta a partir do turno em que ela cai, não de quando foi usada. Confira se o tapa fica com 1 turno de respiro entre um e outro.

Siga para a etapa 6: todas as linhas dos docs de uma vez (arquivo e seção), incluindo as que você guardou nas etapas anteriores (custo por dano, recarga, regen, conjuração, régua de casas, as 4 regras promovidas das Propostas, a tarefa nova e o orçamento 29). Depois liste os JSONs finais.

---

## Claude:

**Ritmo do tapa** (a recarga de 1 turno conta a partir do turno em que ele cai)

| Turno do Spencer | Ele faz | Turno seguinte do jogador |
|---|---|---|
| 1 | Só ataques básicos (`FirstTurn: 2`) | Livre |
| 2 | Conjura o tapa (3 ações) e marca a área | **Sai da área** |
| 3 | O tapa cai no início do turno; a recarga conta aqui; depois, ataques básicos | **Respiro** |
| 4 | Conjura de novo | Sai da área |

✓ Fica 1 turno de respiro entre um tapa e outro.

---

**Etapa 6 · Linhas para os docs** (main `e317a5e`)

**`docs/jogo/mecanicas/combate.md`**

› **Turno**, depois de "Skills custam 1, 2 ou 3 ações…":
```
- Skill de dano custa ações em proporção ao dano: 1, 2 ou 3 (2 ≈ o dobro do ataque básico). Utilidade instantânea custa 1 ação.
- Recarga em turnos = recarga do Albion ÷ 6 s, arredondada (,5 sobe); abaixo de 3 s, sem recarga.
```
› **Turno**, troque a linha "Inimigo: até 4 ações…" por:
```
- Inimigo: até 4 ações por rodada, animadas juntas em até 2 s. A intenção dele aparece antes (casa-alvo), e o ataque comum sai na hora.
- Conjuração (de qualquer um): ao usar, marca a área; ela cai no início do próximo turno de quem conjurou. Quem sair da área escapa. Não rola acerto nem tem crítico: quem está na área leva o dano. Na fase 1, não pode ser interrompida.
- A recarga de skill com conjuração conta a partir do turno em que ela cai.
```
› **Entrada na luta**, troque a linha do HP por:
```
- O HP passa de uma luta para a outra. Fora da luta, o HP regenera 1% do HP máx. por segundo (como o Albion). A ativa do peitoral T1 cura 72% do HP máx., só fora da luta.
- HP do mob: vem pronto em data/, calibrado pelos turnos para matar; o Albion só dá a proporção entre mobs.
```
› **Movimento**, como primeira linha:
```
- 1 casa ≈ 2,5 m (o alcance corpo a corpo do Albion). Toda distância do Albion se converte por essa régua.
```
› **Em Aberto**: apague "HP do mob…", "Valores de base e k…" (agora estão em `data/rules.json`) e "Velocidade do regen…". Acrescente:
```
- Forma exata do tapa do mini boss no grid (provisório: arco de 35° à frente, casas 1 a 2).
```

**`docs/jogo/mecanicas/personagem.md`**

› **Atributos**, depois da linha do peso:
```
- Arma: dá o atributo ofensivo da linha. Botas: dão Astúcia.
- Defesa = 10 + Astúcia + armadura. A armadura vem do peitoral; no T1: pano +1, couro +2, placa +3.
```
› **Progressão**, depois de "Maestria: +1 a +4…":
```
- No T1, maestria +1.
```
› **Propostas**: troque a linha "Atributo por peça…" por `- Atributo por peça: peitoral = do material; capacete = Presença ou Sabedoria Ancestral.` e apague a linha `Defesa = 10 + Astúcia + armadura.`

**`docs/jogo/mecanicas/motor-de-regras.md`**, nova seção depois de **Defesa e CD**:
```
## Conversão do Albion

| # | Passo |
|---|---|
| 1 | Referência no mesmo tier: skill → ataque básico da arma; mob → média dos mobs comuns; utilidade → HP do jogador |
| 2 | Proporção = valor Albion ÷ referência Albion (mob: dano × velocidade de ataque) |
| 3 | Base do Eras: esta tabela de tiers + âncoras (básico T1 1d6; mob comum morre em 2 turnos; placa T1 +3) |
| 4 | Valor = base × proporção, inteiro, arredondado para o mais próximo (,5 sobe) |
| 5 | Dado: XdY+Z com média X×(Y+1)/2+Z perto do valor |
| 6 | Custo em ações = proporção do dano, de 1 a 3. Recarga, régua e conjuração: ver [combate](combate.md) |

- Sem equivalente (ignorar): energia, tempo de animação, armadura e resistência em número, Item Power, qualidade, ameaça, fama, prata e aggro.
```

**`docs/gestao/fase-atual.md`**
- Linha 3: `orçamento: 27 sessões` → `orçamento: 29 sessões`.
- Linha 7 e a última tarefa: `27` → `29` (marco final).
- A tarefa do `data/` passa a ser:
```
- [ ] data/: espada, escudo e couro T1 (3 peças), os inimigos dos 3 encontros, skills.json e rules.json + função equipamento → números (atributo + maestria + item; bônus do mesmo tipo não soma) · 2 sessões
```
- Depois da tarefa do Passo e do ataque de oportunidade:
```
- [ ] Conjuração e o tapa do mini boss: área marcada, cai no próximo turno, empurrão · 2 sessões
```

**`docs/gestao/04-roteiro.md`**, linha 22: `| 1 Luta no terminal | 27 |` → `| 29 |`, e `Marcos nas sessões 5, 10 e 27` → `5, 10 e 29`.

**Contradição que fica aberta** (`02-travessia.md` › Propostas): a tabela ainda diz "2 inimigos T1 + chefe T2". A Beira agora tem 3 encontros. Se quiser registrar até você dar os nomes, a linha para **Em aberto** é:
```
- Inimigos da Beira: 3 encontros da fase 1 (2 tipos T1 + mini boss T2); nomes a dar.
```

**Para o `/estacionar`**: mana · bônus de velocidade de movimento · resistência mágica · Item Power, qualidade e encantamento (fase 3) · imunidade a controle · salvamento básico do PF2e · interromper conjuração.

---

**JSONs finais** (enviados; todo ID referenciado existe em `skills.json`)

| Arquivo | Conteúdo |
|---|---|
| `skills.json` | 10 skills: básico, Heroic Strike, Cleave · mend, barrier, dodge (2 casas, recarga 3) · básico do Scavenger, do Thief e do Spencer · double_slap (`Cast`, recarga 1, 6d8, arco 1–2, empurra 1) |
| `weapons.json` | Pool `sword` (Q com 2 IDs no T1; W e passiva vazios) · `t1_sword` (Força Bruta +1, básico `sword_basic`, E nulo) · `t1_shield` (tudo 0) |
| `armor.json` | Pool `common` (as 3 ativas T1) · pool `leather` vazio · 3 peças (armadura 2 no peitoral, Astúcia +1 nas botas) |
| `enemies.json` | Scavenger 7 HP · Thief 13 · Spencer 39 (Def. 11) · 3 encontros com `RestSecondsBefore` 0 / 30 / 30 |
| `player.json` | O formato que a função equipamento → números lê: 5 slots, a ativa escolhida por slot, o tier da maestria |
| `rules.json` | HP base 28, k 4, Defesa base 10, peso por material, maestria T1 +1, regen 1%/s |

Nas skills do jogador, `HitModifier` é 0, porque a função soma atributo + maestria + item. Nos mobs, ele já vem pronto (+5 no T1, +6 no Spencer).

[Arquivo anexado: skills.json]

[Arquivo anexado: weapons.json]

[Arquivo anexado: armor.json]

[Arquivo anexado: enemies.json]

[Arquivo anexado: player.json]

[Arquivo anexado: rules.json]

---

## Sidinei:

Fechado. Duas correções que apliquei no repo:
- Na seção "Conversão do Albion" do motor-de-regras, "ignorar" ficou só com tempo de animação, fama, prata e aggro. Armadura virou parte da Defesa; mana, qualidade, Item Power e o resto estão no estacionamento, não ignorados.
- As skills do jogador não têm HitModifier no skills.json (a função calcula). Só os mobs têm.

---

## Claude:

Combinado, as duas correções estão certas:

- **Ignorar × estacionar:** estacionar quer dizer que volta depois. Ignorar quer dizer que nunca tem equivalente. Mana, qualidade e Item Power voltam, então ficam no estacionamento. Tempo de animação, fama, prata e aggro não voltam.
- **`HitModifier` só nos mobs:** é melhor assim. Um campo zerado nas skills do jogador dava a entender que existe um valor ali, quando na verdade é a função que calcula.

Os arquivos que te mandei ainda têm `"HitModifier": 0` nas skills do jogador. Vale o que está no seu repo.

Com isso, a tarefa do `data/` da fase 1 está com tudo decidido. A próxima é o estado mínimo do combate.

---
