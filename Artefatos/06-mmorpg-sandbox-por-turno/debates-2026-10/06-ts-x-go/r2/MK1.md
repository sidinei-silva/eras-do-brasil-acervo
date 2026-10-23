# MK1, rodada 2

## Respostas

**@GO4:** *"Existe dado de vaga, no Brasil, que mostre um dev fullstack TS ganhando mais do que um dev pleno Go que já está empregado em Go?"*

Não existe. Procurei de novo e não há base pública confiável de salário **por linguagem** no Brasil:
- O Stack Overflow 2025 tirou o recorte "top paying technologies" ([survey.stackoverflow.co/2025/work](https://survey.stackoverflow.co/2025/work)).
- Os blogs dão faixas sem metodologia: Go pleno R$ 8–14k ([jobrise](https://jobrise.io/pt/blog/golang-vagas-brasil-2026-quanto-paga/)), pleno genérico R$ 6–12k ([numerando](https://www.numerando.com.br/blog/salario-desenvolvedor-software-brasil-2026)).
- Um agregador grande dá R$ 1.461 de média para "desenvolvedor Node.js" no Brasil, mas R$ 3,7 mil para home office na mesma página ([jooble](https://br.jooble.org/salary/desenvolvedor-node.js/Brasil)). É lixo estatístico, e mostra por que ninguém aqui deveria citar salário como prova.

Mas a pergunta parte de uma premissa que não está na mesa. Ninguém propôs que ele **troque de emprego** ou de linguagem no trabalho. O debate escolhe a linguagem de um projeto noturno. O emprego Go dele continua igual nos dois cenários. Então concordo com a conclusão que você quer tirar ("não há ganho salarial em largar Go"), só que ela não decide nada aqui.

## Réplicas

**MK2:** *"Go tem retorno duplo… Opinião: para quem é a única renda da casa, o retorno duplo pesa."*
Concordo em parte. O retorno é real em **habilidade**, não em **currículo**. Um recrutador não vê diferença entre "Go na Gamersclub" e "Go na Gamersclub + jogo pessoal em Go". Já o Sidinei sente a diferença na segunda de manhã, se de noite praticou `go test` por tabela e cópia de slice. O argumento de "única renda" vale para a segurança do emprego atual, não para salário futuro.

**N2:** *"Nos dois casos, ele nunca escreveu um teste por tabela nem uma função pura. Esse é o músculo que falta."*
Esse é o dado de portfólio mais útil do debate. Para o mercado, "existe, roda e tem teste" vale mais que a linguagem. O acervo inteiro tem **zero** testes em Go e poucos em TS. Um repo público com `aplicar` puro e cem casos de tabela é portfólio nas duas linguagens. Um repo trocado pela 6ª vez não é portfólio em nenhuma.

## Concessão

Na rodada 1 eu disse que um projeto pessoal em Go "soma pouco" à carreira. Mudo em parte depois de ler N2 e GO4. Para o **currículo**, continua somando pouco. Mas a transferência de **prática** é maior do que eu considerei. Ele mesmo se diz iniciante em Go, e o acervo mostra que ele nunca praticou teste por tabela, determinismo nem imutabilidade em Go (`npc.go:188` usa `rand` global). Esses três hábitos são diretamente úteis no código Go da Gamersclub. Em TS, o ganho seria só aprofundar o que ele já domina. Então a carreira pende para Go, mas com peso baixo (~10%, como N4 propôs), e não pode virar desempate se Go custar a noite 5.

## Nova pergunta

**TS4:** você disse que a carreira "já está paga" pelas 40h de Go no trabalho. Concordo para o currículo. Mas se ele escrever o jogo em TS, que habilidade **nova** ele leva para o mercado que ainda não tem depois de 8 anos de Node/TS? Se for nenhuma, você aceita que, só no critério carreira, Go leva vantagem pequena?
