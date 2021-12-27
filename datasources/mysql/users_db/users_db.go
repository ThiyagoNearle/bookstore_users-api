package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // we are not using this but we need this to make connections
)

/*const (
	mysql_username = "mysql_username"
	mysql_password = "mysql_username"
	mysql_host     = "mysql_username"
	mysql_schema   = "mysql_username"
)*/

var (
	Client *sql.DB // client is of sql database type
	//username = os.Getenv(mysql_username)
	//password = os.Getenv(mysql_password)
	//host     = os.Getenv(mysql_host)
	//schema   = os.Getenv(mysql_schema)

)

func init() {
	// connection string = "username:password@tcp(127.0.0.1:3306)/test")
	fmt.Println("-----------------------------------------------------------------init function started in terminal")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		"",
		"127.0.0.1:3306",
		"users_db",
	)
	var err error

	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}
	// if dont have any error you are successfully executed but if you want some output use some function to display
	// after making connection just ping the databse
	if err = Client.Ping(); err != nil { // Ping verifies a connection to the database is still alive,
		panic(err)
	}
	log.Println("database successfully configured") // it means log referes date time messgae(entered)
}
