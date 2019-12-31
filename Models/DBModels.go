package Models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Feeling struct {
	ID uint `gorm:"primary_key"`
	UserID uint
	Feeling string
	CreatedAt time.Time
}

type User struct {
	gorm.Model
	Name string
	Hash string
	Feelings []Feeling
}