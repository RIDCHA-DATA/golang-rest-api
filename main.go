package main

import (
	_ "github.com/RIDCHA-DATA/golang-rest-api/docs"
	"github.com/RIDCHA-DATA/golang-rest-api/internal/api"
)

// @title Lab Routing Service with golang
// @Golang API REST
// @version 1.0
// @description IAM API REST in Golang with Gin Framework

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /v1
// @query.collection.format multi

func main() {
	api.Run("")
}
