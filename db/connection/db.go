package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql2 "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	grmsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "user:xxxxxx@tcp(127.0.0.1:3306)/golang?multiStatements=true"
	sqlDB, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	gormDB, err := gorm.Open(grmsql.New(grmsql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	fmt.Println("DB Connected!")

	if err != nil {
		return nil, err
	}

	driver, err := mysql2.WithInstance(sqlDB, &mysql2.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return nil, err
	}

	m.Up()

	// migrate ->
	// migrate -path db/migrations -database "mysql://user:xxxxxx@tcp(localhost:3306)/golang" up

	return gormDB, err
}
