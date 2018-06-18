package main

import (
	"github.com/reactmed/goneurdicom/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"database/sql"
	"github.com/reactmed/goneurdicom/utils"
	"reflect"
	"github.com/reactmed/goneurdicom/dao"
	"github.com/reactmed/goneurdicom/service"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=neurdicom password=neurdicom dbname=neurdicom_test sslmode=disable"

	ctx := utils.GetAppContext()

	ctx.Bind(reflect.TypeOf((*sql.DB)(nil)).Elem(), func() (interface{}, error) {
		return sql.Open("postgres", connStr)
	}, utils.Request)
	ctx.Bind(reflect.TypeOf((*dao.UserDao)(nil)).Elem(), dao.NewUserDao(ctx), utils.Singleton)
	ctx.Bind(reflect.TypeOf((*service.UserService)(nil)).Elem(), service.NewUserService(ctx), utils.Singleton)

	userHandler := handlers.NewUserHandler(ctx)
	router := httprouter.New()
	router.GET("/api/users/:id", userHandler.FindUserById)
	router.PUT("/api/users/:id", userHandler.UpdateUser)
	router.DELETE("/api/users/:id", userHandler.DeleteUser)
	router.GET("/api/users", userHandler.FindUsers)
	router.POST("/api/users", userHandler.SaveUser)
	log.Println("Server is started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
