package service


import (
	_ "../handler"
	"../model"
	"../pkg/errno"
	"github.com/gin-gonic/gin"
	"log"
)

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
	log.Printf("Start search user %s",name)
	if name == ""{
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var user model.User
	if err := user.SelectUserByName(name) ; err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}
	log.Printf("Finish Select.")
	if err := user.Validate(); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	log.Printf("Send back the user %s", user)
	SendResponse(c, nil, user)
}