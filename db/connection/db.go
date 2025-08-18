package db

import (
	"database/sql"
	"fmt"
	"go-micro/config/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql2 "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	grmsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *sql.DB

func ConnectDB() (*gorm.DB, error) {

	cfg := config.Load()

	host := cfg.Database.Host
	port := cfg.Database.Port
	name := cfg.Database.DBName
	user := cfg.Database.User
	pass := cfg.Database.Pass

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		user, pass, host, port, name,
	)

	// dsn := "user:xxxxxx@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	sqlDB, err := sql.Open("mysql", dsn)

	DB = sqlDB

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

func CloseConnection() {
	if DB != nil {
		DB.Close()
	}
}
