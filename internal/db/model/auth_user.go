package model

type AuthUser struct {
	Pk
	UserNo   string `orm:"column(user_no);unique"`
	UserName string `orm:"column(user_name);unique;null"`
	Email    string `orm:"column(email);unique;null"`
	Phone    string `orm:"column(mobile_phone);null"`
	Password string `orm:"column(password)"`
	TableChangeInfo
}
