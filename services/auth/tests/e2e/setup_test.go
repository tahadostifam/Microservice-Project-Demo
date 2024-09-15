package auth_integration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary"
	redis_adapter "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/secondary/redis"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/pkg/otp_manager"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/utils"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/sqlite"
)

var client auth.AuthClient
var otpManager otp_manager.OtpManager

func TestMain(m *testing.M) {
	// signalCh := make(chan os.Signal, 1)
	// signal.Notify(signalCh, os.Interrupt)

	err := os.Setenv("GO_ENV", "test")
	utils.HandleError(err)

	redis_adapter.GetRedisTestInstance(func(redisClient *redis.Client) {
		const grpc_port = 8006

		dialector := sqlite.Open("../../database/test.db")

		// Bootstrap Application (Test ENV)
		go func() {
			primary.Bootstrap(&primary.BootstrapRequirements{
				RedisClient: redisClient,
				Dialector:   dialector,
				Grpc: struct {
					Host string
					Port int
				}{
					Host: "127.0.0.1",
					Port: grpc_port,
				},
				Jwt: struct{ PrivateKey string }{
					PrivateKey: "samplejwtsecret",
				},
			})
		}()

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", "127.0.0.1", grpc_port), grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
		utils.HandleError(err)
		defer conn.Close()

		otpManager = *otp_manager.NewOtpManger(redisClient)

		client = auth.NewAuthClient(conn)

		m.Run()
	})
}
