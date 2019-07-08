package main

import (
	"flag"
	DB "learning/proj1/src/system/db"
	"os"

	"learning/proj1/src/system/app"

	"github.com/joho/godotenv"
)

var port string
var dbhost string
var dbport string
var dbuser string
var dbpass string
var dboptions string
var dbdatabase string

func init() {

	flag.StringVar(&port, "port", "8081", "Assigning the port that the server should listen on.")
	flag.StringVar(&dbhost, "dbhost", "127.0.0.1", "3000")
	flag.StringVar(&dbport, "dbport", "3306", "3000")
	flag.StringVar(&dbuser, "dbuser", "root", "3000")
	flag.StringVar(&dbpass, "dbpass", "", "3000")
	flag.StringVar(&dboptions, "dboptions", "parseTime=true", "3000")
	flag.StringVar(&dbdatabase, "dbdatabase", "test", "3000")

	flag.Parse()
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect(dbhost, dbport, dbuser, dbpass, dbdatabase, dboptions)
	if err != nil {
		panic(err)
	}
	s := app.NewServer()
	s.Init(port, db)
	s.Start()
}
