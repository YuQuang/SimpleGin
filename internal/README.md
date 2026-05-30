# Internal
放置與專案有關的檔案主要為 Go 程式檔。

## api
放置路徑相關檔案 ex: route.go

## app
放置與初始化相關的檔案，伺服器的各種初始化，資料庫的初始化，Swagger的初始化等相關

## controller
有且僅有處理Http請求，解包Query、Body等資料不負責格式與業務邏輯。

## middleware
中介服務，例如：請求來之前可以負責驗證身分，校驗JWT是符合法，不合法即直接回傳401、403等

## model
放置資料庫取出來的資料的 Entity，一方面方便讓後續的 service 好操作，一放面不讓他們直接碰到 DB 的 struct

## repository
負責處理與資料庫溝通的邏輯，所需要的 Dependency 僅僅只是 DB 的connection，方便後續測試可以 mock DB 的 connection

## service
處理業務邏輯的地方，例如：驗證資料是否正確、用戶名稱是否合法、密碼強度是否足夠...都是在這裡處理。