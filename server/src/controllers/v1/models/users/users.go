package users

import (
	"learning/proj1/pkg/db"
)

type Users []User

type User db.Users

func (u *User) TableName() string {
	return "users"
}
