# TS1 — abertura

## 1. Posição
TypeScript para núcleo, simulador e servidor. Go é uma boa linguagem e eu a usaria sem drama: se o Sidinei escolher Go, o projeto não morre por isso (o acervo mostra que nenhuma versão morreu por defeito técnico, `acervo/LINHA-DO-TEMPO.md`). Defendo TS porque é a linguagem em que ele **já pensa** (Node/TS é a linguagem foco dele, briefing) e porque o núcleo pedido é quase todo modelagem de dados, onde TS é mais expressivo. Aceitação de Go: alta. Em distribuição, Go ganha, e eu digo isso abaixo.

## 2. Os 5 problemas

**`aplicar` puro.** `Readonly` faz o compilador recusar a mutação da entrada. Em Go não existe `const` para struct/slice; um `append` no slice da entrada passa calado.
```ts
type Estado = DeepReadonly<{ rodada: number; rng: number; lutadores: Lutador[] }>;
function aplicar(e: Estado, a: Acao): { estado: Estado; eventos: Evento[] } {
  // e.rodada++  -> erro de compilação
  return { estado: { ...e, rodada: e.rodada + 1 }, eventos: [] };
}
```

**Sorteio com semente no estado.** Aqui concedo: `number` é double, não inteiro. Resolve com disciplina: PRNG de 32 bits com `Math.imul` e `>>> 0`, e inteiros abaixo de 2^53. Funciona e é determinístico, mas o tipo não garante; Go com `uint32` garante.
```ts
function d20(rng: number): [valor: number, rng: number] {
  let t = (rng + 0x6D2B79F5) >>> 0;           // mulberry32
  t = Math.imul(t ^ (t >>> 15), t | 1);
  t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
  return [(((t ^ (t >>> 14)) >>> 0) % 20) + 1, (rng + 0x6D2B79F5) >>> 0];
}
```

**Tipos de evento.** É o ponto mais forte de TS. União discriminada com checagem de exaustividade nativa: esqueceu de tratar `reacao`, não compila. Em Go é interface + type switch, e a exaustividade depende de linter externo.
```ts
type Evento =
  | { tipo: "ataque"; de: string; alvo: string; rolagem: number }
  | { tipo: "dano"; alvo: string; valor: number }
  | { tipo: "movimento"; quem: string; para: Casa };
function nunca(x: never): never { throw new Error(String(x)); } // no default do switch
```

**Teste por tabela.** Ele escreve os casos; a IA só o laço. Com Vitest, o laço é uma linha (`it.each`). Ele já usava Vitest com o estilo "Dado/Quando/Então" em `fight.test.ts` (`acervo/Artefatos/05-semi-idle-typescript/.../packages/core/combat/fight.test.ts`). Atenção ao briefing: o histórico mostra que `position.ts` e o playground foram escritos por IA (diário de 14/10); os `types.ts` iniciais e o formato dos testes (commit `da17716`, 11/10) são os que eu leio como dele.
```ts
it.each([
  { caso: "crítico com 20", rolagem: 20, ca: 25, esperado: "critico" },
  { caso: "erra por 1",     rolagem: 10, ca: 21, esperado: "erro" },
])("$caso", ({ rolagem, ca, esperado }) => expect(grau(rolagem, ca)).toBe(esperado));
```

**Servidor local junto do Unity.** Concedo que Go é melhor: `GOOS=windows go build` gera um `.exe` sem runtime ([go.dev/doc/install/source](https://go.dev/doc/install/source#environment)). Em Node existe SEA com `node --build-sea`, mas está em "1.1 Active development" e exige empacotar num arquivo só ([docs Node](https://nodejs.org/download/release/latest/docs/api/single-executable-applications.html)); alternativa: `bun build --compile`. Mitigação real: até um amigo jogar, o servidor roda na máquina dele, então isso só pesa na fase 5. Não é critério para a noite 1.

## 3. Ponto mais fraco de TS para ESTE dono
O ferramental é um cardápio: tsconfig, ESM/CJS, Node×Bun×Deno, bundler para o SEA. Para quem já gastou dias de meta-organização em vez de código (LINHA-DO-TEMPO, linha 48), TS oferece mais botões para mexer. Go tem um jeito só. Minha mitigação: travar uma vez (Node LTS com type stripping + Vitest, sem bundler até a fase 5) e proibir mexer. Também é verdade que o repo TS foi onde a IA escreveu o core; a linguagem não causou isso, mas não ajudou a evitar.

## 4. Perguntas
- **GO1:** em Go, como o compilador impede que `aplicar` faça `append` ou altere um item do slice `lutadores` do estado que entrou? Se a resposta é "copiar à mão e revisar", isso não é exatamente o tipo de disciplina que a IA tende a pular?
- **GO2:** para os tipos de evento, você aceita depender do linter `exhaustive` para ter o que TS dá no compilador? Quem configura e mantém isso: ele, à mão, numa noite de 2h?
