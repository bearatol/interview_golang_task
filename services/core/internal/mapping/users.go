package mapping

import (
	"time"
)

type User struct {
	ID        uint64    `db:"id" json:"id"`
	Login     string    `db:"login" json:"login"`
	Password  string    `db:"password" json:"password"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserAvailableFileds struct {
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
}

type UserToken struct {
	Token string `json:"token"`
}
