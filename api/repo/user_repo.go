package repo

import (
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/internal/database"
)

type UserRepo struct {
	db database.Database
}

func NewUserRepo(db database.Database) UserRepo {
	return UserRepo{db}
}

// GetByUsername returns user by username
func (r UserRepo) GetByUsername(username string) (user *model.User, err error) {
	if err = r.db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create creates a new user
func (r UserRepo) Create(user *model.User) (err error) {
	err = r.db.DB.Create(user).Error
	return
}

func (r UserRepo) Get(id uint) (user *model.User, err error) {
	if err = r.db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
