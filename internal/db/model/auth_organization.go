package model

type AuthOrganization struct {
	Pk
	OrgNo   string `orm:"column(org_no);unique"`
	OrgName string `orm:"column(org_name);unique"`
	TableChangeInfo
}
