package dao

import (
	"github.com/reactmed/goneurdicom/domain"
	"github.com/reactmed/goneurdicom/utils"
	"reflect"
	"database/sql"
	"log"
	"errors"
)

var instance UserDao

func NewUserDao(ctx utils.AppContext) UserDao {
	if instance == nil {
		instance = &userDao{
			ctx: ctx,
		}
	}
	return instance
}

type UserDao interface {
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int) error
	FindOne(id int) (*domain.User, error)
	FindAll() ([]domain.User, error)
	Exists(id int) (bool, error)
	Count() (int, error)
}

type userDao struct {
	ctx utils.AppContext
}

func (dao *userDao) getDb() (*sql.DB, error) {
	db, err := dao.ctx.Get(reflect.TypeOf((*sql.DB)(nil)).Elem())
	return (db).(*sql.DB), err
}

func (dao *userDao) Save(user domain.User) (domain.User, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	stmt, err := db.Prepare(`INSERT INTO users (name, surname, email, password) VALUES ($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	err = stmt.QueryRow(user.Name, user.Surname, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

func (dao *userDao) Update(user domain.User) (domain.User, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	stmt, err := db.Prepare("UPDATE users SET name = $1, surname = $2, email = $3, password = $4 WHERE id = $5")
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	_, err = stmt.Exec(user.Name, user.Surname, user.Email, user.Password, user.Id)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

func (dao *userDao) Delete(id int) error {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (dao *userDao) FindOne(id int) (*domain.User, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := db.Prepare("SELECT name, surname, email, password FROM users WHERE id = $1")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user := &domain.User{}
	err = stmt.QueryRow(id).Scan(&user.Name, &user.Surname, &user.Email, &user.Password)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil
}

func (dao *userDao) FindAll() ([]domain.User, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := db.Prepare("SELECT id, name, surname, email, password FROM users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	users := make(domain.Users, 0)
	for rows.Next() {
		user := domain.User{}
		rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Password)
		users = append(users, user)
	}
	return users, nil
}

func (dao *userDao) Exists(id int) (bool, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	stmt, err := db.Prepare("SELECT EXISTS(SELECT * FROM users WHERE id = $1)")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	var res bool
	err = stmt.QueryRow(id).Scan(&res)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return res, nil
}

func (dao *userDao) Count() (int, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	rows, err := db.Query("SELECT COUNT(*) FROM users")
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var count int
		rows.Scan(&count)
		return count, nil
	}
	return -1, errors.New("empty rows")
}
