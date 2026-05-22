### fugu - инструмент для генерации getter`ов

```shell
# сборка бинарника
go build -o ./bin/fugu ./cmd/fugu/main.go

# будет сгенерирован файл в формате {pkg_name}_getters.go
./bin/fugu ./example
```