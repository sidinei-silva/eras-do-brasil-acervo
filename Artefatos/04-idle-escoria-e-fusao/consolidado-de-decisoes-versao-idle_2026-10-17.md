# Eras do Brasil — Consolidado de decisões

Oct 17, 2026 · @Sidinei

> Documento do Claude Docs escrito pelo Claude em 17/10/2026 na conversa "Consolidado do jogo", exportado pelo Sidinei em 21/10/2026. Resume as decisões da versão idle (set/2026) a partir das conversas.

## Escopo e fontes

Este documento consolida a **versão idle** do Eras do Brasil: servidor Go, cliente Godot, MVP na ilha da Travessia. Ele não reflete o recomeço de outubro de 2026 (TypeScript, camada ativa), que vive no repositório novo.

Fontes lidas, nesta ordem de autoridade:

1. Código e dados do repositório `eras-do-brasil-idle` (último commit em 26/09/2026): `data/`, `design/`.
2. `docs/o-jogo.md`, `decisoes-de-design.md`, `arquitetura-consolidada.md`, `dados-do-mvp.md`, `formulas-e-balanceamento.md`, `direcao-de-arte.md`.
3. `docs/historico-e-estudos.md`, só para o que foi descartado e por quê.
4. Conversas deste Project até 03/10/2026, para decisões fechadas em chat e já registradas no repo.

Convenção usada aqui: **Decisão** é regra escrita nos docs; **Dado** é o que está nos JSONs; **Em aberto** é o que os docs deixam de propósito sem resposta. Onde docs e dados discordam, a última seção lista a divergência em vez de escolher um lado.

## O jogo em uma página

MMORPG idle de fantasia folclórica brasileira, com a gramática de progressão do Albion Online. **Você é o que veste**: não há classe, e a build sai do equipamento.

**Fantasia central:** montar uma build com peças de séculos diferentes. Por isso o jogo precisa de mais de uma era (Decisão).

**O jogador não controla o movimento.** Escolhe uma ação numa zona e observa o personagem executá-la:

```text
escolher zona → escolher ação → observar → recolher o resultado → decidir de novo
```

| Ação | O que acontece |
| --- | --- |
| Coletar | o personagem vai até o recurso e entra no loop de coleta |
| Lutar | escolhe um acampamento e enfrenta uma onda fechada de inimigos |
| Viajar | a viagem consome minutos reais e leva à zona seguinte |
| Produzir | no hub, refina materiais e fabrica equipamento |

Princípios de ritmo: sessões curtas funcionam, sessões longas rendem mais. O jogo **recompensa presença e não pune ausência** — nas conversas isso foi esclarecido como regra da Afinidade (que só cresce conectado), não como regra de persistência.

Sobe **o que você faz**, não um nível global. Não há rebirth: a progressão é acumulativa e a variedade vem de outras árvores, eras e facções.

## Enredo e mundo

A lore é mínima de propósito: **lore é puxada, não empurrada** — escreve-se quando uma fatia precisa (Decisão de método).

### O que está decidido

- **A Raiz do Mundo** é o eixo estrutural do mundo. As eras do Brasil ficam dispostas lado a lado ao redor dela.
- **A Ruptura** distorceu o mundo. É por isso que a geografia não reproduz o mapa do Brasil e as eras não estão nas posições históricas reais.
- **A Raiz puxa gente continuamente.** Não existe evento fundador único. Isso explica quem chegou muito antes do jogador, como o Primeiro Perdido.
- **O jogador é um forasteiro sem origem.** Um entre muitos, puxado pela Raiz. Não é o escolhido, não pertence a nenhuma facção e não tem passado histórico definido. A identidade vem do equipamento, da progressão e da tradição que escolhe.
- **A Travessia** é um espaço liminar: a passagem entre o puxão e o mundo. O nome vale para o continente, o lugar e o ato de embarcar.
- **A Raiz visível (fechada em 20/09/2026):** durante toda a Travessia ela aparece no skybox como silhueta colossal, sempre visível e sem geometria navegável. Some quando o jogador é expulso para o mundo, na Feitoria da Cruz (Costa do Pau-Brasil), e reaparece lá como o Coração da Raiz.
- **Abertura:** fade-in simples dentro do passo 1 do tutorial, sem cutscene nova.
- **Os Perdidos** são gente que a Raiz puxou antes e que nunca atravessou. Não falam e não reconhecem ninguém. Cobrem a ameaça e a prova social do tutorial sem precisar de inimigo novo.
- **Craft não ganha justificativa mística.**
- **Os Encantados** (folclore) são a única facção que cresce conforme se aproxima da Raiz; no centro, o chefe "O que fica na borda" não é de era nenhuma.

### Em aberto, de propósito

- A causa da Ruptura: sem data e sem evento fundador.
- A origem do jogador.
- "O Ventre da Raiz", nome do instante do puxão: reservado, sem uso agora.

### Abandonado (não reintroduzir)

Jogador como escolhido; identidade de jogador derivada de um "Eco" (Ecos é outro sistema, de temporadas); Dom da Revivência; ritual de 1497; Mata Costeira como estrutura do tutorial; Vila de São Tomé. Tudo isso vem do GDD legado com classes por origem.

## Estrutura do mundo

O mundo tem três continentes e a **zona é a unidade de tudo**: recurso, risco, facção, mob e viagem. Cidade também é zona.

&#91;embedded content: estrutura do mundo · 3 continentes\]

### Continentes (Dado: `continents.json`)

| Continente | Tiers | Bandas de risco | Papel |
| --- | --- | --- | --- |
| A Travessia | T1–T2 | segura | ilha tutorial, descartável; não pertence a era nem ao Emaranhado |
| As Eras | T1–T6 | segura, disputada, mortal | mundo principal; eras lado a lado ao redor do Coração da Raiz |
| O Emaranhado | T6–T8 | selvagem | eras misturadas, sem hub fixo, PvP aberto; muito posterior ao lançamento |

### Bandas de risco

