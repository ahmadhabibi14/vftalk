package models

import (
	"fmt"
	"vftalk/conf"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	driverName    = `mysql`
	migrationsDir = `./models/database/schema`
)

func RunMigration() {
	zlog := conf.InitLogger()
	db := conf.ConnectMariaDB()
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	n, err := migrate.Exec(db, driverName, migrations, migrate.Up)
	if err != nil {
		zlog.Fatal().Msg(`Error: ` + err.Error())
	}

	fmt.Printf("Applied %d migrations!\n", n)
	fmt.Println("Migrate")
}
