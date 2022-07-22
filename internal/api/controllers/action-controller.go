package controllers

import (
	"net/http"

	pg "github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/db"
	models "github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/models"
	e "github.com/RIDCHA-DATA/golang-rest-api/pkg/e"
	"github.com/RIDCHA-DATA/golang-rest-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

// GetActions
// @Summary Get Available Actions
// @Description Get available action from db
// @Tags actions
// @ID get-action
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Action
// @Failure 500 int e.ERROR
// @Router /actions [get]
func GetActions(c *gin.Context) {
	var actions = []models.Actions{}
	data := pg.SetupDB()
	err := data.Model(&actions).Select()
	if err != nil {
		e.NewError(c, e.ERROR, e.GetMsg(e.ERROR))
		logger.InlineLog(c, err.Error(), e.GetMsg(e.ERROR))
		return
	}
	c.JSON(e.SUCCESS, actions)
}

// PostAction godoc
// @Summary Post new action
// @Description Post new action to DB
// @Tags actions
// @ID post-action
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Action
// @Failure 400 int e.INVALID_PARAMS
// @Router /actions [post]
func PostAction(c *gin.Context) {
	input := &models.Actions{}
	err := c.BindJSON(input)
	data := pg.SetupDB()

	if err != nil {
		e.NewError(c, e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS))
		logger.InlineLog(c, err.Error(), input)
		return
	}

	_, err = data.Model(input).Insert()

	if err != nil {
		e.NewError(c, http.StatusBadRequest, e.GetMsg(e.ERROR))
		logger.InlineLog(c, err.Error(), err)
		return
	}

	c.JSON(http.StatusOK, input)
}
