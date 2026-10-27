# owl-ontology-versioning

CLI-инструмент для **версионирования OWL-онтологий**: сравнение двух снимков (preview / diff), отчёт о добавленных и удалённых сущностях, конфликтах меток и эвристике совместимости структуры.

Отдельный продукт: кроссплатформенный бинарник, без привязки к конкретному веб-приложению. Сборки публикуются в **GitHub Releases** (Linux, Windows, macOS).

## Возможности

- structural diff двух OWL/TTL/RDF по IRI сущностей;
- сводка: added / removed / common / label conflicts;
- цветной TUI-отчёт в терминале;
- машиночитаемый JSON (`--json`, `--json-stdout`) для скриптов и интеграций;
- эвристика совместимости схемы (object / datatype properties).

## Установка

Скачайте бинарник с [Releases](https://github.com/Pwatermelon/owl-ontology-versioning/releases/latest) под вашу ОС и положите в `PATH`.

Примеры latest:

```text
owl-versioning-linux-amd64
owl-versioning-linux-arm64
owl-versioning-windows-amd64.exe
owl-versioning-windows-arm64.exe
owl-versioning-darwin-amd64
owl-versioning-darwin-arm64
```

macOS / Linux:

```bash
chmod +x owl-versioning-darwin-arm64
sudo mv owl-versioning-darwin-arm64 /usr/local/bin/owl-versioning
owl-versioning version
```

## Быстрый старт

```bash
owl-versioning preview base.owl incoming.owl
owl-versioning preview base.owl incoming.owl --json-stdout
owl-versioning preview base.owl incoming.owl --json report.json
owl-versioning version
```

## Сборка из исходников

```bash
go build -ldflags="-X main.VERSION=0.3.0" -o bin/owl-versioning ./cmd/owl-versioning
./bin/owl-versioning version
```

Кросс-сборка:

```bash
CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o dist/owl-versioning-linux-amd64 ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build -o dist/owl-versioning-linux-arm64 ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/owl-versioning-windows-amd64.exe ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o dist/owl-versioning-windows-arm64.exe ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o dist/owl-versioning-darwin-amd64 ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o dist/owl-versioning-darwin-arm64 ./cmd/owl-versioning
```

## Релизы

Тег `vX.Y.Z` запускает CI и создаёт GitHub Release с бинарниками и `SHA256SUMS.txt`:

```bash
git tag v0.3.0
git push origin v0.3.0
```

| Файл | Платформа |
|------|-----------|
| `owl-versioning-linux-amd64` | Linux x86_64 |
| `owl-versioning-linux-arm64` | Linux ARM64 |
| `owl-versioning-windows-amd64.exe` | Windows x86_64 |
| `owl-versioning-windows-arm64.exe` | Windows ARM64 |
| `owl-versioning-darwin-amd64` | macOS Intel |
| `owl-versioning-darwin-arm64` | macOS Apple Silicon |
| `SHA256SUMS.txt` | контрольные суммы |

Скачать конкретную версию:

```text
https://github.com/Pwatermelon/owl-ontology-versioning/releases/download/v0.3.0/owl-versioning-darwin-arm64
```

## Команды

| Команда | Назначение |
|---------|------------|
| `preview <base> <incoming>` | Diff двух версий онтологии |
| `--json` / `--json-stdout` | JSON-отчёт |
| `--limit N` | Сколько строк показать в TUI |
| `version` | Версия CLI |

## Интеграция

CLI самодостаточен: вызывайте как процесс или в пайплайне CI.

```bash
owl-versioning preview a.owl b.owl --json-stdout | jq '.summary'
```

## Лицензия

См. репозиторий. Обратная связь и issues приветствуются.
