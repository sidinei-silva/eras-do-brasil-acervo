# Compilado: super debate (15 agentes, 5 grupos, 3 rodadas)

**Grupos**
- Técnicos: servidor MMO, engine e arquiteto solo.
- Arte: câmera, artista técnico e identidade brasileira.
- Game design: sistemas, economia e narrativa.
- Jogadores: veterano de MMO, veterano tático e casual.
- IA: opinião de jogadores, praticante e cético.

**Como foi**
1. Rodada 1: cada agente propôs uma resposta para os pontos não decididos e um ponto crítico novo.
2. Rodada 2: crítica cruzada entre grupos e uma lista de obrigatórios.
3. Rodada 3: voto em 9 pares de opções, um cardápio proposto pelo cético de IA para você decidir.

## Consenso
- Mundo em tempo real com clique para andar; turno só na luta.
- Luta no lugar, sem bloquear o mapa.
- Câmera 3/4 (~50°, FOV ~30°), com 2 zooms, fixa e igual no mundo e na luta. Isométrica, vista de cima e de lado foram descartadas.
- Grupo fora do MVP, mas o motor já aceita vários combatentes por lado.
- Servidor único com 20–30 jogadores por zona.
- Mercado único com ordens de compra.
- Unity + TypeScript continua. O servidor precisa de dono por zona, área de interesse e luta como entidade.
- A arte continua; exagerar ombro, cabeça e arma.
- IA:
  - pode explicar, revisar e apontar contradição na lore;
  - não entra em arte, voz, placeholder ou NPC com LLM.

## Votos (o que o debate sugeriu)
Turno por lado · manual livre com bônus · barraca de quem está offline · "firmar o pé" · Enraizados vivem normal · 5 conjuntos de animação · IA escreve o código dos testes a partir da sua tabela de casos.

## O que você decidiu depois (22/10)
| Ponto | Debate sugeriu | Você decidiu |
|---|---|---|
| Offline | cálculo fechado com caçada | **nada offline**; automático só com o jogo aberto, com limite |
| Presença offline | barraca | cai junto (não há atividade offline) |
| Ponto de volta | firmar o pé | **última cidade em que entrou**, como no Albion |
| Combate | bolha no lugar | **instância** com o cenário do lugar; quem passa vê o personagem parado ao lado dos mobs, com símbolo de combate |
| Paisagem da Era 1 | sertão (5) × litoral (1) | em aberto; a Era 1 é 1500–1654, então o litoral e a mata atlântica são os que batem com a época |

## Fatos de IA usados
- **Steam (jan/2026):** declarar o conteúdo de IA do jogo e da loja; ferramenta de desenvolvimento é isenta.
- **Pesquisas de opinião:**
  - Quantic Foundry: 85% dos jogadores negativos;
  - GDC 2026: 52% dos devs veem dano;
  - GameDiscoverCo: 8% não compram.
- **Casos:** Clair Obscur perdeu prêmio; Larian recuou; Arc Raiders regravou vozes.

## Pontos que ficaram para a próxima
- O pilar "pouco tempo não prejudica" e a caçada sem cliente aberto caem com a sua decisão. É preciso reescrever como o jogo atende quem tem pouco tempo.
- A fala do Bartolomeu sobre a morte precisa bater com a regra nova.
- O `03-tecnico.md` ainda diz .NET.
