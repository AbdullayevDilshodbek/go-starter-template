package models


type User struct {
	Id int `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	CreatedAt string `db:"created_at"`
}