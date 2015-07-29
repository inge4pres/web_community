package community

type Community struct {
	Name       string
	Operations map[string]interface{}
}

type User struct {
	Id, Role                        int64
	Uname, Fname, Lname, Email, Pwd string
}

type Role struct {
	Id   int64
	Name string
}

type CompanyAdmin struct {
	User *User
}

type Company struct {
	Id                        int64
	Name, Vat, Address, Phone string
	Admin                     *CompanyAdmin
	Users                     []*User
	Communities               []Community
}
