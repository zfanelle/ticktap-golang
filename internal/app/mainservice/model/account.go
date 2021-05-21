package model

type AccountDTO struct {
	EntityName string `json:"entity_name"`
	EntityType string `json:"entity_type"`
}

func (accountDTO AccountDTO) DTOtoAccount() Account {
	return Account{EntityName: accountDTO.EntityName, EntityType: accountDTO.EntityType}
}

type Account struct {
	Id         int    `json:"id"`
	EntityName string `json:"entity_name"`
	EntityType string `json:"entity_type"`
}

type AccountDB struct {
	Id         int    `db:"id"`
	EntityName string `db:"entity_name"`
	EntityType string `db:"entity_type"`
}

func (accountDB AccountDB) DBtoAccount() Account {
	return Account{Id: accountDB.Id, EntityName: accountDB.EntityName, EntityType: accountDB.EntityType}
}
