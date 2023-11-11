package model

import (
	"strings"

	"github.com/yaswanthsaivendra/prod_mang/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"size:255;not null;uniqueIndex" json:"username"`
	Password  string  `gorm:"size:255;not null;" json:"-"`
	Mobile    string  `gorm:"unique;not null" json:"mobile"`
	Latitude  float64 `gorm:"not null;check:latitude >= -90 AND latitude <= 90" json:"latitude"`
	Longitude float64 `gorm:"not null;check:longitude >= -180 AND longitude <= 180" json:"longitude"`
	Products  []Product
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, nil
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = strings.TrimSpace(user.Username)
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Products").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
