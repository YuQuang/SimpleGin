# SimpleGin
練習專案，實作用戶認證後端以及RestFul API以及專案目錄管理。

## Run database migration
安裝
[golang-migrate](https://github.com/golang-migrate/migrate)
然後執行
``` bash
migrate -database postgres://<user>:<pwd>@<host>:<port>/<dbname>?sslmode=disable -path ./migrations up
```

## How to run
```
go run ./cmd/main.go
```