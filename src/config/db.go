package config

import (
	"log"
	"os"
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User: "linh",
		Password: "",
		Addr: "localhost:5432",
		Database: "todo_db",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
