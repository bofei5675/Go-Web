package service


import(
	"github.com/gin-gonic/gin"
	"../model"
	. "../handler"
	"../pkg/errno"
)

func AddUser(c *gin.Context){
	var r model.User
	// bind check Content-Type to select a binding automatically
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.User{
		UserName: r.UserName,
		Password: r.Password,
	}
	// Validate the data
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Insert the user to the db
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
	var user model.User
	if err := user.SelectUserByName(name) ; err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}
	if err := user.Validate(); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}