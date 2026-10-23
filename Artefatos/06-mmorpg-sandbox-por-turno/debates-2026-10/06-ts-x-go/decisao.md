# Decisão: Go (22/10/2026)

O debate terminou em empate técnico: 7 agentes com Go leve, 3 com TS leve e 4 empates. O desempate veio de duas respostas do dono:

- **Go à noite parece extensão do expediente?** Não. No trabalho ele mexe mais em PHP, pouco em Go, e usa o Claude Code.
- **O projeto também serve para treinar Go?** Ainda em aberto. Ele mexe pouco em Node hoje (só um legado em Nest).

## Por que Go
- A principal vantagem do TS, a fluência, caiu: ele quase não usa Node hoje.
- O projeto pessoal é onde ele escreve à mão. Go escrito à mão treina a linguagem que ele revisa no trabalho.
- No acervo, foi em Go que ele mais escreveu sozinho.
- Para o projeto: binário pequeno junto do Unity, inteiro de verdade e `go test` sem configurar nada.

## Custos aceitos
- Eventos sem união discriminada: struct com campo `Tipo`.
- Nada de `map` dentro do estado.
- Cópia de slice e de map com cuidado; o teste de "entrada intacta" pega esse erro.

## Validade
- Até a fase 2 jogável, pela regra anti-troca.