| Banda | Regra | Onde |
| --- | --- | --- |
| Segura | sem PvP | toda a Travessia e as bordas das eras |
| Disputada | PvP consentido, perda parcial | meio das eras |
| Mortal | PvP consentido, perda total | fundo das eras, perto do centro |
| Selvagem | PvP aberto, perda total | só no Emaranhado |

### Hierarquia: continente → era → região → zona

- **Continente** agrupa eras e define a faixa de tier e as bandas possíveis.
- **Era** é um cluster de zonas de um período histórico. Era 1 = Descoberta e Colonização (1500–1654).
- **Região** é só rótulo de bioma. **Nenhuma regra lê `regionId`** (Decisão).
- **Zona** carrega tier, risco, facção, recursos, acampamentos, chefe e serviços.
- **Conexão** liga duas zonas com `travelMinutes`. Viagem consome minutos reais.

### Regras estruturais (Decisões)

- Continente, era, região, tier e risco são **cinco campos independentes**; nenhum é derivado de outro.
- Toda zona tem **pelo menos dois caminhos**, para não virar corredor.
- **Facção é definida por zona**, não por região. Nenhuma facção tem duas zonas seguidas no mesmo tier.
- Fórmula de zona (do Albion): **fauna em qualquer lugar + uma facção de mob definida pelo bioma**. `factionId` é quem controla; `mobFactionId` é quem nasce. Divergem numa zona tomada.
- **O Coração da Raiz é passagem, não destino**: cidade segura no centro, cercada de zona mortal, e por ela se chega ao segundo continente.
- A geografia **não reproduz o mapa do Brasil**; é consequência da Ruptura.
- Gate de caminho é responsabilidade do tutorial/missão, nunca um campo solto na conexão.

## A Travessia (MVP)

O MVP é a Travessia completa: do fluxo de conta, login e personagem até o embarque final. Quatro zonas, T1 e T2, tudo seguro, modelada no tutorial do Albion. Não é recorte da Era 1; é conteúdo próprio e não se volta a ela.

**Decisões da ilha:** sem PvP e sem perda de equipamento; sem facção política nem reputação; **T1 é craftável**, e pedra, madeira e couro T1 se coletam sem ferramenta; as quatro zonas reutilizam o kit de assets da Costa do Pau-Brasil, sem bioma exclusivo.

### Zonas e conexões (Dado: `data/mvp/zones.json`)

| Zona | Tier | Papel | Equivalente no Albion | Recursos | Acampamentos |
| --- | --- | --- | --- | --- | --- |
| A Beira | T1 | nascimento, primeiro combate; chefe O Primeiro Perdido | The Lighthouse | seixo de praia, pau-mole, couro de capivara | Caranguejal, Cascalho solto |
| Porto de Passagem | T1 | hub: mercado, forja, refino, descanso, armazém, troca | The Cove | — | — |
| Mata Revirada | T2 | primeira coleta T2; árvores de três eras no mesmo terreno | The Forgotten Woods | pau-brasil, couro de veado, seixo | Copa revirada, Acampamento dos Perdidos |
| Pedreira Torta | T2 | segunda coleta T2: fibra e minério | Mountain Fort | algodão-bravo, ferro de brejo, couro de capivara | Perdidos do corte, Boca da fenda |

Conexões: Beira–Porto 6 min; Porto–Mata 8 min; Porto–Pedreira 8 min; Mata–Pedreira 10 min.

### Mobs e NPCs

- **Facção de mob única:** Os Perdidos. Usam o esqueleto e as animações do jogador e custam quase nada de arte.
- **Mobs:** Caranguejo da Beira (T1), Cascalho Vivo (T1), Macaco Revirado (T2), Perdido (T2, humano), Morcego da Torta (T2) e o chefe **O Primeiro Perdido** (T2), que tranca a saída da Beira e solta a bolsa.
- **NPCs com fala (três):** Bartolomeu, o Língua (recepção e tutorial, na Beira); A Barqueira (mercado e travessia, no Porto); O Ferramenteiro (ferramentas, no Porto). Estações são pontos de interação sem fala. Nível 0 de simulação: sem rotina, necessidade nem memória.

### Tutorial em 18 passos (Dado: `tutorial.json`, máquina de estados)

1. Acordar sem nada — falar com o Língua.
2. Pedra e pau — coletar 6 seixos e 6 pau-mole sem ferramenta.
3. A primeira lâmina — craftar a Lâmina Crua e preencher o slot de Ataque.
4. Primeiro combate — vencer uma onda do Caranguejal.
5. Couro de bicho — esfolar 16 couros de capivara.
6. Vestido do jeito que dá — capuz, colete e botina crus; cada peça abre um slot ativo e um passivo.
7. O broquel — off-hand só existe com arma de uma mão.
8. O que não passou — 3 ondas e o Primeiro Perdido; recompensa: bolsa de couro; libera a saída.
9. O Porto — falar com a Barqueira.
10. A mula — comprar a montaria; mercado e carga.
11. Quatro ferramentas — machado, picareta, foice, faca.
12. Mata Revirada — 16 pau-brasil e 16 couros de veado.
13. Pedreira Torta — 16 algodão e 16 ferro; peso da carga.
14. Refinar — tábua, couro curtido, pano e barra.
15. A Árvore do Destino — os quatro nós de aprendiz.
16. A escolha — uma das três armas T2 (tradição), Q, W e E, e ordenar a prioridade.
17. Você é o que veste — três peças T2 da mesma linha; bônus de conjunto e passivas.
18. A travessia — embarcar. A ilha não pode ser revisitada.

Diferença deliberada do Albion: lá várias peças vêm prontas; aqui o jogador coleta e fabrica tudo. Só a bolsa vem do chefe e só a montaria é comprada.

### Cobertura de sistemas no MVP

**Entra:** inventário e equipamento, combate e ataque básico, onda fechada, slots e prioridade, energia, coleta (com e sem ferramenta), craft, off-hand, armadura, mini-chefe, mercado e armazém, montaria e carga, refino, Árvore do Destino, escolha de tradição, conjunto completo.

**Não entra:** PvP, T3+, artefatos, encantamento e Essência, qualidade, reputação, economia entre jogadores, mundo da Era 1.

