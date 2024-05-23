package repository

import "go-crud/internal/domain/entity"

type User interface {
	Insert(u entity.User) (id string, err error)
	Get(id string) (u entity.User, err error)
	GetAll() (sc []entity.User, err error)
	Update(id string, u entity.User) (*entity.User, error)
	Delete(id string) (int64, error)
}
