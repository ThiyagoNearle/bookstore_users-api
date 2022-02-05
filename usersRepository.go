package users

import (
	credentials "bitbucket/Nearle/ariane/Utils/Credentials"
	database "bitbucket/Nearle/ariane/Utils/dbConfig"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	userAuthentication     = "SELECT a.userid,a.roleid,a.configid,b.firstname,b.lastname,b.email,b.contactno,b.dialcode,IFNULL(b.profileimage,'') AS profileimage,IFNULL(b.countrycode,'') AS countrycode,IFNULL(b.currencycode,'') AS currencycode,IFNULL(b.currencysymbol,'') AS currencysymbol,b.status,b.created FROM app_users a, app_userprofiles b WHERE a.userid=b.userid AND a.status ='Active' AND a.userid=?"
	customerAuthentication = "SELECT customerid,firstname,lastname,contactno,email,IFNULL(configid,0) AS configid,postcode  FROM customers WHERE customerid=?"
	createUserwithPassword = "INSERT INTO app_users (authname,password,hashsalt,contactno,dialcode,roleid,configid) VALUES(?,?,?,?,?,?,?)"
	checkUser              = "SELECT IFNULL(userid,0) AS userid,authname,contactno FROM app_users WHERE authname=? AND contactno=? AND configid=?"
	checkUserforUpdate     = "SELECT IFNULL(userid,0) AS userid,authname,contactno FROM app_users WHERE authname=? AND contactno=? AND configid=? AND userid NOT IN (?)"
	createUser             = "INSERT INTO app_users (authname,contactno,dialcode,roleid,configid) VALUES(?,?,?,?,?)"
	createUserProfile      = "INSERT INTO app_userprofiles (userid,firstname,lastname,email,contactno,dialcode,countrycode,currencycode,currencysymbol) VALUES(?,?,?,?,?,?,?,?,?)"
	getUserbyid            = "select userid, firstname,lastname,contactno,dialcode,email,IFNULL(countrycode,'') AS countrycode,IFNULL(currencycode,'') AS currencycode,IFNULL(currencysymbol,'') AS currencysymbol,status,created from app_userprofiles WHERE userid=?"
	updateappuser          = "UPDATE app_users SET authname=? , contactno=?,dialcode=? WHERE userid=?"
	updateuserprofile      = "UPDATE app_userprofiles SET firstname=?,lastname=?,email=?,contactno=?,dialcode=?,profileimage=? WHERE userid=?"
	loginuserresponse      = "SELECT a.userid,a.authname,a.contactno,a.dialcode,a.roleid,a.configid,a.status,a.created,b.firstname,b.lastname,IFNULL(b.profileimage,'') AS profileimage,IFNULL(b.countrycode,'') AS usercountrycode,IFNULL(b.currencycode,'') AS usercurrencycode,IFNULL(b.currencysymbol,'') AS usercurrencysymbol,IFNULL(c.devicetype,'') AS devicetype,IFNULL(c.tenantid,0) AS tenantid,IFNULL(c.tenantname,'') AS tenantname, IFNULL(c.countrycode,'') AS countrycode,IFNULL(c.currencyid,0) AS currencyid,IFNULL(c.currencycode,'') AS currencycode,IFNULL(c.currencysymbol,'') AS currencysymbol,IFNULL(c.tenantimage,'') AS tenantimage,IFNULL(c.tenantaccid,'') AS tenantaccid,IFNULL(c.status,'') AS status FROM app_users a INNER JOIN app_userprofiles b ON a.userid = b.userid LEFT OUTER JOIN tenants c ON a.referenceid=c.tenantid  WHERE   a.userid=?"
	logintenantresponse    = "SELECT a.subscriptionid,IFNULL(a.packageid,0) AS packageid,a.moduleid,a.featureid,a.categoryid,a.subcategoryid,IFNULL(a.validitydate,'') AS validitydate,IF(a.validitydate>=DATE(NOW()), true, false) AS validity,a.paymentstatus,a.taxamount,a.taxpercent,a.totalamount,IFNULL(a.subscriptionaccid,'') AS subscriptionaccid,IFNULL(a.subscriptionmethodid,'') AS subscriptionmethodid,a.status,b.modulename,IFNULL(b.logourl,'') AS logourl,IFNULL(b.iconurl,'') AS iconurl FROM tenantsubscription a,app_module b WHERE a.moduleid=b.moduleid AND a.status='Active'  AND  tenantid=? ORDER BY a.subscriptionid ASC "
	loginlocationresponse  = "SELECT locationid,tenantid,locationname,email,contactno,address,IFNULL(suburb,'') AS suburb,city,state,postcode,IFNULL(latitude,'') AS latitude,IFNULL(longitude,'') AS longitude,IFNULL(opentime,'') AS opentime,IFNULL(closetime,'') AS closetime  FROM tenantlocations  WHERE tenantid=?"
	checkusername          = "SELECT userid,IFNULL(configid,0) AS configid FROM app_users WHERE (authname= ? OR contactno=?) AND configid=?"
	getallusers            = "select userid, firstname,lastname, contactno,dialcode,email,IFNULL(profileimage,'') AS profileimage,IFNULL(countrycode,'') AS countrycode,IFNULL(currencycode,'') AS currencycode,IFNULL(currencysymbol,'') AS currencysymbol, status,created from app_userprofiles"
	updateTenantToken      = "UPDATE tenants SET tenanttoken=?,devicetype=? WHERE tenantid=?"
	insertSessionToken     = "INSERT INTO app_session (userid,sessionname,sessiondate,sessionexpiry) VALUES(?,?,?,?)"
)

