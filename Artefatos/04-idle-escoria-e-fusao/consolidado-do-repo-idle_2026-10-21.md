# Eras do Brasil — Idle · Consolidado das decisões

> Gerado pelo Claude em 21/10/2026, a pedido do Sidinei, no mesmo dia em que o repositório foi trazido para o acervo. A cópia do repositório está em [`eras-do-brasil-idle/`](eras-do-brasil-idle/); os caminhos citados abaixo são relativos a ela.

> Retrato do repositório `sidinei-silva/eras-do-brasil-idle` (branch `main`, commit `cd7409e`), feito em 21/10/2026.
> É um **resumo navegável**, não uma nova fonte de verdade: cada seção aponta para onde a regra mora.

**Legenda**

| Marca | Significa |
|---|---|
| ✅ **Decisão** | registrada em doc normativo ou carregada em `data/` |
| 📐 **Projetado** | existe em `design/`, ainda não entra no runtime |
| ❓ **Em aberto** | o próprio repo diz que não está decidido |
| ⚠️ **Conflito** | duas fontes do repo dizem coisas diferentes (lista no fim) |
| 🗄️ **Descartado** | histórico, não reintroduzir |

---

## 1. O jogo em uma tela

✅ **MMORPG idle de fantasia folclórica brasileira, com a gramática de progressão do Albion Online.**
Fonte: `docs/o-jogo.md`

- **Você é o que veste** — não existe classe.
- **Fantasia central:** montar uma build com peças de séculos diferentes.
- O jogador **não controla movimento**: escolhe ação numa zona e observa.
- Recompensa presença, **não pune ausência**.
- Sobe **o que você faz**, não um nível global.
- **Sem rebirth**: progressão acumulativa.

```text
escolher zona → escolher ação → observar → recolher resultado → decidir de novo
```

Ações: **Coletar · Lutar · Viajar · Produzir** (produção só no hub).

---

## 2. Mundo

### 2.1 Estrutura física

✅ Fonte: `data/shared/continents.json`, `docs/decisoes-de-design.md`

```text
            ┌───────────────────────┐
            │     A TRAVESSIA       │  T1–T2 · só "segura"
            │  ilha tutorial, sem   │  não se volta a ela
            │       retorno         │
            └──────────┬────────────┘
                       │ embarcar (fim do tutorial)
                       ▼
┌──────────────────────────────────────────────┐
│                 AS ERAS                      │  T1–T6
│   eras lado a lado ao redor do               │  segura · disputada · mortal
│        CORAÇÃO DA RAIZ (centro)              │
└──────────────────────┬───────────────────────┘
                       │ passagem pelo centro
                       ▼
            ┌───────────────────────┐
            │     O EMARANHADO      │  T6–T8 · "selvagem"
            │  eras misturadas, sem │  PvP aberto
            │  hub fixo — pós-lanç. │
            └───────────────────────┘
```

| Continente | Tiers | Bandas de risco | Estado |
|---|---|---|---|
| A Travessia | 1–2 | segura | MVP |
| As Eras | 1–6 | segura, disputada, mortal | Era 1 projetada |
| O Emaranhado | 6–8 | selvagem | muito pós-lançamento |

### 2.2 Bandas de risco

✅ Fonte: `continents.json`

| Banda | PvP | Perda |
|---|---|---|
| Segura | não | — |
| Disputada | consentido | parcial |
| Mortal | consentido | total |
| Selvagem | **aberto** | total — só no Emaranhado |

### 2.3 Regras estruturais do mundo

✅ Fonte: `decisoes-de-design.md`

