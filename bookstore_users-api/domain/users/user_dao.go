package users

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/datasources/mysql/users_db"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/date_utils"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
)

func (user *User) Get() *errors.RestErr  {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	// prepare the statement
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// populated auto fields
	user.DateCreated = date_utils.GetNowString()

	// perform the database query
	result, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		// attempt to convert error into SQL error
		return mysql_utils.ParseError(saveErr)
	}

	// get result
	userId, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
