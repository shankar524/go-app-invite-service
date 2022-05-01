package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shankar524/go-app-invite-service/api/service"
	"github.com/shankar524/go-app-invite-service/models"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	service service.ITokenService
}

type ITokenController interface {
	Create(*gin.Context)
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	DisableTokenByID(*gin.Context)
	ValidateToken(*gin.Context)
}

func NewTokenController(tokenService service.TokenService) TokenController {
	return TokenController{
		service: &tokenService,
	}
}

func (controller *TokenController) Create(c *gin.Context) {
	log.Println("TokenController :: Create")

	token, err := controller.service.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, token)
}

func (controller *TokenController) GetAll(c *gin.Context) {
	log.Println("TokenController :: GetAll")

	tokens, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

func (controller *TokenController) GetByID(c *gin.Context) {
	log.Println("TokenController :: GetByID")

	id := c.Param("id")
	token, err := controller.service.GetByID(id)

	if err == nil {
		c.JSON(http.StatusOK, token)
		return
	}

	if err.Error() == "record not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("token(%s) was not found", id)})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (controller *TokenController) DisableTokenByID(c *gin.Context) {
	log.Println("TokenController :: DisableTokenByID")

	id := c.Param("id")
	token, err := controller.service.DisableTokenByID(id)

	if err == nil {
		c.JSON(http.StatusOK, token)
		return
	}

	if err.Error() == "record not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Token(%s) not found", id)})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (controller *TokenController) ValidateToken(c *gin.Context) {
	log.Println("TokenController :: ValidateToken")

	var token models.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := controller.service.ValidateToken(token.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusOK, gin.H{"message": "valid token"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "token not valid. Please provide valid token"})
	}
}
