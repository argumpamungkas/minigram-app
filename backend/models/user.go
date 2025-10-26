package models

import (
	"minigram-app-backend/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string  `gorm:"column:username" json:"username" valid:"required~Username is Required"`
	FullName *string `gorm:"column:full_name" json:"full_name"`
	Email    string  `gorm:"column:email" json:"email" valid:"required~Email is Required, email~Invalid format Email"`
	Password string  `gorm:"column:password" json:"password" valid:"required~Password is Required, minstringlength(8)~Password minimum 8 characters"`
	Avatar   *string `gorm:"column:avatar" json:"avatar"`
	Bio      *string `gorm:"column:bio" json:"bio"`
	ApiKey   string  `gorm:"column:api_key" json:"api_key" `
	// Posting  []Posting `gorm:"constrait:OnUpdate:CASCADE" json:"posts"`
}

// Model sebelum create
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// validasi struct
	if ok, err := govalidator.ValidateStruct(u); !ok {
		return err
	}

	// set created date
	currentTime := time.Now()
	u.CreatedDate = &currentTime

	// hash password
	u.Password = helpers.HashPassword(u.Password)

	return nil
}

type UserLogin struct {
	Username string `gorm:"column:username" json:"username" valid:"required~Username is Required"`
	Password string `gorm:"column:password" json:"password" valid:"required~Password is Required"`
}

type ResponseUser struct {
	Username string  `json:"username"`
	FullName *string `json:"full_name"`
	Email    string  `json:"email"`
	Avatar   *string `json:"avatar"`
	Bio      *string `json:"bio"`
}
