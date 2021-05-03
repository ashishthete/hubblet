package users

import (
	"huddlet/app/db"
	"log"
)

type Services struct{}

func GetServices() Services {
	return Services{}
}

func (service Services) AddUser(user *UserModel) error {
	tx, _ := db.GetPostgresDB().Beginx()
	defer tx.Rollback()
	_, err := GetRepository().insertUser(tx, user)
	if err != nil {
		log.Println("Error inserting data", err)
	}

	err = tx.Commit()
	return err
}

func (service Services) GetUserByField(field string, value string) (*UserModel, error) {
	db := db.GetPostgresDB()
	user, err := GetRepository().GetUserByField(db, field, value)
	if err != nil {
		log.Println("GetUserByField Error ", err)
	}
	return user, err
}
