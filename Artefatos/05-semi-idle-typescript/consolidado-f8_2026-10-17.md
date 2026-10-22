<!-- Consolidado F8 · arquivo original: eras-consolidado-f8.md · escrito pelo Claude em 2026-10-17T14:57 na conversa "Ansiedade no desenvolvimento do Eras do Brasil" (export de 21/10/2026) -->

# Eras do Brasil — Consolidado F8 (até 14/10/2026)

> O que estava decidido no "Eras novo" (F8, 03/10 a 14/10/2026), antes da declaração de
> indefinição de 14/10.
> Fontes: `docs/` do repo `eras-do-brasil` (commit `5658b2c`), os rascunhos de afk/pvp/risco,
> o catálogo `MECANICAS-E-CONTEUDOS.md` e as conversas com o Claude de 03 a 14/10.
>
> Legenda:
> - **Decidido**: estava registrado como decisão.
> - **Rascunho**: em teste no playground, não era decisão.
> - **Aberto**: pergunta sem resposta.
> - ⚠️ marca contradição entre fontes.

---

## 1. O jogo

- MMORPG de turno, híbrido entre ativo e idle. Brasil histórico, folclore e cultura popular.
- **Público:** jogador cansado, com pouco tempo, que pode ser interrompido a qualquer momento.
- **Pilar central:** fácil de entender, profundo para quem quiser.
  - Todo sistema tem um padrão que funciona sem configurar nada.
  - Ninguém precisa de wiki para progredir.
  - A luta é leve de executar.
- **Pilares:**
  - Economia e mercado relevantes.
  - Life skills que importam para a progressão.
  - Craft relevante.
  - Poucas skills na barra, muitas no jogo.
  - Sem classe presa: você é o que veste, com builds híbridas.
  - MMO jogável solo sem penalidade.
- **Não é:**
  - RPG de classe fixa.
  - Jogo que exige wiki.
  - Jogo complicado.
  - Idle puro, nem jogo que pune 1 segundo longe do teclado.
  - Jogo que obriga a jogar em grupo.
- **Princípio do híbrido:** o AFK entrega volume, o ativo entrega qualidade.

### Loops
- **Minuto:** ver inimigo, entrar em combate, ler a intenção, escolher ação, ganhar fama, recursos e itens.
- **Sessão:** preparar build e automação, ir para a zona, caçar ou coletar, carga enche e durabilidade cai, reparar com kit ou voltar à cidade, vender, refinar, craftar, ficar mais forte.
- **Longo prazo:** árvore do destino, subir tier, encantar, virar referência como crafter, zonas mais perigosas, PvP, ocupar espaço na economia.

---

## 2. Personagem e equipamento

- **Decidido:** sem classe. Você é o que veste.
  - A peça dá as skills. Trocar a peça tira as skills dela.
- **Decidido:** sem classe presa e sem alts. Trocar de caminho tem custo, mas nunca joga o investimento fora.
- **Decidido:** barra com 6 skills + ataque normal.
  - 3 da arma principal.
  - 1 do torso.
  - 1 da cabeça.
  - 1 do calçado.
- **Decidido:** atributos derivados do equipamento (vida, energia, resistências).
- **Decidido:** equipamento aparece no personagem. Foi o argumento central para o 3D modular.
- **Tiers T1–T8** (vindo de F6–F7, mantido):
  - T1 é sucata sem dono e não é craftável.
  - T2 são as três tradições.
  - T4 é a base do PvP real.
- **Encantamento até +4, por duas portas:**
  - Craft com material encantado: garantido e mais caro.
  - Encantar item existente: risco visível; falhar reduz o teto; só até +3.
  - O +4 é exclusivo de crafter.
- **Aberto:** qual atributo define a ordem dos turnos (iniciativa).

---

## 3. Durabilidade

- **Decidido:** duas camadas, como no BDO.
  - **Atual:** cai e se repara.
  - **Teto:** o máximo que o reparo alcança.
  - Teto 0 não destrói o item.
- **O que gasta a atual:** ser derrubado e a fuga que falha. Arma não desgasta com o uso comum, igual no Albion.
- **O que derruba o teto:** morte no ativo e falha de encantamento.
  - Sua frase: *"Só não posso deixar o teto descer rápido se durabilidade for zero."*
- **Decidido (10-06):** a fuga que falha custa durabilidade, não teto. O teto fica reservado para morte e upgrade.
- **Kit de reparo por tipo:** recarregável em campo com material refinado.
- **"Descarte de luxo":** item raro ruim vira material para restaurar o teto.
- ⚠️ O catálogo e o rascunho de `risco.md` dizem que a morte no ativo **zera** o teto. Isso conflita com a frase acima. Provável intenção: zerar a atual e reduzir o teto.
- Estudo de 16/10, não decidido: o teto é o ralo de itens que substitui o full loot.
  - **Alavancas:** vida útil por tier; cada reparo tirar um pouco do teto; restaurar caro e reparar barato; o AFK pagar a sua parte; NPC comprador com piso enquanto a população for baixa.

