package accessToken

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
	"github.com/golang-jwt/jwt"
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(userid, configid int) (string, error) {
	var Key string
	Key = "nearleapp"
	SecretKey := []byte(Key) // always secret key should be in byte []
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims) // default structure for create a claim object ( that claim is map type)
	/* Set token claims */ // create & assiging values to that claim objects
	claims["userid"] = userid
	claims["configid"] = configid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // it gives sec of current time from old specific date
	// So token has a claim, in that claim values are stored....
	tokenString, err := token.SignedString(SecretKey) // passing secret key to that token(holding some claims) to create token string..........
	if err != nil {
		errors.UnauthorizedError("Error in Generating key")
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (userid, configid float64, Error error) { // here we getting the token string  => from this token string we can retrieve the values that are stored in claims.
	var Key string // this token string has some key
	Key = "nearleapp"
	SecretKey := []byte(Key) // always secret key should be in byte []

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) { // we need to pass the same secret key inorder to get the token
		return SecretKey, nil

	})
	fmt.Println("-----------------------------step4.6")
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // pass that token  to jwt.mapclaims to get that claims ( that holdiong values previously stored.)

		userid := claims["userid"].(float64) // take the values stored in that claims...

		var tm time.Time
		// specifically they use switch here becas in switch we mentioned .(type) so > switch points to type but iat variable holding the value of claim['exp'] as ususal....
		switch iat := claims["exp"].(type) { // we can't use .type outside of switch statement...
		case float64:
			tm = time.Unix(int64(iat), 0) // we convert that seconds in to time
		case json.Number:
			v, _ := iat.Int64() // in json number we can't ableto convert like above step, first we convert that in to int64
			tm = time.Unix(v, 0)

		}

		fmt.Println("raju", tm) // we just see the token time... thats all.......

		return userid, configid, nil // we return only userid, configid
	} else {

		return 0, 0, err // when we pass token to claims, if we get error, then return this step..
	} // in this error we have given nil...
}
func CheckTokenSession(tokenStr string) (userid, configid float64, expirytime time.Time, status bool, Error error) {
	var Key string
	var expTime time.Time

	Key = "nearleapp"
	SecretKey := []byte(Key)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userid := claims["userid"].(float64)
		configid := claims["configid"].(float64)
		var tm time.Time
		switch iat := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(iat), 0) // from sec get the time, when it going to expiry date time
		case json.Number:
			v, _ := iat.Int64()
			tm = time.Unix(v, 0)

		}

		fmt.Println("raju", tm)

		return userid, configid, tm, ok, nil // we return only userid, configid, token time
	} else {

		return 0, 0, expTime, false, err // inthis, expTime is  =>  0:0:0  when we pass token to claims, if we get error, then return this step..
	}
}
