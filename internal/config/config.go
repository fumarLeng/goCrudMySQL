package config

import (
	"fmt"
	"os"
)

type Config struct {
	MySQL  struct{ DSN string }
	Server struct{ Port string }
}

func Load() Config {
	var c Config
	c.Server.Port = getenv("PORT", "8080")

	// 拆環境變數組 DSN
	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "3306")
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASSWORD", "root")
	name := getenv("DB_NAME", "go_crud")

	c.MySQL.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	return c
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
