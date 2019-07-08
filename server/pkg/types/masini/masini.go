package masini

import (
	"learning/proj1/pkg/db"
)

type Masini []Masina

type Masina db.Masina

func (m *Masina) TableName() string {
	return "Masini"
}
