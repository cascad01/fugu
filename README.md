### fugu - инструмент для генерации getter`ов

```shell
# сборка бинарника
go build -o ./bin/gengetter ./cmd/getter-gen/main.go

# будет сгенерирован файл в формате {pkg_name}_getters.go
./bin/gengetter ./example
```