package usecase

import (
	"myservice/datasource"
	"myservice/repository"
)

type StockUseCaseInterface interface {
}

type StockUseCase struct {
	Repo repository.StockRepository
}

func NewStockUseCase() StockUseCaseInterface {
	repo := repository.NewStockRepository(datasource.GLOBAL_DB, nil) // TODO: 전역변수, 캐시
	return &StockUseCase{Repo: repo}
}
