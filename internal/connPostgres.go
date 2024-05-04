package internal

import (
	"fmt"
	"github.com/clickerTelegram/type/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectPostgres(config db.ConfigEnv) *gorm.DB {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		Default.PostgresUrl, Default.PostgresPort, Default.PostgresUsername, Default.PostgresPassword, Default.PostgresDb, Default.SslMode, Default.TimeZone)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("error 1")
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		panic("error 2")
	}
	sqlDb.SetMaxIdleConns(Default.SetMaxIdleConns)
	sqlDb.SetMaxOpenConns(Default.SetMaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(Default.SetConnMaxLifetime) * time.Hour)

	return dbClient
}
