package service

import (
	"github.com/reactmed/goneurdicom/domain"
	"github.com/reactmed/goneurdicom/dao"
	"github.com/reactmed/goneurdicom/utils"
	"reflect"
)

var service UserService

func NewUserService(ctx utils.AppContext) UserService {
	if service == nil{
		service = &userService{
			Dao: ctx.Get(reflect.TypeOf((dao.UserDao)(nil)).Elem()).(dao.UserDao),
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

func (*userService) Save(user domain.User) (domain.User, error) {
	panic("implement me")
}

func (*userService) Update(user domain.User) (domain.User, error) {
	panic("implement me")
}

func (*userService) Delete(id int) error {
	panic("implement me")
}

func (*userService) FindOne(id int) (*domain.User, error) {
	panic("implement me")
}

func (*userService) FindAll() (domain.Users, error) {
	panic("implement me")
}
