## Sidinei:

Pensando nessa estrutura:

/
├── core/         biblioteca de regras (C# puro)
├── server/       servidor .NET
├── sim/          simulador de console (playground de combate sem tela)
├── tests/        testes do núcleo
├── tools/        exportador do mapa para o servidor etc.
├── data/         itens, skills, mapas (JSON), lidos por todos
└── client/       projeto Unity

Onde docs vai ter obsidian como ficaria o .gitignore?

---

## Claude:

Com essa estrutura, o mais limpo é usar **um `.gitignore` por "dono"**, em vez de um arquivo gigante na raiz. Os padrões de um `.gitignore` aninhado valem relativos à pasta onde ele está. Assim cada ferramenta cuida do próprio lixo, e uma regra pensada para a Unity não pega por acidente algo do `server/`.

Ficariam três arquivos.

### 1. Raiz: o que é comum a tudo

```gitignore
# .NET (core, server, sim, tests, tools)
[Bb]in/
[Oo]bj/
TestResults/
*.user
*.suo

# IDEs
.vs/
.idea/
.vscode/*
!.vscode/extensions.json

# SO
.DS_Store
Thumbs.db

# Segredos locais
.env
*.local.json
```

- **`bin/` e `obj/` sem barra na frente** casam em qualquer nível. Isso é o que você quer, porque cada projeto .NET gera essas pastas.
- **`.vscode/*` com exceção `!extensions.json`** é opcional. Ignora suas configurações pessoais, mas versiona a lista de extensões recomendadas.
- **Não ignore `*.sln` aqui.** O `.sln` da raiz, que junta `core`, `server`, `sim` e `tests`, é seu e precisa ser versionado. Só o `.sln` que a Unity gera é descartável, e ele fica isolado no arquivo do `client/`.

### 2. `client/.gitignore`: Unity

```gitignore
/[Ll]ibrary/
/[Tt]emp/
/[Oo]bj/
/[Bb]uild/
/[Bb]uilds/
/[Ll]ogs/
/[Uu]ser[Ss]ettings/
/[Mm]emory[Cc]aptures/
/[Rr]ecordings/

# Gerados pela Unity, recriados ao abrir o editor
*.csproj
*.sln
*.slnx
*.pidb.meta
*.pdb.meta
*.mdb.meta

# Artefatos de build
*.apk
*.aab
*.unitypackage
```

- **A barra inicial `/`** ancora o padrão na pasta `client/`. `/[Ll]ibrary/` pega só `client/Library`, não uma pasta `Library` qualquer lá dentro de `Assets`.
- **`*.csproj` e `*.sln` aqui são seguros**, porque só valem dentro de `client/`. Os `.csproj` do `core/` e do `server/` continuam versionados.
- **Versione:** `Assets/`, `Packages/` (o `manifest.json` e o `packages-lock.json`), `ProjectSettings/` e **todos os `.meta`**. Se perder os `.meta`, as referências entre assets quebram.

### 3. `docs/.gitignore`: Obsidian

```gitignore
.obsidian/workspace.json
.obsidian/workspace-mobile.json
.trash/
```

Aqui a decisão é **o que fazer com a pasta `.obsidian/`**:

- **Versionar a pasta e ignorar só o `workspace*.json`** (o que está acima). O `workspace` muda toda vez que você abre ou fecha uma nota e só gera ruído nos commits. Já o resto (plugins instalados, atalhos, tema, templates) é configuração do vault e vale levar junto se você abrir em outra máquina.
- **Ignorar `.obsidian/` inteira** é a alternativa, se você não liga para sincronizar as configurações.
- **Atenção a `.obsidian/plugins/*/data.json`.** Alguns plugins guardam token ou API key aí. Dê uma olhada nos que você usa. Se algum guardar segredo, adicione esse caminho específico ao ignore.

### Pontos que dependem de decisões suas

- **Como o `core` chega na Unity?** Se for copiando uma DLL compilada para `client/Assets/Plugins/`, essa cópia é artefato e deve ser ignorada. A fonte da verdade é o `core/`. Se for por `asmdef` apontando para o código-fonte, ou por pacote local no `manifest.json`, não há nada a ignorar.
- **O mesmo vale para `data/`.** Se algum script copiar os JSON para `client/Assets/StreamingAssets/`, a cópia vai no ignore. O original continua em `data/`.
- **Saída do `tools/`.** Se o exportador escreve o mapa em `data/`, o resultado é versionado e conta como dado do jogo. Se escreve num diretório de build do servidor, ele é ignorado.

### Fora do `.gitignore`, mas no mesmo assunto

Com Unity no repo, vale configurar já um **`.gitattributes` com Git LFS** para os binários (`*.png`, `*.psd`, `*.fbx`, `*.wav`, etc.). Também vale marcar `*.unity`, `*.prefab` e `*.asset` como texto, com Force Text Serialization ligado nas configurações da Unity. Migrar para LFS depois que o histórico já tem binários pesados dá bem mais trabalho.

---
