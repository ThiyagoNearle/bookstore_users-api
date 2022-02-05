package users

type User struct {
	Userid             int    `json:"userid"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	Token              string `json:"token"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Mobile             string `json:"mobile"`
	Dialcode           string `json:"dialcode"`
	Profileimage       string `json:"profileimage"`
	CreatedDate        string `json:"created"`
	Status             string `json:"status"`
	Roleid             int    `json:"roleid"`
	Configid           int    `json:"configid"`
	Referenceid        int    `json:"referenceid"`
	LocationId         int    `json:"locationid"`
	Moduleid           int    `json:"moduleid"`
	Modulename         string `json:"modulename"`
	Tenantname         string `json:"tenantname"`
	Tenantimage        string `json:"tenantimage"`
	Opentime           string `json:"opentime"`
	Closetime          string `json:"closetime"`
	From               string `json:"from"`
	Tenantaccid        string `json:"tenantaccid"`
	Countrycode        string `json:"countrycode"`
	Currencyid         int    `json:"currencyid"`
	Currencysymbol     string `json:"currencysymbol"`
	Currencycode       string `json:"currencycode"`
	Devicetype         string `json:"devicetype"`
	Usercountrycode    string `json:"usercountrycode"`
	Usercurrencysymbol string `json:"usercurrencysymbol"`
	UsercurrencyCode   string `json:"usercurrencycode"`
	Tenantstatus       string `json:"tenantstatus"`
	Postcode           string `json:"postcode"`
	Tenanttoken        string `json:"tenanttoken"`

}

type CustomResult struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Tenant struct {
	Subscriptionid       int     `json:"subscriptonid"`
	Moduleid             int     `json:"moduleid"`
	Featureid            int     `json:"featureid"`
	Packageid            int     `json:"packageid"`
	Modulename           string  `json:"modulename"`
	Iconurl              string  `json:"iconurl"`
	Logourl              string  `json:"imageurl"`
	Packagename          string  `json:"packagename"`
	Validiydate          string  `json:"validitydate"`
	Validity             bool    `json:"validity"`
	Subcategoryid        int     `json:"subcategoryid"`
	Categoryid           int     `json:"categoryid"`
	Paymentstatus        bool    `json:"paymentstatus"`
	Subscriptionmethodid string  `json:"subscriptionmethodid"`
	Subscriptionaccid    string  `json:"subscriptionaccid"`
	Taxamount            float64 `json:"taxamount"`
	Taxpercent           string  `json:"taxpercent"`
	Totalamount          float64 `json:"totalamount"`
	Status               string  `json:"status"`
}
type Location struct {
	LocationId   int    `json:"locationid"`
	Tenantid     int    `json:"tenantid"`
	Locationname string `json:"locationname"`
	Email        string `json:"email"`
	Contactno    string `json:"contactno"`
	Address      string `json:"address"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	State        string `json:"state"`
	Postcode     string `json:"postcode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Opentime     string `json:"opentime"`
	Closetime    string `json:"closetime"`
}

type LoginResult struct {
	Status      bool       `json:"status"`
	Code        int        `json:"code"`
	Message     string     `json:"message"`
	Userinfo    User       `json:"userinfo"`
	Tenantinfo  []Tenant   `json:"tenantinfo"`
	Locatoninfo []Location `json:"locationinfo"`
}
type CreateUserResult struct {
	Status   bool       `json:"status"`
	Code     int        `json:"code"`
	Message  string     `json:"message"`
	Userinfo CustomUser `json:"userinfo"`
}
type UsersResult struct {
	Status   bool   `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Userinfo []User `json:"userinfo"`
}
type CustomUser struct {
	Userid         int    `json:"userid"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Dialcode       string `json:"dialcode"`
	Profileimage   string `json:"profileimage"`
	CreatedDate    string `json:"created"`
	Status         string `json:"status"`
	Roleid         int    `json:"roleid"`
	Configid       int    `json:"configid"`
	Countrycode    string `json:"countrycode"`
	Currencysymbol string `json:"currencysymbol"`
	Currencycode   string `json:"currencycode"`
}
type DeliveryPartner struct{
	Userid             int    `json:"userid"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	Email              string `json:"email"`
	Contactno             string `json:"contactno"`
	Dialcode           string `json:"dialcode"`
	Userfcmtoken string `json:"userfcmtoken"`	
	Token              string `json:"token"`
	Devicetype         string `json:"devicetype"`
	Profileid int `json:"profileid"`
	Profileimage       string `json:"profileimage"`
	Roleid             int    `json:"roleid"`
	Configid           int    `json:"configid"`
	Referenceid        int    `json:"referenceid"`
	Userlocationid         int    `json:"userlocationid"`
	Partnerid          int    `json:"partnerid"`
	Usertype string `json:"usertype"`
	Bloodgroup string `json:"bloodgroup"`
	Address string `json:"address"`
	Suburb string `json:"suburb"`
	City string `json:"city"`
	State string `json:"state"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Postcode           string `json:"postcode"`
	Countrycode        string `json:"countrycode"`
	Currencyid         int    `json:"currencyid"`
	Currencysymbol     string `json:"currencysymbol"`
	Currencycode       string `json:"currencycode"`
	CreatedDate        string `json:"created"`
	Status             string `json:"status"`
}
