package api

import (
	"fmt"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/db"
	"github.com/jinzhu/gorm"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/config"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/api/router"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func setConfiguration(configPath string) {
	config.Setup(configPath)
	DB := db.SetupDB()
	fmt.Println("==================> db migration")
	db.Migrate(DB)
	fmt.Println("==================> insert fake data")
	//seeds.Seed()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "config.dev.yml"
	}

	setConfiguration(configPath)
	conf := config.GetConfig()
	println("i am here")
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
