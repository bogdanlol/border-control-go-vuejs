package masini

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (masini Masini, err error) {
	err = db.
		Limit(limit, offset).
		Find(&masini)

	return
}
