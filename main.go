package main

import (
	"test/handler"
	"test/repository"
	"test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	commonRepo := repository.NewCommonRepo(&gorm.DB{})
	commonService := service.NewCommonService(commonRepo)
	commonHandler := handler.NewCommonHandler(commonService)

	r := gin.Default()
	r.POST("/find-pairs", commonHandler.Task1Handler)
	r.Run(":5000")
}
