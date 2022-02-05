package credentials

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // // this function return bcrypt hash of the password that is in [] byte
	return string(bytes), err                                       // we convert that into string inoreder to store into database
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) // here we comppare raw password with string of ( bcrypt hash of the password)
	// that will be return from HashPassword function as well as from database we can retrieve
	// compare rawpassword with

	if err != nil {

		return false
	}

	return true
}

// REFER TO COST IN bcrypt.GenerateFromPassword function
// const (
// 	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
// 	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
// 	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
// )
