package main

import (
	"os"
	"time"

	"reused-modules-api/infrastructure"
	"reused-modules-api/infrastructure/database"
	"reused-modules-api/infrastructure/log"
	"reused-modules-api/infrastructure/router"
	"reused-modules-api/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
