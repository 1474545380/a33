package main

import (
	"staff/config"
	"staff/internal/repository"
)

func main() {
	config.InitConfig()
	repository.InitDB()
}
