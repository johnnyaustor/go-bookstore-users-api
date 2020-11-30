package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/johnnyaustor/go-bookstore-users-api/app/utils/errors"
	"strings"
)

const(
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(),errorNoRows) {
			return errors.NotFound(fmt.Sprintf("error matching given id: %s ", err.Error()))
		}
		return errors.InternalServerError(fmt.Sprintf("error parse database response: %s ", err.Error()))
	}

	fmt.Println(sqlError.Number)
	fmt.Println(sqlError.Message)
	switch sqlError.Number {
	case 1062:
		return errors.BadRequest("Duplicate Data")
	}
	return errors.InternalServerError("error processing request")
}
