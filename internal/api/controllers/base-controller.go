package controllers

import (
	"net/http"

	models "github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/models"
	"github.com/gin-gonic/gin"

	variable "github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/config"

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
		Version:  variable.Version,
		Deployed: variable.BuildTime,
	}
	c.JSON(http.StatusOK, info)
}
