package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Value     string         `json:"token" gorm:"uniqueIndex;type:string;size:12"`
	Disabled  bool           `json:"disabled" gorm:"default:false"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName -> Overriding default table name
func (Token) TableName() string {
	return "tokens"
}
