package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string `gorm:"type:BINARY(36);primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()

	return
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = uuid.Parse(b.ID)

	return
}