**Andamento do backlog (26/09/2026):** fatias 1 a 3 (conta, login, personagem) concluídas; fatia 4 (entrar no mundo) iniciada.

## Era 1 — Descoberta e Colonização (1500–1654)

A Era 1 está **projetada e não promovida**: vive em `design/era-1/` (22 zonas, 43 mobs, 17 NPCs, 3 facções) e só entra em `data/era1/` quando houver código que a leia. Vai da costa (T1, segura) ao centro, A Raiz (T6, mortal).

### Regiões

| Região | Bioma | Tiers | Zonas |
| --- | --- | --- | --- |
| Costa do Pau-Brasil | mata atlântica litorânea | T1–T2 | Feitoria da Cruz (hub), Praia de Coqueiros, Manguezal do Sal, Roça de Mandioca, Mata do Pau-Brasil, Aldeia Potiguara, Baixio das Garças |
| Mata dos Engenhos | mata atlântica | T2–T3 | Canavial Velho, Engenho da Cruz, Mata Fechada, Olho de Boitatá, Serra da Barriga |
| Sertão de Dentro | caatinga | T4 | Caatinga Branca, Serrote das Pedras, Curral do Gado, Lajedo Rachado |
| Rio das Almas | várzea fluvial | T5 | Várzea Alagada, Porto de Canoas, Corredeira Negra |
| Mata Sem Fim | floresta amazônica | T6 | Castanhal Antigo, Igarapé de Anhangá, A Raiz |

Risco por zona nos dados: T1–T2 segura, T3–T4 disputada, T5–T6 mortal. São 44 conexões; as viagens crescem de 8–12 min na costa a 70–85 min no fundo. A Raiz carrega `passage` para o Emaranhado, exigindo T6.

### Facções políticas (pertencem à era, não ao jogo)

| Facção | Apelido | Linha de armadura | Árvore preferida | Presença | Material exclusivo |
| --- | --- | --- | --- | --- | --- |
| A Coroa | colonizador | pesada | Físico | domina a costa e some para dentro | ferro de brejo |
| Os Potiguara e aliados | indígena | média | Projétil | em toda faixa T2–T6, nunca dominante | couro de anta |
| Os Encantados | folclórico | leve | Mágico | raro na praia, dono da Mata Sem Fim | pedra de raiz |

Rivalidades: a Coroa é rival das outras duas. Cada facção tem voz própria (a Coroa é contábil; os Encantados falam em sinal, não em frase).

**Grupos independentes (não jogáveis, sem reputação):** Contrabandistas franceses (inimigo humano de tier baixo, hostil a todos) e O Mocambo da Serra (zona neutra com descanso na Serra da Barriga; candidato a quarta facção se a reputação entrar).

### Chefes e elites

| Nome | Tier | Zona | Comportamento | Atestação |
| --- | --- | --- | --- | --- |
| Curupira | T3 | Mata Fechada | fuga volta ao mesmo lugar; anti-fuga dobrada | Anchieta, 1560 |
| Boitatá | T3 | Olho de Boitatá | queimadura persiste até achar água | Anchieta, 1560 |
| Mula sem cabeça | T4 | Curral do Gado | só na segunda metade do dia; foge e volta | lenda ibérica trazida pelos colonos |
| Sucuri anciã (elite) | T5 | Várzea Alagada | anti-fuga alta na várzea | — |
| Ipupiara | T5 | Corredeira Negra | anula tentativas de fuga periodicamente | Gandavo, 1576 |
| Castanheira velha (elite) | T6 | Castanhal Antigo | não persegue; só não deixa passar | — |
| Anhangá | T6 | Igarapé de Anhangá | copia o loadout do jogador | cartas jesuíticas, séc. XVI |
| O que fica na borda | T6 | A Raiz | nenhuma fuga funciona; não há retirada | sem era |

Arquétipos de mob: fauna (anti-fuga baixa), humano (média), elite (entra em grupo por sorteio), chefe (fixo na zona, anti-fuga altíssima).

### NPCs

Na Feitoria da Cruz: Bartolomeu, o Língua; quatro artesãos (Cosme Ferreiro, Inês da Ponta, Dona Firmina, Aparício); Dâmaso (mercado); Sebastião (armazém); cinco mestres de refino (ainda sem nome próprio). Fora dela: Poti (troca, Aldeia Potiguara), Brás do Engenho (refino avançado), O Vigia (Mocambo), O Barqueiro (Porto de Canoas) e A Voz na Borda (passagem para o Emaranhado, em A Raiz).

## Identidade e equipamento

O personagem é o que veste: arma, off-hand e armadura decidem habilidades, papel e estilo. Três conceitos separados sustentam isso e **não se misturam**:

| Conceito | O que é | Escopo |
| --- | --- | --- |
| Árvore de habilidade | kit de Q, W e passivas (Físico, Projétil, Mágico) | global; era nova **nunca** traz árvore nova |
| Tradição de equipamento | estilo de arma e armadura (Colonial, Indígena, Encantada) | global; atravessa o jogo inteiro |
| Facção política | território, reputação, conflito | pertence a uma era; a Era 2 terá outras |

### Slots

Sete slots de equipamento: mão principal, off-hand, cabeça, torso, botas, montaria, bolsa.

Regras de mão: arma de uma mão libera o off-hand; arma de duas mãos ocupa ambas, sem off-hand, em troca de mais Poder de Item; off-hand exige arma de uma mão da mesma árvore.

### Loadout: seis ativos, quatro passivos

| Slot ativo | Fonte | Pool |
| --- | --- | --- |
| Ataque | arma | Q, escolhido entre os da arma |
| Utilidade | arma | W, escolhido entre os da arma |
| Especial | arma | E, fixo pela arma; é o que diferencia duas armas da mesma árvore |
| Cabeça | armadura de cabeça | ativa da peça |
| Torso | armadura de torso | ativa da peça |
| Botas | armadura de botas | ativa da peça |

