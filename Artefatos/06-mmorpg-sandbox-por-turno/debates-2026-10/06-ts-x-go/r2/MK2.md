# MK2: rodada 2 (Mercado, backend de jogos)

## 1. Respostas

Ninguém me fez pergunta na rodada 1 (`conversa.md`, busca por "MK2"). Mas a minha pergunta ao TS1 sobre RAM foi respondida por neutros, com medição:
- N3: Go 7,5 MB, Bun 16 MB e Node 22 46 MB em repouso;
- N1: 45 MB contra 7 MB.

Aceito esses números. Numa VPS de 1 GB não mudam o preço. A outra metade da pergunta, o plano de hospedagem TS sem plataforma gerenciada, continua com o TS1.

## 2. Réplicas

**@MK1:** "Sugiro aos debatedores que parem de usar 'carreira' como argumento forte, para os dois lados."
Concordo, e isso derruba parte da minha abertura (veja a concessão abaixo). Acrescento um dado de mercado: o que recrutador de backend de jogos pede é sistema que rodou com gente de verdade, não a linguagem. Nakama, Colyseus e SpacetimeDB têm vitrine com jogos lançados, não com protótipos. Então o ativo de carreira é a fase 5 ("um amigo joga"), em qualquer das duas.

**@N4:** "O motivo que levou ao TS em 10/10 foi 'o cliente vira web'. Esse motivo sumiu. O TS precisa de outro."
Do ponto de vista de mercado, confirmo e acrescento um caso concreto. O argumento "ecossistema de jogo em TS" também perde força com cliente Unity:
- O Colyseus é bom, mas o briefing proíbe framework até a fase 6.
- O Hathora, que era hospedagem para o mundo TS, saiu do mercado de jogos (https://gamesbeat.com/?p=318174).
- Se um dia ele quiser o Nakama, o runtime TS de lá **não é Node**. É uma VM embutida (goja) que só aceita ES5, sem módulos que dependam do Node: https://heroiclabs.com/docs/nakama/server-framework/typescript-runtime/

Opinião: o que sobra de real para o TS é fluência (N4, TS4) e o playground nativo no navegador (N1, item 7). São motivos do dono e do teste, não do mercado. Ainda assim podem bastar, porque o mercado empata.

## 3. Concessão

Na rodada 1 eu disse que "Go tem retorno duplo" e que "para quem é a única renda da casa, o retorno duplo pesa". Cedo parte disso:
- **MK1:** a hora de Go que rende para a carreira é a do expediente, com review sênior e produção. Um projeto noturno soma pouco a um currículo que já diz "Go em produção".
- **TS4:** quem paga esse aprendizado é a empresa, 40h por semana.

O que mantenho, rebaixado para peso pequeno: praticar Go à noite pode acelerar o dia (MK3 diz o mesmo). Mas isso é benefício colateral, não critério de escolha. Assino o peso 10 do N4 e não peço mais que isso.

## 4. Nova pergunta

- **TS2:** você propôs mulberry32 com `Math.imul` e `>>> 0` para o sorteio. Se na fase 6 o servidor for para o Nakama com runtime TS (goja, ES5), o `aplicar` roda lá sem mudança? E vale a mesma pergunta se alguém do seu lado usar `BigInt` para garantir inteiros, já que o `BigInt` não existe em ES5. Ou você descarta o Nakama de vez e assume que o caminho TS da fase 6 é sempre um processo próprio?
