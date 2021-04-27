package model

type AuthUser struct {
	Pk
	UserNo   string `orm:"column(user_no);unique"`
	UserName string `orm:"column(user_name);unique"`
	Email    string `orm:"column(email);unique"`
	Phone    string `orm:"column(mobile_phone);unique"`
	TableChangeInfo
}
