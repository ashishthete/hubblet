package users

type UserModel struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
}
