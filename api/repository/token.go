package repository

import (
	"fmt"
	"log"

	"github.com/shankar524/go-app-invite-service/lib"
	"github.com/shankar524/go-app-invite-service/models"
)

type TokenRepository struct {
	db    lib.Database
	cache lib.ICache
}

type ITokenRepository interface {
	Migrate() error
	Save(token models.Token) (models.Token, error)
	GetAll() (tokens []models.Token, err error)
	GetByID(string) (models.Token, error)
	DisableTokenByID(string) (models.Token, error)
	ValidToken(string) (bool, error)
	InvalidateToken(int) error
}

func NewTokenRepository(db lib.Database, cache lib.ICache) TokenRepository {
	return TokenRepository{db, cache}
}

func (t *TokenRepository) Migrate() error {
	log.Print("TokenRepository :: Migrate")

	return t.db.DB.AutoMigrate(&models.Token{})
}

func (t *TokenRepository) Save(token models.Token) (models.Token, error) {
	log.Print("TokenRepository :: Save")

	err := t.db.DB.Create(&token).Error
	if err != nil {
		log.Printf("error on saving token. Error: %s", err.Error())
	}
	t.cache.Save(token.Value)
	return token, err
}

func (t *TokenRepository) GetAll() (tokens []models.Token, err error) {
	log.Print("TokenRepository :: GetAll")

	result := t.db.DB.Find(&tokens)

	return tokens, result.Error
}

func (t *TokenRepository) GetByID(id string) (token models.Token, err error) {
	log.Print("TokenRepository :: GetByID")

	err = t.db.DB.Where("id = ?", id).First(&token).Error

	return
}

func (t *TokenRepository) DisableTokenByID(id string) (token models.Token, err error) {
	log.Print("TokenRepository :: DisableTokenByID")

	token, err = t.GetByID(id)
	if err != nil {
		return
	}

	err = t.db.DB.Model(&models.Token{}).Where("id = ?", id).Update("disabled", true).Error
	if err != nil {
		return
	}

	err = t.db.DB.Model(&token).Update("disabled", true).Error
	if err == nil {
		t.cache.Delete(token.Value)
	}

	return
}

func (t *TokenRepository) ValidToken(token string) (bool, error) {
	log.Print("TokenRepository :: ValidToken")
	return t.cache.Exists(token)
}

// InvalidateToken updates all token that were created provided days ago
// and sets disabled property as true
func (t *TokenRepository) InvalidateToken(days int) error {
	var tokens []models.Token
	result := t.db.DB.Where("disabled = ? AND created_at::DATE < DATEADD(day, ?, GETDATE())", false, -1*days).Find(&tokens)
	err := result.Error

	if result.Error != nil {
		return err
	}

	for _, token := range tokens {
		_, err = t.DisableTokenByID(fmt.Sprintf("%d", token.ID))
		if err != nil {
			break
		}
	}
	return err
}
