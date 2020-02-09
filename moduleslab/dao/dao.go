package dao

import "github.com/iproduct/coursego/moduleslab/model"

type UserRepo interface {
	Find(start, count int) ([]model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(id int) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	DeleteByID(id int) (*model.User, error)
}
