package models

import (
	"github.com/gofiber/fiber/v3/log"
	"github.com/nglab-dev/nglab/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (u *User) IsAdmin() bool {
	return u.Username == "admin"
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Get().Where("username = ?", username).First(&user).Error
	return user, err
}

func CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return db.Get().Create(user).Error
}

func InitAdminUser() {

	u, _ := GetUserByUsername("admin")
	if u.ID != 0 {
		return
	}

	user := User{
		Username: "admin",
		Password: "admin",
		Email:    "admin@localhost",
	}
	if err := CreateUser(&user); err != nil {
		log.Warnf("Failed to create admin user: %v", err)
	}

	log.Info("Admin user created")
}
