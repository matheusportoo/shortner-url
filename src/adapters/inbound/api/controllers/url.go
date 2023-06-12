package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusportoo/shortner-url/src/domain/use_cases"
)

func GetURL(c *gin.Context) {
	id := c.Param("id")
	params := use_cases.RetrieURLParamsDTO{Id: id}

	err, url := use_cases.RetrieURL(params)
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusMovedPermanently, url.Link)
}

func CreateURL(c *gin.Context) {
	params := use_cases.SaveURLParamsDTO{}
	c.BindJSON(&params)

	err, url := use_cases.CreateURL(params)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, url)
}
