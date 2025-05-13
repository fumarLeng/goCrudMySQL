// @title Go CRUD API
// @version 1.0
// @description  GO語言  CRUD_TEST
// @host localhost:8088
// @BasePath /api/v1

package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goCrudMySQL/docs" // 匯入 swag init 產生的 Swagger 文件

	"context"
	"errors"
	"gorm.io/gorm"
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

	var gormDB *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		log.Printf("[%d/10] 正在連接資料庫...", i+1)
		gormDB, err = db.TryConnect(cfg.MySQL.DSN)
		if err == nil {
			log.Println("成功連接資料庫")
			break
		}
		log.Printf("連線失敗: %v", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("多次重試仍無法連接資料庫: %v", err)
	}

	// 初始化用戶儲存庫和服務
	userRepo := user.NewRepository(gormDB)
	userSvc := user.NewService(userRepo)

	// 創建路由器
	r := router.New(gin.Default(), user.NewHandler(userSvc))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 設定 HTTP 伺服器參數
	srv := &http.Server{
		Addr:           ":" + cfg.Server.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 啟動 HTTP 伺服器
	go func() {
		log.Printf("HTTP 伺服器啟動於 :%s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("伺服器啟動失敗：%s", err)
		}
	}()

	//  監聽中斷或終止訊號
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 嘗試關閉伺服器（最多等待 5 秒）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
