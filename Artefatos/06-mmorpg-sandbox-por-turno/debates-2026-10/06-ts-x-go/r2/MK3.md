# MK3, rodada 2

## 1. Respostas

Ninguém me fez pergunta na rodada 1: procurei "MK3" em `conversa.md`. As minhas perguntas ao TS1 (ferramental congelado) e ao GO2 (autoria e eventos) continuam abertas.

## 2. Réplicas

**@GO2:** você disse: "Até a Microsoft desistiu de rodar o compilador do TypeScript em Node e o reescreveu em Go [...] por que o servidor do Sidinei deveria?"

- Discordo do raciocínio, não do fato. O compilador do TS processa projetos de milhões de linhas, depende de CPU e de memória compartilhada entre threads.
- A própria equipe explicou que escolheu Go porque a estrutura do código existente se parecia com Go, o que facilitava portar arquivo por arquivo, e não porque Node seja ruim para servidor (https://github.com/microsoft/typescript-go/discussions/411).
- O servidor do Eras é o caso oposto. O N1 mediu microssegundos por `aplicar` e a CPU parada mais de 99% do tempo.
- O TS 7 serve a outro argumento, que é o meu: ele mudou os padrões do `tsconfig` e removeu opções em 2026 (https://devblogs.microsoft.com/typescript/announcing-typescript-7-0-beta). Isso mostra que o ferramental do TS se mexe. Não mostra que o Node é inadequado.

**@TS2:** você disse: "O TS facilita colar código de IA, porque é a linguagem em que a IA mais tem exemplo."

- Concordo com a metade sobre volume. TS é a nº 1 em contribuidores no GitHub (https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/).
- Mais exemplo não virou mais acerto:
  - No Multi-SWE-bench, "TS e JS consistentemente" tiveram as menores taxas de resolução (https://arxiv.org/abs/2504.02605).
  - No Multi-LCB, Go e TS empatam (https://arxiv.org/pdf/2606.20517).
- Para este dono, a IA não escreve o núcleo, então isso pesa pouco. Concordo com a regra do N4: "a IA escreve mais rápido" conta contra.
- Um dado conta a favor do TS: 94% dos erros de compilação em código gerado por LLM são erros de tipo (Octoverse). Se a IA só revisa, um compilador de tipos mais rico também acusa os erros dele. Isso só vale se o `tsc --noEmit` rodar de fato, coisa que não acontecia no semi-idle (GO2, TS4).

## 3. Concessão

- Na rodada 1, eu disse que, em TS, "cada decisão de ferramental é uma porta para a troca de novo". Depois de ler TS1 e TS4, cedo em parte.
- Se o ferramental for congelado na noite 1, o risco cai para algo próximo do Go:
  - Node LTS, que roda TS sem build: a remoção de tipos ficou estável no 24.12 (https://x-cmd.com/blog/251218);
  - Vitest;
  - `tsc --noEmit` com `strict` no `npm test`.
- O que sobra é externo e tem data: a versão grande do TS (7.0 → 7.1) e a troca de LTS do Node. São eventos de calendário, não noites perdidas.
- Ressalva: o acervo mostra que o congelamento não aconteceu em outubro (monorepo, Vite, mise). O plano funciona, mas o histórico de cumpri-lo é ruim.

## 4. Nova pergunta

**@GO4:** você disse que Go tem "um ferramental que não tem nada para configurar". Só que GO1, GO2 e GO3 admitiram que a exaustividade dos eventos depende de `exhaustive` ou `gochecksumtype`. Isso entra como `golangci-lint` com `.golangci.yml` e versão fixa, como `go vet` com analisador próprio, ou não entra? Se não entrar, quem pega o `case` esquecido de `reacao`: o teste por tabela ou a IA revisora?
