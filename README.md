# SimpleGin
練習專案，實作用戶認證後端以及RestFul API以及專案目錄管理。


## Run database migration
Install
[golang-migrate](https://github.com/golang-migrate/migrate)
then run
``` bash
migrate -database postgres://<user>:<pwd>@<host>:<port>/<dbname>?sslmode=disable -path ./migrations up
```
For rollback please run
``` bash
migrate -database postgres://<user>:<pwd>@<host>:<port>/<dbname>?sslmode=disable -path ./migrations down
```


## How to run
``` bash
go run ./cmd/main.go
```

## Reference
參考資料:
1. Database Migration Tool
[golang-migrate](https://github.com/golang-migrate/migrate)

2. 參考如何建立專案檔案架構
[gin-directory-structure](https://www.compilenrun.com/docs/framework/gin/gin-fundamentals/gin-directory-structure/)

3. Swagger for gin
[swaggo](https://github.com/swaggo/swag)

4. JWT for go
