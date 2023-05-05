package service

import (
	"test-echo/model"
	"test-echo/repository"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(user model.CreateUserRequest) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(user model.UpdateUserRequest, id int) (model.User, error)
	DeleteUser(id int) error
}
type UserService struct {
	userRepository repository.IUserRespository
}

func NewUserService(userRepository repository.IUserRespository) *UserService {
	return &UserService{userRepository}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (service UserService) CreateUser(user model.CreateUserRequest) (model.UserResponse, error) {
	hash, err := HashPassword(user.Password)

	u := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
		Age:      user.Age,
	}

	users, err := service.userRepository.Store(u)
	if err != nil {
		return model.UserResponse{}, err
	}
	userRes := model.UserResponse{
		ID:    users.ID,
		Name:  users.Name,
		Email: users.Email,
		Age:   users.Age,
	}
	return userRes, nil
}
func (service UserService) GetAllUser() ([]model.UserResponse, error) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userRes := []model.UserResponse{}

	//maping with copier
	copier.Copy(&userRes, &users)
	return userRes, nil
}
func (service UserService) GetUserById(id int) (model.UserResponse, error) {
	user, err := service.userRepository.FindById(id)
	if err != nil {
		return model.UserResponse{}, err
	}
	userRes := model.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil
}

func (service UserService) UpdateUser(userRequest model.UpdateUserRequest, id int) (model.UserResponse, error) {
	user, err := service.userRepository.FindById(id)
	if err != nil {
		return model.UserResponse{}, err
	}
	hash, err := HashPassword(userRequest.Password)
	userRequest.Password = hash
	// user.Name = userRequest.Name
	// user.Email = userRequest.Email
	// user.Password = hash
	// user.Age = userRequest.Age

	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true})
	user, err = service.userRepository.Update(user)
	userRes := model.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil
}
func (service UserService) DeleteUser(id int) error {
	_, err := service.userRepository.FindById(id)
	if err != nil {
		return err
	}
	err = service.userRepository.Delete(id)
	return err
}
