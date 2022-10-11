package dao

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env"
	"log"
	"time"
)

type Config struct {
	Drive    string        `env:"DATABASE_DRIVER" envDefault:"postgres" `
	Hots     string        `env:"DATABASE_HOTS,required"`
	DbName   string        `env:"DATABASE_NAME,required"`
	Password string        `env:"DATABASE_PASSWORD,required"`
	User     string        `env:"DATABASE_USER,required"`
	Port     string        `env:"DATABASE_PORT"  envDefault:"5432"`
	SSLMode  string        `env:"DATABASE_SSL"  envDefault:"disable"`
	Timeout  time.Duration `env:"DATABASE_TIMEOUT" envDefault:"1s"`
}

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Error on configure Database client. Error : %s", err)
	}

	return NewDatabaseWithSettings(config)
}

func NewDatabaseWithSettings(config Config) *Database {

	conn, err := connection(config)
	if err != nil {
		return nil
	}

	return &Database{db: conn}

}

func connection(config Config) (*sql.DB, error) {

	dataSourceName := getDataSourceName(config)
	db, err := sql.Open(config.Drive, dataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err

}

func getDataSourceName(config Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Hots,
		config.Port,
		config.User,
		config.Password,
		config.DbName,
		config.SSLMode)
}
