# Rodada 1: resumo das 6 propostas

A. **Unity + .NET + núcleo C#** (proposta atual). Um núcleo, três consumidores (sim, servidor, cliente). Unity resolve "você é o que veste" (malha no esqueleto, Asset Store). C# é próximo de TS/Java/Kotlin. Risco: editor engolir as noites; mitiga com fases 0–1 só .NET. Custo até fase 2: ~2–2,5 meses.

B. **Unity + Go (núcleo e servidor em Go)**. Regras moram no servidor de qualquer jeito (caçada offline); cliente só anima; servidor manda a prévia de chance de acerto junto do estado do turno. Go ele usa no trabalho; ~1.460 linhas Go já existiram. Diz que C# não garante resultado idêntico entre Mono/IL2CPP e .NET. Risco: contrato de mensagens duplicado (mitiga com esquema gerando tipos C#). Custo: 10–14 semanas.

C. **TypeScript total** (núcleo TS, Node, Three.js, Electron). Linguagem dele; a F8 caiu por IA escrevendo o núcleo e design indefinido, não por TS. Risco: arte/animação sem editor (Blender como editor de cena). Gatilho: teste do personagem modular em 2 semanas; se falhar, só o cliente cai. Custo: ~30–40h de Three.js.

D. **Godot 4 com C# + servidor .NET**. Jogo pede pouco da engine; licença MIT; ele já montou pipeline Blender→Godot em set/2026. Risco: retargeting/personagem modular menos polido que Unity; teste de 2 noites com saída para Unity levando o núcleo. Custo: ~35–45h.

E. **Unity + SpacetimeDB (módulo C#)**. Scheduled reducers rodam caçada offline; persistência e sync prontos; usado pela BitCraft (MMO sandbox). Licença BSL 1.1 (uma instância em produção ok). Risco: lock-in; mitiga com núcleo C# puro. Ponto a verificar: módulo WASM referenciar o núcleo .NET Standard. Custo: ~2,5–3 meses.

F. **Cético: TypeScript até a fase 3, engine 3D só na fase 5** (cliente placeholder 2D em canvas). Jogo é por turno e autoritativo: cliente não precisa do núcleo. Nenhuma linguagem nova até ver a luta. Risco: placeholder 2D não empolga; Unity/Blender como estudo paralelo fora do repo. Regra: stack só se rediscute com a fase em curso jogável.

## Pontos de choque
- Precisa compartilhar o núcleo com o cliente? (A, D, E: sim; B, F: não, o servidor manda prévia e eventos)
- C# novo vs linguagem que ele domina (TS/Go).
- Engine desde a fase 2 ou só no fim.
- Backend pronto (E) vs escrever servidor.
