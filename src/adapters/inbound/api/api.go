package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusportoo/shortner-url/src/adapters/inbound/api/controllers"
)

func Run() {
	r := gin.Default()

	r.GET("/:id", controllers.GetURL)
	r.POST("/url", controllers.CreateURL)

	r.Run()
}
