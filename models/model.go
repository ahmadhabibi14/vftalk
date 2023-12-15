package models

import (
	"fmt"
	"os"
	"vftalk/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	MySQLTables = []string{ // Define SQL tables here
		`Users`,
	}
	migrationsDir = `file://models/database/migration`
)

func RunMigration() {
	zlog := conf.InitLogger()
	db := conf.ConnectMariaDB()
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, os.Getenv(`MARIADB_NAME`), driver)
	if err != nil {
		zlog.Fatal().Msg(`Error: ` + err.Error())
	}
	// Run the migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		zlog.Fatal().Msg(`Error: ` + err.Error())
	}

	// Create or modify table schema, so then SQLC could generate code
	for _, table := range MySQLTables {
		query := fmt.Sprintf("SHOW CREATE TABLE %s", table)
		rows, err := db.Query(query)
		if err != nil {
			zlog.Fatal().Msg(`Error: ` + err.Error())
		}
		defer rows.Close()
		var tableNameResult, createTableStatement string
		for rows.Next() {
			err := rows.Scan(&tableNameResult, &createTableStatement)
			if err != nil {
				zlog.Fatal().Msg(`Error: ` + err.Error())
			}
		}
		sqlSchemaFile := fmt.Sprintf("models/database/schema/%s.sql", table)
		file, err := os.Create(sqlSchemaFile)
		if err != nil {
			zlog.Fatal().Msg(`Error: ` + err.Error())
		}
		defer file.Close()
		_, err = file.WriteString(createTableStatement)
		if err != nil {
			zlog.Fatal().Msg(`Error: ` + err.Error())
		}
	}

	fmt.Println("Migration successful !!")
}
