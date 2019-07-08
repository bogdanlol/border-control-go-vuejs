package verificari

import (
	"learning/proj1/pkg/db"
)

type Verificari []Verificare

type Verificare db.Verificare

func (v *Verificare) TableName() string {
	return "verificari"
}
