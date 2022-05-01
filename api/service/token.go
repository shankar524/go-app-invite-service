package service

import (
	"github.com/shankar524/go-app-invite-service/api/repository"
	"github.com/shankar524/go-app-invite-service/models"
	"github.com/shankar524/password-generator/src/text"
)

type TokenService struct {
	tokenRepository repository.ITokenRepository
	textgenerator   text.Generator
}

type ITokenService interface {
	Create() (models.Token, error)
	GetAll() ([]models.Token, error)
	GetByID(string) (models.Token, error)
	DisableTokenByID(id string) (models.Token, error)
	ValidateToken(token string) (bool, error)
	InvalidateToken(days int) error
}

func NewTokenService(repo repository.TokenRepository, textgenerator text.Generator) TokenService {
	return TokenService{
		tokenRepository: &repo,
		textgenerator:   textgenerator,
	}
}

func (ts *TokenService) Create() (models.Token, error) {
	token := models.Token{Value: ts.textgenerator.Generate(), Disabled: false}

	return ts.tokenRepository.Save(token)
}

func (ts *TokenService) GetAll() ([]models.Token, error) {
	return ts.tokenRepository.GetAll()
}

func (ts *TokenService) GetByID(id string) (models.Token, error) {
	return ts.tokenRepository.GetByID(id)
}

func (ts *TokenService) DisableTokenByID(id string) (models.Token, error) {
	return ts.tokenRepository.DisableTokenByID(id)
}

func (ts *TokenService) ValidateToken(token string) (bool, error) {
	return ts.tokenRepository.ValidToken(token)
}

func (ts *TokenService) InvalidateToken(days int) error {
	return ts.tokenRepository.InvalidateToken(days)
}
