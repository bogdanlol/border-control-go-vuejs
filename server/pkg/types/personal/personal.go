package personal

import (
	"learning/proj1/pkg/db"
)

type Pers []Personal

type Personal db.Personal

func (p *Personal) TableName() string {
	return "Personal"
}
