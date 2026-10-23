# Rodada 2: resumo

## Consensos (todos aceitaram)
- O cliente NÃO precisa do núcleo de regras. Servidor manda estado + prévia (casas alcançáveis, chance de cada ação) + eventos. Núcleo no cliente vira opcional (A, C) ou sai (B, D, E, F).
- Regra anti-troca: a stack só se rediscute quando a fase em curso estiver jogável.
- Teste do personagem modular (malha no osso, troca de arma e capacete, 1 animação) cedo, com prazo, como gatilho de saída só do cliente (A, C, D, F).
- Determinismo: núcleo só com inteiros; aí C#, Go e TS dão o mesmo resultado (B concedeu).

## Fatos novos
- Unity 6.x roda em Mono: C# 9 e .NET Standard 2.1. O 6.7 LTS é o último com Mono; Unity 7 vai para CoreCLR (.NET 10), otimização no LTS de 2027. Afeta só quem compartilha código C# com o cliente.
- Godot 4 com C# usa .NET 8+; exportação desktop estável.
- SpacetimeDB tem módulos em TypeScript oficiais; toolchain C#→WASM está em transição (experimental).

## Como o debate se reorganizou
Com o cliente separado do núcleo, viraram 3 decisões independentes:
1. **Cliente:** Unity (A, B, E) × Three.js (C, F) × Godot (D).
2. **Linguagem do núcleo e do servidor:** C# (A, D) × Go (B) × TypeScript (C, F, E agora).
3. **Servidor próprio × backend pronto** (SpacetimeDB com módulo TS, proposta E).

## Mudanças declaradas
- A: núcleo no cliente opcional; só inteiros; teste de arte na fase 2.
- B: fase 1 só Go no console; esquema JSON gera tipos C#.
- C: Electron só na fase 6; prévia oficial do servidor; se o teste do personagem falhar, cliente vai para Unity/Godot e núcleo continua TS.
- D: cliente só anima; teste do personagem antes da fase 2; se o servidor for Go, cliente em GDScript.
- E: núcleo sai do cliente; módulo SpacetimeDB em TS; gatilho: tutorial Unity + SpacetimeDB em 2–3 noites.
- F: canvas 2D vira Three.js na fase 2; une-se a C com a regra anti-troca.
