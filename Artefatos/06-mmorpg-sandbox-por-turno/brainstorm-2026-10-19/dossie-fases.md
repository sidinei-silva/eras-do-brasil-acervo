<!-- Brainstorm multiagente de 19/10 · arquivo original: dossie-fases.md · escrito pelo Claude em 2026-10-19T23:41 na conversa "RavenIdle melee vs ranged mechanics" (export de 21/10/2026) -->

# Dossiê das fases do Eras do Brasil (preparação do debate)

| # | Período | Fase | O que foi feito | Por que mudou |
|---|---|---|---|---|
| 1 | jul/2023 | "1500 Caminhos do Brasil" | Conceito de RPG digital por turno sobre a história do Brasil (inspirado em Pillars of Eternity, Wartales, Baldur's Gate). Linhas do tempo históricas | Parou por ~19 meses |
| 2 | fev–ago/2025 | RPG de mesa persistente | "Eras do Brasil" nasce (mar/2025): d20, classes de figuras históricas e folclóricas (3 origens: Colonizador, Indígena, Ser Folclórico; 12 classes), craft, economia, mundo reativo. Livro de Regras, Livro de Classes, Enredo Principal (Raiz do Mundo, Ruptura de 1497, Dom da Revivência, 18 mini-campanhas, Linha Eco-Histórica) | Hiato de ~3 meses |
| 3 | nov/2025–fev/2026 | Planejamento (Gemini) | Ticks, facções, plano em três fases (mesa + digital), stack, simulação de NPC por eventos | Virou planejamento sem jogo |
| 4 | fev/2026 | Unity tático cooperativo | GDD, ADRs, Unity cooperativo P2P, combate estático que evoluiria para grid; 38 PoCs planejadas e quase nenhum código | Diagnóstico de 07/03: "GDD 100%, código 0%" |
| 5–6 | mar–abr/2026 | MUD moderno | MMORPG textual, servidor Go autoritativo, cliente web. Tese "o mundo não espera por herói" (17/03). Temporadas (Tensão → Instabilidade → Apogeu → Legado), Bandeirantes de Sangue como facção antagonista do servidor, Guardião da Fenda, Restauradores × Fatalistas, causa da Ruptura (ritual folclórico fracassado para proteger a terra de um presságio). Reset de 03/04: código feito por IA apagado para ser reescrito à mão | Fadiga, dúvida recorrente de linguagem, ideia de idle |
| 7–8 | mai–ago/2026 | Spin-off idle "A Escória" | Idle inspirado no Albion, em Go: Têmpera, Litania, full loot, lore própria (Deuses-Armas) para não diluir o Eras. Eras congelado | Viável para dev solo, mas sem mundo próprio |
| 9 | set/2026 | Fusão idle | Esqueleto da Escória + mundo do Eras. Go + Godot, 3D low poly. Travessia como ilha tutorial (18 passos), Era 1 projetada (1500–1654: 22 zonas, 3 facções políticas Coroa / Potiguara / Encantados, chefes folclóricos com atestação), forasteiro sem origem, Raiz puxa gente continuamente, Perdidos, sem Dom e sem escolhido | 60 trocas de meta-organização; quarta fatia travada |
| 10 | 27/09–02/10/2026 | Jogos pequenos | Trilha de jogos menores em Godot como estudo | Visto como sintoma de falta de loop jogável |
| 11 | 03–14/10/2026 | "Eras novo" (F8) | Combate instanciado por turno, modo ativo + caçada automática, 3 filas, fuga, intenção visível, TypeScript + Electron + Three.js, core determinístico. Código do core escrito em parte por IA | 14/10: "o jogo está totalmente indefinido" |
| 12 | desde 15/10/2026 | **Estudo de design (fase atual)** | Estudo solto de mecânicas, sem código. Decisões tomadas nesta fase (abaixo) | Em curso |

## Fase atual: decisões de mecânica (valem para o debate)

- MMORPG sandbox para PC (Steam), dev solo, projeto "for fun". Público: pouco tempo, interrupções.
- Resolução por d20 + atributo + proficiência + item contra Defesa/CD, com graus de sucesso (estilo Pathfinder 2e), em combate, coleta, craft e conversa.
- Combate por turno em grid visível: 3 ações + 1 reação, penalidade de ataque múltiplo, ataque de oportunidade. Lutas curtas, poucos oponentes. Grupo de inimigos visível no mapa abre a luta no próprio lugar.
- Modo manual (pilar; chefes; rende mais por luta) e automático (lista de prioridade + intenção de movimento; rende mais por hora). Coleta offline limitada pela durabilidade da ferramenta.
- Você é o que veste: sem classes; skills e passivas na arma e na armadura; atributos (Força Bruta, Astúcia, Vigor, Conhecimento, Sabedoria Ancestral, Presença) só do equipamento. Equipamento visível.
- Progressão por uso: maestria da linha de arma (acerto e gate de tier) e especialização do item (dano). Igual para ofícios.
- Durabilidade em duas camadas (atual e teto). Nada se destrói por azar. Morte no manual: durabilidade, um pouco de teto e parte do que carrega; no automático só durabilidade. Sem full loot.
- PvP assíncrono contra o eco (cópia da build e da automação) de outro jogador, sem perda de itens.
- Mundo persistente em zonas, facções controlando zonas (algumas divididas), eventos periódicos com marcas permanentes.
- Expansões = eras; cada era tem tiers 1 a 6 e soma opções, não poder.
- NPCs da vila (comércio, tutores de estação, tarefas) e viajantes que só levam e trazem fofocas.
- Craft de peça de facção: afinidade com mestre artesão da facção + artefato (missão, drop raro, mercado).
- MVP: área tutorial (equivalente à ilha tutorial do Albion) com itens até o tier dela. Depois do tutorial vem a Era 1.
- Stack recomendada: Unity (cliente 3D low poly, câmera fixa inclinada) + servidor .NET + núcleo C#.

## Fase atual: decisões e preferências de lore

- O jogador é pessoa comum, nunca o escolhido. Reconhecimento vem das ações e de como o mundo e os outros o veem.
- Mundo persistente com eventos que deixam marcas.
- Brasil e folclore como inspiração indireta, sem reprodução histórica, sem nomes fortes (povos, pessoas, religiões, lugares reais), evitando polêmica.
- Era 1 inspirada no período de 1500–1654 (chegada e colonização), de forma indireta.
- Eras novas somam opções, não poder.
- O criador quer história com começo, meio e fim, motivação clara e uma introdução que ambiente. Mistério demais não motiva.
- Considerou (não fechado): Ruptura como processo contínuo; "Eternos" (gente que não morre), "Enraizados" (puxados com a terra, as facções), "Perdidos" (quem não conseguiu se montar ou ficou); origem moderna do jogador (isekai) como hipótese.
- Do debate anterior (18–19/10), mais votado: chefe negociado com trato semanal da zona; varanda com "o que aconteceu enquanto você esteve fora"; Visagem de morte com escolha de devolver ou ficar com 1/3; compadrio como ponto de retorno; terra "desassuntada" com Donos e Trato; costume nos objetos. Essas ideias podem ser usadas, cortadas ou ignoradas.

## Regras do novo debate

- Não poetizar mecânica: sistemas usam nomes comuns de MMORPG (durabilidade, maestria, modo automático, eco).
- Lore, história e mundo sóbrios: pouco floreio, frases diretas.
- Respeitar a Era 1 e as mecânicas acima.
