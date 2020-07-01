package users

import (
	"fmt"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/datasources/mysql/users_db"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/logger"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES (?, ?, ?, ?, ?, ?)"
	queryGetUser      = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id=?"
	queryUpdateUser   = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser   = "DELETE FROM users WHERE id=?"
	queryFindByStatus = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status=?"
)

func (user *User) Get() *errors_utils.RestErr  {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors_utils.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return errors_utils.NewNotFoundErr("could not find user with given id")
	}
	return nil
}

func (user *User) Save() *errors_utils.RestErr {
	// prepare the statement
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors_utils.NewInternalServerError("database error. couldn't save the user")
	}
	defer stmt.Close()

	// perform the database query
	result, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save user" , err)
		return errors_utils.NewInternalServerError("database error. couldn't save the user")
	}

	// get result
	userId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last user insert id after creating a new user" , err)
		return errors_utils.NewInternalServerError("database error. couldn't save the user")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update statement" , err)
		return errors_utils.NewInternalServerError("database error. couldn't update the user")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error(fmt.Sprintf("error trying to update user with id %d", user.Id) , err)
		return errors_utils.NewInternalServerError("database error. couldn't update the user")
	}
	return nil
}

func (user *User) Delete() *errors_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement" , err)
		return errors_utils.NewInternalServerError("database error. couldn't delete the user")
	}
	defer stmt.Close()
	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user" , err)
		return errors_utils.NewInternalServerError("database error. couldn't delete the user")
	}
	return nil
}

func (user *User) FindByStatus() ([]User, *errors_utils.RestErr)  {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare search user by status statement" , err)
		return nil, errors_utils.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Status)
	if err != nil {
		logger.Error("error when trying to query search user by status statement" , err)
		return nil, errors_utils.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
			logger.Error("error when trying to parse search results into user object" , err)
			return nil, errors_utils.NewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors_utils.NewNotFoundErr(fmt.Sprintf("no users found with status %s", user.Status))
	}
	return results, nil
}
