# Resultado do debate geral (15 agentes, 3 rodadas)

## Consenso
- Mundo em tempo real com clique para andar; turno só na luta.
- A luta acontece no lugar, numa "bolha", sem bloquear o mapa.
- Câmera 3/4 (~50°, FOV ~30°) com 2 zooms, a mesma no mundo e na luta, sem rotação no MVP.
- Grupo fica fora do MVP, mas o motor já aceita N combatentes por lado.
- Servidor único, ~20–30 jogadores por zona, sem canais.
- O offline é um cálculo fechado ao logar, com teto em horas e cota própria.
- Mercado único com ordens de compra.
- Stack: Unity no cliente; TypeScript no núcleo e no servidor (Node, WebSocket, Postgres).
- Arte: low poly sombrio + contorno; exagerar ombro, cabeça e arma.
- IA: nunca em placeholder, arte, voz ou NPC com LLM. Na lore, só aponta contradição. Os casos de teste são escritos por ele.

## Votos (A / B / tanto faz)
| # | Ponto | A | B | Tanto faz | Maioria |
|---|---|---|---|---|---|
| 1 | Turno: individual × por lado | 4 | 10 | 1 | B por lado |
| 2 | Manual: raro × livre com bônus | 1 | 14 | 0 | B livre |
| 3 | Offline: coleta/refino × + caçada | 5 | 10 | 0 | B + caçada |
| 4 | Offline visível: barraca × nada | 14 | 1 | 0 | A barraca |
| 5 | Retorno: firmar o pé × fixo | 15 | 0 | 0 | A firmar |
| 6 | Enraizados: vivem × parados | 15 | 0 | 0 | A vivem |
| 7 | Era 1: litoral/mata × sertão/caatinga | 1 | 5 | 9 | dono escolhe |
| 8 | Animação: 5 conjuntos × por arma | 15 | 0 | 0 | A 5 conjuntos |
| 9 | IA: restrito × escreve código de teste da tabela dele | 1 | 13 | 1 | B |

## Obrigatórios do MVP
1. Offline com cota, teto e a tela "enquanto você esteve fora".
2. Grid marcado na greybox e o teste das 3 arenas.
3. Comprador NPC com cota diária, simulado no sim/.
4. Intenção visível do inimigo.
5. Contrato de IA no AGENTS e registro de origem dos assets.

## Ajustes na doc
- 03-tecnico: trocar .NET por TypeScript.
- Bartolomeu: "sem o que carregava" vira "parte do que carregava", porque não existe full loot.
