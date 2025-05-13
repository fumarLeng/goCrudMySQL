# Go CRUD API 練習

使用 Go 語言、Gin 框架與 GORM ORM 實作的簡易 RESTful API，並整合 MySQL 與 Swagger。使用 Docker 容器化

##  使用技術

- Go 1.23
- Gin Web Framework
- GORM ORM
- MySQL 5.7
- Swagger (via swaggo)
- Docker + Docker Compose

##  專案結構

```
goCrudMySQL/
├── internal/
│   └── user/               # 使用者 CRUD 模組
├── router/                 # 路由設定
├── docs/                   # Swagger 自動生成文件（swag init）
├── init/init.sql           # MySQL 初始化 SQL（建表/預設資料）
├── main.go                 # 主程式進入點
├── Dockerfile              # Go 應用的容器設定
├── docker-compose.yml      # 正式用 Docker Compose
├── docker-compose.dev.yml  # 本地開發用設定（可選）
└── README.md               # 本說明文件
```

##  啟動方式

---

### 1. 在本機直接開發/測試

```bash
go mod tidy            # 下載 Go 依賴
go run main.go         # 啟動 API（需建立 MySQL 資料）
```

### 2. 開發者（本地 build image）
```bash
git clone https://github.com/fumarLeng/goCrudMySQL.git
cd goCrudMySQL
docker-compose -f docker-compose-dev.yml build --no-cache
docker-compose -f docker-compose-dev.yml up
```

> 透過 `Dockerfile` 建立本地 Go image，然後啟動 Go + MySQL 容器。

---

###  3. zip檔案透過 docker-compose 啟動

參考 `go-crud-deploy.zip`，使用者可於任何有安裝 Docker 的環境啟動容器並測試 API。

📦📦 [下載壓縮檔](https://github.com/user-attachments/files/20184588/go-crud-deploy.zip)📦📦

```bash
# 解壓 zip
unzip go-crud-deploy.zip
cd go-crud-deploy

# 啟動容器（會從 Docker Hub 拉下 fumarleo/go-crud-api:latest）
docker compose up
```

##  Swagger API 文件

啟動完成打開：  
[http://localhost:8088/swagger/index.html](http://localhost:8088/swagger/index.html)

若 zip 啟動後 DB 沒有資料，執行 `docker compose down -v` 清除舊的資料 volume，再重新載入 init.sql 的假資料。

---
