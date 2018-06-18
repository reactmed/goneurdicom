package service

import (
	"github.com/reactmed/goneurdicom/domain"
	"github.com/reactmed/goneurdicom/dao"
	"github.com/reactmed/goneurdicom/utils"
	"reflect"
)

var service UserService

func NewUserService(ctx utils.AppContext) UserService {
	if service == nil {
		userDao, err := ctx.Get(reflect.TypeOf((*dao.UserDao)(nil)).Elem())
		if err != nil {
			panic(err)
		}
		service = &userService{
			Dao: userDao.(dao.UserDao),
		}
	}
	return service
}

type UserService interface {
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int) error
	FindOne(id int) (*domain.User, error)
	FindAll() (domain.Users, error)
}

type userService struct {
	Dao dao.UserDao
}

func (service *userService) Save(user domain.User) (domain.User, error) {
	return service.Dao.Save(user)
}

func (service *userService) Update(user domain.User) (domain.User, error) {
	return service.Dao.Update(user)
}

func (service *userService) Delete(id int) error {
	return service.Dao.Delete(id)
}

func (service *userService) FindOne(id int) (*domain.User, error) {
	return service.Dao.FindOne(id)
}

func (service *userService) FindAll() (domain.Users, error) {
	return service.Dao.FindAll()
}
