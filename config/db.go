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
		User:     "Yogesh",
		Password: "yogi21",
		Addr:     "localhost:5050",
		Database: "karnatakaPolls",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")
	controllers.InitiateDB(db)
	controllers.CreateVoterTable(db)
	controllers.CreateCandidateTable(db)
	controllers.CreateConstituencyTable(db)
	return db
}
