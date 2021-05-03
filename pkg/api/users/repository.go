package users

import (
	"database/sql"
	"fmt"
	"huddlet/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repository struct{}

func GetRepository() Repository {
	return Repository{}
}

func (repo Repository) insertUser(tx *sqlx.Tx, user *UserModel) (sql.Result, error) {
	user.Password = utils.HashAndSalt(user.Password)
	log.Println("password", user.Password)
	insertStmt := `
	INSERT INTO huddlet.user_account(name, email, username, password)
	VALUES(:name, :email, :username, :password)
	`
	return tx.NamedExec(insertStmt, user)
}

func (repo Repository) GetUserByField(db *sqlx.DB, field, value string) (*UserModel, error) {
	query := `
	SELECT id, name, email, username, password
	FROM huddlet.user_account
	WHERE %s=$1
	`
	query = fmt.Sprintf(query, field)
	var user UserModel
	err := db.Get(&user, query, value)
	return &user, err
}
