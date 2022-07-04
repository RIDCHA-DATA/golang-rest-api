package controllers

import (
	"net/http"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/models"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/config"

	"github.com/containers/image/v5/pkg/docker/config"
	"github.com/onsi/ginkgo/config"
)

func GetVersion(c *gin.context) {
	info := models.AppInfo{
		Version:  config.Version,
		Deployed: config.BuildTime,
	}
	c.json(http.StatusOK, info)
}
