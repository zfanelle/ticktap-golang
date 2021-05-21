package repsitory

import (
	"fmt"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	account "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
)

func CreateAccount(appConfig config.AppConfig, newAccount account.Account) {

	db := appConfig.DB

	sql := fmt.Sprintf("INSERT INTO ticktap.account (entity_name,entity_type) VALUES (\"%s\", \"%s\")", newAccount.EntityName, newAccount.EntityType)

	db.MustExec(sql)

}

func GetAccount(appConfig config.AppConfig, accountID int) account.Account {
	db := appConfig.DB

	accounts := []account.AccountDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.account WHERE ID = %d", accountID)

	db.Select(&accounts, sql)

	// Add in error checking

	account := accounts[0].DBtoAccount()

	return account
}

func GetAllAccounts(appConfig config.AppConfig) []account.Account {
	db := appConfig.DB

	dbAccounts := []account.AccountDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.account")

	db.Select(&dbAccounts, sql)

	accounts := []account.Account{}

	// Add in error checking

	for _, account := range dbAccounts {
		accounts = append(accounts, account.DBtoAccount())
	}

	return accounts
}