1. **Zona é a unidade.** Recurso, risco, facção, mob e viagem moram na zona. Cidade também é zona.
2. **Região é só rótulo de bioma.** Nenhuma regra lê `regionId`.
3. **Cinco campos independentes:** continente, era, região, tier e risco. Nenhum deriva do outro.
4. **Toda zona tem pelo menos dois caminhos** — senão vira corredor.
5. **Facção é por zona**, e nenhuma facção tem duas zonas seguidas no mesmo tier.
6. **A geografia não reproduz o mapa do Brasil** — é distorção da Ruptura.
7. **Coração da Raiz é passagem, não destino:** cidade segura no centro, cercada de zona mortal ⚠️ (ver conflito C1).
8. **Precisa de mais de uma era** para a fantasia de misturar equipamento funcionar.
9. **Viagem consome minutos reais** e faz parte do ritmo idle.

### 2.4 Anatomia de uma zona (modelo de dados)

✅ Fonte: `data/mvp/zones.json`, `dados-do-mvp.md`

```text
zona
 ├── tier, risk, factionId (quem controla), mobFactionId (o que nasce)
 ├── resources → materialId + abundância (primária/secundária)
 ├── camps     → mobIds + tamanho min/máx + elite + facção (nula = fauna)
 ├── bossId
 └── connections (fora da zona) → travelMinutes
```

- **Ferramenta exigida mora no material**, não no nó da zona.
- **Bestiário define o mob; o mapa define onde e em que grupo aparece.**
- **Gate não mora na conexão**: quem libera caminho é o tutorial (no futuro, missão).
- Fórmula herdada do Albion: **zona = fauna em qualquer lugar + uma facção de mob pelo bioma.**

---

## 3. Enredo e lore

✅ Regras de método (`decisoes-de-design.md`):
- **Lore é puxada, não empurrada** — escreve-se quando uma fatia precisa.
- **Mecânica não se poetiza.** Item e habilidade podem ter nome próprio; sistema não.
- Nome de sistema global não vem de era nem de facção.
- Texto de IA é mockup; o que sobrevive é reescrito à mão.

### 3.1 O que está decidido

| Elemento | Decisão |
|---|---|
| **A Raiz do Mundo** | Elemento estrutural do mundo; as eras ficam ao redor dela. |
| **O puxão** | A Raiz puxa gente **continuamente** — não há evento fundador único. Explica quem chegou muito antes do jogador. |
| **O jogador** | **Forasteiro**, um entre muitos. Sem origem fixa. **Não é o escolhido.** |
| **A Travessia** | Espaço de passagem (liminar). O nome cobre o continente, o lugar e o ato de embarcar. |
| **A Raiz no céu** | Durante toda a Travessia aparece no skybox, colossal e não navegável. Some ao ser expulso para o mundo (Feitoria da Cruz) e reaparece como o Coração da Raiz. |
| **Os Perdidos** | Gente puxada antes que não chegou do outro lado. Não falam, não reconhecem ninguém. Única facção de mob da ilha. |
| **O Primeiro Perdido** | Mini-chefe da Beira; chegou antes de todos. |

❓ **Em aberto de propósito:** a causa da Ruptura — sem data e sem evento fundador.
❓ **Reservado sem uso:** "O Ventre da Raiz" (instante do puxão antes da Travessia).

### 3.2 Personagens do tutorial

✅ `data/mvp/npcs.json`

| NPC | Papel |
|---|---|
| **Bartolomeu, o Língua** | recepção e tutorial na Beira |
| **A Barqueira** | mercado e travessia, no Porto |
| **O Ferramenteiro** | ferramentas, no Porto |

Falas-chave já escritas nos passos do tutorial (ex.: *"O que você trouxe de lá não vale nada aqui. Nem o nome."*).

### 3.3 Descartado

🗄️ Jogador escolhido · identidade especial por "Eco" · Dom da Revivência · Mata Costeira como tutorial · Vila de São Tomé · ritual de 1497.
🗄️ Tudo de **A Escória** que não foi adotado: Lastro, D20, classes, gank adjacente, zonas A Ressaca / A Bigorna / O Verde Surdo / A Costela, panteões ligados às armas.

---

## 4. Eras, regiões e zonas

### 4.1 Como os conceitos se encaixam

