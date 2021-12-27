package mysqlutils

import (
	"strings"

	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorMessageGetRows   = "no rows in result set"
	errorMessageSaveEmail = "email"
)

func ParseError(err error) *errors.RestErr {
	// converting error to mysql error type ( it contains struct of number , message)

	sqlErr, ok := err.(*mysql.MySQLError)

	// ok returns true, [if err gives 2 values ( number, messgae) & we can change this err type to MysqleRROR type )
	// ok return false, [ if err is string type]
	if !ok { // ok = false ( that means err is string type) // we check that that error message contains our defined string
		if strings.Contains(err.Error(), errorMessageGetRows) {
			return errors.NewNotFoundError("No record matching for given id")
		} //we check that that error message doesn't contains our defined string so it returns
		return errors.NewInternalServerError("error parsing database response")
	}
	// if the control comes to this point means ok = true & we successfully converted this ( that means err is type of mysqlerror type, this mysqlerror will give number & message )
	// sqlErr holds Number & Message

	// starting switch loop

	switch sqlErr.Number { // switch number ( whatever we get from the error)
	case 1062:
		return errors.NewsBadRequestError("invalid data given")
		//case 5455:
		//return "somthing"
	}
	return errors.NewInternalServerError("error processing request") // this return for function level( in this scenario either be error or nil)

}
