package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct{
	
	Status int `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
}

type response struct{

	Data interface{} `json:"data"`
}

func Failure(c *gin.Context, status int, message string) {
	response := errorResponse{Status: status, Code: http.StatusText(status), Message: message}
	c.JSON(status,response)
}

func Success(c *gin.Context, data interface{}){
	c.JSON(200,data)
}