//Authentication
func (user *User) UserAuthentication(id int64) (*User, bool, error) {

	print(id)
	var data User
	stmt, err := database.Db.Prepare(userAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	print("qwe")
	row := stmt.QueryRow(id)
	err = row.Scan(&data.Userid, &data.Roleid, &data.Configid, &data.Firstname, &data.Lastname, &data.Email, &data.Mobile, &data.Dialcode, &data.Profileimage,
		&data.Countrycode, &data.Currencycode, &data.Currencysymbol, &data.Status, &data.CreatedDate)
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
func (c *User) CustomerAuthentication(id int64) (*User, bool, error) {

	fmt.Println("enrty in customergetbyid")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(customerAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	err = row.Scan(&data.Userid, &data.Firstname, &data.Lastname, &data.Mobile, &data.Email, &data.Configid, &data.Postcode)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")

			return &data, false, err
		} else {

			fmt.Println("nodata")

			return &data, false, err
		}

	}

	return &data, true, nil
}
func (c *User) Getzeroauth(custid, config int64) *User {

	fmt.Println("enrty in zeroauth")

	var data User
	data.Configid = int(config)
	data.Userid = int(custid)
	fmt.Println("completed 0 auth")
	return &data
}
func (user *User) CheckUserName(config int) bool {
	print("config=", config)
	var data User
	statement, err := database.Db.Prepare(checkusername)

	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(user.Firstname, user.Firstname, config)
	err = row.Scan(&data.Userid, &data.Configid)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			return false

		}
	}
	user.Userid = data.Userid
	user.Configid = data.Configid

	return true
}

