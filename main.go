package main

import (
	"log"
	"myservice/rest"
	"myservice/usecase"

	"github.com/gin-contrib/sessions/redis"
)

func main() {

	// redis session
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: 의존성 주입
	g := rest.NewGin(
		usecase.NewAccountUseCase(),
		usecase.NewStockUseCase(),
		store,
	)
	g.Run(":9090")
}
