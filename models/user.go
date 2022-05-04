package models

import (
	"time"
)

// User ...
type User struct {
	ID        int       `db:"id" json:"uid"`
	UserName  string    `db:"userName" json:"userName"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	PhoneNum  string    `db:"phoneNum" json:"phoneNum"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
