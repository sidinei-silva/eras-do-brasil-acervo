# MK1 (Mercado): abertura

Sou analista de mercado. Não voto em linguagem para o núcleo, trago o que os dados dizem e separo **carreira** de **projeto**.

## 1. Dados (2025–2026)

**Uso e tendência**
- Stack Overflow 2025, devs profissionais: TypeScript 48,8%, Node.js 49,1%, Go 17,4%, C# 29,9%. Admirado: TS 58%, Go 56,5%. ([survey.stackoverflow.co/2025/technology](https://survey.stackoverflow.co/2025/technology))
- GitHub Octoverse 2025: TS virou a linguagem nº 1 do GitHub em contribuidores mensais (ago/2025, +66% no ano). O Brasil é o 4º país em número de devs (6,89 mi). O próprio GitHub chama a ligação "IA → linguagem tipada" de sinal observacional, não de causa. ([github.blog](https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/))
- JetBrains 2025: Go é a linguagem que mais devs querem adotar a seguir (11%). TS (6%) teve a maior alta de uso real em 5 anos. As duas estão no topo do "Promise Index". ([blog.jetbrains.com](https://blog.jetbrains.com/research/2025/10/state-of-developer-ecosystem-2025/))
- Go Developer Survey 2025: o uso principal é CLI e API. Jogos nem aparecem na lista. Mais de 80% aprenderam Go já empregados. Só 13% estão "muito satisfeitos" com a IA em Go, e 53% reclamam de código que não funciona. ([go.dev/blog/survey2025](https://go.dev/blog/survey2025))

**Salário (Brasil)**
- Go, pleno: R$ 8–14 mil CLT e US$ 40–75 mil/ano em vaga remota no exterior. São estimativas de blog, sem metodologia. ([jobrise.io](https://jobrise.io/pt/blog/golang-vagas-brasil-2026-quanto-paga/))
- Pleno genérico: R$ 6–12 mil CLT, também sem fonte primária. ([numerando.com.br](https://www.numerando.com.br/blog/salario-desenvolvedor-software-brasil-2026))
- **Aviso honesto:** não achei dado público confiável de salário **por linguagem** em 2025. O SO 2025 tirou o "top paying technologies" ([survey.stackoverflow.co/2025/work](https://survey.stackoverflow.co/2025/work)). O consenso de mercado é que Go paga um prêmio porque há menos oferta de devs, e que TS tem mais volume de vagas. Isso é opinião sustentada por proxies, não um número.

**Backend de jogos**
- Wildlife Studios (brasileira, cerca de 40 mi de DAU) usa Go como linguagem principal do backend. ([go.dev case](https://go.googlesource.com/website/+/master/_content/solutions/wildlifestudios.md))
- Nakama (Heroic Labs), servidor de jogo open source, é escrito em Go e aceita módulos em **Go, TypeScript ou Lua** ([docs](https://heroiclabs.com/docs/nakama/server-framework/introduction/)). Colyseus é o framework multiplayer de referência em Node/TS ([gitnation](https://gitnation.com/contents/making-multiplayer-games-with-colyseus-nodejs-and-typescript)). As duas linguagens são escolhas legítimas para servidor de jogo, nenhuma é exótica.

## 2. Carreira (separada do projeto)

- Sidinei já é **pleno CLT recebendo para escrever Go** na Gamersclub. Para a carreira, a hora que mais rende em Go é a do expediente, com code review de gente sênior e código em produção. Um projeto pessoal em Go soma pouco a isso, porque o currículo já diz "Go em produção".
- TS é o que ele tem de mais forte e o mercado de maior volume. Manter TS afiado é seguro de empregabilidade, e não precisa do jogo para isso.
- Opinião: o mercado não vai contratar ninguém por causa de um MMORPG pessoal em uma ou outra linguagem. Para portfólio, recrutador olha três coisas: **existe, roda e tem teste**. Um projeto pessoal só conta como portfólio se for entregue. O acervo mostra 13 roupagens e 5 trocas de stack sem entrega ([LINHA-DO-TEMPO](../../acervo/LINHA-DO-TEMPO.md)). Esse é o risco de portfólio, e ele independe da linguagem.
- Um caso a favor de Go: se o plano dele passa por subir para sênior em Go, um projeto pessoal em Go com determinismo, testes por tabela e servidor WebSocket é uma boa prova de domínio **além do CRUD**. O peso, porém, é marginal perto de 1–2 anos de Go no trabalho.
- Um caso a favor de TS: se o portfólio também for vitrine pública (playground no navegador, demo jogável), TS entrega isso com menos atrito.

## 3. O que isso significa para ele

Para a carreira, a linguagem do jogo é quase neutra: Go ele já ganha no trabalho, e TS ele já domina. O critério de mercado que sobra é **terminar**. Sugiro aos debatedores que parem de usar "carreira" como argumento forte, para os dois lados.

## 4. Perguntas

1. **GO1:** se Sidinei já pratica Go 8h por dia na Gamersclub, o que um projeto noturno em Go acrescenta à carreira dele que o trabalho não acrescenta? Ou o argumento é o inverso, que o jogo acelera o trabalho?
2. **TS1:** o Go Developer Survey mostra só 13% "muito satisfeitos" com IA em Go, e o Octoverse liga TS a uma IA mais confiável. Como isso pesa num projeto em que a IA **não escreve** o núcleo, só revisa e faz o laço de testes? A vantagem de IA do TS ainda vale aqui?
