## Sidinei:

Eu quero agora começar a fazer o combate usando o TDD estou na fase atual sobre combate. 
Queria entender um coisa, qual o papel do apply.go, combat.go nisso? É o ideial mesmo usar o apply para que ele serve ? qual o conceito de tras dele?

---

## Claude:

Conferi o `main`: **`apply.go` e `combat.go` ainda não existem**. Em `core/` só há `dice.go`. Então eles são o que você vai criar agora. O papel de cada um:

**O conceito: uma porta só**

O `03-tecnico` já decidiu `aplicar(estado, ação, agora) → (estado, eventos)`. O `Apply` é essa função. Ele é a **única porta** por onde o jogo muda.

- Entra: o estado atual + o que o jogador quer fazer (`Strike`).
- Sai: o estado novo + o que aconteceu (eventos: "rolou 14", "acertou", "tirou 6 PV").
- Não lê relógio, não guarda nada fora, não altera o estado que entrou.

O nome do padrão é **reducer** (Redux, Elm) ou *functional core, imperative shell*: o núcleo é puro e quem está em volta (sim, servidor) cuida de entrada, tela e rede.

**Por que serve para o Eras**

| Ganho | Onde aparece na fase 1 |
|---|---|
| Mesma entrada → mesma saída | "a mesma semente repete a luta" |
| Teste é tabela: estado + ação → esperado | TDD sem mock |
| O sim e o servidor chamam a mesma função | 1.000 lutas agora, servidor na fase 2 |
| Eventos separados do estado | o cliente anima pelos eventos |

**Divisão dos arquivos**

- `apply.go` → só o `switch` no tipo de ação, que chama a função certa. Fica fino, sem regra dentro.
- `combat.go` → a regra: `strike(estado, golpe)` com 1d20 + bônus vs Defesa, usando o seu `degree`.

Isso evita que o `Apply` vire uma função gigante quando entrarem Andar, Passo, coleta etc.

**Para o seu primeiro teste vermelho**

A tarefa diz "estado criado pelo teste". Então o primeiro teste chama `Apply(State{...}, Strike{...})` sem que nada exista. O que não compilar te diz o campo mínimo que falta criar.

**Um ponto para checar:** hoje o `Dice` é `*Dice` e guarda a `seed` dentro dele. Pela regra do técnico, semente e contador moram **no estado**, como inteiros, sem ponteiro. Vale decidir se o `Roll` passa a receber `(seed, counter)` direto.

---
