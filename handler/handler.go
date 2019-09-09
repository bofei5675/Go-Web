package handler

import (
	"Go-Web/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}){
	code, message := errno.DecodeErr(err)
	// always return status OK here

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}