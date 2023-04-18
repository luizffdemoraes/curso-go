package main

import "github.com/devfullcycle/goexpert/9-APIS/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
