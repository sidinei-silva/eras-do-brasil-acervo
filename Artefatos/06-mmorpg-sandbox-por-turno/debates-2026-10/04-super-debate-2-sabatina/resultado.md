# Resultado da sabatina: para o dono marcar

Marque [x] no que aceita. O resto vai para o `estacionamento.md`.

## O que muda de verdade
- [ ] **Ordem do trabalho:** primeiro a luta no terminal (1 personagem × 2 mobs, marcos nas noites 5, 10 e 20). Depois o cliente Unity. O jogo roda local, com save em arquivo, até um amigo jogar a Travessia sem explicação. Só então entram rede, login e banco. O contrato de mensagens existe desde a noite 1.
- [ ] **Orçamento de noites por fase:** estourou 1,5×, corta escopo, nunca troca de stack. Registro de 1 linha por noite.
- [ ] **Travessia com 10 passos** (hoje são 20).
- [ ] **Era 1 até T4 no lançamento**, com os 8 marcos mantidos.

## Arquitetura
- [ ] **Paradigma:** a luta é uma função pura `(estado, ação, semente) → (estado, eventos)`, com log, porque assim dá replay, reconexão e teste por tabela. O resto do servidor é organizado por pasta de domínio, sem ECS. Nomes do domínio em português.
- [ ] **Servidor e framework:** servidor próprio em TypeScript, um processo só. Sem framework (Colyseus fica como reserva). Sem "dono por zona" por enquanto.
- [ ] **Navegador:** não é produto. Serve só para playground e playtest, se quiser.

## Jogo
- [ ] **Automático:** limitado por recurso (bolsa cheia, durabilidade), sem timer de inatividade. Rende MENOS que o manual. *Isso muda o pilar atual, que diz "rende mais por hora".*
- [ ] **Descanso:** tempo deslogado numa cidade vira bônus só no manual, com teto de 3 noites. *É sistema novo; 2 agentes pediram para deixar para depois.*
- [ ] **Combate:**
  - turno por lado;
  - inimigo com até 4 ações por rodada;
  - 2x/3x e pular o turno inimigo;
  - luta manual pausa sem limite;
  - acerto e "andar custa ação" decididos jogando no terminal.
- [ ] **Morte:** a pé até a última cidade; nenhum nó a mais de ~4 min de uma cidade.
- [ ] **Era 1:** litoral, Mata Atlântica e engenho. Facções só depois da luta jogável, sem povo real nem religião como lado.

## Arte
- [ ] **Personagem:** troca de peça do pack, tier pela cor. Nada de modelar antes do MVP.
- [ ] **Teste de 1 noite:** 3 Sidekicks no grid, na câmera real.
- [ ] **Cenário e bichos:** kit da Era 1 depois da Travessia; bichos só comprados já animados.

## IA
- [ ] **Teste do núcleo:** você escreve os casos e o resultado esperado; a IA escreve só o laço que roda a tabela.
- [ ] **Onde a IA pode e onde não pode:** pode no simulador, nos scripts e nos testes de rede. Nunca no que o jogador vê, nem na regra.
- [ ] **Página da Steam:** só depois do teste de arte.

## Docs
- [ ] **1 noite para corrigir as contradições, depois congelar:**
  - `03-tecnico` (.NET);
  - pilar "pouco tempo";
  - Era 1 T6 × marcos;
  - craft com falha × "nada por azar";
  - fala do Bartolomeu × morte na cidade.
