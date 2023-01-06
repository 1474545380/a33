package main

import (
	"store/config"
	"store/internal/repository"
)

func main() {
	config.InitConfig()
	repository.InitDB()
}
