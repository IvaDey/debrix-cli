<p style="text-align: center" align="center">
 <img src="docs/logo.svg" height="120" alt="Debrix logo"/>
</p>

<h1 style="text-align: center" align="center">Debrix CLI</h1>

## Configuration

Debrix supports custom configuration via a YAML file. By default, it looks for a file named .debrix.yml in the project root. You can also provide a custom path using the --config flag.

| Field        | Type     | Default                                                                                                                                                                                                                          | Description                                                                                                                                                                                    |
|--------------|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| pattern      | String[] | `['todo']`                                                                                                                                                                                                                       | A list of TODO patterns to search for in source files                                                                                                                                          |
| outFile      | String   | `TODO.md`                                                                                                                                                                                                                        | The output file path for the generated TODO list relative to the root                                                                                                                          |
| language     | String   | `en`                                                                                                                                                                                                                             | The language for CLI messages and generated output                                                                                                                                             |
| layout       | String   | `plain`                                                                                                                                                                                                                          | Output layout format                                                                                                                                                                           |
| exclude      | String[] | - `.git`<br/>- `.idea`<br/>- `.vscode`<br/>- `node_modules`<br/>- `dist`<br/>- `build`<br/>- `out`<br/>- `bin`<br/>- `vendor`<br/>- `third_party`<br/>- `venv`<br/>- `__pycache__`<br/>- `target`<br/>- `coverage`<br/>- `cache` | A list of paths or directories to exclude from scanning.                                                                                                                                       |
| linkTemplate | String   | `{{filePath}}:{{lineNumber}}`                                                                                                                                                                                                    | A template for generating clickable links to TODOs.<br/>You can use the following placeholders:<br/>- `{{filePath}}` – file path relative to project root<br/>- `{{lineNumber}}` – line number |

**Supported locales**
- English (en)
- Spanish (es)
- French (fr)
- German (de)
- Ukrainian (ua)
- Russian (ru)

**Supported layout values**
- plain
- table

## Supported languages

- JavaScript/TypeScript
- Go
- Swift
- Objective-C
- C/C++
- C#
- Java
- Kotlin
- Dart
- Bash shell
- Python
- HTML
- PHP
- CSS
- Lua
- Yaml files
- Ruby
- Rust
- Scala
- Elixir
- Perl