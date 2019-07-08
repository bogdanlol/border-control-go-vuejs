package jwt

import (
	Users "learning/proj1/pkg/types/users"
	ORM "learning/proj1/src/system/db"

	"errors"

	"github.com/go-xorm/xorm"
)

func GetUserFromToken(db *xorm.Engine, tokenVal string) (user Users.User, err error) {
	if tokenVal == "" {
		err = errors.New("No token present.")
		return
	}

	userId, err := IsTokenValid(tokenVal)
	if err != nil {
		err = errors.New("Token is invalid.")
		return
	}

	if userId < 1 {
		err = errors.New("Token missing required data.")
		return
	}

	user = Users.User{Id: userId}
	err = ORM.FindBy(db, &user)
	if err != nil || user.Id < 1 {
		err = errors.New("User in token does not exist in system.")
		return
	}

	return
}
