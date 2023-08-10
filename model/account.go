package model

import "myservice/entity"

type AccountInput struct {
	ID   string
	PW   string
	Name string
}

type AccountOutput struct {
	ID   string
	Name string
}

func (output *AccountOutput) FromEntity(accountEntity *entity.Account) {
	output.ID = accountEntity.ID
	output.Name = accountEntity.Name
}

type LoginInput struct {
	ID string
	PW string
}