Passivas: arma, cabeça, torso, botas. Montaria e bolsa têm passiva fixa por item. Fora da prioridade: ataque básico (sempre ativo, sem custo) e poção (por limiar de vida). Os rótulos são por **fonte**, não por tecla, porque num idle não há teclado.

**O slot é a restrição** (Decisão): não existe lista de habilidades desabilitadas. Trocar uma peça troca uma ativa e uma passiva.

### Árvores e papéis

| Árvore | Papéis | Estação de craft | Linha de armadura afim |
| --- | --- | --- | --- |
| Físico | dano corpo a corpo; controle e impacto | A Forja | pesada |
| Projétil | precisão; área e desgaste | Posto de Caça | média |
| Mágico | dano; suporte | Terreiro | leve |

Família de armas: Q, W e passivas formam a linguagem compartilhada da árvore; a arma concreta traz seu E próprio. Isso evita explosão de árvores a cada era.

### Armas

| Arma | Árvore | Tradição | Tier | Mãos | E |
| --- | --- | --- | --- | --- | --- |
| Lâmina Crua | Físico | — | T1 | uma | só Q |
| Broquel Cru | Físico | — | T1 | off-hand | sem habilidade |
| Espada de lado | Físico | Colonial | T2 | uma | Estocada de Fé |
| Zarabatana | Projétil | Indígena | T2 | duas | Dardo de Curare |
| Cabaça de Boitatá | Mágico | Encantada | T2 | uma | Boitatá |
| Rodela | Físico | Colonial | T2 | off-hand | Encostar o Escudo |
| Patuá | Mágico | Encantada | T2 | off-hand | Abrir o Patuá |
| Tacape | Físico | Indígena | T3 | duas | Quebra-Crânio (design) |
| Besta de mão | Projétil | Colonial | T3 | uma | Virote Perfurante (design) |
| Maracá | Mágico | Encantada | T3 | uma | Chamado dos Espíritos (design) |
| Tocha de breu | Projétil | Colonial | T3 | off-hand | Atear (design) |

### Armaduras

| Linha | Tradição | Materiais | Identidade | Bônus de conjunto (3 peças) |
| --- | --- | --- | --- | --- |
| Crua | — | couro | só o que dá para vestir | — |
| Pesada | Colonial | minério, couro | armadura, autocura instantânea, resistir a controle | armadura e resistência a controle |
| Média | Indígena | couro, fibra | evasão, crítico, viagem, escape | evasão e escape |
| Leve | Encantada | fibra | recurso, escudo, cura ao longo do tempo, sumir | recurso, escudo e cura |

Peças T2: Pesada = Morrião, Couraça, Grevas; Média = Cocar, Escaupil, Alpercatas; Leve = Carapuça, Manto de Retalhos, Sandálias de Piaçava. Papel do slot: cabeça = utilidade e percepção; torso = a grande defensiva; botas = peso, viagem e fuga.

### Montarias, bolsas e ferramentas

- Montarias: Mula (MVP; velocidade e carga médias), Cavalo de sela (rápido, pouca carga), Boi de carga (lento, muita carga).
- Bolsas: Bolsa de couro (MVP), Cesto de cipó, Surrão de sal.
- Ferramentas: machado (madeira), picareta (minério), foice (fibra), faca de esfolar (couro). A ferramenta pode estar **no máximo um tier abaixo** do recurso; para qualidade acima de normal, precisa ser do mesmo tier.

## Combate

Auto-battler **sem posicionamento como regra de jogo**: "área" significa número de alvos, não formato nem alcance. O jogador decide antes (loadout, prioridade, limiares); o servidor resolve a luta inteira e o cliente só anima.

### Dois relógios e prioridade (sem turno, sem tick)

O próximo evento é o menor entre o próximo ataque básico e a próxima habilidade pronta. A luta vira uma linha do tempo marcada em segundos.

- **Ataque básico:** sempre ativo, vem da arma, não ocupa slot nem custa energia; intervalo = 1 / velocidade de ataque; poder 20.
- **Habilidades:** cada uma tem recarga e custo de energia. Bases: Ataque (Q) 5 s / 8; Utilidade (W) 15 s / 15; Especial (E) 25 s / 25; Armadura 35 s / 20. Tags ajustam (área e cura encarecem).
- **Prioridade:** quando mais de uma está pronta e há energia, dispara a de maior prioridade entre os seis slots ativos.
- **Energia:** recurso comum às três árvores (não é "mana"). Máxima 100 × multiplicador de IP; regeneração 3,5/s × multiplicador.
- **Ordem de verificação a cada evento:** condição de parada → limiar de retirada → limiar de cura → habilidade a disparar.

### Onda fechada

1. O jogador escolhe um acampamento da zona, como escolhe um nó de coleta.
2. O grupo nasce inteiro: 2 a 4 inimigos, 12% de chance de elite (ajustada pelo estado da zona).
3. O servidor resolve a luta completa e envia a linha do tempo.
4. A próxima onda só nasce quando a anterior morre (8 s entre ondas).

O jogador foca um alvo por vez; os inimigos vivos atacam. **Teto de 50 rodadas:** estourou, é **impasse** — ninguém morre, sem loot dos sobreviventes, volta ao terreno. **Parar no meio:** o servidor corta a linha do tempo no ponto certo; o que não foi animado não aconteceu.

### Mobs

Atributos derivam de um loadout equivalente ao tier. Multiplicadores: normal 0,35× vida e 0,45× dano; elite 1,5× e 0,85×; chefe 5,0× e 1,15×. Anti-fuga: fauna 0,5; humano 1,5; elite 2,5; chefe 4,0.

### Fuga e retirada

- Chance de fuga = 0,30 + 0,10 por ponto de fuga − 0,08 por ponto de anti-fuga, entre 5% e 75%. **Mesma fórmula em PvE e PvP.**
- Pontos de fuga vêm do equipamento (árvore Projétil, botas, torso médio, montaria, bolsa). Um slot, uma escolha: duas habilidades de fuga do mesmo slot não somam.
- Grupo: anti-fuga efetiva = a do mais forte + 0,5 por inimigo vivo além do primeiro.
- Falhar custa durabilidade, perde o loot da luta e leva golpe de saída; a chance **não sobe** com tentativas. Recarga de 20 s.
- **Largar a carga:** +4 de fuga, uma vez por combate, descarta toda a carga. Você troca o que juntou por sair vivo.

