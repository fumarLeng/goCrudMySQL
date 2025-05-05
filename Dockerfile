# 使用官方 Go 映像作為基礎映像
FROM golang:1.23 AS builder

# 設置工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum，並下載依賴
COPY go.mod go.sum ./
RUN go mod tidy

# 複製源代碼
COPY . .

# 編譯 Go 應用
RUN go build -o myapp.exe .

# 使用較小的映像來運行應用
FROM alpine:latest

# 安裝需要的庫（例如，運行 Go 應用可能需要的一些庫）
RUN apk --no-cache add ca-certificates

# 從 builder 映像中複製已編譯的應用
COPY --from=builder /app/myapp.exe /myapp.exe

# 設置環境變數（如需要）
ENV DB_HOST=mysql
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=root
ENV DB_NAME=go_crud

# 開放容器端口
EXPOSE 8080

# 設置容器啟動命令
CMD ["/myapp.exe"]