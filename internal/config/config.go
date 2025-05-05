package config

import "os"

type Config struct {
	MySQL  struct{ DSN string }
	Server struct{ Port string }
}

func Load() Config {
	var c Config
	c.MySQL.DSN = getenv("MYSQL_DSN", "root:root@tcp(localhost:3308)/go_crud?charset=utf8mb4&parseTime=True&loc=Local")
	c.Server.Port = getenv("PORT", "8080")
	return c
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