### Regras de combate (configuráveis)

- Limiar de cura: 45% de vida (poção cura 35%, 8 por padrão).
- Limiar de retirada: 20% de vida.
- Paradas fixas (não se desligam, só se ajustam): carga cheia, poções esgotadas, durabilidade mínima 30%.
- **Parar é parar no lugar, nunca voltar para a cidade.** A viagem de volta continua custando.
- A parada por "sinalizado na zona" foi removida: com flag mútua, quem não sinalizou não pode ser atacado.

## Progressão e números

Princípio: num idle **fixa-se o tempo desejado e deriva-se o dano**, não o contrário. Toda constante mora em `data/shared/balance/` (`combat`, `economy`, `progression`), nunca no código; servidor e simulador leem o mesmo arquivo. As **fórmulas** abaixo são o modelo adotado; vários **valores** (velocidade de ataque, recargas, custos, regeneração, W e passivas) estão marcados como `guessedValues` e não foram validados.

### Poder de Item (IP)

IP(peça) = base do tier + encantamento + qualidade. Não existe atributo separado de poder de ataque; o conceito de AP morreu.

| Componente | Valores |
| --- | --- |
| Base por tier | T1 100 · T2 200 · T3 350 · T4 550 · T5 800 · T6 1100 |
| Encantamento | +0 0 · +1 80 · +2 170 · +3 270 · +4 380 |
| Qualidade | normal 0 · bom 20 · excepcional 50 · excelente 90 · obra-prima 140 |
| Peso no loadout | arma 0,35 (duas mãos 0,45) · off-hand 0,10 · cabeça 0,15 · torso 0,25 · botas 0,15 |

A sobreposição é intencional: um T4+3 (820) supera um T5 limpo (800).

### Curva multiplicativa

```latex
m_{IP} = 1.0918^{(IP - 100)/100}
```

Normalizada em IP 100 (T1 = 1,00×; T6 ≈ 2,41×). Dano, vida, armadura, custo de energia, barra e regeneração passam pela mesma curva. **A recarga não escala.** Motivo (Decisão): com escala linear o tier domina tudo, encantamento vira ruído e peça de era antiga deixa de competir — o oposto da fantasia central.

### Atributos e dano

- vida = 600 × m; armadura = 240 × m médio de cabeça, torso e botas.
- mitigação = ARM / (ARM + 200 + 2,0 × IP do atacante), teto 75%. A penetração por IP impede que tier alto fique estritamente mais seguro.
- fator relativo = 1 + (IP atacante − IP alvo) / 2000, entre 0,5 e 1,8.
- dano = poder da habilidade × m × (1 − mitigação) × fator relativo. Poderes: Q 35, dano de W 45, E 90, básico 20.

### Fama, Afinidade e Árvore do Destino

- **Fama por ação** cresce 2,5× por tier (mob 20, coleta 15, refino 10 no T1).
- **Fama por nível** cresce 6× por nível; nos dados, 3.600 → 22.000 → 132.000 → 800.000 → 4.800.000.
- **Afinidade:** multiplica a fama de 1,0 a 1,5 ao longo de 60 ações com a mesma arma; congela offline; reseta só na troca. É custo de oportunidade, não punição.
- **Árvore do Destino:** nó central Forasteiro (fama de qualquer ação) e quatro ramos independentes. No MVP só os quatro nós de aprendiz (3.600 de fama cada) são alcançáveis.

| Ramo | Um nó por | A maestria dá | Ritmo de fama |
| --- | --- | --- | --- |
| Combate | linha de arma e de armadura | +12 de IP por nível, 10 níveis por tier | 1,0× |
| Coleta | tipo de recurso | +1,5% de rendimento por nível; acesso a tier | 0,35× |
| Refino | tipo de refinado | retorno de material 15% + 0,8% por nível | 0,15× |
| Craft | família de item | chance de qualidade 10% + 1,2% por nível | 0,20× |

**Maestria entra no IP da arma** (Decisão): dois personagens no mesmo tier rendem diferente pela especialização.

**Especialização ampla** (Decisão de intenção): investir numa arma não pode tornar a troca para outra do mesmo arquétipo um recomeço do zero. A forma exata está adiada, junto com a revisão do reset da Afinidade — hoje Afinidade e maestria por arma punem a troca ao mesmo tempo, o que vai contra o pitch.

Tier vai de 1 a 8 no jogo completo; o balanceamento cobre T1 a T6.

## Economia

Coleta, refino, craft e equipamento formam um ciclo; prata **entra pelos mobs e sai pelo conserto, refino e craft**. Material de tier baixo continua útil porque desce em cascata nas receitas de tier alto.

### Recursos e coleta

- Cinco tipos: madeira (machado), fibra (foice), couro (faca de esfolar), minério (picareta), pedra. **O requisito de ferramenta pertence ao material, não ao nó da zona.**
- Nó de tier alto demora mais a extrair (6 s no T1 a 78 s no T6) e é mais raro (10 s a 160 s entre nós). É o que impede equipamento de tier alto de sair de graça.
- 2 unidades por nó; estado da zona multiplica (abundante 1,5×, normal 1,0×, escassa 0,6×).
- Nó encantado tem chance pequena (8% para +1 até 0,1% para +4).

### Refino e craft

```text
1 refinado T(N) = 3 brutos T(N) + 1 refinado T(N−1)
```

Peças por refinado: arma de duas mãos 20, de uma mão 16, torso 16, off-hand 8, cabeça 8, botas 8. Conjunto completo considerado: arma + off-hand + cabeça + torso + botas. Linhas de armadura usam materiais diferentes, então vestir um conjunto exige rodar mais de uma zona (overhead de viagem de 25%).

### Estações

Cada estação tem tier próprio e só produz até ele; o dono pode cobrar taxa (padrão 10%); a cidade dá bônus de produção de 18%.

