package models

import (
	"fmt"
	"os"
	"vftalk/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	MySQLTables = []string{ // Define SQL tables here
		`Users`,
	}
	migrationsDir = `file://database/migration`
	logger        = configs.InitLogger()
)

func init() {
	configs.LoadEnv()
}

func RunMigrationUp() {
	db, err := configs.ConnectMariaDB()
	if err != nil {
		logger.Fatal().Str("error", err.Error()).Msg("failed when try to connect database")
	}
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, os.Getenv(`MARIADB_NAME`), driver)
	if err != nil {
		logger.Fatal().Str("error", err.Error()).Msg("error when create db instance")
	}
	// Run the migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Fatal().Str("error", err.Error()).Msg("error when migrate database")
	}

	// Create or modify table schema, so then SQLC could generate code
	for _, table := range MySQLTables {
		query := fmt.Sprintf("SHOW CREATE TABLE %s", table)
		rows, err := db.Query(query)
		if err != nil {
			logger.Fatal().Str("error", err.Error()).Msg("error when query database")
		}
		defer rows.Close()
		var tableNameResult, createTableStatement string
		for rows.Next() {
			err := rows.Scan(&tableNameResult, &createTableStatement)
			if err != nil {
				logger.Fatal().Str("error", err.Error()).Msg("cannot scan table")
			}
		}
		sqlSchemaFile := fmt.Sprintf("database/schema/%s.sql", table)
		file, err := os.Create(sqlSchemaFile)
		if err != nil {
			logger.Fatal().Str("error", err.Error()).Msg("cannot create file")
		}
		defer file.Close()
		_, err = file.WriteString(createTableStatement)
		if err != nil {
			logger.Fatal().Str("error", err.Error()).Msg("cannot write file")
		}
	}

	fmt.Println("Migration up successful !!")
}

func RunMigrationDown() {
	db, err := configs.ConnectMariaDB()
	if err != nil {
		logger.Fatal().Str("error", err.Error()).Msg("cannot connect to database")
		os.Exit(1)
	}
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, os.Getenv(`MARIADB_NAME`), driver)
	if err != nil {
		logger.Fatal().Str("error", err.Error()).Msg("error when create db instance")
	}
	// Run the migration
	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		logger.Fatal().Str("error", err.Error()).Msg("error when migrate database")
	}

	// Create or modify table schema, so then SQLC could generate code
	for _, table := range MySQLTables {
		query := fmt.Sprintf("SHOW CREATE TABLE %s", table)
		rows, err := db.Query(query)
		if err != nil {
			logger.Fatal().Str("error", err.Error()).Msg("error when query database")
		}
		defer rows.Close()
		var tableNameResult, createTableStatement string
		for rows.Next() {
			err := rows.Scan(&tableNameResult, &createTableStatement)
			if err != nil {
				logger.Fatal().Str("error", err.Error()).Msg("cannot scan table")
			}
		}
		sqlSchemaFile := fmt.Sprintf("database/schema/%s.sql", table)
		file, err := os.Create(sqlSchemaFile)
		if err != nil {
			logger.Fatal().Str("error", err.Error()).Msg("cannot create file")
		}
		defer file.Close()
		_, err = file.WriteString(createTableStatement)
		if err != nil {
			logger.Fatal().Str("error", err.Error()).Msg("cannot write file")
		}
	}

	fmt.Println("Migration down successful !!")
}
