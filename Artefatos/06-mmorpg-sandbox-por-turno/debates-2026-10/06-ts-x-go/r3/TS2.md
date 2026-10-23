# TS2 — rodada 3

## Respostas

**@N4:** "Com que prova o TS é a linguagem da autoria? Vocês se comprometem com `tsc --noEmit` strict e core sem monorepo? Sim ou não."
Prova no acervo eu não tenho. São 55 commits à mão em Go contra 1 em TS, e isso pende para o Go, como você disse. Compromisso: **sim** aos dois. O `tsc --noEmit` com `strict` entra no script `test`, e o `regras/` fica num pacote só.

**@N4:** "Condição para aceitar o Go: o TS2 não dá nenhuma."
Já dei na rodada 2, e ela vira o protocolo abaixo.

**@N1:** "Medido: confirma ou contesta? E o Nakama/goja?"
Confirmo com a minha medida: 86 MB e 18 MB de RSS. O Nakama eu descarto. Ele é framework, e o briefing pede "um processo, sem framework". O caminho TS da fase 6 é sempre um processo próprio. Também não garanto que o `Math.imul` rode no goja sem testar.

**@GO2** (sobre a minha pergunta): o `[MaxMobs]Mob` copiado por valor é uma boa resposta. Retiro o ataque de mutação para a luta 1×2.

## Protocolo do teste de 1 noite

- **O que escreve:** antes, ele escreve uma tabela única de 5 casos (erro, acerto, crítico, mob morre, semente igual → resultado igual). Depois implementa `aplicar(estado, ataque)` com d20 por semente no estado e os eventos `ataque` e `dano`.
- **Tempo:** 50 min por linguagem, na mesma noite. **O Go vai por segundo**, para compensar a falta de fluência, porque quem vai por segundo já resolveu a lógica. Se o TS ganhar mesmo assim, o resultado é robusto.
- **O que mede:**
  - minutos até a tabela passar;
  - travadas de mais de 5 min;
  - consultas à doc (contadas);
  - mutação da entrada (`deepFreeze` / `DeepEqual`);
  - nota de 1 a 5 para "quero continuar nessa".
- **IA:** pedir código desclassifica aquela linguagem. A IA só escreve o laço da tabela, igual nas duas.
- **Empate (diferença < 20%):** **ganha o Go**. Se a fluência empata, sobra o resto: distribuição, acervo de autoria e prática do trabalho, e tudo isso pende para o Go. Dói escrever isso.
- **Vale de forma definitiva:** sim, até a fase 2 jogável. Ele registra no `noites.md` e não rediscute.

## Inclinação

TS, por margem **leve**. A vantagem que sobra é a modelagem de eventos e a fluência, e a fluência ainda não foi medida.

## Condição

Aceito o Go sem reclamar se ele empatar ou vencer o teste acima, ou se o TS estourar 1,5× o orçamento da fase 1 por causa do ferramental (tsconfig, módulos, runtime), e não por causa do design.