```text
CONTINENTE ─┬─ ERA (cluster de zonas, com facções próprias)
            │    └─ REGIÃO (rótulo de bioma, faixa de tier)
            │         └─ ZONA (unidade de gameplay)
            └─ TIER e RISCO são campos da zona, não derivados
```

- ✅ **Eras ficam lado a lado**, não em linha cronológica.
- ✅ **Era nova soma conteúdo nos sistemas existentes** — nunca traz árvore nova.
- ✅ **Facção política pertence à era**; a Era 2 terá outras.
- ✅ **Tradição de equipamento atravessa o jogo todo** (não confundir com facção).

### 4.2 A Travessia — MVP

✅ `data/mvp/zones.json`

```text
                 ┌──────────────┐
                 │   A BEIRA    │ T1 · spawn · mini-chefe
                 └──────┬───────┘
                        │ 6 min
                 ┌──────▼───────┐
                 │   PORTO DE   │ T1 · hub (mercado, estações)
                 │   PASSAGEM   │
                 └──┬────────┬──┘
              8 min │        │ 8 min
        ┌───────────▼┐      ┌▼────────────┐
        │   MATA     │10min │  PEDREIRA   │
        │  REVIRADA  ├──────┤   TORTA     │
        └────────────┘      └─────────────┘
           T2                  T2
```

| Zona | T | Função | Recursos | Acampamentos |
|---|---|---|---|---|
| **A Beira** | 1 | início, 1º combate | seixo, pau-mole, couro de capivara | Caranguejal, Cascalho solto · chefe: O Primeiro Perdido |
| **Porto de Passagem** | 1 | vila, mercado, estações | — | — |
| **Mata Revirada** | 2 | 1ª coleta T2 — "árvore de três eras no mesmo terreno" | pau-brasil, couro de veado, seixo | Copa revirada, Perdidos da mata |
| **Pedreira Torta** | 2 | 2ª coleta T2 | algodão-bravo, ferro-de-brejo, couro de capivara | Perdidos do corte, Boca da fenda |

Ilha toda **segura**: sem PvP, sem perda de equipamento.
Restrição de arte: as quatro zonas reaproveitam o kit de assets da Costa do Pau-Brasil.

### 4.3 Era 1 — Descoberta e Colonização (1500–1654)

📐 `design/era-1/` — projetada, **não é requisito do MVP**.

| Região | Bioma | Tier | Zonas |
|---|---|---|---|
| **Costa do Pau-Brasil** | mata atlântica litorânea | 1–2 | Feitoria da Cruz (hub), Praia de Coqueiros, Manguezal do Sal, Roça de Mandioca, Mata do Pau-Brasil, Aldeia Potiguara, Baixio das Garças |
| **Mata dos Engenhos** | mata atlântica | 2–3 | Canavial Velho, Engenho da Cruz, Mata Fechada, Olho de Boitatá, Serra da Barriga |
| **Sertão de Dentro** | caatinga | 4 | Caatinga Branca, Serrote das Pedras, Curral do Gado, Lajedo Rachado |
| **Rio das Almas** | várzea fluvial | 5 | Várzea Alagada, Porto de Canoas, Corredeira Negra |
| **Mata Sem Fim** | floresta amazônica | 6 | Castanhal Antigo, Igarapé de Anhangá, **A Raiz** (passagem p/ Emaranhado, exige T6) |

22 zonas, 44 conexões. Risco sobe de segura (T1–2) → disputada (T3–4) → mortal (T5–6).

**Facções da Era 1** 📐

| Facção | Apelido | Linha de armadura | Árvore preferida | Material exclusivo |
|---|---|---|---|---|
| **A Coroa** | colonizador | Pesada | Físico | ferro-de-brejo |
| **Os Potiguara e aliados** | indígena | Média | Projétil | couro de anta |
| **Os Encantados** | folclórico | Leve | Mágico | pedra de raiz |

O jogador **não pertence a nenhuma facção** — é forasteiro.
Os Encantados se apoiam em registros do séc. XVI (Anchieta, Gandavo): Curupira, Boitatá, Ipupiara, Anhangá, Jurupari.

