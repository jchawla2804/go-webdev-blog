package domain

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBHost     string
	DBPort     int
	DBUsername string
	DBPassword string
	DBName     string
}

func SetUpDBConnection(config DBConfig) (*gorm.DB, error) {
	dbString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	log.Print("DB Connected")
	gormDb, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := gormDb.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	return gormDb, nil

}
