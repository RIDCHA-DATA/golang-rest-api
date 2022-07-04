package db

import (
	"context"
	"path"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB
var SQLITE *gorm.DB
var err error

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	configuration := config.GetConfig()

	database := configuration.Database.Dbname
	host := configuration.Database.Host
	port := configuration.Database.Port
	username := configuration.Database.Username
	password := configuration.Database.Password
	Dialect := configuration.Database.Dialect

	if Dialect == "sqlite3" {
		SQLITE, err = gorm.Open("sqlite3", path.Join(".", "app.db"))
	} else {
		DB = pg.Connect(&pg.Options{
			Addr:     host + ":" + port,
			User:     username,
			Password: password,
			Database: database,
		})

		// Check if database is connected
		ctx := context.Background()

		if err := DB.Ping(ctx); err != nil {
			panic(err)
		}
	}
}
