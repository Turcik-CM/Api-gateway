package main

import (
	"api-gateway/api"
	config2 "api-gateway/pkg/config"
	"api-gateway/pkg/logger"
	"api-gateway/pkg/minio"
	redis2 "api-gateway/service/redis"
	"github.com/casbin/casbin/v2"
	"log"
	"os"
)

func main() {
	appLogger := logger.NewLogger()
	redis := redis2.NewRedisStorage(redis2.ConnectDB(), appLogger)
	config := config2.Load()

	err := minio.InitUserMinio()
	if err != nil {
		log.Fatal(err)
	}

	path, err := os.Getwd()
	if err != nil {
		appLogger.Error("Failed to get current working directory")
		return
	}

	casbinEnforcer, err := casbin.NewEnforcer(path+"/pkg/config/model.conf", path+"/pkg/config/policy.csv")
	if err != nil {
		panic(err)
	}

	controller := api.NewRouter(&config, appLogger, casbinEnforcer, redis)
	controller.Run(config.API_GATEWAY)
}
