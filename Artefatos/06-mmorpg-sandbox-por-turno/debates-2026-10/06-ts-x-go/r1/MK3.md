# MK3 (Mercado): ecossistema, estabilidade e IA

## 1. Dados atuais

**Adoção e tendência**
- Stack Overflow 2025: TS é usado por 43,6% (48,8% dos profissionais), Go por 16,4%. Na pergunta "quer usar", TS tem 29,2% e Go 19,4%. Na pergunta "usou e quer continuar", TS tem 72,4% e Go 55,8%. https://survey.stackoverflow.co/2025/technology
- JetBrains 2025: Go é a linguagem que mais devs pretendem adotar (11%; TS 6%). TS, Go e Rust lideram o Language Promise Index. https://blog.jetbrains.com/research/2025/10/state-of-developer-ecosystem-2025/
- GitHub Octoverse 2025: TS passou a ser a linguagem nº 1 em contribuidores mensais (2,64 milhões, +66% no ano). O relatório liga isso à IA: 94% dos erros de compilação em código gerado por LLM são erros de tipo. https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/

**Trajetória e estabilidade**
- O TypeScript 7 é o compilador portado para Go, cerca de 10× mais rápido. O beta saiu em 21/04/2026 e a versão final em julho. O pacote mudou padrões: `strict` agora vem ligado e `module` vem como `esnext`. Removeu ou deixou como erro `moduleResolution node10`, `baseUrl` e `target es5`. A API programática só fica estável na 7.1. https://devblogs.microsoft.com/typescript/announcing-typescript-7-0-beta · https://adtmag.com/articles/2026/07/10/microsoft-bets-typescript-future-on-a-native-compiler.aspx
- Go 1.26 (fev/2026) traz `new(expr)`, genéricos que se referem a si mesmos, o coletor de lixo Green Tea como padrão e um `go fix` que moderniza o código sozinho. Tudo segue a promessa de compatibilidade do Go 1. https://go.dev/doc/go1.26
- Runtimes JS:
  - O Node 24.12 LTS tornou estável a remoção de tipos: roda `.ts` direto, mas só com sintaxe que dá para apagar (sem `enum`) e sem checar tipos. https://x-cmd.com/blog/251218
  - A Anthropic comprou o Bun em 03/12/2025, que segue com licença MIT. https://anthropic.com/news/anthropic-acquires-bun-as-claude-code-reaches-usd1b-milestone
  - O Deno 2 continua como terceira opção.

**Assistência de IA (benchmarks)**
- Multi-SWE-bench (corrigir issues em repositórios reais): TS e JS "consistentemente" tiveram as menores taxas. A média que calculei pela Tabela 4 é parecida: Go ~3,5% e TS ~3,4%. https://arxiv.org/abs/2504.02605
- Multi-LCB (2026): Go e TS ficam no mesmo patamar, abaixo de Python. No GPT-OSS-120B, Go fez 69,9 e TS 70,3. https://arxiv.org/pdf/2606.20517
- Leitura: nos benchmarks, as duas linguagens empatam.

**Backend de jogos**
- Nakama (Heroic Labs) é um servidor de jogo em Go que roda módulos em TS, JS e Lua. https://heroiclabs.com/blog/announcing-nakama/
- Agones (Google, servidores dedicados) é feito em Go. https://www.mintlify.com/googleforgames/agones/introduction
- Colyseus é o equivalente em Node/TS. https://colyseus.io
- As duas linguagens têm ecossistema de backend de jogo.

**Distribuição junto do Unity**
- Go compila em um binário só, com cross-compile nativo.
- No lado JS existem `bun build --compile` (https://bun.sh/docs/bundler/executables) e o executável único do Node (SEA), que ainda é experimental (https://nodejs.org/api/single-executable-applications.html).

## 2. O que isso significa para este dono

1. **Longevidade não decide.** As duas linguagens estão em alta e têm dono forte (Microsoft e Google). Quem usar "risco de morrer" como argumento está forçando.
2. **A estabilidade é de tipos diferentes.**
   - Em Go, a linguagem e o ferramental são uma coisa só (`go.mod` + `go test`). Os dois Go antigos do acervo tinham uma dependência direta só: `coder/websocket` em `.../03-mud-moderno-pre-fusao/.../server/go.mod`.
   - Em TS, a linguagem é estável, mas o redor muda: padrões de tsconfig a cada versão grande, ESM×CJS, a escolha entre Node, Bun e Deno, o runner de teste. O semi-idle (`.../05-semi-idle-typescript/eras-do-brasil-semi-idle`) já tinha workspaces npm, Vitest 5, Vite/React, três tsconfig e mise. Parte desse projeto foi escrita por IA, então ele mostra o tamanho do ferramental, não as escolhas dele.
   - Opinião: para quem tem 2h por noite e trocou de stack 5 vezes, cada decisão de ferramental é uma porta para a "troca de novo".
3. **A IA empata.** O que muda é o contexto. Em TS, a IA erra mais pela configuração (módulo, runtime) do que pelo código. Em Go, erra mais por estilo verboso. Nenhuma das duas justifica a escolha.
4. **Carreira.**
   - TS tem o mercado maior (Stack Overflow, Octoverse) e é a especialidade dele.
   - Go é a linguagem que mais cresce em intenção de adoção e é a do emprego atual.
   - Opinião: praticar Go à noite rende no trabalho de dia. Praticar TS rende em vagas futuras. Os dois casos são defensáveis. A diferença é que um tem uso hoje e o outro é uma aposta.
5. **Remoção de tipos favorece um estilo.** Ela não aceita `enum`, o que empurra o TS para uniões discriminadas com literais. Para os eventos de luta, isso é bom e cabe sem build.

## 3. Perguntas

- **TS1:** O TS 7 mudou padrões e removeu opções (`baseUrl`, `node10`), e o `tsgo` só ganha API estável na 7.1. Qual combinação exata você congela na noite 1 (runtime, versão do TS, runner de teste, número de arquivos de config)? Como garante que isso não vire uma "noite de ferramental" na noite 15?
- **GO2:** O Sidinei se diz iniciante em Go, e boa parte do acervo Go foi escrita por IA. Que evidência do próprio acervo mostra código Go escrito por ele? E sem união discriminada nativa, como você modela os eventos de luta sem depender de `interface` + `type switch`, onde esquecer um caso só aparece em tempo de execução?
