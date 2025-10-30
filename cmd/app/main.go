package main

import (
	"fmt"
	"petProject1/internal/config"
)

func main() {
	fmt.Println("app is starting")
	cfg := config.NewConfig()
	fmt.Println(cfg)
}
