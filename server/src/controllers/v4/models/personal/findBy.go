package personal

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (personal Pers, err error) {
	err = db.
		Limit(limit, offset).
		Find(&personal)

	return
}
