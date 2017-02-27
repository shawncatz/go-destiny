package models

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/maxwellhealth/bongo"
)

var db *Database

type Database struct {
	config     *bongo.Config
	connection *bongo.Connection
}

func init() {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		panic("DATABASE_HOST not set")
	}

	dbname := os.Getenv("DATABASE_NAME")
	if dbname == "" {
		panic("DATABASE_NAME not set")
	}

	config := &bongo.Config{
		ConnectionString: host,
		Database:         dbname,
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", host, dbname, err))
	}

	db = &Database{config: config, connection: connection}
}
