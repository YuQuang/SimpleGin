# SimpleGin
1. 實作 JWT 用戶認證後端
2. RestFul API
3. 專案目錄管理
4. Swagger
5. 資料庫 Migration
6. 自動化測試


## Index
- [Setup PostgreSQL](#SetupPostgreSQL)
- [Migration](#Migration)
- [Swagger](#Swagger)
- [Start](#Start)
- [Testcase](#Testcase)
- [Reference](#Reference)
---


## SetupPostgreSQL
拉取 [PostgreSQL Image](https://hub.docker.com/_/postgres) 然後 Run 一個 container
``` bash
# Pull image
docker pull postgres:19beta1-trixie

# Run image
docker run -itd -p 5432:5432 --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword postgres:19beta1-trixie
```

## Migration
安裝
[golang-migrate](https://github.com/golang-migrate/migrate)
然後執行
``` bash
migrate -database postgres://<user>:<pwd>@<host>:<port>/<dbname>?sslmode=disable -path ./migrations up
```
如果需要回滾則執行
``` bash
migrate -database postgres://<user>:<pwd>@<host>:<port>/<dbname>?sslmode=disable -path ./migrations down
```
如果需要做資料庫的更動，先開 Issue 選擇 migration 範本
``` bash
migrate create -dir .\migrations\ -ext sql -seq <Your_Migration_Script>
```
這會在 ./migrations 底下添加 up 跟 down 的腳本，接著就可以直接寫了


## Swagger
使用的套件是 swaggo
``` bash
swag init -g ./cmd/main.go
```
這會根據 controller 內的註釋產生 swagger 文件接著啟動伺服器
``` bash
go run ./cmd/main.go
```
開啟網頁 http://localhost:81/swagger/index.html

詳細關於 swaggo 的套件可以參考最後附上的參考資料


## Start
``` bash
go run ./cmd/main.go
```


## Testcase
如果需要執行測試的話需要先準備環境
``` bash
# 測試環境檔路徑
/configs/config.test.yaml
# 設定好 PostgreSQL 設定之後便可以開始測試
```
執行 Go test
``` bash
# -count=1 取消快取
# -v 顯示更多資訊
# 執行 tests 路徑底下所有測試
go test -count=1 -v ./tests/
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
