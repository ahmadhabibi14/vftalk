package configs

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

func ConnectMariaDB(l *zerolog.Logger) *sql.DB {
	DbDriver := "mysql"
	DbHost := os.Getenv("MARIADB_HOST")
	DbPort := os.Getenv("MARIADB_PORT")
	DbName := os.Getenv("MARIADB_NAME")
	DbUser := os.Getenv("MARIADB_USER")
	DbPassword := os.Getenv("MARIADB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open(DbDriver, dsn)
	if err != nil {
		l.Fatal().
			Str("ERROR", err.Error()).
			Msg("cannot connect to " + DbDriver)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
