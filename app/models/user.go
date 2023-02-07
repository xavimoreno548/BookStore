package models

import (
	"fmt"

	"github.com/xavimoreno548/BookStore/internal/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Purchases []Purchase
	Favorite []Favorite
}

func (user *User) Save(db *gorm.DB) (*User, error){
	err := db.Create(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) BeforeSave() error {
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func LoginCheck(email string, password string, db *gorm.DB) (string, error){
	var err error
	user := User{}

	err = db.Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword{
		return "", err
	}

	token, err := services.GenerateToken(user.ID)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (user *User) PrepareGive(){
	user.Password = ""
}

func VerifyPassword(password string, hashedPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserByID(uid uint, db *gorm.DB) (User, error){
	var user User

	if err := db.First(&user, uid).Error; err != nil {
		return user, fmt.Errorf("user not found: %v", err)
	}
	user.PrepareGive()
	return user, nil
}

