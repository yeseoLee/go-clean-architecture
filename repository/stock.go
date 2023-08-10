package repository

import (
	"database/sql"
	"myservice/repository/cache"
)

type StockRepository interface {
}

type stockRepository struct {
	db    *sql.DB
	cache cache.StockCacheRepository
}

func NewStockRepository(dbEngine *sql.DB, cache cache.StockCacheRepository) StockRepository {
	return &stockRepository{
		db:    dbEngine,
		cache: cache,
	}
}
