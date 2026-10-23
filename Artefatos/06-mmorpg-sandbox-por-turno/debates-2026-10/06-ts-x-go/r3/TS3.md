# TS3 – Rodada 3

## 1. Respostas
**@N4:** *"Com que prova o TS é a linguagem da autoria? Vocês se comprometem com `tsc --noEmit` strict e sem monorepo?"*
Prova eu não tenho. O acervo favorece o Go, com 55 commits à mão. O meu argumento é fluência, e o teste vai confirmar ou derrubar isso. Sobre o compromisso: **sim**, com `tsc --noEmit` strict, sem monorepo e com o lint abaixo, todos dentro do `npm test` desde a noite 1.

**@N1:** *"Qual regra exata? Pega `x / 2` num helper importado de fora?"*
Testei agora (ESLint 9 + `typescript-eslint`, 2 dependências). O mesmo arquivo deu 4 erros: import de fora, `/`, `Math.random` e `Date`.
```js
{ files: ["regras/**/*.ts"], languageOptions: { parser: tseslint.parser }, rules: {
  "no-restricted-syntax": ["error", { selector: "BinaryExpression[operator='/']" },
                                    { selector: "AssignmentExpression[operator='/=']" }],
  "no-restricted-properties": ["error", { object: "Math", property: "random" }],
  "no-restricted-globals": ["error", "Date", "performance"],
  "no-restricted-imports": ["error", { patterns: ["../*"] }] } }
```
O lint olha um arquivo por vez. O helper de fora não é pego pela regra do `/`, mas o import em si é bloqueado, e por isso `regras/` não importa nada de fora.

**@N3:** *"Versão do Node, MB do `node.exe`, servidor órfão?"*
A versão seria o Node 24 LTS. O tamanho não consegui medir, porque o proxy bloqueou o download. Estimo que seja da ordem dos 86 MB do Bun, e é estimativa. Para o servidor órfão, o Unity abre o processo com `RedirectStandardInput`. O servidor faz `process.stdin.on("end", () => process.exit(0))` e, a cada 2 s, roda `process.kill(pidDoPai, 0)` para ver se o Unity ainda existe. Concordo com você: em Go é igual.

## 2. Protocolo
- **Antes, separado:** ele escreve a tabela de 5 casos (crítico, acerto, erro, falha crítica, alvo morto), sem linguagem e sem cronômetro.
- **Ordem:** a escolha é por moeda, e as duas versões ficam em **noites diferentes**. Quem vai em segundo herda o raciocínio da primeira. A moeda espalha esse viés, e as noites separadas evitam o cansaço.
- **Tempo:** 60 min por linguagem. O setup entra no cronômetro.
- **Medidas** (aceito a régua do GO4):
  1. minutos até a tabela ficar verde;
  2. perguntas à IA, contando as de sintaxe;
  3. código pedido à IA, inclusive setup e `Clonar()`, que reprova;
  4. teste de entrada intacta;
  5. no dia seguinte, ele explica o próprio código sem consultar nada.
- **Empate:** diferença menor que 20% em (1), com (2) parecido. Nesse caso **ganha o Go**. A fluência é o último argumento do TS e, se não aparecer, sobram o binário, o `uint32` e a prática do trabalho.
- **Validade:** é definitiva até a fase 5, e o resultado fica registrado no `noites.md`. A regra de 1,5× corta escopo, não troca de linguagem.

## 3. Inclinação
TS, por margem **leve** e só até o teste.

## 4. Condição
Aceito Go sem reclamar se ele empatar ou for mais rápido em Go, ou se precisar pedir à IA o setup do TS.
