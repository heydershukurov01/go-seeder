package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Surname  string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string
}
type UsersDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UsersDSO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
