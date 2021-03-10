package db

import (
	"fmt"
	"github.com/Irbi/citrouser/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var Connection *gorm.DB

func Init(host, port, dbname, user, password string) (db *gorm.DB) {

	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)

	isMigrationNeeded, err := strconv.ParseBool(os.Getenv("DB_MIGRATION"))
	if err != nil {
		log.Error("Failed to parse migration flag")
		log.Fatal(err)
	}

	Connection, err = Connect(connString, isMigrationNeeded)

	if err != nil {
		log.Errorf("Failed on connection to DB. Conn: %s", connString)
		log.Fatal(err)
	}

	isDebug, err := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	if isDebug {
		Connection = Connection.Debug()
	}

	return Connection
}

func Connect(connString string, isMigrationNeeded bool) (connection *gorm.DB, err error) {
	connection, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if isMigrationNeeded {
		err = connection.AutoMigrate(
			&model.User{},
			&model.Profile{},
			&model.TempPassword{},
			&model.FinancialAssetCategory{},
			&model.FinancialAsset{},
			&model.FinancialInstrument{},
			&model.ClientAdvisor{},
		)

		if err != nil {
			log.Error("DB migration failed")
			log.Fatal(err)
		}
	}

	return
}
