# GO4: rodada 1

## 1. Posição

Sou dev de Go e quero migrar para TypeScript na minha carreira: uma linguagem no front e no back, mais vagas fullstack, o ecossistema do npm. Aceito bem o TS: ele é uma boa linguagem e, para a carreira de muita gente, a escolha certa. Só que a pergunta aqui não é "o que é bom para um dev". A pergunta é "o que é bom para este núcleo, deste dono, com 2h por noite". Para isso defendo Go, e explico a diferença:

- **Bom para a carreira:** amplia o mercado no futuro. TS ganha.
- **Bom para o projeto:** reduz atrito e decisões hoje e não esconde erro de regra. Go ganha por três motivos: `int64` de verdade, um binário só e um ferramental que não tem nada para configurar.
- **Detalhe que inverte o meu caso pessoal:** o Sidinei *já está* em Go no trabalho (Gamersclub) e se diz iniciante. Cada noite de Go no jogo também é treino para o emprego que paga a casa. Para ele, carreira e projeto apontam para o mesmo lado. Para mim, não.

## 2. Os 5 problemas em Go

**`aplicar` puro.** Go passa struct por valor, então campos simples já chegam copiados.
```go
func Aplicar(e Estado, a Acao) (Estado, []Evento) {
    e.Herois = slices.Clone(e.Herois) // slice e map NÃO copiam sozinhos
    // ... regras sobre e (cópia) ...
    return e, eventos
}
```

**Sorteio com semente no estado.** Inteiro sem sinal, com overflow definido e igual em toda máquina.
```go
type Rng struct{ S uint64 } // fica dentro de Estado, vai no JSON
func (r Rng) D20() (Rng, int) {
    r.S = r.S*6364136223846793005 + 1442695040888963407
    return r, int(r.S>>33)%20 + 1
}
```
Em TS, `number` é float: a multiplicação acima perde precisão. Seria preciso usar `BigInt` ou `Math.imul` em 32 bits.

**Tipos de evento.** Go não tem união discriminada. Uso uma interface marcadora mais um `switch` por tipo:
```go
type Evento interface{ tipo() string }
type Dano struct{ Alvo, Valor int }
func (Dano) tipo() string { return "dano" }
switch ev := e.(type) { case Dano: ...; case Movimento: ... }
```

**Teste por tabela.** É o idioma padrão do `testing`, sem lib nenhuma. O Sidinei escreve a tabela e a IA escreve só o `for`, exatamente o contrato do AGENTS.md.
```go
casos := []struct{ nome string; ca, rolagem int; acerta bool }{
    {"20 natural", 30, 20, true}, {"1 natural", 5, 1, false},
}
for _, c := range casos { t.Run(c.nome, func(t *testing.T) { /* ... */ }) }
```

**Servidor local junto do Unity.** `GOOS=windows go build` gera um `.exe` estático, sem runtime, que vai na pasta do jogo e o Unity abre com `Process.Start`. Com TS teríamos Node junto do jogo, ou `bun build --compile`, ou o Node SEA, que ainda é experimental. WebSocket vem de [coder/websocket](https://pkg.go.dev/github.com/coder/websocket@v1.8.15), que é uma lib e não um framework.

## 3. O ponto mais fraco do Go para ESTE dono

São dois, e sou honesto sobre eles:

1. **Imutabilidade não é garantida pelo compilador.** O `slices.Clone` acima existe porque slice e map compartilham memória. Se ele esquecer, a regra "não mexer no estado que entrou" quebra em silêncio. O `readonly` do TS pegaria parte disso na compilação. O código Go escrito à mão depois do reset de 03/04 mostra o risco: o MUD usa `rand.Intn` global, nada determinístico (`acervo/Artefatos/03-mud-moderno-pre-fusao/eras-do-brasil-game-arquivado/server/npc/npc.go:188`). Ou seja, Go não deu disciplina sozinho.
2. **Go não salvou o projeto antes.** Foram três fases em Go (MUD, Albion Idle/Escória, idle fundido) e todas caíram por troca de gênero e por IA no núcleo (`acervo/LINHA-DO-TEMPO.md`, marcos 6–10). Linguagem nenhuma resolve isso. Quem disser que resolve está vendendo.

## Quando eu mudaria de lado

- Se o Sidinei disser que pretende sair de Go no trabalho e fechar a carreira em TS. Aí o argumento de carreira passa a apontar para o TS e o meu "bônus duplo" some.
- Se o playground no navegador deixar de ser só teste e virar cliente de verdade.
- Se, no marco da noite 5, a falta de união discriminada e as cópias manuais custarem noites medidas no `noites.md`. O culpado tem que ser a linguagem, não o design.

## 4. Perguntas

- **TS2 (ou quem defender o TS no núcleo):** em TS, como vocês garantem o sorteio com semente em inteiros de 64 bits sem `BigInt` e sem limitar o gerador a 32 bits? E quanto o `Readonly<>` protege de fato, sendo raso?
- **MK1:** existe dado de vaga, no Brasil, que mostre um dev fullstack TS ganhando mais do que um dev pleno Go *que já está empregado em Go*? A troca de linguagem custa o tempo de reaprender o que ele usa todo dia.
