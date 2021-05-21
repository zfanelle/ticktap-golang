package service

import (
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	account "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	repository "github.com/zfanelle/ticktap-golang/internal/app/mainservice/repository"
)

func CreateAccount(appConfig config.AppConfig, newAccount account.Account) {

	repository.CreateAccount(appConfig, newAccount)

}

func GetAccount(appConfig config.AppConfig, accountID int) account.Account {

	account := repository.GetAccount(appConfig, accountID)

	return account

}

func GetAllAccounts(appConfig config.AppConfig) []account.Account {

	accounts := repository.GetAllAccounts(appConfig)

	return accounts
}
