# Rodada 3 (final): voto

Leia `rodada-2.md`. Escreva em `scratchpad/sabatina/r3/<SEU-ID>.md`, com até ~250 palavras:

## Parte 1: consensos (responda "aceito" ou "rejeito + motivo de 1 linha")
- C1. MMO é o destino, mas até existir mundo o jogo roda local: servidor TS no mesmo processo, save em arquivo e contrato de mensagens desde o dia 1. A rede vem depois.
- C2. Nada offline. Existe o **descanso**: tempo deslogado numa cidade vira bônus de progressão com teto de 3 noites, sem produzir nada.
- C3. O automático é limitado por recurso (bolsa cheia, durabilidade), sem timer de inatividade.
- C4. Cliente Unity fino. Ele só entra depois que a luta rodar no terminal.
- C5. A luta é função pura com log de eventos, e o resto do servidor é organizado por domínio. Nomes do domínio em português.
- C6. Primeiro entregável: luta no terminal (1 personagem × 2 mobs), com marcos nas noites 5, 10 e 20.
- C7. A Travessia cai para 10 passos.
- C8. A morte leva de volta a pé à última cidade em que entrou. Nenhum nó fica a mais de ~4 min de uma cidade.
- C9. Arte: forma pela troca de peça do pack e tier pela cor. O kit da Era 1 vem depois da Travessia. Bichos só os que dá para comprar já animados.
- C10. A paisagem da Era 1 é litoral, Mata Atlântica e engenho.
- C11. Cada fase tem orçamento de noites. Se estourar 1,5×, corta escopo, nunca troca de stack. Registro de uma linha por noite.
- C12. Turno por lado, acerto de 65–70%, velocidade 2x/3x e pular o turno inimigo.

## Parte 2: pares empatados (vote A, B ou tanto faz)
- P1. **Código do teste do núcleo:** A, ele escreve a tabela e o código × B, ele escreve a tabela de casos e o resultado esperado, e a IA escreve só o laço que roda a tabela.
- P2. **Ações do inimigo por rodada:** A, até ~6 × B, até 4, animadas juntas em no máximo 2 s.
- P3. **Luta manual parada:** A, passa ao automático depois de X min × B, fica pausada sem limite (o mob ganha uma cópia por jogador).
- P4. **Automático:** A, rende mais por hora, como no pilar atual × B, rende menos que o manual e é limitado por recurso.
- P5. **Descanso vale:** A, em tudo × B, só no manual.
- P6. **Página da Steam:** A, na fase 2 × B, depois do teste de arte, com a cápsula feita por gente.
- P7. **Andar custa ação:** A, mantém × B, o simulador decide (gatilho de 60%).
- P8. **Era 1 e marcos:** A, Era 1 até T6 com 8 marcos (como está) × B, Era 1 até T4 no lançamento, sem mexer nos 8 marcos.
- P9. **Facções da Era 1:** A, definir agora × B, depois da luta jogável, só com a regra de não pôr povo indígena nem quilombo como lado a ser atacado.
- P10. **Docs:** A, congelar já × B, uma noite corrigindo as contradições que o M2 achou e depois congela.

## Parte 3
Em 1 linha, a coisa que o dono não pode esquecer.

Devolva só os votos: C1–C12 em uma linha (ex.: "todos aceito, exceto C7"), P1–P10 em uma linha e a frase da parte 3.

**Mediadores:** não votem. M1 vai receber a tarefa final depois.