//Create or Insert
func (user *User) CreateUserWithPassword() (int64, error) {

	statement, err := database.Db.Prepare(createUserwithPassword)

	if err != nil {

		return 0, err
	}
	defer statement.Close()
	hashedPassword, err := credentials.HashPassword(user.Password)

	res, err1 := statement.Exec(&user.Email, &user.Password, &hashedPassword, &user.Mobile, &user.Dialcode, &user.Roleid, &user.Configid)
	if err1 != nil {

		// log.Fatal(err1)
		return 0, err1

	}
	id, err2 := res.LastInsertId()
	if err2 != nil {

		return 0, err2
	}
	log.Print("Row inserted!")
	return id, nil
}
func (user *User) CreateUser() (int64, error) {

	statement, err := database.Db.Prepare(createUser)

	if err != nil {
		print(err)
		return 0, err

	}
	defer statement.Close()

	res, err1 := statement.Exec(&user.Email, &user.Mobile, &user.Dialcode, &user.Roleid, &user.Configid)
	if err1 != nil {

		fmt.Println(err1)

		return 0, err1

	}
	id, err2 := res.LastInsertId()
	if err2 != nil {
		log.Fatal("Error:", err2.Error()) // If an error occurs it calls log.Fatal to print the error message and stop the programm and doesn't show anything as response
		return 0, err
	}

	return id, nil
}
func (user *User) CreateUserProfile() int64 {
	statement, err := database.Db.Prepare(createUserProfile)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&user.Userid, &user.Firstname, &user.Lastname, &user.Email, &user.Mobile, &user.Dialcode, &user.Countrycode,
		&user.Currencycode, &user.Currencysymbol)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()  // we are passing userid to app_usersprofile table and dont want take that same userid # better we can skip this step also
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id
}
func (user *User) InsertToken(token string) {

	statement, err := database.Db.Prepare(insertSessionToken)
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	sessiondate := time.Now()
	sessionexpiry := time.Hour.Hours()
	res, err := statement.Exec(&user.Userid, token, sessiondate, sessionexpiry)
	if err != nil {
		log.Fatal(err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("token inserted!")

}

//Get
func (user *User) Getuserbyid(id int64) (*User, error) {
	var data User
	stmt, err := database.Db.Prepare(getUserbyid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.Userid, &data.Firstname, &data.Lastname, &data.Mobile, &data.Dialcode, &data.Email, &data.Countrycode, &data.Currencycode, &data.Currencysymbol, &data.Status, &data.CreatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
		} else {
			log.Fatal(err)
		}
	}

	user.Userid = data.Userid
	user.Firstname = data.Firstname
	user.Lastname = data.Lastname
	user.Email = data.Email
	user.CreatedDate = data.CreatedDate
	user.Mobile = data.Mobile
	user.Status = data.Status
	user.Dialcode = data.Dialcode
	fmt.Println("completed")
	return &data, nil
}
func GetAllUsers() []User {
	stmt, err := database.Db.Prepare(getallusers)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Userid, &user.Firstname, &user.Lastname, &user.Mobile, &user.Dialcode, &user.Email, &user.Profileimage,
			&user.Countrycode, &user.Currencycode, &user.Currencysymbol,
			&user.Status, &user.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}
func Tenantresponse(userid int) []Tenant {
	stmt, err := database.Db.Prepare(logintenantresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var list []Tenant
	for rows.Next() {
		var t Tenant
		err := rows.Scan(&t.Subscriptionid, &t.Packageid, &t.Moduleid, &t.Featureid, &t.Categoryid, &t.Subcategoryid, &t.Validiydate, &t.Validity, &t.Paymentstatus, &t.Taxamount, &t.Taxpercent, &t.Totalamount, &t.Subscriptionaccid, &t.Subscriptionmethodid, &t.Status, &t.Modulename, &t.Logourl, &t.Iconurl)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
func Locationresponse(userid int) []Location {
	stmt, err := database.Db.Prepare(loginlocationresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var list []Location
	for rows.Next() {
		var t Location
		err := rows.Scan(&t.LocationId, &t.Tenantid, &t.Locationname, &t.Email, &t.Contactno, &t.Address, &t.Suburb, &t.City,
			&t.State, &t.Postcode, &t.Latitude, &t.Longitude, &t.Opentime, &t.Closetime)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
func (user *User) LoginResponse(id int64) (*User, error) {
	var data User
	stmt, err := database.Db.Prepare(loginuserresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	err = row.Scan(&data.Userid, &data.Email, &data.Mobile, &data.Dialcode, &data.Roleid, &data.Configid, &data.Status,
		&data.CreatedDate, &data.Firstname, &data.Lastname, &data.Profileimage, &data.Usercountrycode,
		&data.UsercurrencyCode, &data.Usercurrencysymbol, &data.Devicetype, &data.Referenceid, &data.Tenantname,
		&data.Countrycode, &data.Currencyid, &data.Currencycode, &data.Currencysymbol, &data.Tenantimage,
		&data.Tenantaccid, &data.Tenantstatus)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
		} else {
			log.Fatal(err)
		}
	}

	user.Userid = data.Userid
	user.Firstname = data.Firstname
	user.Lastname = data.Lastname
	user.Email = data.Email
	user.CreatedDate = data.CreatedDate
	user.Mobile = data.Mobile
	user.Status = data.Status
	user.Referenceid = data.Referenceid
	user.Roleid = data.Roleid
	user.Configid = data.Configid
	user.Tenantname = data.Tenantname
	user.Tenantimage = data.Tenantimage
	user.Profileimage = data.Profileimage
	user.Tenantaccid = data.Tenantaccid
	user.Countrycode = data.Countrycode
	user.Currencycode = data.Currencycode
	user.Currencyid = data.Currencyid
	user.Currencysymbol = data.Currencysymbol
	user.Devicetype = data.Devicetype
	user.Usercountrycode = data.Usercountrycode
	user.UsercurrencyCode = data.UsercurrencyCode
	user.Usercurrencysymbol = data.Usercurrencysymbol
	user.Dialcode = data.Dialcode
	user.Tenantstatus = data.Tenantstatus

	return &data, nil
}
func (user *User) CheckUserforupdate() *User {
	var data User
	statement, err := database.Db.Prepare(checkUserforUpdate)

	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(user.Email, user.Mobile, user.Configid, user.Userid)
	err = row.Scan(&data.Userid, &data.Email, &data.Mobile)
	print(err)

	if err != nil {
		if err == sql.ErrNoRows {
			data.Userid = 0
			return &data
		} else {
			data.Userid = 0
			return &data

		}
	}
	// fmt.Println(user)
	user.Userid = data.Userid
	user.Email = data.Email
	user.Mobile = data.Mobile
	return &data
}
func (user *User) CheckUser() *User {
	var data User
	statement, err := database.Db.Prepare(checkUser)

	if err != nil {
		log.Fatal(err) // Log.Fatal is not logger, it’s just another kind of panic but it just print the error and exit from the execution.
	}

	row := statement.QueryRow(user.Email, user.Mobile, user.Configid) // if we creating for the first there will be no row () => that means at program side it is error (throwing error)
	err = row.Scan(&data.Userid, &data.Email, &data.Mobile)           // In database it is authname  & here it is Email (we storing authname in Email)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			data.Userid = 0
			return &data
		} else {
			data.Userid = 0 // in query we gave if null, we save 0
			return &data

		}
	}
	// fmt.Println(user)
	user.Userid = data.Userid // this step the user already exist

	return &data
}

//update
func (u *User) Updateappuser() (bool, error) {
	stmt, err := database.Db.Prepare(updateappuser)
	if err != nil {
		fmt.Println(err)
		return false, err

	}

	_, err1 := stmt.Exec(&u.Email, &u.Mobile, &u.Dialcode, &u.Userid)
	if err1 != nil {
		fmt.Println(err1)
		return false, err1

	}
	return true, nil

}
func (u *User) Updateuserprofile() (bool, error) {
	stmt, err := database.Db.Prepare(updateuserprofile)
	if err != nil {
		fmt.Println(err)
		return false, err

	}

	_, err1 := stmt.Exec(&u.Firstname, &u.Lastname, &u.Email, &u.Mobile, &u.Dialcode, &u.Profileimage, &u.Userid)
	if err1 != nil {
		fmt.Println(err1)
		return false, err1

	}
	return true, nil

}
func UpdateTenantToken(token, devicetype string, tenantid int) bool {
	stmt, err := database.Db.Prepare(updateTenantToken)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(token, devicetype, tenantid)
	if err != nil {
		log.Fatal(err)
	}
	return true

}
