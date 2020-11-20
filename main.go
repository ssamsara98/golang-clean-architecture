package main

import (
	"github.com/joho/godotenv"
	"github.com/ssamsara98/golang-clean-architecture/src/bootstrap"
)

func main() {
	_ = godotenv.Load()
	_ = bootstrap.RootApp.Execute()
}