| Tipo | Estação | Faz |
| --- | --- | --- |
| Refino | Fundição | barra (minério) |
| Refino | Curtume | couro curtido |
| Refino | Tecelagem | pano (fibra) |
| Refino | Serraria | tábua (madeira) |
| Refino | Canteiro | bloco (pedra) |
| Craft | A Forja | Físico e armadura pesada |
| Craft | Posto de Caça | Projétil e armadura média |
| Craft | Terreiro | Mágico e armadura leve |
| Craft | Ferramentaria | ferramentas |
| Craft | Forja de Artefato | sobe encantamento gastando Essência |

### Encantamento e Essência

Dois caminhos, como no Albion: craftar com recurso já encantado, ou subir o nível de um item pronto gastando Essência do mesmo tier na Forja de Artefato (16, 36, 80 e 180 por nível). **Essência não se coleta: cai de mob**, por tier (equivale a runas/almas/relíquias). Encantamento é nível de item; qualidade é outro sistema e soma separado. Fora do MVP.

### Durabilidade

- Desgaste: morte tira 5%; uso tira 1 ponto a cada 400 ações; fuga fracassada também desgasta.
- Efeito: 100–50% nenhum; 49–25% −100 IP; 24–10% −200 IP; abaixo de 10% inutilizável. Ferramentas não perdem IP; a bolsa desgasta.
- Conserto por ponto cresce \~2,4× por tier (10 no T1 a 840 no T6), mesma taxa da prata que os mobs soltam (15 a 1.194).

### Carga, viagem e teleporte

- Capacidade base 100; bruto pesa 1,5, refinado 1,0. Mula +150, Boi +400, Cavalo +50; bolsas +50 a +120.
- Viagem: a pé 1,0×; Cavalo 0,6×; Mula 0,8×; Boi 1,0×.
- Teletransporte cobra por peso (pessoal até 150 de peso) e, na cidade, também por IP equipado.

Refino T1 é de graça, como no Albion. Custos em prata de refino e craft sobem por tier; os valores dependem do simulador novo para valer.

## PvP, risco, reputação e sistemas adiados

**PvP é consentido e mútuo.** Só sinalizado ataca sinalizado, nas bandas disputada e mortal; a banda decide a perda (parcial ou total).

### Regras de PvP (Decisões)

- A flag **liga e desliga só na cidade**. Desligar exige voltar, então sinalizar é compromisso, não invisibilidade sob demanda.
- **Artefato bruto só cai para sinalizados.** Num idle, o risco precisa de recompensa exclusiva.
- Cada zona mostra **quantos** sinalizados há nela; o **Faro** (habilidade) diz **quem** são.
- A busca de PvP acontece dentro da própria zona; **não existe pareamento** que mova o jogador.
- Não sinalizado não pode ser atacado. Por isso não é preciso parada extra no farm de quem não sinalizou.
- Removido do histórico: regras que deixavam alguém interferir no farm alheio sem assumir risco; gank por adjacência.

### Reputação (adiada)

- Destino planejado: artefato cai **não purificado**; purificar exige reputação com a facção dele. O jogador sobe com uma baixando a rival, ou compra purificado no mercado — o que cria o purificador como profissão.
- Níveis previstos: Inimigo, Desconfiado, Tolerado (padrão), Conhecido, Aliado.
- **Critério de corte:** se até o T4 nada depender dela, ela não entra. Até lá, a Aldeia Potiguara usa um booleano (negocia ou não).

### Adiado ou cortado

| Sistema | Estado | Nota |
| --- | --- | --- |
| NPCs com rotina, necessidades, `knowledgeBase`, fofoca, dia e noite, maré | adiado | MVP só tem o nível de simulação que o mundo precisa |
| Companheiros e mercenários | cortado | — |
| Migração de mobs | possibilidade | dentro da faixa de tier da zona ou com sinal visível |
| Temporadas | pós-lançamento | expansão traz era; temporada traz evento; eventos deixam mudanças permanentes; Ecos transformam conteúdo passado em instâncias repetíveis; sem reset; sem data anunciada antes de haver base |
| Teto de Poder de Item | reservado | não implementar agora |
| Personagem offline | esboço, não decidido | ao desconectar, sai do GameState com atividade e horário persistidos; ao voltar, ganho recalculado com teto. Decidir quando existir a primeira Activity de coleta |

### Fora do design (não reintroduzir sem nova decisão)

Combate D20; grid isométrico; posicionamento e cobertura; modo RPG de mesa; 18 mini-campanhas; variáveis procedurais por save; jogador escolhido; companheiros e mercenários; rebirth; receita desbloqueada por NPC.

## Arquitetura do servidor e dos dados

Servidor Go **autoritativo**: o cliente expressa intenção, o servidor decide. Arquitetura definitiva em escala reduzida — pouco implementado, fronteiras já certas.

### Os cinco princípios

1. O servidor é autoritativo; nada de gameplay é calculado no cliente.
2. O jogo é **Command → State → Event**; os três não se misturam.
3. Tempo é do servidor: ação é `Activity` com `StartedAt` e `EndsAt`; o loop só compara `agora >= EndsAt`.
4. Concorrência segue ownership: primeiro quem possui o estado, depois mutex, channel ou loop.
5. Escopo pequeno, fronteiras definitivas.

### As onze decisões

