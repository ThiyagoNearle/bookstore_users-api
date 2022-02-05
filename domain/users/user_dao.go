package users

import (
	"database/sql"
	"fmt"
	"log"

	users_db "github.com/ThiyagoNearle/bookstore_users-api/dbConfig"
	"github.com/ThiyagoNearle/bookstore_users-api/logger"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/credentials"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser                 = "INSERT INTO users_db.users(first_name, last_name, email, date_created, status, password, config_id) VALUES(?,?,?,?,?,?,?);"
	queryGetUser                    = "SELECT first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser                 = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser                 = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus           = "SELECT id, first_name, last_name, email, date_created, status from users WHERE status = ?;"
	queryFindUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status from users WHERE email=? AND password=? AND status=?;"
	checkusername                   = "SELECT id, first_name, last_name, email, password, config_id from users WHERE email = ? AND config_id=?;"
	loginuserresponse               = "SELECT a.id, IFNULL(a.first_name,'') AS first_name, IFNULL(a.last_name,'') AS last_name,  a.email, a.config_id, b.age, b.shopname, b.location FROM users a, users_profiles b WHERE a.id =b.id AND a.id =?;"
	queryInsertUserProfile          = "INSERT INTO users_profile(id, age, shopname, location) VALUES(?,?,?,?);"
	userAuthentication              = "SELECT a.id, IFNULL(a.first_name,'') AS first_name, IFNULL(a.last_name,'') AS last_name,  a.email, a.config_id, a.status, a.date_created b.age, b.shopname, b.location FROM users a, users_profiles b WHERE a.id =b.id AND a.status =active AND a.id =?;"
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.DB.Prepare(queryGetUser) // preparing query and checking whether the query is right & it throws err , as mentioned in Prepare function

	// users_db = packAGE , Client is a object for connection_string ( database connectivity string), so we access the [ package.object ] , then by using Client object we can able to call all built in functions
	// like result := &users.User{Id: userId}
	//	fmt.Println("err", err)
	//	fmt.Printf("err type %T", err) // so err is a error type , SO we can't pass this as a sting type
	//	fmt.Println("-----------------------------------------------------")
	//	fmt.Println("err.Error()", err.Error())
	//	fmt.Printf("err type %T", err.Error()) // err.Error() is a string type
	if err != nil {
		//logger.Error("error when trying to prepare get user statement", err)
		log.Fatal("errrrrrrrrrrrrror", err.Error())
		return errors.NewInternalServerError("database error") // err.Error() => gives long error message as a string, we dont want this, instaed we directly give our own message
	}
	logger.Info("Get user successfully created.............................")
	fmt.Println("Got.....................................................")
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

	if getErr := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil { //Scan copies the columns values from the matched row & this values stored in address

		logger.Error("error when trying to get user by id", getErr) // internal log

		return errors.NewInternalServerError("database error") // audience or client display

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

func (user *User) SaveUser() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(queryInsertUser) // take the query and validate whether it is right or wrong
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// otherway for above code is
	// result, err := users_db.Client.Exec(queryInsertUser,user.FirstNanme, user.LastName, user.Email, user.DateCreated )

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password) // passing values from struct to database
	if saveErr != nil {                                                                                                         // stmt.Exec(value) insert values to the database and while executing database automatically generate some numbers for id column
		// and insertResult has the result that means row of values
		logger.Error("error when trying to execute save user ", saveErr)

		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId() // we take the id from the inserted row with database, when we use LastInsertId(), it will give the id for the row
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user ", err)

		return errors.NewInternalServerError("database error")
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
	stmt, err := users_db.DB.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement ", err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// stmt.Exec => returns 2 values 1 is result & other one is err
	// returns a Result summarizing the effect of the statement ( result = 1 row affected)
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to Execute the update user statement", err)

		return errors.NewInternalServerError("database error")
	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare Delete user statement", err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close() // always need to close the statement once it made

	_, err = stmt.Exec(user.Id) // in this _ is result that hold the result row ( internally database would be updated based on this query)  & we just need to update in database not taking that result
	if err != nil {             // returns a Result summarizing the effect of the statement ( result = 1 row affected)
		logger.Error("error when trying to Delete user", err)

		return errors.NewInternalServerError("database error")

	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) { // []User means valid list of users to retrieve
	stmt, err := users_db.DB.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare FindBy status user staement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status) // returns pointer to sql.rows  & err
	if err != nil {
		logger.Error("error when trying to find the user by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0) // we dont know how many users we get (rows) , 0 is length  of results, so we can adding values only by append method

	for rows.Next() { // Next passes each result row values & scan those values and save it each field to user struct field orderly
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user) // for the first row, values stored in struct get appended to slice , similarly for all other rows
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users status %s", status))
	}
	return results, nil

}

func (user *User) Login() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(queryFindUserByEmailAndPassword)
	if err != nil {
		logger.Error("Error when trying to prepare the query", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when scan user row into user struct", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) CheckUserName(login_param LoginRequest) (*User, bool) {
	var data User
	stmt, err := users_db.DB.Prepare(checkusername)
	if err != nil {
		return nil, false

	}
	row := stmt.QueryRow(login_param.Email) //"SELECT userid, IFNULL(configid,0) AS configid FROM app_users WHERE (authname= ? OR contactno=?) AND configid=?"
	err = row.Scan(&data.Id, &data.FirstName, &data.LastName, &data.Email, &data.Password, &data.Configid)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		} else {
			return nil, false
		}
	}
	status := credentials.CheckPasswordHash(login_param.Password, data.Password)
	if status == false {
		return nil, false
	}

	return &data, true

}

func (user *User) LoginResponse(id int64) (*User, *errors.RestErr) {
	var data User
	stmt, err := users_db.DB.Prepare(loginuserresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	err = row.Scan(&data.Id, &data.FirstName, &data.LastName, &data.Email, &data.Configid, &data.Age, &data.Shopname,
		&data.Location)

	if err != nil {
		return nil, errors.NewNotFoundError("there is no profile records for the given id ")
	}

	return &data, nil
}

func (user *User) SaveUserProfile() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(queryInsertUserProfile) // take the query and validate whether it is right or wrong
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// otherway for above code is
	// result, err := users_db.Client.Exec(queryInsertUser,user.FirstNanme, user.LastName, user.Email, user.DateCreated )

	_, saveErr := stmt.Exec(user.Id, user.Age, user.Shopname, user.Location) // passing values from struct to database
	if saveErr != nil {                                                      // stmt.Exec(value) insert values to the database and while executing database automatically generate some numbers for id column
		// and insertResult has the result that means row of values
		logger.Error("error when trying to execute save user ", saveErr)

		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) UserAuthentication(id int64) (*User, bool, error) {

	print(id)
	var data User
	stmt, err := users_db.DB.Prepare(userAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	print("qwe")
	row := stmt.QueryRow(id)
	err = row.Scan(&data.Id, &data.FirstName, &data.LastName, &data.Email, &data.Configid, &data.Status, &data.DateCreated, &data.Age, &data.Shopname,
		&data.Location)
	print(err)
	print("qwedd")
	if err != nil {
		print("bbb")
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")

			return &data, false, err // eventhough the data is empty we just need to display that also
		} else {
			log.Fatal(err) // it prints the message (present in bracket) and Exit the current program with the given status code.
			fmt.Println("nodata")

			return &data, false, err // // eventhough the data is empty we just need to display that also
		}

	}

	return &data, true, nil
}
