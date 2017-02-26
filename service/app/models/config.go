package models

import (
	"fmt"

	"github.com/maxwellhealth/bongo"
)

var db *Database

type Database struct {
	config     *bongo.Config
	connection *bongo.Connection
}

func init() {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "destiny",
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		panic(fmt.Sprintf("bongo error: %s", err))
	}

	db = &Database{config: config, connection: connection}
}
