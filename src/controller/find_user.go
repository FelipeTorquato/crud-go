package controller

import (
	"crud/src/configuration/logger"
	"crud/src/configuration/rest_err"
	"crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate id",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findById services",
			err,
			zap.String("journey", "findUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "FindUserByEmail"),
	)

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findByEmail services",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