---

## 4. Combate

### Entrada e arena
- **NPC no ativo:** clicar no grupo de inimigos visível no mapa abre a luta.
- **Arena:** um recorte instanciado do mapa. No mundo aberto, os outros só veem um ícone de combate.
- **Três filas** (uma central e duas laterais):
  - O jogador começa na central, à esquerda.
  - Os inimigos vêm da direita.
  - Elite e chefe vêm sempre pela central.
  - O spawn dos demais é definido pelo designer.
- **Rascunho:** campo de 9 passos.
  - `[sobra 2] [jogador] — distância 3 — [mobs da frente] [mobs de trás] [sobra 2]`
  - Ninguém nasce na borda, para sobrar espaço de empurrão e recuo.
  - A posição inicial vem dos dados (formação por arquétipo).
- **Aberto (14/10):** 1 fila ou 3 filas. A ideia era decidir no checkpoint, jogando.

### Formato
- **Turno contra NPC:** sem tempo limite.
- **Turno no PvP:** 30 s. Se estourar, o auto assume sem perder a vez.
- **Ordem:** iniciativa como peso num sorteio. Grupos de inimigos são sorteados separadamente.
- **Estrutura do turno** (sua proposta de 10/10): **andar + item + ação de combate**. Andar não gasta o turno.
  - Proposto e não confirmado: ordem fixa andar → item → combate; o item com recarga própria para não virar cura grátis.
  - Fuga e skill de movimento gastam a ação de combate.
  - O `combate.md` diz: 1 item por turno, com recarga.

### Movimento
- **Decidido:** o movimento comum é automático, 1 passo por turno.
  - O melee anda até ter o alvo ao alcance.
- **Decidido:** distância = diferença de passos. Fila vizinha +0; fila oposta +1.
- **Decidido:** quem está na fila bloqueia quem anda nela. O mob bloqueia, e isso dá função ao tank.
- **Decidido:** só skill de movimento muda de fila, e ela não é bloqueada.
  - Exceções: a fuga com fila fechada e o frente a frente (este, rascunho).
- **Decidido:** arma à distância perde dano quando está colada no alvo.
- **Skill de movimento:** o jogador escolhe o alvo, não o caminho, e acerta quem estiver no trajeto. Onde termina:
  1. Ao lado do alvo, na fila dele, do lado de quem atacou.
  2. Se ocupado, no espaço livre mais próximo que mantenha o alvo no alcance.
  3. Se não houver, não pode ser usada naquele alvo.
  - Desempate: fila do alvo, depois fila vizinha mais perto da origem, depois passo mais perto de quem atacou.
- **Rascunho do playground (13–14/10):**
  - O ranged anda até o limite do alcance e foge quando fica perto demais.
    - Margem = alcance − 2.
    - Alcances: curto 1, médio 3, longo 5.
  - Na parede, vira se do outro lado não houver mais inimigos capazes de atacá-lo. *(Decidido 10-13.)*
  - Fila fechada: foge pela diagonal para uma fila vizinha. *(Decidido 10-13.)*
  - Frente a frente: o melee entra na fila do alvo ou na diagonal da frente. Trocar de fila não gasta o turno.
  - Isso contradiz a regra antiga "ranged não anda, ninguém recua andando", que foi testada de propósito.
- **Aberto:**
  - Posição em grupo (proposta: um por fila).
  - Quanto dano o ranged perde quando está colado.
  - Onde o alcance longo passa a valer.
  - Os ataques grátis do ranged contra os turnos que ele leva para matar.
  - Disputa de casa no frente a frente.

### Habilidades e ações
- **Recarga:** cada skill tem a sua, que cai a cada turno. Passivas podem reduzir. O ataque normal não tem recarga.
- **Item:** cura, buff ou debuff, com recarga própria.
- **Intenção do inimigo:** visível só no ativo PvE. É a informação que paga o risco maior do ativo.
- **Mob:** executa a intenção que anunciou. Se ficar impossível, decide de novo. Prefere ser previsível a ser esperto.

### Fuga
- É a ação do turno. Não garante escapar: é escapar pagando menos que morrer.
- **Chance:** fuga contra anti-fuga, com piso e teto.
  - A fuga vem de equipamento, montaria e habilidades.
  - A anti-fuga vem do tipo de inimigo e das habilidades dele.
  - Mais inimigos vivos dificultam. Estar longe facilita; estar colado num melee dificulta.
