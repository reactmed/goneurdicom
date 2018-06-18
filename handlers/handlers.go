package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/reactmed/goneurdicom/service"
	"github.com/reactmed/goneurdicom/utils"
	"reflect"
	"strconv"
	"encoding/json"
	"github.com/reactmed/goneurdicom/domain"
)

var userHandlerInstance UserHandler

func NewUserHandler(ctx utils.AppContext) UserHandler {
	if userHandlerInstance == nil {
		userService, err := ctx.Get(reflect.TypeOf((*service.UserService)(nil)).Elem())
		if err != nil {
			panic(err)
		}
		userHandlerInstance = &userHandler{
			service: userService.(service.UserService),
		}
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

func (handler *userHandler) SaveUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	user := domain.User{}
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err = handler.service.Save(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (handler *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	decoder := json.NewDecoder(r.Body)
	user := domain.User{}
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Id = id
	user, err = handler.service.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (handler *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	err := handler.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *userHandler) FindUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _ := handler.service.FindAll()
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (handler *userHandler) FindUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	user, _ := handler.service.FindOne(id)
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
