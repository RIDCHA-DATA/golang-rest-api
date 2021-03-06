package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/config"
	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/models"
	"github.com/go-pg/pg/v10"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config.Setup("config.dev.yml")

	database_config := config.Config.Database

	db, err := sql.Open("postgres",
		"postgres://"+database_config.Username+":"+database_config.Password+"@"+database_config.Host+":5432/"+database_config.Dbname+"?sslmode=disable")
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:./seeds/migrations",
		"postgres", driver)

	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	m.Steps(2)

	fmt.Println("Migration completed")
	fmt.Println("Seeding Database")

	pgdb := pg.Connect(&pg.Options{
		User:     database_config.Username,
		Password: database_config.Password,
		Database: database_config.Dbname,
	})
	defer pgdb.Close()

	// Check if DB is connected
	ctx := context.Background()
	if err := pgdb.Ping(ctx); err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	models := []interface{}{
		(*models.Actions)(&models.Actions{
			Action: "PostRegistration",
		}),
	}
	for _, model := range models {
		r, e := pgdb.Model(model).Count()
		fmt.Println(r)
		fmt.Println(e)
		_, err := pgdb.Model(model).SelectOrInsert()
		if err != nil {
			fmt.Print(err.Error())
			panic(err)
		}
	}

	fmt.Print("Database Seeded")
}
