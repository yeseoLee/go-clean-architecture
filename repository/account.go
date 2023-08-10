package repository

import (
	"database/sql"
	"fmt"
	"myservice/entity"
	"myservice/model"
	"myservice/repository/cache"
)

type AccountRepository interface {
	Create(model.AccountInput) (*entity.Account, error)
	FindById(string) (*entity.Account, error)
	Delete(string) error
}

type accountRepository struct {
	db    *sql.DB
	cache cache.AccountCacheRepository
}

func NewAccountRepository(dbEngine *sql.DB, cache cache.AccountCacheRepository) AccountRepository {
	return &accountRepository{
		db:    dbEngine,
		cache: cache,
	}
}

func (repo *accountRepository) Create(accountInput model.AccountInput) (*entity.Account, error) {
	id := accountInput.ID
	pw := accountInput.PW
	name := accountInput.Name
	result, err := repo.db.Exec("INSERT INTO `account` (`id`, `pw`, `name`) VALUES(?, ?, ?)", id, pw, name)
	if err != nil {
		return nil, err
	}

	seq, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	account := &entity.Account{
		Seq:  seq,
		ID:   id,
		PW:   pw,
		Name: name,
	}

	return account, nil
}

func (repo *accountRepository) FindById(id string) (*entity.Account, error) {
	account := &entity.Account{}

	err := repo.db.QueryRow("SELECT * FROM `account` WHERE id = ?", id).
		Scan(&account.Seq, &account.ID, &account.PW, &account.Name)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repo *accountRepository) Delete(id string) error {
	result, err := repo.db.Exec("DELETE FROM `account` WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}
