package db

import (
	"cart-api/config"
	"cart-api/pkg/log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

const dbUrl = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func newDB(cfg *config.DbConf, log *log.Logger) *gorm.DB {
	const op = "pkg.db.newDB"
	connectionString := fmt.Sprintf(dbUrl, cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	return connection
}

func MustStartDB(cfg *config.DbConf, log *log.Logger) *gorm.DB {
	const op = "pkg.db.MustStartDB"
	if db == nil {
		db = newDB(cfg, log)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	if err = sqlDB.Ping(); err != nil {
		log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	log.InfoLog.Printf("%s: Connect to db\n", op)

	return db
}

func MustCloseDB(db *gorm.DB, log *log.Logger) {
	const op = "pkg.db.MustCloseDB"

	sqlDB, err := db.DB()

	if err != nil {
		log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	if err = sqlDB.Close(); err != nil {
		log.ErrorLog.Panicf("%s: can't close connection: %s", op, err.Error())
	}

	log.InfoLog.Printf("%s: Success close connection\n", op)
}
