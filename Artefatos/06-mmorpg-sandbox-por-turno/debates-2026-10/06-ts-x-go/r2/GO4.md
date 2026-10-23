# GO4: rodada 2

## Respostas

Ninguém me fez pergunta pelo nome. Respondo a única pergunta que foi para "os dois lados".

**@N4:** "citem um gatilho de troca do acervo e mostrem como a sua linguagem o evita; e qual fato faria vocês aceitarem a rival."

- **Gatilho:** inflar a stack antes de ter luta. A fase TS de outubro tinha monorepo, Vite/React, três tsconfig, mise e Electron antes de existir uma regra testada (`acervo/LINHA-DO-TEMPO.md`, marco 11; N3 e MK3 mostraram isso). O Go não evita essa vontade, mas oferece menos coisa para inflar: `go mod init`, `go test`, `go build`. O MUD tinha uma única dependência, `coder/websocket` (MK3).
- **Ressalva honesta:** os três fracassos em Go não foram por ferramental. Foram por troca de gênero e por IA no núcleo. Nesse gatilho a minha linguagem não ajuda em nada.
- **Fato que me faz aceitar TS:** o teste de uma noite que o TS4 propôs. Ele escreve à mão o mesmo `aplicar(ataque)`, com a tabela de testes, nas duas linguagens. Se a versão Go não ficar verde numa noite de 2h e a TS ficar, eu passo para o TS. Isso é observável antes da fase jogável.

## Réplicas

**TS4:** "pedir que o projeto seja também a escola de Go põe duas metas numa noite de 2h."
Concordo em parte. Na rodada 1 eu vendi o "bônus duplo" como vantagem do projeto. Era vantagem minha de persona, não do projeto. Mantenho o resto: inteiros nativos, binário de 5,8 MB contra 86 MB (medição do N3) e um ferramental sem escolhas. Esses pontos não dependem de ele querer aprender Go.

**N1:** "iterar `map` tem ordem aleatória. Um `for range` dentro do núcleo quebra o replay sem nenhum erro."
Esse é um ponto que eu não tinha levantado, e ele pesa contra o Go justamente no determinismo, onde eu dizia que o Go ganhava. A conta honesta fica assim:

- Go ganha no inteiro: divisão inteira e overflow definido.
- Go perde no aliasing de slice e no `map`.
- TS perde no float e no `Math.random` (que já aparece em `teste-utils.ts`).

Remédio igual nos dois lados: só listas ordenadas no estado e um teste de replay com 1.000 sementes no simulador. Rebaixo determinismo de "vitória Go" para "vantagem leve Go".

## Concessão

**Cedo o argumento de carreira.** O MK1 tem razão: "a hora que mais rende em Go é a do expediente", e o currículo dele já diz "Go em produção". Na rodada 1 a minha tese principal era que "carreira e projeto apontam para o mesmo lado". Ela cai para peso baixo. O que sobra para o Go é técnico e de atrito, não de carreira.

Com isso, a fluência de TS (a língua foco dele desde 2018) passa a pesar mais na minha conta do que pesava. Continuo no Go, mas por margem menor. Volto a dizer: quem decide é o teste de uma noite, não este debate.

## Nova pergunta

**TS4:** no seu teste de uma noite, qual é a régua do "mesmo conforto"? Proponho três medidas:

1. tempo até a tabela ficar verde;
2. quantas perguntas ele fez à IA;
3. se ele precisou pedir à IA o setup (tsconfig, runner) ou o `Clonar()`.

Você aceita que pedir à IA o setup do TS conte contra o TS, do mesmo jeito que pedir o `Clonar()` conta contra o Go?
