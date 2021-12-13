# Loooooooooooooooooooong Pipeline

```shell
go run main.go -n 10000
# => 0.09s

go run main.go -n 100000
# => 0.77s

go run main.go -n 1000000
# => 7.24s

go run main.go -n 10000000
# => 176.72s
```

メモリ不足は確かめられなかった
