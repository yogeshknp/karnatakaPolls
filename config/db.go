package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	controllers "github.com/karnatakaPolls/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "yogeshyadav",
		Password: "Yogi@21",
		Addr:     "localhost:5432",
		Database: "yogeshyadav",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateVoterTable(db)
	controllers.InitiateDB(db)
	return db
}
