package model

type AuthUserOrganizationV struct {
	Pk
	UserNo   string `orm:"column(user_no);unique"`
	UserName string `orm:"column(user_name);unique;null"`
	Email    string `orm:"column(email);unique;null"`
	Phone    string `orm:"column(mobile_phone);unique;null"`
	Password string `orm:"column(password)"`
	OrgId    int    `orm:"column(org_id)"`
	TableChangeInfo
}
