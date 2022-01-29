package users

import (
	"bitbucket/Nearle/ariane/Controllers/middleware"
	"bitbucket/Nearle/ariane/Domain/users"
	"bitbucket/Nearle/ariane/Utils/accessToken"
	"context"
	"errors"
	"net/http"
)

func CreateUser(u users.User) (*users.CreateUserResult, error) {

	var user users.User
	var Result users.CreateUserResult  // display purpose
	var userid int64
	var err error
	user.Firstname = u.Firstname
	user.Lastname = u.Lastname
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Roleid = u.Roleid
	user.Configid = u.Configid
	user.Countrycode = u.Countrycode
	user.Currencycode = u.Currencycode
	user.Currencysymbol = u.Currencysymbol
	user.Dialcode = u.Dialcode
	res1 := u.CheckUser()

	if res1.Userid != 0 {  // consider there is some data in id column in the databse, so we copy and paste
		if res1.Email == u.Email {
			return &Result, errors.New("Email Already Exists")
		} else if res1.Mobile == u.Mobile {
			return &Result, errors.New("Contactno Already Exists")
		}

	} else {
		
			userid, err = u.CreateUser()


	}

	if err != nil {

		Result.Userinfo.Userid = 0        // i think these steps are not needed
		Result.Code = http.StatusBadRequest
		Result.Status = false
		Result.Message = "Unsuccess"
		return &Result, err
	}
	user.Userid = int(userid)
	user.CreateUserProfile()
	res, _ := user.Getuserbyid(userid)
	Result.Userinfo.Userid = res.Userid
	Result.Code = http.StatusCreated
	Result.Status = true
	Result.Message = "Success"
	Result.Userinfo.Firstname = res.Firstname
	Result.Userinfo.Lastname = res.Lastname
	Result.Userinfo.Email = res.Email
	Result.Userinfo.Mobile = res.Mobile
	Result.Userinfo.Roleid = user.Roleid
	Result.Userinfo.Configid = user.Configid
	Result.Userinfo.Dialcode = res.Dialcode
	Result.Userinfo.Status = res.Status
	Result.Userinfo.CreatedDate = res.CreatedDate
	Result.Userinfo.Countrycode = res.Countrycode
	Result.Userinfo.Currencycode = res.Currencycode
	Result.Userinfo.Currencysymbol = res.Currencysymbol

	return &Result, nil
}

func UpdateUser(u users.User) (bool, error) {
	var user users.User
	user.Firstname = u.Firstname
	user.Lastname = u.Lastname
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Profileimage = u.Profileimage
	user.Dialcode = u.Dialcode
	res1 := u.CheckUserforupdate()
	if res1.Userid != 0 {
		if res1.Email == u.Email {
			return false, errors.New("Email Already Exists")
		} else if res1.Mobile == u.Mobile {
			return false, errors.New("Contactno Already Exists")
		}
	} else {

			status, err := u.Updateappuser()
			if err != nil {
				return false, err
			}
			if status == true {
				_, err1 := u.Updateuserprofile()
				if err1 != nil {
					return false, err1
				}
			}
		

	}

	return true, nil
}

func Login(c users.User) (users.LoginResult, bool, error) {
	var user users.User
	var Result users.LoginResult
	var token string
	var err error
	user.Firstname = c.Firstname
	user.Devicetype = c.Devicetype
	user.Tenanttoken = c.Tenanttoken
	var config int
	config = 1

	status := user.CheckUserName(config)
	if status == false {
		return Result, status, nil
	}
	if user.Userid != 0 && user.Configid != 0 {
		token, err = accessToken.GenerateToken(user.Userid, user.Configid)
		if err != nil {
			return Result, false, err
		}
	}
		user.LoginResponse(int64(user.Userid))
		user.InsertToken(token)
		user.Token = token
		var tenantResult []users.Tenant
		var locationResult []users.Location
		if user.Referenceid != 0 {

			if c.Tenanttoken != "" || c.Devicetype != "" {

				status := users.UpdateTenantToken(c.Tenanttoken, c.Devicetype, user.Referenceid)
				print("tentokenupdate=", status)
			}
			tenantResult = users.Tenantresponse(user.Referenceid)
			locationResult = users.Locationresponse(user.Referenceid)
			Result.Userinfo = user
			Result.Tenantinfo = tenantResult
			Result.Locatoninfo = locationResult

		} else {

			Result.Userinfo = user
			Result.Tenantinfo = nil
			Result.Locatoninfo = nil

		}
	

	return Result, true, nil
}
func Getuser(ctx context.Context) (users.CreateUserResult, error) {
	print("stp1")
	var Result users.CreateUserResult

	print("stp2")
	user, _ := middleware.GetCustomContext(ctx)
	print("stp3")
	print("usid==", user.Userid)
	if user.Userid != 0 {
		Result.Status = true
		Result.Code = http.StatusOK
		Result.Message = "Success"
		Result.Userinfo.Userid = user.Userid
		Result.Userinfo.Firstname = user.Firstname
		Result.Userinfo.Lastname = user.Lastname
		Result.Userinfo.Email = user.Email
		Result.Userinfo.Mobile = user.Mobile
		Result.Userinfo.Roleid = user.Roleid
		Result.Userinfo.Configid = user.Configid
		Result.Userinfo.Dialcode = user.Dialcode
		Result.Userinfo.Status = user.Status
		Result.Userinfo.CreatedDate = user.CreatedDate
		Result.Userinfo.Countrycode = user.Countrycode
		Result.Userinfo.Currencycode = user.Currencycode
		Result.Userinfo.Currencysymbol = user.Currencysymbol
	} else {
		Result.Status = false
		Result.Code = http.StatusBadRequest
		Result.Message = "UnSuccess"

	}

	return Result, nil
}

func GetUsers() (users.UsersResult, error) {
	var result users.UsersResult
	var res []users.User
	res = users.GetAllUsers()
	result.Code = http.StatusOK
	result.Message = "Success"
	result.Status = false
	result.Userinfo = res
	return result, nil
}