- **Falhou:** perde o turno, os inimigos agem e a durabilidade cai.
- **Conseguiu:** leva um golpe de despedida de quem o alcança, perde o loot da luta e fica protegido de qualquer luta por um tempo.
- **Cada árvore foge do seu jeito:**
  - Projétil: aumenta a chance.
  - Físico: reduz o custo.
  - Mágico: ignora a posição.
  - Calçado, torso e montaria somam fuga para qualquer árvore.
- Alguns inimigos anulam, dificultam ou impedem a fuga.
- **Decidido (10-06):**
  - A chance não sobe a cada falha.
  - A anti-fuga cresce com o número de inimigos.
  - Tentar fugir é a ação do turno.
  - Todas as árvores fogem.
  - Quem foge fica protegido por um tempo.
- **Aberto:**
  - "Preso" impede a fuga ou só dificulta?
  - Como evitar que aproximar e prender trave a fuga para sempre?
  - A fuga precisa de recarga entre tentativas?

### Ativo × automatizado
- **Decidido:** qualquer ação manual (inclusive tentar fugir) faz a luta contar como ativa no preço da morte. A interface avisa.
- **Decidido:** o botão de pânico foi removido. No PvE o turno já espera; no PvP o auto assume.

---

## 5. Automação

- **Termos:**
  - **Build:** o equipamento vestido.
  - **Automação:** o sistema que luta pelo jogador.
  - **Preset:** a configuração por tipo de luta.
- **Skills:** lista de prioridade. Usa a primeira fora de recarga; se todas estiverem em recarga, usa o ataque normal.
- **Alvo:** mais perto, mais longe, menos vida, mais vida, mais forte ou mais fraco.
- **Itens:** lista de prioridade, respeitando a recarga.
- **Fuga:** por limiar de vida. Ordem de verificação: fuga, cura, habilidade.
- **Presets:** comum, chefe e PvP, escolhidos pelo jogo. Vêm iguais por padrão.
- **Decidido (10-07):**
  - A automação não enxerga a intenção.
  - Um preset por tipo de luta.
  - Os presets vêm iguais por padrão.
  - "Loadout" saiu do vocabulário.
- **Decidido:** o mob usa o mesmo decisor do jogador (estado + regras + rng → ação), com regras em `data/`.
  - Arquétipos: bruto, atirador, curandeiro e tank.
  - Chefe tem fases.
- **Aberto:**
  - Elite usa o preset de comum ou de chefe?
  - Condição por skill baseada em vida.
  - "Cura" é item, skill ou os dois?
  - O que acontece com uma skill que deixou de existir.

---

## 6. AFK (caçada automatizada)

- O jogador fala com o NPC de caçada da zona e escolhe a caçada automatizada. Os encontros vêm de tempos em tempos, com inimigos únicos ou grupos, conforme a configuração.
- **AFK faz de tudo, mas reduzido:** menos recurso e drop raro bem menor que no ativo.
- **Sem chefes** por enquanto.
- **Não consome** os encontros visíveis de quem está no ativo.
- **Peso e carga limitam a caçada.** Parar significa parar, sem teleporte de graça.
- **Assumir o controle:** fica ativo naquela luta e, no fim, volta ao mundo no ativo. Para voltar ao auto, precisa ir de novo ao NPC.
- **Idle com presença:** não progride com o jogo fechado. Offline congela e não decai.
- **Aberto:** condições de parada do auto (inventário cheio, vida, durabilidade) e caçada de chefes.

---

## 7. Morte e risco

- **Decidido:** a morte no auto é mais barata. Protege quem não estava presente para reagir.
- **Decidido:** o ativo pune mais a morte e compensa com fuga melhor e com a intenção visível.
- **Preço:**
  - Ativo: atinge o teto (⚠️ "zera", ver seção 3).
  - Auto: só desgasta.
- **Inventário:** perde parte do que carrega, com limite. Substitui o full loot.
- **Aberto:** zonas por risco.

---

## 8. PvP

- **Ativo contra ativo:** em zona não pacífica, com PvP ligado, clica no outro jogador.
  - O atacado recebe um aviso (não um convite) e escolhe ativo ou auto. Com a contagem esgotada, entra no auto.
  - O atacante sempre entra no ativo.
- **Auto contra auto:** jogadores em caçada na mesma zona não pacífica podem se enfrentar. Nessa luta não há troca para o ativo.
- **Ativo nunca ataca quem está em caçada automática,** e vice-versa. O atacado nunca fica sem defesa; ausência não é derrota.
- **Fuga:**
  - O atacado pode fugir já no 1º turno.
  - O atacante não pode fugir nos primeiros turnos.
  - O preset de PvP pode pôr a fuga como primeira ação.
