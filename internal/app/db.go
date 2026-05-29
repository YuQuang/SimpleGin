package app

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/royxu/simplegin/v2/configs"
)

func InitDB(config *configs.Configuration) *sql.DB {
	cfg := pq.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		Database: config.DBName,
		SSLMode:  "disable",
	}

	// Or: create a new Config from the defaults, environment, and DSN.
	// cfg, err := pq.NewConfig("host=postgres dbname=pqgo")
	// if err != nil {
	//     log.Fatal(err)
	// }

	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Create connection pool.
	db := sql.OpenDB(c)
	defer db.Close()

	// Make sure it works.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