| # | Decisão | Por quê, em uma linha |
| --- | --- | --- |
| 1 | Uma goroutine possui o `GameState`; a rede põe comandos num channel | jogo é sequencial; dois crafts do mesmo jogador não veem o mesmo saldo |
| 2 | Command e Tick entram no mesmo loop por portas diferentes; o engine recebe `now` | tempo é do servidor; núcleo testável e determinístico |
| 3 | `data/` → `gamedata` → `bootstrap` → `game`; um arquivo por tipo de entidade; `data/` na raiz do monorepo | definição estática e estado mutável têm ciclos de vida diferentes |
| 4 | `bootstrap` é o composition root; `main` quase vazio | montar dependências é uma responsabilidade real |
| 5 | Conta não é personagem; conta fica fora do Game Core | identidade é CRUD; jogo é transição de estado |
| 6 | HTTP até entrar no mundo; depois uma única conexão WebSocket | antes é request/response; depois é contínuo e bidirecional |
| 7 | Coordenada não é estado de jogo; o servidor nunca manda posição | sem posicionamento, posição é apresentação |
| 8 | Combate em onda fechada; a luta inteira vai de uma vez; teto de 50 rodadas | lote com começo e fim, menos tráfego |
| 9 | Tópicos, não salas; uma conexão por jogador | conexão pertence ao jogador, não ao lugar |
| 10 | Persistir em pontos significativos, nunca a cada tick | `UPDATE` por tick vira RAM lenta |
| 11 | Estilo Go: interface pequena no consumidor; tipo perto do conceito; sem `models/` global; `network` em vez de `network/http` | evitar camadas e abstrações antecipadas |

Fluxos separados:

```text
Criar conta:  HTTP → Handler → Service → Repository → Postgres
Jogar:        WS → Command → Game Core → State → Event → WS
```

### Dados: `design/` e `data/`

- `data/` é o que o servidor carrega; `design/` é o que foi projetado e ainda não entra no runtime.
- `data/shared/` = sistemas e catálogos globais; `data/mvp/` = Travessia; `data/era1/` = Era 1 quando promovida. O escopo é a pasta; o nome do arquivo se repete (`mvp/zones.json`, `era1/zones.json`).
- Chaves em inglês, valores de conteúdo em português.
- Bestiário define o mob; o mapa define onde ele aparece (zona → acampamento → grupo → mob).
- Tutorial é máquina de estados em dado, não sequência no código.
- `balance/` fica inteiro em `data/` e pode citar ids ainda não carregados.

### Deliberadamente não decidido

Particionamento e sharding; transação e outbox; formato de wire do WebSocket; versionamento do conteúdo estático no cliente; reconexão no meio de um lote de combate; personagem offline.

**Descartado para o escopo atual:** mutex como dono do estado, Actor Model, microserviços, Redis, Kafka/NATS, Kubernetes, sharding, Event Sourcing (eventos são contrato, não fonte de verdade).

## Direção de arte

3D low poly estilizado, Godot, **web primeiro** (depois desktop e mobile), câmera isométrica fixa 3/4, Blender para modelo e animação. Regra de ouro: simples de modelar, fácil de animar e difícil de confundir com outro jogo.

```text
3D LOW POLY + SILHUETA FORTE + CORES CHAPADAS DE PALETA + SOMBRA SIMPLES
+ EQUIPAMENTO VISÍVEL + ANIMAÇÃO CURTA + CENÁRIO MODULAR
```

### Orçamento técnico

| Item | Orçamento |
| --- | --- |
| Personagem completo com equipamento | até 3.000 tris (o corpo nu fica na casa das centenas) |
| Peça de equipamento | 150–500 tris |
| Fauna | 800–1.500 tris |
| Chefe | até 4.000 tris |
| Prop | 50–300 tris |
| Textura | atlas de paleta 256×256, um por família; sem PBR, normal map ou metal/rough |

### Regras principais

- 1 unidade Godot = 1 metro; personagem 1,8 u; tile 2×2 u; grade de 0,5 u.
- Proporção de \~6 cabeças. **Em aberto:** chibi de 2,5–3 cabeças, após feedback de amigos (15/09/2026). Até decidir, 6 cabeças é a referência.
- Leitura: silhueta → cor → equipamento → detalhe. Um esqueleto humano compartilhado; sockets para mão principal, secundária, cabeça, torso, pés e costas.
- Progressão de tier por forma, proporção, material e acabamento, sem remodelar a peça. Raridade por ícone, borda, brilho e acento.
- Criaturas sobrenaturais com linguagem própria, nunca monstro de fantasia medieval europeia. Nada de castelos nem UI dourada.
- Servidor não manda posição: layout de zona é conteúdo estático do cliente. A cena mostra parado, coletando, lutando, produzindo e viajando (anda → para → executa → repete).
- UI 2D em blocos independentes sobre a cena; leitura: personagem → ação → zona → consequência. Tipografia sem serifa, com suporte completo a acentos.
- Arquivos: `.glb` para o jogo; fonte em `art/source/`, exports em `art/exports/`, mais `palettes/`, `fonts/`, `references/`. Nomes em inglês, minúsculas e underscore.

### Ordem de produção

Esqueleto e personagem base → atlas de paleta → poucos equipamentos representativos → um conjunto de armadura → kit de cenário da primeira área → duas criaturas → um NPC → ícones derivados dos modelos. Só depois repetir o método. **Modele como se tivesse que modelar a mesma peça cem vezes.**

## Como adicionar coisas novas

A regra-mãe: **sistema é compartilhado por definição**. Conteúdo novo usa os sistemas que já existem; não cria cópia paralela de árvore, linha de armadura ou mecânica. As listas abaixo juntam as regras escritas nos docs; onde a lista é só a consequência prática delas, isso está marcado como *derivado*.

### Promover de `design/` para `data/`

`data/` só recebe conteúdo para o qual já existe código que o carregue. A cada fatia do backlog, pergunte "que dado precisa existir para esta fatia funcionar?" e promova exatamente isso:

1. Recortar de `design/` só o necessário.
2. Colocar no arquivo correspondente de `data/`.
3. Conferir se as referências do que foi promovido existem em `data/`.
4. Remover de `design/` o que foi promovido, para não haver duas versões.

`balance/` é a exceção: fica inteiro em `data/` e pode citar ids ainda não carregados. Quando a Era 1 chegar: `design/era-1/` → `data/era1/` e o catálogo T3–T6 de `design/catalog/` → `data/shared/`.

### Nova era

