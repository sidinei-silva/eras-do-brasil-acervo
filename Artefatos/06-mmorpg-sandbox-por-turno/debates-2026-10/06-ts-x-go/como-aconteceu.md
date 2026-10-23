# Como foi o debate TypeScript × Go

## Quem participou (15)
- **TypeScript:**
  - TS1 não tem problema com Go;
  - TS2 odeia Go;
  - TS3 vê Go com bons olhos;
  - TS4 já pretendia migrar para Go.
- **Go:**
  - GO1 não tem problema com TS;
  - GO2 odeia TS e Node;
  - GO3 vê TS com bons olhos;
  - GO4 já pretendia migrar para TS.
- **Neutros:** N1 (performance), N2 (dev solo e conforto), N3 (integração com o Unity), N4 (moderador).
- **Mercado:** MK1 (trabalho), MK2 (backend de jogos), MK3 (ecossistema e IA).

Todos leram o briefing com quem você é, o repo, a LINHA-DO-TEMPO e o seu código antigo em Go (MUD e idle) e em TS (semi-idle).

## Como a conversa funcionou
1. **Rodada 1 (abertura):** cada um defendeu a sua posição, mostrou como a sua linguagem resolve 5 problemas do projeto e fez perguntas diretas a agentes nomeados.
2. **Rodada 2 (conversa):** todos leram a transcrição completa (`conversa.md`), responderam pelo nome a cada pergunta recebida, rebateram 2 falas, cederam em algo e fizeram uma pergunta nova.
3. **Rodada 3 (fechamento):** responderam às perguntas da rodada 2, propuseram o protocolo do teste e disseram a inclinação. O N4 juntou tudo no protocolo final.

A transcrição inteira, com as 45 falas em ordem, está em `conversa.md`.

## Rodada 1: o que cada um trouxe
- **Medições que mudaram o tom:**
  - **N1:** mediu o mesmo `aplicar` nas duas: 1.000 lutas em 29 ms (Node) e 2 ms (Go), com resultados idênticos. Conclusão: "performance é argumento de fachada".
  - **N3:** mediu o servidor para Windows: 5,8 MB (Go) contra 86 MB (Bun). O Node SEA ainda está em desenvolvimento ativo.
- **N2 olhou o seu código:**
  - **Go:** foi onde você mais escreveu à mão (7 semanas depois do reset de abril), mas sem testes e com estado guardado nos objetos.
  - **TS:** no único commit seu à mão, a função de luta muda a entrada, e dois dias depois a IA tomou o núcleo.
- **Ataques:**
  - **TS2:** sem união discriminada, Go é ruim para eventos.
  - **GO2:** o compilador do TS 7 foi reescrito em Go, e o seu núcleo TS antigo nunca teve os tipos checados.
- **Autocrítica:**
  - **TS1 e TS3:** `number` é float, e o `Math.random` apareceu nos seus testes TS.
  - **GO1 e GO3:** copiar slice e map compartilha dados, e o seu MUD usava `rand.Intn` global.
- **N4:** lembrou que a dúvida Go × TS já foi reaberta 7 vezes e propôs pesos, com autoria e risco de troca em 25%.

## Rodada 2: onde a conversa virou
- **TS1:** respondeu às 13 perguntas que recebeu e propôs travar o TS com `deepFreeze`, um checador de "só inteiros" e `tsc --noEmit` no teste. Concedeu que a vantagem de IA do TS não pesa aqui.
- **GO1:** admitiu que a versão pura do seu Go veio do Copilot e que você chegaria à noite 5 na noite 6 ou 7 em Go.
- **GO2:** concedeu que percorrer `map` em Go sai em ordem aleatória (o seu `score.go` quebra empates assim) e estimou 5 noites em TS contra 6–7 em Go.
- **GO3:** cedeu ao N2 que "Go te força a ser simples" não se confirmou no seu histórico.
- **TS2, TS3 e TS4:** mediram o Bun e admitiram que Go ganha a distribuição. TS4 admitiu que o conforto em TS não impediu você de mexer na entrada.
- **N3:** corrigiu o GO1 (a doc do Node SEA fala, sim, em compilação cruzada) e rebaixou o contrato com C# para empate.
- **MK1 e MK2:** cederam que carreira e mercado não desempatam.
- **De onde veio o desempate:** GO2, TS4, N2 e GO3 propuseram, cada um do seu jeito, decidir por um teste de 1 noite escrito por você. Virou consenso.
- **N4:** listou o que empatou e o que ainda decide: autoria, noite 5, ferramental e distribuição.

## Rodada 3: fechamento
- **TS1:** admitiu que não tem prova de autoria em TS, porque os commits à mão pendem para o Go.
- **TS2:** aceitou `tsc` strict e um pacote só, e propôs que empate abaixo de 20% dê Go.
- **TS4:** propôs empate abaixo de 25% dando Go, e "se ele não fizer o teste em 2 noites, fica Go".
- **Os 4 do Go:** aceitaram TS sem reclamar se o Go levar 1,5× o tempo ou se só a versão Go precisar de IA.
- **N4:** fechou o protocolo. Placar: Go leve 7, TS leve 3, empate 4, e ninguém com margem clara.
- **Pontos sem resposta:**
  - o modelo de eventos em Go (GO1 propõe struct achatada, GO4 interface + registro);
  - o tamanho do `node.exe`;
  - não há prova de que você escreve função pura em nenhuma das duas.

## Quem mudou de ideia
| Agente | Mudou em quê |
|---|---|
| GO3 | Desceu: Go não te levou a função pura nem a testes |
| GO4 | Determinismo virou "vantagem leve"; cedeu carreira |
| TS1 | De TS para empate |
| TS4 | Cedeu carreira e distribuição; ficou "TS leve, quase empate" |
| N1 | Distribuição de vantagem clara para moderada do Go |
| N2 | O peso técnico subiu de 10% para ~20%, puxando para o Go |
| N3 | Contrato e "sem framework" empatam |
| MK1, MK2 | Carreira e mercado deixaram de desempatar |

## Incidente
Na rodada 3, o limite da sessão caiu enquanto os agentes devolviam os resumos. Todos já tinham salvo as respostas completas antes do corte, então nada se perdeu. A rodada do moderador rodou depois que o limite voltou.
