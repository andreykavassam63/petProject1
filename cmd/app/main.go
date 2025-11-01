package main

import (
	"fmt"
	"petProject1/internal/config"
)

func main() {
	cfg := config.NewConfig() //Парсим конфиг
	fmt.Println(cfg)

}
