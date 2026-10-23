# Compilado: debates de stack e de arte

## Debate 1: stack (6 agentes, 3 rodadas)

**Como foi.** Cada agente defendeu uma stack diferente. A rodada 2 foi de ataque cruzado e a rodada 3 de posição final.

**Propostas iniciais**
| | Proposta | Argumento principal |
|---|---|---|
| A | Unity + núcleo C# + servidor .NET | um núcleo para sim, servidor e cliente |
| B | Unity + Go | regra mora no servidor; Go ele usa no trabalho |
| C | TypeScript total (Three.js + Electron) | é a linguagem dele |
| D | Godot C# + .NET | MIT; ele já montou pipeline Blender→Godot |
| E | Unity + SpacetimeDB | backend pronto, usado no BitCraft |
| F | Cético: TS até a fase 3, engine só no fim | nada de linguagem nova antes de ver a luta |

**A virada (rodada 2).** Todos aceitaram que o cliente não precisa do núcleo de regras. O servidor manda o estado, a prévia (casas alcançáveis e chance de cada ação) e os eventos. Com isso, a escolha virou 3 decisões separadas:
1. **Cliente:** Unity, Godot ou Three.js.
2. **Linguagem do núcleo e do servidor:** C#, Go ou TS.
3. **Servidor próprio ou backend pronto.**

**Resultado**
- Núcleo e servidor em TypeScript: 5 de 6 (Go ficou como alternativa).
- Servidor próprio com Node, WebSocket e Postgres: 5 de 6 (SpacetimeDB ficou como dissidência).
- Cliente em Unity, com Godot de reserva.
- Núcleo só com inteiros, para o resultado ser determinístico em qualquer linguagem.
- Regras anti-troca:
  - a stack só se rediscute com a fase em curso jogável;
  - cada rediscussão troca uma camada só, nunca o núcleo;
  - o cliente só troca se reprovar no teste do personagem modular.

**O que ficou mal resolvido.** O debate escolheu linguagem sem discutir como o jogo funciona (movimento, zona, luta visível, grupo, AFK), e você apontou isso. O `03-tecnico.md` ainda diz .NET.

## Debate 2: arte (6 agentes, 2 rodadas)

**Propostas iniciais**
| | Proposta |
|---|---|
| A | Low poly sombrio (Valheim, Wartales) |
| B | Pintado à mão adulto |
| C | Retrô PS1 |
| D | Xilogravura de cordel |
| E | Pixel art isométrico |
| F | Cético: tudo comprado de um fornecedor (Synty + Mixamo) |

**Consenso**
- Nada de chibi. O infantil vem de proporção (~6–7 cabeças) e de saturação.
- Base comprada: Synty Sidekick + Mixamo; nada de modelagem até o MVP.
- Material muda a silhueta e o tier muda só cor e acabamento: 3 faixas de forma (T1–3, T4–6, T7–8).
- Peças T6+ ganham detalhe com a luz do Rasgo, a única cor viva.
- Teste de 1 noite: 3 personagens no grid, na câmera real. Se não der para ler material e tier, o tratamento cai.
- Pixel 2D foi abandonado pelo próprio defensor.

**Resultado.** Low poly sombrio + contorno com 2 tons.

**O que ficou mal resolvido.** A câmera e a perspectiva não foram discutidas, e você apontou isso também.
