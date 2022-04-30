package repository

import (
	"log"

	"github.com/shankar524/go-app-invite-service/lib"
	"github.com/shankar524/go-app-invite-service/models"
)

type TokenRepository struct {
	db    lib.Database
	cache lib.Cache
}

func NewTokenRepository(db lib.Database, cache lib.Cache) TokenRepository {
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
