package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} { // marshall is a user define function
	results := make([]interface{}, len(users)) // lets take 5 for len(users) & the length for results is 5
	for index, user := range users {
		results[index] = user.Marshall(isPublic) // user.Marshall(isPublic) receives either { 1 date active} or {1 thi yag t@mail date active}
	}
	return results
}

/// if we have different key in json like for user { Id `json:"user_id`} & for public struct we have { Id `json:"id`}  => we have different json id, so we can we go with below method
/// ( this method also can be used for same id in user struct & public struct)

func (user *User) Marshall(isPublic bool) interface{} { // interface{}   => we can return any data type ( whether it may be string, int, struct)
	// though we have given interface{} as a type for return, if you return a variable that hold struct means that variable is of type struct
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	/// if we have same key in json like for user { Id `json:"id`} & for public struct we have { Id `json:"id`}  => we have same json id, so we can we go with below method
	/// ( this method can't be used for different id in user struct & public struct)

	// Marshal returns the JSON encoding of user
	// String values encode as JSON strings
	// Floating point, integer, and Number values encode as JSON numbers.

	userJson, _ := json.Marshal(user) // MARSHAL TAKES THE INTERFACE(any type) AS THE INPUT & GIVE BYTES & ERROR AS OUTPUT

	var PrivateUser PrivateUser

	// Unmarshal uses the inverse of the encodings that Marshal uses.
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	// If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

	json.Unmarshal(userJson, &PrivateUser)
	return PrivateUser

}

/*
	func (user *User) Marshall(isPublic bool) interface{} { // interface{}   => we can return any data type ( whether it may be string, int, struct)
	// though we have given interface{} as a type for return, if you return a variable that hold struct means that variable is of type struct
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	var PrivateUser PrivateUser
	return PrivateUser{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateCreated: user.DateCreated,
		Status:      user.Status,
	}

}

*/
