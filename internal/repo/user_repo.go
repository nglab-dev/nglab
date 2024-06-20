package repo

import "github.com/nglab-dev/nglab/internal/database"

type UserRepo struct {
	db database.Database
}

func NewUserRepo(db database.Database) UserRepo {
	return UserRepo{db: db}
}
