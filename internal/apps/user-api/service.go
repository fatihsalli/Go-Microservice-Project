package user_api

import (
	"OrderUserProject/internal/models"
	"OrderUserProject/internal/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	userService := &UserService{Repository: repository}
	return userService
}

type IUserService interface {
	Insert(user models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetBookById(id string) (models.User, error)
	Update(user models.User) (bool, error)
	Delete(id string) (bool, error)
}

func (b UserService) Insert(user models.User) (models.User, error) {

	// to create id and created date value
	user.ID = uuid.New().String()
	user.CreatedDate = primitive.NewDateTimeFromTime(time.Now())

	result, err := b.Repository.Insert(user)

	if err != nil || result == false {
		return user, err
	}

	return user, nil
}

func (b UserService) GetAll() ([]models.User, error) {
	result, err := b.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b UserService) GetBookById(id string) (models.User, error) {

	result, err := b.Repository.GetBookById(id)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (b UserService) Update(user models.User) (bool, error) {
	// to create updated date value
	user.UpdatedDate = primitive.NewDateTimeFromTime(time.Now())

	result, err := b.Repository.Update(user)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func (b UserService) Delete(id string) (bool, error) {
	result, err := b.Repository.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}