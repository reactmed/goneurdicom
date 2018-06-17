package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/reactmed/goneurdicom/service"
)

var userHandlerInstance UserHandler

func NewUserHandler() UserHandler {
	if userHandlerInstance == nil{
		userHandlerInstance = &userHandler{}
	}
	return userHandlerInstance
}

type UserHandler interface {
	FindUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindUserById(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	SaveUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type userHandler struct {
	service service.UserService
}

func (*userHandler) FindUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("implement me")
}

func (*userHandler) FindUserById(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("implement me")
}

func (*userHandler) SaveUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("implement me")
}

func (*userHandler) UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("implement me")
}

func (*userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("implement me")
}
