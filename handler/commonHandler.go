package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/models"
	"test/service"

	"github.com/gin-gonic/gin"
)

type CommonHandler struct{
	CommonService service.CommonService
}

func NewCommonHandler(commonService service.CommonService)CommonHandler{
	return CommonHandler{CommonService: commonService}
}

func (ch *CommonHandler) Task1Handler(c *gin.Context){
	var taskModel models.Task1Model
	body, err := io.ReadAll(c.Request.Body)
	if err != nil{
		return
	}
	fmt.Println("The Request Body is : ",string(body))
	if err := json.Unmarshal(body, &taskModel); err!=nil{
		fmt.Println("Unable to Marshal data")
		return
	}

	result, err := ch.CommonService.Task1Service(taskModel)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to complete task1"})
	}
	c.JSON(http.StatusOK, result)
}