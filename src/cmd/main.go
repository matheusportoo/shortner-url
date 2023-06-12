package main

import (
	"github.com/matheusportoo/shortner-url/src/adapters/inbound/api"
	"github.com/matheusportoo/shortner-url/src/adapters/outbound/redis_service"
)

func main() {
	rds := redis_service.CreateConnection()
	defer rds.Close()

	api.Run()
}
