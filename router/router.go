package router
import (
	"Go-Web/router/middleware"
	"Go-Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(g *gin.Engine){
	middlewares := []gin.HandlerFunc{}
	// middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)
	// 404 handler

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "You sent the request to an incorrect API route.")
	})

	// The health check handlers
	router := g.Group("/user")
	{
		router.POST("/addUser",service.AddUser) // add user
		router.POST("/selectUser", service.SelectUser)
	}

}