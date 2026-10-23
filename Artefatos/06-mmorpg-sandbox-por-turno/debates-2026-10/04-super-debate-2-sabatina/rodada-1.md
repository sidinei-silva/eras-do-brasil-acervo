# Rodada 1: o que aconteceu

**Como foi.** Os 22 agentes leram o briefing, os dois repos e os compilados, e cada um escreveu sozinho, sem ver os outros. Os mediadores fizeram tarefa própria:
- **M1:** auditoria de tempo;
- **M2:** mapa do acervo.

As respostas completas estão em `r1/<ID>.md`.

## Os dados que mudaram o tom
- **M1 (tempo):** o roteiro atual leva de 205 a 355 noites até o MVP, ou seja, 10 a 18 meses. As duas últimas versões com código foram abandonadas em cerca de 2 semanas, e só a fase 1 já passa disso sem mostrar nada na tela. Com cortes, cai para 110–170 noites. Critérios propostos para julgar tudo: custo em noites, tempo até aparecer algo na tela, risco de abandono, se dá para voltar atrás, se ele aprende e quanto custa validar.
- **M2 (acervo):** quase tudo o que foi decidido hoje já apareceu antes: automático com o jogo aberto, bolha de combate, grid d20, Travessia, low poly, "eu escrevo e a IA guia". Nada morreu por defeito técnico. Tudo caiu por três motivos:
  1. trocar de gênero antes de entregar;
  2. a IA escrevendo o núcleo, seguida de reset;
  3. documentação no lugar de código.

  Só 4 coisas são de fato novas: a trava anti-troca, a arte comprada, o enredo à mão e "ele escreve o código".
- **O2 (processo):** 13 roupagens, 4 recomeços, cerca de 6.200 linhas jogadas fora. O repo atual tem 0 linhas de código, e esta rodada sozinha produziu mais texto que toda a doc do jogo.

## Posições por grupo

**Técnicos**
- **T1 (servidor):** servidor próprio em Node/TS, um processo, WebSocket com JSON e Postgres, abaixo de R$ 60/mês.
  - SpacetimeDB sai por causa da licença.
  - Nakama sai porque o runtime JS é ES5.
  - Colyseus fica só como reserva.
- **T2 (Unity):** cliente fino em Unity. O Synty Sidekick só existe para Unity, então trocar de engine derruba também a arte. Cortar o UDP do 03-tecnico.
- **T3 (web):** Babylon.js + Electron, com o navegador só para desenvolver e para playtest. Admite que o Sidekick prende o jogo ao Unity e propõe o mesmo teste de 1 noite: se o Babylon reprovar, fica o Unity.
- **T4 (funcional) × T5 (OO):** primeira disputa de paradigma.
  - **T4:** núcleo puro, `(estado, ação, semente) → (estado, eventos)`, com o teste virando tabela e o replay de graça.
  - **T5:** domínio em classes (Command, State), com função pura só no cálculo. A prova seria o núcleo funcional do semi-idle: 1.173 linhas e uma ação só.
  - **Onde concordam:** inteiros, semente e monólito.
  - **Onde discordam:** T4 quer o núcleo como pacote separado; T5 quer um módulo do servidor, compartilhando só o contrato.
- **T1 e T4 atacam juntos** o "dono por zona + área de interesse" do debate anterior: é cedo demais para 20–30 jogadores. Basta um loop só.

**Arte**
- **A1:** arena pequena (~10×8 casas) que cabe inteira na tela; intenção do inimigo desenhada no chão. Sem offline, o jogo fica aberto num canto, então pede modo compacto e aviso na barra de tarefas.
- **A2:** ninguém modela, faz rig ou anima antes do MVP. Ataca as "3 faixas de forma por material", que exigem modelagem: no lugar, trocar a peça do pack, mostrar o tier só pela cor e usar uma aura única. Lembra que o Mixamo só anima humanoides, e que 4 dos 6 inimigos da Travessia são bichos.
- **A3:** Era 1 = litoral, Mata Atlântica e várzea de engenho; o sertão só foi ocupado por volta de 1684. A Travessia deve ser o "chão de ninguém", sem cor, para que cruzar a bruma traga a primeira cor do jogo. O Synty não serve para o cenário da Era 1, que precisa de um kit próprio depois do MVP. Usar pintura de domínio público da época como referência (Frans Post).