---

## 5. Personagem, equipamento e identidade

### 5.1 Slots

✅ `o-jogo.md`, `trees.json`

```text
EQUIPAMENTO (7)          LOADOUT DE COMBATE
mão principal ──┐        6 ATIVOS            4 PASSIVOS
off-hand        │        ├ Ataque (Q)  arma  ├ passiva da arma
cabeça ─────────┼──────► ├ Utilidade (W) arma├ passiva da cabeça
torso           │        ├ Especial (E) arma ├ passiva do torso
botas           │        │   └ fixo pela arma└ passiva das botas
montaria        │        ├ Cabeça   (ativa)
bolsa ──────────┘        ├ Torso    (ativa)
                         └ Botas    (ativa)
```

- ✅ **O slot é a restrição**: cada slot só aceita habilidade da sua fonte. Não há lista geral nem "desabilitadas".
- ✅ Trocar uma peça troca uma ativa e uma passiva → **você é o que veste**.
- ✅ Depois de preencher, o jogador **ordena a prioridade** dos seis ativos.
- ✅ Montaria e bolsa também têm uma ativa e uma passiva próprias.

### 5.2 Árvores de habilidade

✅ `data/shared/trees.json` — **três, para sempre**:

| Árvore | Papéis |
|---|---|
| **Físico** | dano corpo a corpo · controle e impacto |
| **Projétil** | precisão · área e desgaste |
| **Mágico** | dano · suporte |

- Q, W e passivas = **linguagem compartilhada da árvore**.
- E = **assinatura da arma** (1:1); é o que diferencia duas armas da mesma árvore.
- Mãos: **uma mão** libera off-hand · **duas mãos** sem off-hand, mais IP · **off-hand** exige arma de uma mão da mesma árvore.

### 5.3 Tradições de equipamento

✅ `traditions.json` — atravessam todas as eras:

| Tradição | Linha de armadura | Identidade da armadura |
|---|---|---|
| **Colonial** | Pesada | armadura, autocura instantânea, resistir a controle |
| **Indígena** | Média | evasão, crítico, viagem, escape |
| **Encantada** | Leve | recurso, escudo, cura no tempo, sumir |
| *(nenhuma)* | Crua | o que dá para vestir com couro cru |

Bônus de conjunto: 3 peças da mesma linha.

### 5.4 Equipamento carregado (T1–T2)

✅ `weapons.json`, `armor.json`, `equipment.json`

| Tipo | Itens |
|---|---|
| Armas T1 | Lâmina Crua (físico, só Q) · Broquel Cru (off-hand, sem habilidade) |
| Armas T2 | Espada de lado (Colonial/Físico) · Zarabatana (Indígena/Projétil) · Cabaça de Boitatá (Encantada/Mágico) |
| Off-hands T2 | Rodela (Colonial) · Patuá (Encantada) |
| Armadura Crua T1 | Capuz, Colete, Botina |
| Pesada T2 | Morrião, Couraça, Grevas |
| Média T2 | Cocar, Escaupil, Alpercatas |
| Leve T2 | Carapuça, Manto de Retalhos, Sandálias de Piaçava |
| Montaria / bolsa | Mula · Bolsa de couro |
| Ferramentas | Machado, Picareta, Foice, Faca de esfolar |

📐 T3–T6 em `design/catalog/` (ex.: Tacape, Besta de mão, Maracá, Tocha de breu).

---

## 6. Combate

✅ Fontes: `decisoes-de-design.md`, `formulas-e-balanceamento.md`, `arquitetura-consolidada.md`

**Auto-battler, sem posicionamento.** "Área" = número de alvos, não forma.

```text
jogador escolhe acampamento
        ↓
nasce a ONDA FECHADA (grupo inteiro)
        ↓
servidor resolve a luta toda  ──►  linha do tempo pronta
        ↓                                 ↓
próxima onda só depois             cliente só anima
que a anterior morre
```

### 6.1 Regras

