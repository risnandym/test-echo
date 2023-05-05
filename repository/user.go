package repository

import (
	"test-echo/model"

	"gorm.io/gorm"
)

type IUserRespository interface {
	Store(user model.User) (model.User, error)
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(id int) error
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
func (r UserRepository) Store(user model.User) (model.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (r UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r UserRepository) FindById(id int) (model.User, error) {
	var user model.User
	if err := r.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (repo UserRepository) Update(user model.User) (model.User, error) {
	if err := repo.db.Debug().Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (repo UserRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
