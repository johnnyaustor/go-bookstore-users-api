package users

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/johnnyaustor/go-bookstore-users-api/app/datasources/mysql/users_db"
	"github.com/johnnyaustor/go-bookstore-users-api/app/logger"
	"github.com/johnnyaustor/go-bookstore-users-api/app/utils/errors"
	"github.com/johnnyaustor/go-bookstore-users-api/app/utils/time_utils"
)

const (
	sqlInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	sqlGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	sqlUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	sqlDeleteUser = "DELETE FROM users WHERE id=?;"
	sqlFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (user *User) Get() *errors.RestError {

	stmt, err := users_db.Client.Prepare(sqlGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.InternalServerError("sql error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying get user by id", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(sqlInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.InternalServerError("sql error")
	}
	defer stmt.Close()

	user.DateCreated = time_utils.GetNowString()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.InternalServerError("database error")
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating new user", err)
		return errors.InternalServerError("database error")
	}

	user.Id = insertId
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(sqlUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.InternalServerError("sql error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err!= nil {
		logger.Error("error when trying to update user", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(sqlDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.InternalServerError("sql error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err!= nil {
		logger.Error("error when trying to delete user", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(sqlFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.InternalServerError("sql error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, errors.InternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err:= rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.InternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NotFound(fmt.Sprintf("no users found with status %s", status))
	}
	return results, nil
}
