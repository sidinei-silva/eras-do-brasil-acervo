# Decisões do dono (22/10/2026)

Tomadas depois da sabatina, a partir do `resultado.md` do super debate 2. O que muda em cada arquivo está em `mudancas-docs.md`.

## Funcionamento
- Nada offline. O automático só roda com o jogo aberto, é limitado por recurso (inventário, durabilidade), não tem timer e rende menos que o manual.
- **Descanso:**
  - só começa a contar depois de 4h deslogado, em qualquer lugar;
  - 1 noite de bônus a cada 20h fora;
  - é guardado em pontos, como o rested XP do WoW;
  - dobra o ganho das 2 escadas, só no manual;
  - o tempo novo soma ao que sobrou;
  - teto de 3 noites;
  - menos de 4h fora não soma nem gasta.
- **Morte:** volta a pé à última cidade em que entrou. Nenhum nó a mais de ~4 min de uma cidade. Sem carroça.
- **Combate:**
  - instância com o cenário do lugar;
  - turno por lado;
  - no máximo 2 golpes por turno (o 2º com -5), e a 3ª ação é tática;
  - inimigo com até 4 ações, animadas juntas em até 2 s;
  - 2x/3x e pular o turno inimigo;
  - a luta manual pausa sem limite;
  - o acerto se calibra jogando;
  - andar custa ação: o simulador decide, e se passar de 60% de turnos iguais, testa o Passo grátis.

## Mundo
- Era 1: 1500–1654, T1–T4 no lançamento, com 8 marcos (a interface mostra só os que existem). Litoral, Mata Atlântica e engenho.
- Facções só depois da luta jogável. Nenhum povo real nem religião como lado.
- Enraizados vivem normalmente, mas sempre dentro da sua era, como realidade paralela.
- Travessia com 15 passos (lista em `docs/02-travessia.md`).

## Ofícios
- A falha crítica de craft e de refino não perde material: não produz nada.
- O foco fica para a Era 1.
- Bolsa como no Albion; sem bolsa ainda existe inventário.

## Técnico e processo
- **Ordem:**
  1. luta no terminal;
  2. Unity;
  3. local com save em arquivo até um amigo jogar;
  4. rede, login e banco.
- Servidor próprio, um processo, sem framework, WebSocket. Navegador só como playground.
- **Linguagem:** Go no núcleo, no simulador e no servidor (veja `06-ts-x-go/decisao.md`).
- **Núcleo:** função pura `aplicar(estado, ação, semente)`, com estado de dado puro e sem estado escondido.
- **Orçamento de noites por fase:** estourou 1,5×, corta escopo. Noite = 30 minutos ou mais.
- **IA:** você escreve os casos e o resultado esperado, e a IA escreve só o laço. A IA pode no simulador, em scripts e em testes de rede; nunca na regra nem no que o jogador vê.
- Página da Steam só depois do teste de arte.
- Docs: 1 noite de correção e depois congela.
