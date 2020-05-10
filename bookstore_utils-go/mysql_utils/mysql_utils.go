package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr  {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundErr("no record found with given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestErr("duplicate etnry")
	}
	return errors.NewBadRequestErr("error processing request")
}