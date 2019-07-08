package persoane

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (persoane Persoane, err error) {
	err = db.
		Limit(limit, offset).
		Find(&persoane)

	return
}
