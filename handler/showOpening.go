package handler

import (
	"net/http"

	"github.com/GiovannaK/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id").Error())
		return
	}

	opening := schemas.Opening{}

	// find opening by id

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error finding opening: %v", err)
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, "ShowOpeningHandler", opening)
}
