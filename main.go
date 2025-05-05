package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"goCrudMySQL/internal/config"
	"goCrudMySQL/internal/db"
	"goCrudMySQL/internal/user"
	"goCrudMySQL/router"
)

func main() {
	// 加載環境變數
	_ = godotenv.Load()

	// 讀取環境配置
	cfg := config.Load()

	// 連接到 MySQL
	gormDB := db.NewMySQL(cfg.MySQL.DSN)

	// 初始化用戶儲存庫和服務
	userRepo := user.NewRepository(gormDB)
	userSvc := user.NewService(userRepo)

	// 創建路由器
	r := router.New(gin.Default(), user.NewHandler(userSvc))

	// 設置 HTTP 伺服器
	srv := &http.Server{
		Addr:           ":" + cfg.Server.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("server listening on :%s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 優雅關閉伺服器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
