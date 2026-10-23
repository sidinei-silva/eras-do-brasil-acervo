# Rodada 1 (15 agentes): resumo

## Consensos (praticamente todos)
- Movimento: clique para andar (J3: + "ir até" nó/vila/zona). Mundo em tempo real, turno só na luta. WASD não.
- A luta no lugar NÃO bloqueia o mapa: aparece como "bolha" (grid no chão, ícone de espadas); quem está fora atravessa. Grupo de inimigos fica preso a quem encostou primeiro (operação atômica no servidor).
- AFK: solo sem relógio (luta espera; após muitos minutos, auto). Em grupo, relógio por turno; estourou, o automático do próprio jogador joga. Sem punição.
- Câmera 3/4 inclinada (~45–55°), perspectiva com FOV baixo (~30°, quase isométrica), zoom limitado, MESMA câmera no mundo e na luta (luta só aproxima e acende o grid).
- Stack: cliente Unity (API runtime do Sidekick só existe na Unity) + TypeScript/Node/WebSocket/Postgres serve. Tirar UDP do plano. O que muda é o desenho do servidor: dono por zona, área de interesse, luta como entidade da zona, relógio de turno no servidor, reconexão, luta salva como semente + log. 03-tecnico.md ainda diz .NET: corrigir.
- Arte: low poly sombrio + contorno 2 tons serve na 3/4. Teste da silhueta cinza no ângulo e zoom reais. Exagerar ombro, cabeça e arma (de cima, peito e bota quase não aparecem). PS1 e render 480×270 reprovam (A1: personagem cairia a 25–40 px).
- IA: sim em testes/explicação/revisão/ferramentas internas/simulação; não em arte final, voz, música, lore final, loja; NPC com LLM não (declaração de conteúdo ao vivo, custo, quebra determinismo, exploit de Where Winds Meet). Fofoca por modelos de frase com eventos reais do servidor.

## Fatos de IA levantados (com fonte nos agentes)
Steam jan/2026: declara-se conteúdo de IA no jogo e na loja (arte, som, texto, localização, marketing) e conteúdo gerado ao vivo; ferramentas de dev isentas. Quantic Foundry (dez/2025, 1.799, público core de PC): 85% negativos; diálogo 83% e missão 77% de rejeição. GDC 2026: 52% dos devs veem dano (64% artistas). GameDiscoverCo (jul/2026, ~3.800): 31% negativo, 8% não compram, código é o uso mais aceito. Game Oracle: jogos com selo de IA ~53% menos reviews. Casos: Clair Obscur perdeu Indie Game Awards por placeholder de IA que vazou; Larian recuou; Crimson Desert se desculpou; Vapor World 36% aprovação; Arc Raiders regravou vozes; Where Winds Meet NPC LLM explorado. Licenças de geração 3D variam (Hunyuan3D bloqueado em UE/UK/Coreia; TRELLIS.2 MIT). Asset Store proíbe treinar IA com assets. Brasil: Lei 9.610 só pessoa física é autor. Estudo Anthropic 2026: delegar código reduz compreensão.

## Disputas
1. Rotação: fixa sem rotação (maioria) × 4 ângulos de 90° (G1, T2, J2: grid com altura esconde casas).
2. Grupo: fora do MVP (T3, G2, G3, A3) × até 3 ou 4 desde cedo; turno por lado/simultâneo (T2, J2, J3, A1) × individual; relógio 30, 45 ou 60 s.
3. Jogadores por zona: 20 / 30–50; canais × servidor único.
4. Offline: simulação real × cálculo fechado ao logar com teto (T1, T3, G2). Corpo do jogador offline no mundo para povoar (J1) × nunca (G3).
5. Testes: IA escreve os testes (roteiro atual) × ele escreve os casos e a IA só sugere bordas (I3).
6. Placeholder de IA: sim com trava `ia_` no build (I2) × nunca (A3, I3, T2).
7. Lore: IA revisora de coerência (maioria) × nenhuma IA (I3).
8. Mercado: único com ordens de compra (G2) × locais como o Albion.

## Pontos críticos novos (um por agente)
- T1/T3/G2: offline não pode ser simulação ao vivo; cota própria que não gasta nós/mobs do mundo; teto diário; senão CPU e mercado quebram.
- G2: economia com pouca gente; comprador NPC vira torneira de prata → cota diária, preço decrescente, taxas, teto de ganho automático, simular 50 jogadores por 30 dias no sim/.
- J1: população baixa (DreamWorld, pico 56, retirado): mostrar gente; promessa de modo local se fechar.
- J3: a volta depois da interrupção: tela "enquanto você esteve fora" em 10 s; limite offline em HORAS; regras de saída do automático; notificação do sistema; trocar manual→auto no meio da luta não pode fugir da penalidade de morte.
- J2: "manual rende mais por luta" vira grind de lutas fáceis → bônus pela dificuldade; acelerar animação, pular turno inimigo, encerrar luta ganha; pouca luta com decisão.
- G1: intenção visível do inimigo (próxima ação anunciada, como Into the Breach).
- A1/T2: mundo construído sobre o grid desde o dia 1 (terreno em casas, objetos marcados bloqueia/cobertura/decoração, clareiras para luta, exportador e validação a cada build; teste: recortar 3 arenas e ver se leem). Shader de oclusão é custo escondido.
- A2: explosão de animações → 5 conjuntos de empunhadura; skills reaproveitam clipes.
- I2: orçamento de conteúdo da Travessia; variedade pelas regras, não por modelo novo.
- A3: Brasil nas pessoas (variedade real de rostos e tons de pele) e grafismo próprio por facção (não usar grafismo de povo real nem objeto de culto).
- G3: contradição da morte na fala do Bartolomeu; "firmar o pé" = ponto de retorno; fala do passo 1 longa; Enraizados devem viver normalmente.
- I1: não chamar o automático de "IA" na loja.
- I3: o próprio debate é IA gerando design → entregar ao dono um cardápio com 2 opções por ponto; ele decide.
