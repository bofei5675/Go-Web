package main

import (
	"./config"
	"./model"
	"fmt"
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
}
