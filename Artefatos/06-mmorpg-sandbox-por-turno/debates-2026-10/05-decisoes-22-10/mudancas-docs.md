# O que muda nos docs (noite de correção)

Uma linha por mudança. Você escreve à mão; isto é só a lista.

## 01-o-jogo.md
- Pilar "Pouco tempo não prejudica" → nada acontece offline. O automático só roda com o jogo aberto e é limitado por recurso. Quem ficou fora ganha descanso.
- Pilar "Manual rende mais por luta, automático rende mais por hora" → "O automático rende menos que o manual."
- Novo pilar: "A luta manual pausa sem limite."

## 02-travessia.md (+ canvas)
- Tabela de passos: de 20 para os 15 combinados.
- Fala do Bartolomeu:
  - "firmar o pé" sai;
  - "volta a cair aqui, sem o que carregava" → volta à última cidade em que entrou e perde parte do que carregava.
- Proposta "passos 1–8 na Beira, instanciados e solo" → passos 1–6.
- Inimigos: os bichos dependem de existir um pack já animado. Marcar como provisório.

## 03-tecnico.md
- Stack:
  - **cliente:** Unity fino, só depois da luta no terminal;
  - **servidor:** um processo, sem framework, WebSocket + JSON;
  - **banco e login Steam:** só depois que um amigo jogar;
  - **até lá:** save em arquivo.
  - **Linguagem:** aguarda o debate TS × Go.
- Sai: .NET, .NET Standard, "núcleo como pacote no Unity", a seta Núcleo → Cliente, UDP/LiteNetLib, "servidor só manda o que está perto".
- Servidor "roda lutas sem cliente" → só com o cliente conectado.
- Núcleo: `aplicar(estado, ação, semente) → {estado, eventos}`, com 3 regras:
  - nada que muda fica guardado no objeto;
  - estado é dado puro, sem métodos;
  - não mexer no estado que entrou.
- Contrato de mensagens desde a noite 1: id em cada comando e número de sequência em cada evento.
- Log de luta com a versão das regras e o hash de `data/`.
- Combate é instância: usa o cenário do lugar; no mapa aparece o personagem parado com o ícone.
- Pastas: `regras/` (no lugar de `core/`), `contrato/`, `sim/`, `server/`, `client/`, `data/`, `docs/`. `tools/` só quando precisar.
- Navegador: só playground.
- "Em aberto: confirmar a stack" → fica só a linguagem.

## 04-roteiro.md (+ canvas)
- Fases:
  0. Base
  1. Luta no terminal (marcos nas noites 5, 10 e 20)
  2. Luta na tela (teste de 1 noite com 3 Sidekicks)
  3. Mundo e ofícios local
  4. Travessia (15 passos)
  5. MVP: um amigo joga local
  6. Online: rede, login e banco
- Regras novas:
  - orçamento de noites por fase; estourou 1,5×, corta escopo, nunca troca de stack;
  - noite = sessão de 30 minutos ou mais;
  - registro de 1 linha por noite trabalhada.
- "Como trabalhar": "IA escreve os testes" → você escreve os casos e o resultado esperado, e a IA escreve só o laço que roda a tabela.
- Sai: "referência da IA num repositório separado".
- Página da Steam: só depois do teste de arte.

## mecanicas/combate.md
- Turno por lado. Inimigo com até 4 ações por rodada, animadas juntas em até 2 s, e intenção visível (casa-alvo).
- "-5 no 2º, -10 no 3º" → no máximo 2 golpes por turno, só o 2º com -5; a 3ª ação é tática.
- 2x/3x e pular o turno do inimigo.
- Manual pausa sem limite. Mob ganha uma cópia por jogador.
- Andar continua custando ação. Gatilho: se o simulador medir mais de 60% de turnos "anda + bate + bate", testar o Passo grátis.
- Acerto: um alvo inicial, calibrado jogando.
- Modos: automático rende menos que o manual, é limitado por recurso e só roda com o jogo aberto.
- Proposta "servidor recorta o grid" → instância com o cenário do lugar.

## mecanicas/personagem.md
- 8 marcos mantidos; a interface só mostra os marcos dos tiers que existem (até T4).
- Nova seção, Descanso:
  - guardado em pontos;
  - dobra o ganho das 2 escadas em ação manual;
  - conta depois de 4h fora;
  - cada 20h fora = 1 noite de bônus;
  - soma ao que sobrou;
  - teto de 3 noites;
  - menos de 4h fora não soma nem gasta.

## mecanicas/oficios-e-economia.md
- Coleta: "automática e offline" → só automática. Para com o inventário cheio ou a ferramenta sem durabilidade.
- Craft, falha crítica: "perde material" → perde foco e não ganha nada.
- Bolsa: como no Albion, mais capacidade de peso e status. Sem bolsa ainda há inventário.
- "Trabalhadores NPC para auto-coleta" sai: produzem enquanto você está fora.

## mecanicas/durabilidade-e-morte.md
- Ponto de retorno: a pé até a última cidade em que entrou. Nenhum nó fica a mais de ~4 min de uma cidade. Sai do "Em aberto".

## mecanicas/mundo.md
- Era 1: T1–T6 → T1–T4 no lançamento (no diagrama também).
- Paisagem: litoral, Mata Atlântica e engenho.
- Facções: só depois da luta jogável. Nenhum povo real nem religião como lado.
- "Arena no lugar ou fixa por bioma" → resolvido: instância com o cenário do lugar.
- Regra de mapa: nenhum nó a mais de ~4 min de uma cidade.

## mecanicas/motor-de-regras.md
- Determinismo: a semente e o contador do sorteio ficam dentro do estado.

## AGENTS.md
- Contrato de IA:
  - casos e resultado esperado são seus; a IA escreve só o laço;
  - a IA pode no simulador, em scripts e nos testes de rede;
  - nunca na regra nem no que o jogador vê.

## Arquivo novo
- `noites.md`: registro de 1 linha por noite (data · fase · usadas/orçamento · código ou design · o que fez).

## Ainda em aberto (não decidido)
- **Foco:** é recurso novo. Precisa dizer o que é, como volta e para que serve além do craft.
- **Refino:** a falha crítica também perde material hoje. Segue a regra do craft?
- **Enraizados:** param no tempo ou vivem normalmente (no debate anterior, 15 a 0 por "vivem")?
- **Linguagem:** sai do debate TS × Go.
