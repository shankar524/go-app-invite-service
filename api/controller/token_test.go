package controller

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shankar524/go-app-invite-service/models"
	"gopkg.in/go-playground/assert.v1"
)

type mockTokenService struct {
	token              models.Token
	createError        error
	modelList          []models.Token
	listError          error
	getByIdError       error
	disableTokenError  error
	validToken         bool
	validateTokenError error
	invalidateError    error
}

func (m mockTokenService) Create() (models.Token, error) {
	return m.token, m.createError
}

func (m mockTokenService) GetAll() ([]models.Token, error) {
	return m.modelList, m.listError
}

func (m mockTokenService) GetByID(string) (models.Token, error) {
	return m.token, m.getByIdError
}

func (m mockTokenService) DisableTokenByID(id string) (models.Token, error) {
	return m.token, m.disableTokenError
}

func (m mockTokenService) ValidateToken(id string) (bool, error) {
	return m.validToken, m.validateTokenError
}

func (m mockTokenService) InvalidateToken(days int) error {
	return m.invalidateError
}

func Text_Create(t *testing.T) {
	t.Log("When there is an error creating token")
	{
		t.Run("Should return error", func(t *testing.T) {
			c := &gin.Context{}
			tokenService := mockTokenService{
				createError: errors.New("error creating token"),
			}
			tc := TokenController{tokenService}
			tc.Create(c)
			response := c.Request.Response
			assert.Equal(t, 400, response.StatusCode)
		})
	}
}
