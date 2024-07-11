package service

import (
	"errors"

	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/query"
	"golang.org/x/crypto/bcrypt"
)

var q = query.Q

func CreateUser(user *model.User) error {
	if u, err := q.User.Where(q.User.Username.Eq(user.Username)).First(); err == nil && u != nil {
		return errors.New("username already exists")
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPwd)

	return q.User.Create(user)
}
