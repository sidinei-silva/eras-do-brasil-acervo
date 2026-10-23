# Rodada 2 · I3 (cético de IA e autoria)

## 1. Disputas
1. **Rotação:** fixa. Se o grid com altura esconder casas, a correção é baixar o relevo, não girar a câmera.
2. **Grupo:** fora do MVP. Grupo traz o relógio por turno, que é a parte mais difícil de escrever à mão.
3. **Jogadores por zona:** tanto faz. Servidor único e 20 por zona no MVP.
4. **Offline:** cálculo fechado ao logar, com teto. Sem corpos offline povoando o mundo: fingir população é desonesto.
5. **Testes:** ele escreve os casos (entrada e saída esperada); a IA só sugere bordas. Num motor d20, o teste é a regra.
6. **Placeholder de IA:** nunca. O caso *Clair Obscur* mostra que trava de build falha. Cubo cinza não vaza.
7. **Lore:** IA aceitável só como caçadora de furo, que é o que ele já faz desde 20/10, sem propor texto. Mantenho que nome e frase não saem dela.
8. **Mercado:** único com ordens de compra. Com pouca gente, mercados locais morrem.

## 2. Ataque e defesa
- **Defendo T1/T3/G2 (offline como cálculo fechado).** Além da CPU, é a única versão que ele consegue escrever e testar à mão: uma função pura de tempo, cota e ferramenta. Simulação ao vivo de lutas offline é o tipo de sistema que ele acabaria pedindo pronto à IA, e isso reabre o ciclo de abril e de 14/10.
- **Ataco em parte A1/T2 (exportador e validação a cada build desde o dia 1).** O mundo sobre o grid, sim. Mas pipeline de exportação e validação no Unity, com 2h por noite e sem C#, vira semanas de ferramenta antes de existir jogo. Na fase 4 bastam casas marcadas à mão na greybox e as 3 arenas recortadas a olho. O exportador é ferramenta descartável: se vier, pode ser gerado por IA, fora do jogo.

## 3. Obrigatórios no MVP (máx. 5)
1. Offline com cálculo fechado, cota própria e teto em horas (T1/T3/J3).
2. Tela "enquanto você esteve fora" em 10 segundos (J3).
3. Torneira do comprador NPC com cota e preço decrescente (G2).
4. Grid marcado no terreno da greybox (A1/T2, versão mínima).
5. Contrato de IA escrito no AGENTS.md antes da fase 1. Sem ele, o MVP não sobrevive ao próximo reset.

**Podem esperar:** grupo, canais, PvP contra eco, intenção visível do inimigo, simulação de 50 jogadores, variedade de rostos (entra na fase de arte, não antes).

## 4. Cardápio
Concordo, e defendo: 2 opções, cada uma com o custo dela, e a escolha registrada como dele.

**Par no tema IA (o que entra no jogo é igual nas duas: zero conteúdo gerado):**
- **A. Contrato restrito.** IA só explica, revisa e caça furo. Ele escreve código, casos de teste e textos. IA gera só ferramentas fora do jogo. *Custo:* mais lento.
- **B. Contrato com testes.** Igual ao A, mas a IA escreve o código dos testes a partir de uma tabela de casos que ele ditou. *Custo:* precisa conferir se o teste diz o que a tabela diz.
