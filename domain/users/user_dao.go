package users

import (
	"fmt"

	"github.com/ThiyagoNearle/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
	mysqlutils "github.com/ThiyagoNearle/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users_db.users(first_name, last_name, email, date_created, status password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, date_created, status from users WHERE status = ?;"
)

func (user *User) Get() *errors.RestErr {
	fmt.Println("we are in the starting in the before connecting database")
	stmt, err := users_db.Client.Prepare(queryGetUser) // preparing query and checking whether the query is right & it throws err , as mentioned in Prepare function

	// users_db = packAGE , Client is a object for connection_string ( database connectivity string), so we access the [ package.object ] , then by using Client object we can able to call all built in functions
	// like result := &users.User{Id: userId}
	//	fmt.Println("err", err)
	//	fmt.Printf("err type %T", err) // so err is a error type , SO we can't pass this as a sting type
	//	fmt.Println("-----------------------------------------------------")
	//	fmt.Println("err.Error()", err.Error())
	//	fmt.Printf("err type %T", err.Error()) // err.Error() is a string type
	if err != nil {
		return errors.NewInternalServerError(err.Error()) // err.Error => gives error message as a string ,
	}
	defer stmt.Close()

	// if you retrive 1 row use => QueryRow & it return 1 orw
	// if you retrive more than 1 rows use=> Query  & it returns err & rows

	// if many rows
	/*	 results, err := stmt.Query(user.Id)
	 if err !=nil {
		 return errors.NewInternalServerError(err.Error())
	 }
	 defer results.Close() // if you get many rows as a results, in this case need to close that results that holds fetched rows values from db.

	*/

	// single row
	result := stmt.QueryRow(user.Id) //in queryRow have some condition as we given condition in  WHERE condition , so need to pass value for that.
	// that value has been taken from the url value , we already stored in user variable

	// in result there may be 1 row or 2 many rows base don the query

	// If an error occurs during the execution of the above statement, that error will be returned by a call to Scan on the returned *Row,

	if getErr := result.Scan(&user.FirstNanme, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil { //Scan copies the columns values from the matched row & this values stored in address

		return mysqlutils.ParseError(getErr)
	}
	return nil // for single row results no need to close that

}

// below required only for mock database for original database code thase are not required       fmt.Println("result", result)
/*	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstNanme = result.FirstNanme
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
*/

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser) // take the query and validate whether it is right or wrong
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// otherway for above code is
	// result, err := users_db.Client.Exec(queryInsertUser,user.FirstNanme, user.LastName, user.Email, user.DateCreated )

	insertResult, saveErr := stmt.Exec(user.FirstNanme, user.LastName, user.Email, user.DateCreated, user.Status, user.Password) // passing values from struct to database
	if saveErr != nil {                                                                                                          // stmt.Exec(value) insert values to the database and while executing database automatically generate some numbers for id column
		// and insertResult has the result that means row of values
		return mysqlutils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId() // we take the id from the inserted row with database, when we use LastInsertId(), it will give the id for the row
	if err != nil {
		return mysqlutils.ParseError(saveErr)
	}
	user.Id = userId
	return nil
}

/*current := usersDB[user.Id]
if current != nil {
	if current.Email == user.Email {
		return errors.NewsBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	}
	return errors.NewsBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
}

user.DateCreated = date_utils.GetNowString()

usersDB[user.Id] = user // current = user */

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// stmt.Exec => returns 2 values 1 is result & other one is err
	// returns a Result summarizing the effect of the statement ( result = 1 row affected)
	_, err = stmt.Exec(user.FirstNanme, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)

	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // always need to close the statement once it made

	_, err = stmt.Exec(user.Id) // in this _ is result that hold the result row ( internally database would be updated based on this query)  & we just need to update in database not taking that result
	if err != nil {             // returns a Result summarizing the effect of the statement ( result = 1 row affected)
		return mysqlutils.ParseError(err)

	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) { // []User means valid list of users to retrieve
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status) // returns pointer to sql.rows  & err
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0) // we dont know how many users we get ( rows) , so we value as 0

	for rows.Next() { // Next passes each result row values & scan those values and save it each field to user struct field orderly
		var user User
		if err := rows.Scan(&user.Id, &user.FirstNanme, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlutils.ParseError(err)
		}
		results = append(results, user) // for the first row, values stored in struct get appended to slice , similarly for all other rows
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError((fmt.Sprintf("no users matching status %s", status)))
	}
	return results, nil

}
