# owl-ontology-versioning

CLI **версионирования OWL-онтологий**: сравнение двух снимков (preview / diff).

**Без Docker.** Сборки — бинарники Linux и Windows, публикация в **GitHub Releases**.  
Каталог асан (`asana-services`) при сборке API-контейнера **скачивает** `owl-versioning-linux-amd64` из нужного Release.

## Локальная сборка

```bash
go build -ldflags="-X main.VERSION=0.3.0" -o bin/owl-versioning ./cmd/owl-versioning
./bin/owl-versioning version
./bin/owl-versioning preview base.owl incoming.owl
./bin/owl-versioning preview base.owl incoming.owl --json-stdout
```

Кросс-сборка (как в CI):

```bash
CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o dist/owl-versioning-linux-amd64 ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build -o dist/owl-versioning-linux-arm64 ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/owl-versioning-windows-amd64.exe ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o dist/owl-versioning-darwin-amd64 ./cmd/owl-versioning
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o dist/owl-versioning-darwin-arm64 ./cmd/owl-versioning
```

## Релизы

1. Создай репозиторий на GitHub и запушь код.
2. Выпусти тег:

```bash
git tag v0.3.0
git push origin v0.3.0
```

3. Workflow `Release CLI binaries` положит в Release:

| Файл | Платформа |
|------|-----------|
| `owl-versioning-linux-amd64` | Linux x86_64 (тянет asana-services) |
| `owl-versioning-linux-arm64` | Linux ARM64 |
| `owl-versioning-windows-amd64.exe` | Windows x86_64 |
| `owl-versioning-windows-arm64.exe` | Windows ARM64 |
| `owl-versioning-darwin-amd64` | macOS Intel |
| `owl-versioning-darwin-arm64` | macOS Apple Silicon (M1/M2/…) |
| `SHA256SUMS.txt` | контрольные суммы |

Скачать latest (примеры):

```text
https://github.com/Pwatermelon/owl-ontology-versioning/releases/latest/download/owl-versioning-linux-amd64
https://github.com/Pwatermelon/owl-ontology-versioning/releases/latest/download/owl-versioning-darwin-arm64
https://github.com/Pwatermelon/owl-ontology-versioning/releases/latest/download/owl-versioning-windows-amd64.exe
```

Конкретная версия:

```text
https://github.com/Pwatermelon/owl-ontology-versioning/releases/download/v0.3.0/owl-versioning-darwin-arm64
```

## Связь с asana-services

В `asana-ontology-merge-api` Dockerfile:

```dockerfile
ARG OWL_VERSIONING_VERSION=latest   # или v0.3.0
# curl …/releases/…/owl-versioning-linux-amd64 → /usr/local/bin/owl-versioning
```

Переменная compose/CI: `OWL_VERSIONING_VERSION`.

## Команды

| Команда | Назначение |
|---------|------------|
| `preview <base> <incoming>` | Diff двух OWL/TTL |
| `--json` / `--json-stdout` | JSON-отчёт |
| `version` | Версия CLI |

## Тема НИР

Методы / средства версионирования OWL-онтологий.
