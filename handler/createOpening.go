package handler

import (
	"net/http"

	"github.com/GiovannaK/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	//request := struct { // anonymous struct
	//	Role string `json:"role"`
	//	}{} // empty struct literal

	request := CreateOpeningHandlerRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Infof("Request received: %+v", request)

	sendSuccess(ctx, "CreateOpeningHandler", opening)
}
