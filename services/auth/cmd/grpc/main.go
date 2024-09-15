package main

import (
	"fmt"

	"github.com/geekengineers/Microservice-Project-Demo/services/auth/config"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary"
	redis_adapter "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/secondary/redis"
	"gorm.io/driver/sqlite"
)

func main() {
	cfg := config.Read()

	redisClient := redis_adapter.GetRedisDBInstance(&redis_adapter.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	dialector := sqlite.Open(fmt.Sprintf("./database/%s.db", config.CurrentEnv.String()))

	// Bootstrap Application
	primary.Bootstrap(&primary.BootstrapRequirements{
		RedisClient: redisClient,
		Dialector:   dialector,
		Grpc: struct {
			Host string
			Port int
		}{
			Host: cfg.Grpc.Host,
			Port: cfg.Grpc.Port,
		},
		Jwt: struct{ PrivateKey string }{
			PrivateKey: cfg.Jwt.PrivateKey,
		},
	})
}
