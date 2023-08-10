package usecase

import (
	"myservice/datasource"
	"myservice/model"
	"myservice/repository"

	"golang.org/x/crypto/bcrypt"
)

type AccountUseCaseInterface interface {
	Create(model.AccountInput) (model.AccountOutput, error)
	FindById(string) (model.AccountOutput, error)
	Delete(string) error
	Login(model.LoginInput) error
}

type AccountUseCase struct {
	Repo repository.AccountRepository
}

func NewAccountUseCase() AccountUseCaseInterface {
	repo := repository.NewAccountRepository(datasource.GLOBAL_DB, nil) // TODO: 전역변수, 캐시
	return &AccountUseCase{Repo: repo}
}

func (uc *AccountUseCase) Create(accountInput model.AccountInput) (model.AccountOutput, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(accountInput.PW), bcrypt.DefaultCost)
	if err != nil {
		return model.AccountOutput{}, err
	}
	accountInput.PW = string(pwHash)

	newAccountEntity, err := uc.Repo.Create(accountInput)
	if err != nil {
		return model.AccountOutput{}, err
	}

	newAccountOutput := model.AccountOutput{}
	newAccountOutput.FromEntity(newAccountEntity)
	return newAccountOutput, nil
}

func (uc *AccountUseCase) FindById(id string) (model.AccountOutput, error) {
	accountEntity, err := uc.Repo.FindById(id)
	if err != nil {
		return model.AccountOutput{}, err
	}

	accountOutput := model.AccountOutput{}
	accountOutput.FromEntity(accountEntity)
	return accountOutput, nil
}

func (uc *AccountUseCase) Delete(id string) error {
	return uc.Repo.Delete(id)
}

func (uc *AccountUseCase) Login(loginInput model.LoginInput) error {
	accountEntity, err := uc.Repo.FindById(loginInput.ID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(accountEntity.PW), []byte(loginInput.PW))
	return err
}