**Game design**
- **G1:** turno por lado, como no Into the Breach. O automático só aceita mobs até o gate de maestria do jogador, e o inimigo faz no máximo ~6 ações por rodada.
- **G2:** a economia precisa fechar com 1 jogador online.
  - Comprador NPC como o Mercado Negro do Albion, com cota diária; taxa de mercado de 2%.
  - Gate de 6 marcos, porque a Era 1 vai até T6.
  - O teto de durabilidade cai com o uso nos dois modos.
  - Critério de pronto: um bot simulado jogando 15h.
- **G3:** a fala do Bartolomeu contradiz 3 regras; o abrigo dele vira a "cidade 0".
  - Cortar a lore para 10 linhas.
  - Facções Coroa × Companhia, inspiradas em 1630–1654, com facção fixa por zona.
  - Glossário congelado.

**Jogadores**
- **J1:** precisa funcionar com 5 online, e tudo o que depende de outro jogador precisa de saída por NPC. Ataca a volta na cidade, porque perder minutos andando é o pior castigo para quem tem pouco tempo: propõe uma carroça paga. Quer evento fixo às 21h.
- **J2:** o risco é a luta ficar chata na 50ª vez. Pede acelerar a luta, pular o turno inimigo, acerto de 65–70% e luta com mediana de 90 s. Ataca "andar custa ação", que faz todo turno virar "anda + bate + bate"; propõe ação e movimento separados.
- **J3:** "nada offline" atende quem é interrompido, não quem tem pouco tempo. Propõe:
  - offline curto, só de coleta e refino, até 8h;
  - ou um "descanso", como o rested XP do WoW, sem produzir nada.

**IA**
- **I1:** o jogador pune a IA que vê, não a do código. O acervo tem 39 imagens geradas no ChatGPT, então deixá-lo privado ou pôr um aviso.
- **I2:** código que decide regra ele escreve à mão; a IA pode escrever teste, simulação e ferramentas. Propõe uma cota de conversa de design com IA.
- **I3:** contrato rígido no AGENTS e diário de decisões escrito à mão.

**Defesa e oposição**
- **D1:** defende o formato atual, o primeiro que junta os 4 desejos de 2023. Propõe abrir a página da Steam na fase 2.
- **O1:** o erro é o MMO; propõe single-player persistente agora e o MMO depois. Ataca o online antes do mundo. Exige plano de fim de vida para o servidor.
- **O2:** travas contra a 14ª roupagem:
  - doc congelada e um `estacionamento.md` para ideias;
  - esta sabatina com no máximo 10 decisões;
  - primeiro entregável: uma luta no terminal em 4 semanas.

## Disputas abertas para a rodada 2
1. **Formato:** MMO agora × single-player persistente com MMO depois.
2. **Pouco tempo sem offline:**
   - nada offline;
   - coleta e refino offline até 8h;
   - descanso (rested).

   E o limite do automático: por recurso ou por relógio; rende mais ou menos que o manual.
3. **Cliente:** Unity × Babylon/Electron.
4. **Paradigma:** núcleo funcional × domínio OO; núcleo como pacote × módulo do servidor.
5. **Ordem do roteiro:** online antes do mundo? MVP menor (luta no terminal; Travessia com 10 passos)? Limite de noites por fase?
6. **Combate:**
   - andar custa ação × movimento separado;
   - teto de duração da luta;
   - o automático só contra mobs até o gate.
7. **Morte:** a pé até a cidade × carroça paga.
8. **Arte:**
   - 3 faixas de forma × só trocar peça e cor;
   - kit próprio para a Era 1;
   - bestiário só com o que se compra animado.
9. **Lore e mundo:**
   - lore com 10 linhas;
   - Travessia sem bioma;
   - facção fixa por zona;
   - Coroa × Companhia.
10. **IA e processo:**
    - quem escreve o código do teste;
    - cota de design com IA;
    - travas do O2;
    - página da Steam cedo;
    - acervo privado.