- **Sem turno e sem tick fixo:** dois relógios — ataque básico (pela velocidade da arma) e habilidades (pela recarga). O próximo evento é o menor dos dois.
- **Ataque básico:** sempre ativo, não ocupa slot, não custa energia, mais fraco que habilidade.
- **Energia:** recurso único das três árvores. Barra e regeneração seguem a curva de IP.
- **Prioridade:** dispara a habilidade pronta de maior prioridade com energia suficiente.
- **Ordem de checagem a cada evento:** parada → retirada → cura (poção por limiar de vida, fora da prioridade) → habilidade.
- **Foco:** o jogador ataca um alvo por vez; os inimigos vivos atacam.
- **Teto de 50 rodadas** por onda → **impasse**: ninguém morre, sem loot, volta ao terreno.
- **Parar no meio:** o servidor corta a linha do tempo; o que não foi animado não aconteceu.

### 6.2 Slots e custos base

| Slot | Recarga | Custo |
|---|---|---|
| Ataque (Q) | 5 s | 8 |
| Utilidade (W) | 15 s | 15 |
| Especial (E) | 25 s | 25 |
| Armadura | 35 s | 20 |

❓ Esses valores, velocidades, regeneração e o catálogo de habilidades são **`guessedValues`** — não validados.

### 6.3 Mobs

| Tipo | HP | Dano | Anti-fuga |
|---|---|---|---|
| normal (fauna) | 0,35× | 0,45× | 0,5 |
| normal (humano) | 0,35× | 0,45× | 1,5 |
| elite | 1,5× | 0,85× | 2,5 |
| chefe | 5,0× | 1,15× | 4,0 |

Atributos derivam de um loadout equivalente ao tier do mob.

---

## 7. Poder de Item e escala

✅ `formulas-e-balanceamento.md` · números em `data/shared/balance/`

```text
IP(peça) = base(tier) + encantamento + qualidade
multiplicadorIP = 1.0918 ^ ((IP − 100) / 100)      → T1 = 1,00× · T6 = 2,41×
```

| Tier | 1 | 2 | 3 | 4 | 5 | 6 |
|---|---|---|---|---|---|---|
| IP base | 100 | 200 | 350 | 550 | 800 | 1100 |

- Encantamento +1…+4 = +80 / 170 / 270 / 380 · Qualidade = +20 / 50 / 90 / 140.
- **Sobreposição intencional:** T4+3 (820) > T5 limpo (800). Tier não domina sozinho.
- IP do loadout: arma 0,35 · off 0,10 · cabeça 0,15 · torso 0,25 · botas 0,15 (duas mãos = 0,45).
- **Não existe AP** separado. Custo de energia escala com a arma; recarga não.

```text
hp       = 600 × mult
armadura = 240 × mult médio (cabeça, torso, botas)
mitig.   = ARM / (ARM + 200 + 2·IP_atacante), teto 75%
fatorRel = clamp(1 + (IP_atk − IP_alvo)/2000, 0.5, 1.8)
dano     = poderSkill × mult × (1 − mitig.) × fatorRel
```

❓ Teto de Poder de Item: reservado, não implementar agora.

---

## 8. Progressão

### 8.1 Árvore do Destino

✅ `destiny.json`, `formulas-e-balanceamento.md`

```text
                  [ Forasteiro ]  ← nó central, fama de qualquer ação
       ┌─────────────┬──┴──────────┬─────────────┐
   COMBATE        COLETA        REFINO         CRAFT
  1 nó por       1 nó por      1 nó por      1 nó por
  linha de       tipo de       tipo de       família
  arma/armad.    recurso       refinado      de item
     ↓              ↓             ↓             ↓
  Poder de      tier e        retorno de    chance de
  Item          rendimento    material      qualidade
```

- Quatro ramos **independentes** desde o começo.
- No MVP só os quatro nós de **aprendiz** são alcançáveis.
- Coleta, refino e craft sobem mais rápido (multiplicadores 0,35 · 0,15 · 0,20 sobre a curva).
- **Maestria entra no Poder de Item** → mesmo tier, desempenho diferente.

