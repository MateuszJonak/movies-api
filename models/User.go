package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `form:"email" json:"email" binding:"required" gorm:"unique;not null"`
	Password  string `form:"password" json:"password" binding:"required" gorm:"not null"`
	FirstName string `form:"firstName" json:"firstName"`
	LastName  string `form:"lastName" json:"lastName"`
}

func (User) TableName() string {
	return "user"
}
