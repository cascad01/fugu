### fugu - инструмент для генерации getter`ов

```shell
# сборка бинарника
go build -o ./bin/gengetter ./cmd/getter-gen/main.go

# будут сгенерированы 2 файла с геттерами в указанных пакетах в формате {pkg_name}_getters.go
./bin/gengetter ./example
```