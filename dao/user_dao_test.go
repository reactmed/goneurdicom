package dao

import (
	"testing"
	"github.com/reactmed/goneurdicom/utils"
	"github.com/reactmed/goneurdicom/domain"
	"sort"
	"database/sql"
	_ "github.com/lib/pq"
	"reflect"
)

func CreateDao() UserDao {
	connStr := "user=neurdicom password=neurdicom dbname=neurdicom_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Can not connect to db")
	}
	ctx := utils.GetAppContext().Bind(reflect.TypeOf((*sql.DB)(nil)).Elem(), db, utils.Singleton)
	db.Exec(
		`
	TRUNCATE users CASCADE;
	INSERT INTO users (id, name, surname, email, password) VALUES 
	(1, 'Name 1', 'Surname 1', 'email1@mail.ru', 'password1'),
	(2, 'Name 2', 'Surname 2', 'email2@mail.ru', 'password2');
	SELECT setval('users_id_seq'::REGCLASS, 3);
	`)
	dao := NewUserDao(ctx)
	return dao
}

func TestUserDao_Save(t *testing.T) {
	dao := CreateDao()
	user := domain.User{
		Name:     "Name 3",
		Surname:  "Surname 3",
		Email:    "email3@mail.ru",
		Password: "password3",
	}
	user, err := dao.Save(user)
	utils.AssertNotErr(err, t)
	if user, err := dao.FindOne(user.Id); err != nil || user == nil {
		t.Error("User not found")
	}
	utils.AssertEqual(user.Name, "Name 3", t)
	utils.AssertEqual(user.Surname, "Surname 3", t)
	utils.AssertEqual(user.Email, "email3@mail.ru", t)
	utils.AssertEqual(user.Password, "password3", t)
}

func TestUserDao_Update(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "Name New",
		Surname:  "Surname New",
		Email:    "email_new@mail.ru",
		Password: "password_new",
	}
	dao := CreateDao()
	user, err := dao.Update(user)
	utils.AssertNotErr(err, t)
	user2, err := dao.FindOne(1)
	utils.AssertNotNil(user2, t)

	utils.AssertEqual(user.Name, "Name New", t)
	utils.AssertEqual(user.Surname, "Surname New", t)
	utils.AssertEqual(user.Email, "email_new@mail.ru", t)
	utils.AssertEqual(user.Password, "password_new", t)
}

func TestUserDao_Delete(t *testing.T) {
	dao := CreateDao()
	dao.Delete(1)
	res, _ := dao.Exists(1)
	utils.AssertFalsef(res, t, "User should be removed")
}

func TestUserDao_FindOne(t *testing.T) {
	dao := CreateDao()
	user, err := dao.FindOne(1)
	utils.AssertNilf(err, t, "Can not find user")
	utils.AssertNotNilf(user, t, "User not found")
	utils.AssertEqual(user.Name, "Name 1", t)
	utils.AssertEqual(user.Surname, "Surname 1", t)
	utils.AssertEqual(user.Email, "email1@mail.ru", t)
	utils.AssertEqual(user.Password, "password1", t)
}

func TestUserDao_FindAll(t *testing.T) {
	dao := CreateDao()
	users, _ := dao.FindAll()
	users2 := make([]interface{}, len(users))
	for i, v := range users {
		users2[i] = v
	}
	utils.AssertHasLen(users2, 2, t)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})
	user := users[0]
	utils.AssertEqual(user.Name, "Name 1", t)
	utils.AssertEqual(user.Surname, "Surname 1", t)
	utils.AssertEqual(user.Email, "email1@mail.ru", t)
	utils.AssertEqual(user.Password, "password1", t)
	user = users[1]
	utils.AssertEqual(user.Name, "Name 2", t)
	utils.AssertEqual(user.Surname, "Surname 2", t)
	utils.AssertEqual(user.Email, "email2@mail.ru", t)
	utils.AssertEqual(user.Password, "password2", t)
}

func TestUserDao_Exists(t *testing.T) {
	dao := CreateDao()
	res, _ := dao.Exists(1)
	utils.AssertTruef(res, t, "User should exist")
}

func TestUserDao_Count(t *testing.T) {
	dao := CreateDao()
	c, _ := dao.Count()
	utils.AssertEqual(c, 2, t)
}
