package verificari

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (verificari Verificari, err error) {
	err = db.
		Limit(limit, offset).
		Find(&verificari)

	return
}
