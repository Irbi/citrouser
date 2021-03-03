package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(connString string) (db *gorm.DB) {
	connection, err := connect(connString)

	if err != nil {
		log.Errorf("Failed on connection to DB. Conn: %s", connString)
		log.Fatal(err)
	}

	return connection
}

func connect(connString string) (connection *gorm.DB, err error) {
	connection, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return
}
