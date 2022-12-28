package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
} //@name User

type LoginForm struct {
	Username string
	Password string
} //@name LoginForm
