package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors_utils.RestErr  {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors_utils.NewNotFoundErr("no record found with given id")
		}
		return errors_utils.NewInternalServerError("error parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return errors_utils.NewBadRequestErr("duplicate etnry")
	}
	return errors_utils.NewBadRequestErr("error processing request")
}