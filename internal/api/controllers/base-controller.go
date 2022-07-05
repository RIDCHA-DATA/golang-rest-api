package controllers

import (
	"net/http"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/models"
	"github.com/gin-gonic/gin"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/config"

	_ "github.com/containers/image/v5/pkg/docker/config"
	_ "github.com/onsi/ginkgo/config"
)

// isActive godoc
// @Summary Get Current deployed version
// @Description Get deployed version of the api from db
// @Tags sample
// @Accept  json
// @Produce  json
// @Router /isActive [get]
func GetVersion(c *gin.Context) {
	info := models.AppInfo{
		Version:  config.Version,
		Deployed: config.BuildTime,
	}
	c.JSON(http.StatusOK, info)
}
