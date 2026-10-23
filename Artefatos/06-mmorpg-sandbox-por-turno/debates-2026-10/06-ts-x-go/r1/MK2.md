# MK2: abertura (Mercado, backend de jogos)

Persona: analista de tendências de backend de jogos. Trago fatos com link e digo quando é opinião.

## 1. Dados atuais

**Adoção geral (Stack Overflow Survey 2025, uso entre profissionais)**, https://survey.stackoverflow.co/2025/technology

| Linguagem | Usa | Admirada |
|---|---|---|
| TypeScript | 48,8% | 58% |
| C# | 29,9% | 55,8% |
| Go | 17,4% | 56,5% |
| Rust | 14,5% | 72,4% |
| Elixir | 2,8% | 65,9% |

- TS virou a linguagem nº 1 do GitHub em contribuidores (ago/2025), puxada por código gerado com IA: https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/
- O compilador do TS 7 está sendo reescrito **em Go** pela Microsoft: https://infoworld.com/article/4100582/microsoft-steers-native-port-of-typescript-to-early-2026-release.html. As duas linguagens não competem por sobrevivência. As duas vão durar.

**Backend de jogos: quem usa o quê**
- **Go, Nakama (Heroic Labs).** É o servidor de jogo open source mais conhecido. O núcleo é em Go, e a lógica do jogo pode ser escrita em Go, TypeScript ou Lua, com a mesma API de partida: https://heroiclabs.com/docs/nakama/concepts/multiplayer/authoritative/llm.md. A empresa diz que tem "centenas de títulos", mas o número vem dela mesma: https://www.gamedeveloper.com/game-platforms/embracing-the-future-of-open-source-game-servers-with-nakama
- **TS, Colyseus.** A 0.17 saiu em fev/2026, com reconexão automática e tipos compartilhados entre cliente e servidor: https://colyseus.io/blog/colyseus-017-is-here/
- **TS, Hathora.** Foi comprada pela Fireworks AI em mar/2026 e saiu do mercado de jogos: https://gamesbeat.com/?p=318174. Por causa disso, o online do Stormgate desligou em abr/2026: https://www.techspot.com/news/111969-stormgate-servers-go-dark-following-ai-focused-hosting.html
- **C#.** O Albion, que é a referência de progressão do dono, roda com Unity + Photon Server: https://blog.photonengine.com/?p=3622. O estúdio diz que "o núcleo da simulação é totalmente independente da Unity": https://gamedev.net/news/3258-architecting-albion-online-how-sandbox-interactive-built-a-pvp-mmo-to-scale/
- **Rust, SpacetimeDB 1.0** (mar/2025). Junta banco e servidor num só sistema e roda o backend do MMO BitCraft: https://spacetimedb.com/blog/introducing-spacetimedb-1-0
- **Node, Screeps.** É um MMO de estratégia vendido na Steam, com o servidor em Node aberto: https://github.com/screeps/screeps

**Padrão de servidor por turno.** A doc do Nakama separa dois tipos:
- **turno ativo**, com os jogadores conectados e poucas mensagens. Nesse caso o tick é baixo;
- **turno passivo**, em que o servidor valida, grava e para o loop até a próxima jogada.

Fonte: o mesmo link do Nakama acima. O Eras é turno ativo dentro de uma instância. É o mesmo desenho do `aplicar` puro com o servidor autoritativo, que é o padrão da indústria.

**Carreira (Brasil).** As faixas salariais de Go pleno CLT giram em torno de R$ 8–14k, mas a fonte não cita de onde tirou os números: https://jobrise.io/pt/blog/golang-vagas-brasil-2026-quanto-paga/. Trate como estimativa.

## 2. O que isso significa para ESTE dono

1. **O mercado não decide esta escolha.** Go e TS são mainstream, bem admirados e têm ecossistema de jogo de verdade (Nakama × Colyseus). Nenhum dos dois é aposta arriscada.
2. **Framework não conta agora.** O briefing já decidiu "sem framework". Nakama e Colyseus só entram na fase 6 (online), se entrarem. Opinião: quem usar "o ecossistema de jogo é maior" como argumento hoje está vendendo o que só vai servir na fase 6.
3. **A lição do Hathora favorece servidor próprio, em binário, local.** Depender de plataforma gerenciada mata jogo pequeno. Os dois lados conseguem gerar um executável. Fica para TS2 e GO provarem o custo de cada um.
4. **C# fica de fora, mas é o que o mercado de Unity usa no servidor (Photon/Albion).** O mercado mostra que núcleo em C# compartilhado com o cliente é o caminho "padrão". Opinião: o dono rejeitou isso com razão, porque não sabe C# e o histórico dele mostra que trocar de stack mata projeto.
5. **Carreira.** Go tem retorno duplo: ele já usa no trabalho, e cada noite no Eras também treina para a Gamersclub. Em TS, o retorno de carreira é ganhar profundidade, não ganhar mercado novo. Opinião: para quem é a única renda da casa, o retorno duplo pesa.
6. **IA.** TS ter o maior volume de código gerado por IA no GitHub é faca de dois gumes para quem quer recuperar autoria. Esse é o mesmo ponto que TS2 admitiu.

## 3. Perguntas

- **GO1:** o Nakama aceita lógica em Go ou TS com a mesma API. Na fase 6, você recomendaria migrar o servidor próprio para dentro do Nakama ou manter o processo sem framework? Se for migrar, o `aplicar` em Go entra sem reescrita?
- **TS1:** o Hathora era o maior nome de hospedagem do mundo TS e saiu do mercado. Na fase 6, qual é o plano de hospedagem TS que não depende de plataforma gerenciada? Quanto de RAM um processo Node ou Bun ocioso ocupa numa VPS de 1 GB, comparado a um binário Go?
