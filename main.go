package main

import (
	"Go-Web/config"
	"Go-Web/model"
	"Go-Web/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	fmt.Println("Configuration Initialized ...")
	if err := model.Init(); err != nil {
		panic(err)
	}

	fmt.Println("Model Initialized ...")

	gin.SetMode(viper.GetString("runmode"))

	// create new engine
	g := gin.New()

	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http %s", viper.GetString("addr"))
	if err := g.Run(viper.GetString("addr")); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
