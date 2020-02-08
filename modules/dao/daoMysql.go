package dao

import (
	"database/sql"
	"fmt"
	"github.com/iproduct/coursego/modules/model"
	"log"
)

type userRepoMysql struct {
	db *sql.DB
}

//FindAll returns all users
func (r *userRepoMysql) Find(start, count int) ([]model.User, error) {
	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := r.db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
	//return nil, errors.New("Not implemented")
}

//FindById return users by user ID or error otherwise
func (r *userRepoMysql) FindByID(id int) (*model.User, error) {
	user := &model.User{}
	statement := fmt.Sprintf("SELECT id, name, age FROM users WHERE id=%d", id)
	err := r.db.QueryRow(statement).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
	//return nil, errors.New("Not implemented")
}

//Create creates and returns new user with autogenerated ID
func (r *userRepoMysql) Create(user *model.User) (*model.User, error) {
	statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", user.Name, user.Age)
	result, err := r.db.Exec(statement)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	user.ID = int(id)
	//err = r.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
	//return nil, errors.New("Not implemented")
}

//Update updates existing user data
func (r *userRepoMysql) Update(user *model.User) (*model.User, error) {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", user.Name, user.Age, user.ID)
	_, err := r.db.Exec(statement)
	return user, err
	//return nil, errors.New("Not implemented")
}

//DeleteById removes and returns user with specified ID or error otherwise
func (r *userRepoMysql) DeleteByID(id int) (*model.User, error) {
	user, err := r.FindByID(id)
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", id)
	_, err = r.db.Exec(statement)
	return user, err
	//return nil, errors.New("Not implemented")
}

// NewMysql is a UserRepo constructor
func NewMysql(user, password, dbname string) UserRepo {
	repo := &userRepoMysql{}
	var err error
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	repo.db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
