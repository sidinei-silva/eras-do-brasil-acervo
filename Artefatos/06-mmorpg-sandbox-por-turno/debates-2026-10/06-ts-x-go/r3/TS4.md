# TS4, rodada 3

## Respostas

**@GO2:** "Tabela de 5 casos, sem código de IA, ganha quem terminar primeiro? Se empatar, Go leva?"
- Aceito os 5 casos e o "sem código de IA".
- Sobre "quem terminar primeiro": a segunda versão já sai com a lógica pensada. Por isso proponho Go **por último**, o que favorece vocês.
- Empate leva Go. Aceito: inteiros e binário são vantagens medidas, não adjetivos.

**@GO4:** "Pedir à IA o setup do TS conta contra o TS, como o `Clonar()` conta contra o Go?"
Sim, e de forma simétrica. O setup (`package.json`, `tsconfig`, `go mod init`) entra no cronômetro. Se a IA escrever o setup, isso conta como código de IA.

**@MK1:** "Que habilidade nova o TS leva ao mercado?"
Quase nenhuma. Função pura, teste por tabela e determinismo valem em qualquer linguagem. Aceito: no critério carreira, Go tem vantagem pequena.

## Protocolo do teste de 1 noite

| Item | Proposta |
|---|---|
| O quê | `aplicar(estado, {tipo:"atacar", alvo}) → {estado, eventos}`. Usa d20 com semente no estado, CA, dano e PV, e emite os eventos `ataque`, `dano` e `morte`. A especificação fica escrita no papel antes. |
| Casos | 5 casos, escritos por ele: 20 natural, 1 natural, acerto, erro e morte. Mais um teste obrigatório: a entrada é igual antes e depois. A IA só escreve o laço, como manda a regra do projeto. |
| Tempo | Duas noites, uma por linguagem: TS primeiro, Go depois. Teto de 90 minutos em cada, contando o setup. Não pode abrir a primeira versão. |
| O que mede | 1. Minutos até a tabela ficar verde. 2. Se mutou a entrada. 3. Travadas de mais de 10 minutos. 4. Consultas: documentação é livre, mas conta; pergunta de código à IA conta; código vindo da IA desclassifica a versão. 5. No fim, uma nota de 1 a 5 para "sinto que é meu". |
| Quem ganha | Só uma ficou verde: ela ganha. As duas verdes com diferença acima de 25% no tempo: a mais rápida ganha. Diferença menor que isso: Go. |
| Validade | Definitiva até a fase jogável, conforme a regra do N4. O resultado fica registrado num arquivo de decisões. Não se reabre por "curva de iniciante" nem por "faltou configurar". |

## Inclinação atual

**TS, leve.** Quase empate. Cedi carreira e distribuição. Também cedi que o conforto em TS não protegeu a entrada no `da17716`. Ficam a fluência e a união discriminada, e isso o teste mede.

## Condição

Aceito Go sem reclamar em três casos:
- se o teste der Go ou empate;
- se ele declarar que o projeto também serve para treinar Go;
- se ele ficar 2 noites sem fazer o teste. Não fazer o teste também é debate no lugar de código, e aí vale o padrão com menos portas: Go.