### 8.2 Fama

```text
Fama(N→N+1) ≈ base × 6^(N−1)        ⚠️ base 3000 no doc, 3600 nos dados (C3)
mob = 20 × 2,5^(N−1) · coleta = 15 × 2,5^(N−1) · refino = 10 × 2,5^(N−1)
```

### 8.3 Afinidade

✅ Multiplica a fama de 1,0 a 1,5 em 60 ações com a mesma arma. Congela offline, reseta só na troca. É custo de oportunidade, não punição.
❓ A regra de reset pode ser revista por causa da maestria por arma.

### 8.4 Especialização ampla

✅ Investir numa arma não pode tornar a troca para outra do mesmo arquétipo um recomeço do zero.
❓ Forma exata e valores não definidos.

---

## 9. Economia e produção

✅ `formulas-e-balanceamento.md`

```text
coletar ──► refinar ──► fabricar ──► equipar ──► gastar durabilidade ──► consertar
   ▲                                                                       │
   └────────────────────── prata e mercado ◄───────────────────────────────┘
```

- **Cascata de refino:** `1 refinado T(N) = 3 brutos T(N) + 1 refinado T(N−1)` → tier baixo nunca morre.
- **Ferramentas:** no máximo um tier abaixo do recurso; qualidade acima de normal exige ferramenta do mesmo tier.
- **T1 se coleta sem ferramenta** (pedra, madeira, couro).
- **Essência** não se coleta: cai de mob.
- **Estações têm tier próprio** e limitam o que se produz:
  - Refino: Fundição, Curtume, Tecelagem, Serraria, Canteiro
  - Craft: Forja, Posto de Caça, Terreiro, Ferramentaria

**Durabilidade** — desgasta por uso, morte (principal) e fuga fracassada:

| Faixa | Efeito |
|---|---|
| 100–50% | nenhum |
| 49–25% | −100 IP |
| 24–10% | −200 IP |
| < 10% | inutilizável |

Ferramentas não perdem IP.

❓ Fluxo de prata, custo de conserto, custo por item e tier: dependem do simulador novo.

---

## 10. Risco e PvP

✅ `decisoes-de-design.md`

| Regra | Decisão |
|---|---|
| Natureza | **consentido e mútuo** |
| Flag | liga e desliga **só na cidade** |
| Não sinalizado | **não pode ser atacado** |
| Incentivo | **artefato bruto só cai para sinalizados** |
| Contador | cada zona mostra quantos sinalizados há |
| Faro | mostra **quem** são |
| Encontro | dentro da própria zona, sem pareamento entre zonas |
| Consequência | definida pela banda de risco |

❓ **Reputação adiada.** Destino: purificação de artefatos e facções. Se até T4 nada depender dela, não entra antes disso.

---

## 11. Tutorial — A Travessia passo a passo

✅ `data/mvp/tutorial.json` (máquina de estados em dado, não em código)

Diferença deliberada do Albion: **o jogador coleta e fabrica tudo.** Só a bolsa vem do mini-chefe e só a mula é comprada.

| # | Passo | Ensina |
|---|---|---|
| 1 | Acordar sem nada | onde está |
| 2 | Pedra e pau | nó de recurso, T1 sem ferramenta |
| 3 | A primeira lâmina | craft, equipar, slot de Ataque |
| 4 | Primeiro combate | auto-battler, onda fechada |
| 5 | Couro de bicho | esfolar mob morto |
| 6 | Vestido do jeito que dá | armadura = 1 ativa + 1 passiva |
| 7 | O broquel | off-hand só com arma de uma mão |
| 8 | O que não passou | mini-chefe → **bolsa**; libera o Porto |
| 9 | O Porto | vila, estações, armazém |
| 10 | A mula | mercado, montaria, carga |
| 11 | Quatro ferramentas | ferramenta e regra de tier |
| 12 | Mata Revirada | tier de recurso, viagem real |
| 13 | Pedreira Torta | peso da carga |
| 14 | Refinar | 4 estações e a cascata |
| 15 | A Árvore do Destino | 4 nós de aprendiz |
| 16 | A escolha | **tradição**, Q/W/E e **prioridade** |
| 17 | Você é o que veste | conjunto T2 e passivas |
| 18 | A travessia | embarcar — ilha não volta |

