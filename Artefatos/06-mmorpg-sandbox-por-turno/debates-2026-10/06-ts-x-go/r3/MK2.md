# MK2: rodada 3 (Mercado, backend de jogos)

## 1. Respostas

Ninguém me fez pergunta na rodada 2. O MK1 fez uma réplica à minha fala e eu concordo com ele: o retorno do Go é de **habilidade**, não de currículo (`conversa.md`, linha 1424). A minha pergunta ao TS2 sobre o Nakama, o ES5 e o `BigInt` continua aberta. Ela pesa pouco, porque o GO1 e o TS1 concordam que o servidor fica próprio até existir necessidade.

## 2. Protocolo do teste de 1 noite

Aceito o teste do TS4 com o critério fechado do GO2 (linha 1122). Acrescento duas coisas que vêm de prática de benchmark: **viés de ordem** e **tabela comum**.

| Item | Proposta |
|---|---|
| Antes (fora do cronômetro) | Ele escreve **à mão** 5 casos de tabela em um `casos.json`, usado pelas duas versões: acerto, erro, crítico, dano que zera PV e sorteio com semente fixa. O mesmo arquivo serve de semente para o contrato. |
| O que escreve | `aplicar(estado, ataque)` puro e o laço de teste lendo o `casos.json`. O laço pode ser da IA, como o briefing permite. |
| Tempo | 50 min por linguagem, com 10 min de intervalo. |
| Ordem | Sorteada na moeda. Quem vai em segundo já resolveu a lógica uma vez, então a ordem fixa enviesa o resultado. |
| Mede | 1) minutos até 5/5 verdes; 2) quantas travadas de mais de 5 min; 3) consultas: documentação vale, código de IA não vale; 4) se mutou a entrada, conferido por um caso que compara a cópia do estado antes e depois. |
| Empate | Diferença menor que 20% no tempo, ou os dois sem fechar. Nesse caso leva **Go**, pelos desempates que sobraram mensuráveis: binário de 5,8 MB × 86 MB (N3) e um toolchain só. |
| Validade | Definitiva até a fase jogável, pela regra anti-troca (N4). Uma linha com data no `docs/03-tecnico.md`. |

Rejeito uma variação: "terminou primeiro ganha", sem margem. Com 2h e um bebê, uma interrupção de 5 min decide o resultado. A margem de 20% absorve isso.

## 3. Inclinação atual

**Go, leve.** O mercado empata em adoção, ecossistema, IA e hospedagem (MK1, MK3, N4). Inclino por dois fatos que não empataram:
- **Distribuição sem plataforma:** a lição do Hathora/Stormgate favorece o binário próprio, e o de Go é o menor e o mais simples.
- **Prática transferível:** a de Go vale para o trabalho, mas é peso pequeno (concessão da rodada 2).

## 4. Condição

Aceito TS sem reclamar se ele:
- fechar TS com 5/5 e Go não fechar em 50 min; ou
- fechar TS com margem de 20% ou mais.

Nesse caso a fluência vence a noite 5, e para um dev solo é isso que o mercado premia: jogo entregue, em qualquer linguagem.
