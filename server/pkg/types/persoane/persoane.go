package persoane

import (
	"learning/proj1/pkg/db"
)

type Persoane []Persoana

type Persoana db.Persoana

func (p *Persoana) TableName() string {
	return "persoane"
}
