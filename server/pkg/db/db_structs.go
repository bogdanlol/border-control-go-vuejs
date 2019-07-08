package db

import (
	"time"
)

type Masina struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Numar string	`xorm:"numar" json:"numar" schema:"numar"`
	Culoare string	`xorm:"culoare" json:"culoare" schema:"culoare"`
}

func (t Masina) TableName() string {
	 return "masina"
}

func (t Masina) SetId(id int64) {
	t.Id = id
}

func (t Masina) GetId() int64 {
	return t.Id
}

type Migrations struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Name string	`xorm:"name" json:"name" schema:"name"`
	RunOn time.Time	`xorm:"run_on" json:"run_on" schema:"run_on"`
}

func (t Migrations) TableName() string {
	 return "migrations"
}

func (t Migrations) SetId(id int64) {
	t.Id = id
}

func (t Migrations) GetId() int64 {
	return t.Id
}

type Persoana struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	NUME string	`xorm:"NUME" json:"NUME" schema:"NUME"`
	PREN string	`xorm:"PREN" json:"PREN" schema:"PREN"`
	Varsta int64	`xorm:"varsta" json:"varsta" schema:"varsta"`
	CNP string	`xorm:"CNP" json:"CNP" schema:"CNP"`
	Codm int64	`xorm:"codm" json:"codm" schema:"codm"`
}

func (t Persoana) TableName() string {
	 return "persoana"
}

func (t Persoana) SetId(id int64) {
	t.Id = id
}

func (t Persoana) GetId() int64 {
	return t.Id
}

type Personal struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	NUME string	`xorm:"NUME" json:"NUME" schema:"NUME"`
	PREN string	`xorm:"PREN" json:"PREN" schema:"PREN"`
	RANG string	`xorm:"RANG" json:"RANG" schema:"RANG"`
}

func (t Personal) TableName() string {
	 return "personal"
}

func (t Personal) SetId(id int64) {
	t.Id = id
}

func (t Personal) GetId() int64 {
	return t.Id
}

type Users struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	First string	`xorm:"first" json:"first" schema:"first"`
	Last string	`xorm:"last" json:"last" schema:"last"`
	Email string	`xorm:"email" json:"email" schema:"email"`
	Password string	`xorm:"password" json:"password" schema:"password"`
}

func (t Users) TableName() string {
	 return "users"
}

func (t Users) SetId(id int64) {
	t.Id = id
}

func (t Users) GetId() int64 {
	return t.Id
}

type Verificare struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Codp int64	`xorm:"codp" json:"codp" schema:"codp"`
	Codm int64	`xorm:"codm" json:"codm" schema:"codm"`
	Codo int64	`xorm:"codo" json:"codo" schema:"codo"`
	Orav time.Time	`xorm:"orav" json:"orav" schema:"orav"`
	Orap time.Time	`xorm:"orap" json:"orap" schema:"orap"`
}