Cobertura: o MVP exercita inventário, combate, onda, slots, energia, coleta, ferramentas, craft, armadura, mini-chefe, mercado, montaria, refino, Destino, tradição e conjunto.
**Fora do MVP:** PvP, T3+, artefatos, encantamento/Essência, qualidade, reputação, economia entre jogadores, mundo da Era 1.

---

## 12. Arquitetura técnica

✅ `docs/arquitetura-consolidada.md` (normativo) — **definitiva em escala reduzida**

```text
HTTP: conta · login · refresh · personagem · entrar no mundo
                                        │
                                        ▼
WS 1 ─┐                           1 WebSocket por jogador
WS 2 ─┼─► commands chan ─► GAME LOOP ─► GameState ─► eventos ─► tópicos (zona/guilda/global)
WS 3 ─┘                     ▲
                     ticker ┘  (só pergunta: now >= EndsAt?)
```

| # | Decisão |
|---|---|
| 1 | **Single-owner:** uma goroutine possui o `GameState` |
| 2 | Comando e tick entram no mesmo loop por portas diferentes; engine recebe `now` |
| 3 | `data → gamedata → bootstrap → game`; nenhuma camada invade a outra |
| 4 | `bootstrap` é o composition root; `main` quase vazio |
| 5 | Conta não é personagem; conta fica fora do Game Core |
| 6 | HTTP até entrar no mundo, WebSocket depois (decisão de transporte, reversível) |
| 7 | **Servidor nunca manda posição** — layout e animação são do cliente |
| 8 | Combate em onda fechada, linha do tempo inteira de uma vez |
| 9 | Tópicos, não salas |
| 10 | Persistir em pontos significativos, nunca por tick |
| 11 | Go idiomático: interface no consumidor, tipo perto do conceito, sem camadas globais |

Princípios: servidor autoritativo · Command → State → Event · tempo é `Activity{StartedAt, EndsAt}` · ownership antes de mutex/channel.

🗄️ Descartado agora: mutex como dono do estado, actor por jogador/zona, microserviços, Redis, Kafka/NATS, Kubernetes, sharding, event sourcing.

❓ Não decidido: particionamento, outbox/transação, formato de wire do WS, versionamento do conteúdo do cliente, reconexão no meio do combate, **personagem offline** (esboço com teto de ganho, decidir na primeira Activity de coleta).

**Estado do backlog:** fatias 1–3 (conta, login, personagem) concluídas; fatia 4 (entrar no mundo) ⚠️ status divergente (C4).

---

## 13. Direção de arte

✅ `docs/direcao-de-arte.md`

> **3D low poly + silhueta forte + cores chapadas de paleta + sombra simples + equipamento visível + animação curta + cenário modular**

| Item | Decisão |
|---|---|
| Estilo | 3D low poly estilizado, nunca realista |
| Engine / ferramenta | Godot · Blender · `.glb` |
| Plataforma | web primeiro |
| Câmera | isométrica fixa 3/4 |
| Cor | atlas de paleta 256×256, sem PBR nem normal map |
| Escala | 1 u = 1 m · personagem 1,8 u · grade 0,5 u · tile 2×2 |
| Orçamento | personagem ≤ 3.000 tris · peça 150–500 · chefe ≤ 4.000 |
| Esqueleto | humano compartilhado (jogador, NPCs, Perdidos) + sockets de equipamento |
| UI | 2D em blocos sobre a cena; cena é observacional |
| Evitar | castelo/torre medieval europeu, neon, bloom pesado |

❓ Proporção: 6 cabeças é a referência; chibi (2,5–3 cabeças) em avaliação desde 15/09/2026.

