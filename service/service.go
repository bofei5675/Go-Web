package service


import (
	"Go-Web/model"
	"Go-Web/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

func AddUser(c *gin.Context){
	var r model.User
	// bind check Content-Type to select a binding automatically
	log.Printf("Start adding user %s", r)
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.User{
		UserName: r.UserName,
		Password: r.Password,
	}
	log.Printf("Parse user from request %s", u)
	// Validate the data
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Insert the user to the db
	log.Printf("Insert into DB")
	if _, err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase,nil)
		return
	}
	SendResponse(c, nil, u)

}

func SelectUser(c *gin.Context){
	name := c.Query("user_name")
	if name == ""{
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var  user model.User
	if err := user.SelectUserByName(name);nil != err {
		fmt.Println(err)
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// Validate the data.
	if err := user.Validate(); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}