package main

import (
	"goCrudMySQL/controller"
	"goCrudMySQL/model"
	"goCrudMySQL/repository"
	"goCrudMySQL/router"
	"goCrudMySQL/service"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userSvc)

	r := router.NewRouter(userCtrl)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
