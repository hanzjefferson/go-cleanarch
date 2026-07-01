package entity

type User struct {
	ID       uint   `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}