- **Expansão traz era; temporada traz evento** (Decisão). O jogo precisa de várias eras para a fantasia de misturar séculos.
- A era entra no continente As Eras, **lado a lado** com as outras, ao redor do Coração da Raiz.
- Usa o **mesmo esquema** de dados da Era 1: pasta própria em `data/` com `zones`, `mobs`, `npcs`, `factions`.
- **Não cria árvore nova** nem linha de armadura nova. Dá contexto visual, histórico e material às armas e peças.
- Traz **suas próprias facções políticas**; as tradições de equipamento (Colonial, Indígena, Encantada) seguem globais.
- Nome de sistema nunca vem de uma era ou facção.
- *Derivado:* uma nova tradição de equipamento seria acréscimo em `data/shared/traditions.json`; os docs não dizem se uma era deve ou pode trazer uma.

### Nova zona

- Pelo menos **dois caminhos** (conexões).
- Fauna em qualquer lugar + **uma** facção de mob definida pelo bioma; nenhuma facção com duas zonas seguidas no mesmo tier.
- Tier, risco, região, era e continente preenchidos de forma independente; nenhuma regra pode depender de `regionId`.
- Recursos declarados por material (a ferramenta exigida já vem do material); acampamentos aninhados na zona, referenciando mobs do bestiário.
- Gate de acesso pertence ao tutorial ou a missão, nunca à conexão.

### Nova arma, peça ou habilidade

- Arma entra numa das três árvores e num papel; acessa Q e W da árvore e traz **um E próprio** (1:1).
- Peça de armadura entra numa linha existente e num slot; traz uma ativa e uma passiva.
- Habilidade vai para `skills.json` com `source` (árvore, armadura, montaria, bolsa) e tags; custo e recarga vêm das bases do slot ajustadas por tag.
- Constantes no `balance/`, nunca no código. Item e habilidade podem ter nome próprio; sistema não.

### Novo mob ou chefe

- Definição no bestiário com arquétipo (fauna, humano, elite, chefe); os multiplicadores do arquétipo derivam os números.
- Para conteúdo folclórico, registrar a atestação histórica ou dizer que não há. Não inventar origem.

### Nova facção

- Facção política pertence à era e é definida por zona. Grupos independentes (como o Mocambo) ficam fora do sistema de reputação até ela existir.
- Reputação só entra se houver conteúdo que dependa dela até o T4.

### Temporada (pós-lançamento)

Evento com mudança permanente em zona; Ecos transformam conteúdo passado em instâncias repetíveis; sem reset; sem data anunciada antes de haver base.

### Reabrir uma decisão

Decisão registrada é decisão. Reabrir exige dizer o que mudou e por quê; a regra nova vai para o doc normativo, e a antiga vai para `historico-e-estudos.md`.

## Vocabulário

Mecânica não se poetiza: sistema usa termo comum de MMORPG e não vem de era nem de facção.

| Conceito | Nome vigente | Nome antigo (resíduo) |
| --- | --- | --- |
| Progressão global | Árvore do Destino | Litania |
| Uso recorrente da mesma arma | Afinidade | Têmpera |
| Regras de execução do combate | Regras de combate | Doutrina de combate |
| Estruturas de produção | Estações | Barracas |
| Estação de craft físico | A Forja | A Bigorna |
| Recurso de encantamento | Essência | Sopro |
| Recurso das habilidades | Energia | mana |
| Bandas de risco | Segura, Disputada, Mortal (+ Selvagem no Emaranhado) | firme, trêmula, rachada |
| Conteúdo de memória do mundo | Ecos | — |
| Escala de dano | multiplicador de IP | AP |

**Mortos de vez** (herança de *A Escória*): Lastro; panteões e deuses dentro de armas; gank por adjacência; as zonas A Ressaca, A Bigorna, O Verde Surdo e A Costela; combate D20; classes como classes. Arquitetura e aprendizado técnico de A Escória podem ser reaproveitados; lore, nomes e mecânicas, não.

## Divergências e pontos em aberto

Oito pontos onde docs e dados não batem, ou onde a decisão ficou explicitamente aberta. Nenhum foi resolvido aqui.

### Divergências encontradas na leitura

| # | Onde | O que diverge |
| --- | --- | --- |
| 1 | Fama por nível | `formulas-e-balanceamento.md` usa 3.000 × 6^(N−1) (3.000, 18.000…); `progression.json` usa \~3.600 × 6^(N−1) (3.600, 22.000…) |
| 2 | Coração da Raiz | a Decisão diz cidade segura no centro, cercada de zona mortal; em `design/era-1/zones.json`, A Raiz é zona T6 **mortal**, sem serviços, com chefe e passagem para o Emaranhado (os serviços estão no NPC A Voz na Borda) |
| 3 | Tier inicial da Costa | a Decisão da Travessia diz que a Costa "pode começar em T2"; os dados da Era 1 ainda têm quatro zonas T1 na Costa |
| 4 | Resíduo de turno | Ipupiara anula fuga "a cada três turnos" e a poção tem `cooldownTicks`, mas o combate não tem turno nem tick |
| 5 | Simulador | `tools/sim.py` ainda usa escala linear, tick e ciclo fixo; nenhum número de calibração atual vale até ele ser reescrito |
| 6 | Tier máximo | o jogo fala em T1–T8, mas o balanceamento cobre só T1–T6 (T7–T8 seriam do Emaranhado) |
| 7 | Tier por profundidade | a Era 1 vai de T1 na borda a T6 no centro, mas isso só aparece nos dados, não como regra escrita |
| 8 | O Primeiro Perdido | a nota da facção diz que Os Perdidos reutilizam o esqueleto do jogador; o chefe tem arquétipo `chefe`, não `humano` |

### Em aberto, de propósito

- Causa da Ruptura e origem do jogador.
- Uso do nome "O Ventre da Raiz".
- Proporção chibi × 6 cabeças.
- Forma da especialização ampla e revisão do reset da Afinidade.
- Personagem offline (teto de ganho ao reconectar).
- Reputação (critério de corte no T4).
- Se `Camp` pertence a `game.World` ou a um pacote de encontro/combate.
- Nomes próprios dos cinco mestres de refino; se a cidade da Raiz tem mais gente além da Voz na Borda.
- Catálogo completo de habilidades (W, passivas, custos, recargas) e calibração de grupo.