- **"Se atacar e perder, perde mais":** risco assimétrico do atacante.
- **Atacar quem já luta contra mob:** a vida fica como está, e a luta contra o mob termina sem loot.
- **Marcação de quem mata, lista de vingança, janela de proteção** para o atacado que venceu (contra o abutre).
- **Aberto:**
  - Vantagem do atacante.
  - Detalhes da marcação.
  - Escolher "ativo" no aviso, sem agir, conta como ação manual?
- **Ideia considerada, não decidida:** PvP assíncrono com um "eco" (cópia da build do jogador). O desafiado só recebe o resultado, como no Lootborne.

---

## 9. Mundo

- O servidor só conhece zonas e pontos de interesse. O caminho é do cliente.
- Grupos de mobs visíveis no mapa, como no Dofus.
- Mobs comuns são pessoais; elites e chefes são compartilhados.
- Sem tempo de viagem cronometrado no ativo.
- Mundo e encontros iam para um `mundo.md`, que nunca foi escrito.

---

## 10. Arquitetura (decidido 10-09)

- **Servidor autoritativo em TypeScript.** Regras puras num `core` compartilhado. Cliente web, empacotado para a Steam via Electron.
- **Princípios:**
  - O servidor decide e o cliente mostra.
  - Regra é função pura: estado + ação → estado + eventos.
  - RNG com semente: mesma semente + mesmas ações = mesma luta.
  - Constantes em `data/`.
- **Monorepo:**
  - `packages/core`: regras, sem rede, banco ou arquivo.
  - `apps/server`: Fastify + ws + Zod + Postgres.
  - `apps/client`: Three.js + HTML/CSS.
  - `apps/playground`.
  - `data/`.
- **Estado:**
  - Estático em `data/`.
  - Persistente no Postgres, que é a fonte da verdade.
  - Execução em memória, num único event loop.
  - Economia: banco primeiro, memória depois.
- **Sem** Redis, Colyseus ou framework pesado.
- **Reabrir:**
  - TypeScript: após as fatias 1–4, com uma lista escrita de travas.
  - Cliente web: se o personagem modular no Three.js falhar.
- **Não decidido de propósito:**
  - Formato das mensagens.
  - Framework de UI.
  - Reconexão no meio da luta.
  - Ganho offline.
  - Escala.

---

## 11. Direção de arte (decidido 10-09)

- 3D low poly estilizado, cor chapada de atlas de paleta, sem PBR e sem normal map.
- Câmera 3/4 fixa. Equipamento visível por sockets:
  - mão principal e mão secundária;
  - cabeça, torso e pés;
  - costas.
- Brasil histórico e folclórico. Nada de castelo, goblin ou orc.
- Interface em HTML/CSS sobre o canvas.
- **No combate:**
  - Faixa com a ordem de turno.
  - Intenção do inimigo com cartão do alvo.
  - 6 skills + ataque normal + item + fugir.
  - Seletor manual/auto.
  - Diário curto.
  - Sem inventário completo.
- **Produção:** Blender → glb → Three.js; animação base do Mixamo; pacote modular comprado vale como base.
- **Primeiro teste:** personagem base + 2 peças trocáveis + 1 animação do Mixamo.
- **Aberto:**
  - Proporção (~6 cabeças ou chibi).
  - Linguagem do ícone de intenção.
  - Como a fila aparece no chão.

---

## 12. Processo

- **Fatias:**
  1. 1v1 numa fila.
  2. Playground.
  3. Skills, energia e recarga.
  4. Intenção.
  5. Auto.
  6. 3 filas e área.
  7. Movimento forçado.
  8. Fuga.
  9. Atributos e fórmulas.
- **Checkpoint** depois de ativo + auto + intenção: jogar noites de verdade e responder "gosto do híbrido?".
- O doc segue o código. Ideias novas vão para o `ideias.md` até o checkpoint.
  - Já anotada lá: IA com gatilho no estilo Fire Emblem.
- **Estado do código em 14/10:**
  - Core TS com spawn, posição/movimento e luta (`initializeCombat`, `resolveAction`, `simulateCombat`).
  - 41 testes no Vitest.
  - Playground React com depurador de combate.
  - Parte foi escrita com Claude Code.
  - JSONs do `eras-do-brasil-idle` em `data/` e `design/`, ainda sem uso.

---

## 13. Depois do reinício (14–16/10) — fora do consolidado F8

- **14/10:** jogo declarado totalmente indefinido. É o mesmo ciclo do reset de 03/04.
- **Norte em rascunho** (palavras do Sidinei, 2023 + 14/10):
  - Turno.
  - Sua classe é o que você usa: troque quando quiser, sem alt.
  - O seu trabalho importa no mundo: life skill relevante, com mercado vivo.
  - Poder parar sem perder nada.
- **15/10:** projeto parado. Estudo solto de mecânicas, sem decidir nada.