---

## 14. Como o jogo cresce

✅ Regras que já valem para conteúdo futuro.

### 14.1 Adicionar uma era

```text
era nova
 ├── zonas, regiões, mobs, NPCs, facções próprias   → conteúdo novo
 ├── armas novas nas TRÊS árvores existentes        → só o E é novo
 ├── peças nas linhas Pesada / Média / Leve         → mesmas tradições
 └── NÃO cria árvore, linha de armadura nem sistema paralelo
```

- Mesmo esquema de dados: `data/era1/zones.json`, `data/era2/zones.json`… O escopo é a pasta.
- Facção política é da era; tradição é global.
- Tier é por zona, então equipamento de eras diferentes convive e se mistura.

### 14.2 Promover conteúdo: `design/` → `data/`

```text
fatia do backlog pergunta: "que dado ESTA fatia precisa?"
        ↓
recorta só isso de design/  →  cola em data/  →  confere referências  →  apaga de design/
```

- `data/` só tem o que existe código para carregar.
- Exceção: `balance/` fica inteiro em `data/` (T1–T6), mesmo citando ids não carregados.
- Quando a Era 1 entrar: `design/era-1/` → `data/era1/` e catálogo T3–T6 → `data/shared/`.

### 14.3 Temporadas (pós-lançamento)

- **Expansão traz era; temporada traz evento.**
- Eventos deixam **mudança permanente** em zonas.
- **Ecos** podem transformar conteúdo passado em instâncias repetíveis.
- **Sem reset.** Sem data anunciada antes de haver base.

### 14.4 Sistemas adiados ou cortados

| Situação | Itens |
|---|---|
| Adiado | rotinas de NPC, necessidades, `knowledgeBase`, fofoca, dia/noite, maré, reputação, teto de IP, temporadas |
| Possível | migração de mobs dentro da faixa de tier da zona |
| Cortado | companheiros, mercenários, rebirth, receita desbloqueada por NPC, D20, grid isométrico, posicionamento/cobertura, modo RPG de mesa, 18 mini-campanhas, variáveis procedurais por save |

---

## 15. Conflitos encontrados no repo

Não resolvi nenhum — só aponto.

| # | Onde | Divergência |
|---|---|---|
| **C1** | `decisoes-de-design.md` × `design/era-1/zones.json` | Decisão: Coração da Raiz é **cidade segura** no centro e leva ao "segundo continente". Dados: a zona **A Raiz** é T6, **mortal**, sem hub, dentro da Mata Sem Fim, com passagem para o Emaranhado — que é o **terceiro** continente, já que a Travessia conta como um. |
| **C2** | `decisoes-de-design.md` × `design/era-1/zones.json` | Decisão diz que a Costa do Pau-Brasil **pode começar em T2**; os dados põem a Costa em T1–2, com Feitoria da Cruz e mais três zonas em T1. |
| **C3** | `formulas-e-balanceamento.md` × `balance/progression.json` e `destiny.json` | Fama por nível: **3000** × 6^(N−1) no doc; **3600** nos dados (e 3600 nos nós de aprendiz). |
| **C4** | `BACKLOG.md` × `04-entrar-no-mundo.md` | Fatia 4 aparece como **iniciada** na tabela e **não iniciada** no texto e no arquivo da fatia. |
| **C5** | `o-jogo.md` × `balance/` | Tier "de 1 a 8" no texto; balanceamento cobre só T1–T6 (T7–8 seria o Emaranhado, sem números). Lacuna, não contradição direta. |
| **C6** | `arquitetura-consolidada.md` | A árvore de estrutura cita `data/zones.json, dialogs.json, objectives.json` — formato antigo; hoje é `data/shared/`, `data/mvp/`. |
| **C7** | `design/era-1/factions.json` | Nota diz que "no MVP a Aldeia Potiguara usa um booleano" — o MVP hoje é a Travessia, sem a Aldeia. Nota desatualizada. |
