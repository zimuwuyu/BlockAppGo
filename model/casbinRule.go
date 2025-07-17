package model

type CasbinRule struct {
	ID            int    `gorm:"column:id;type:integer;primaryKey"`
	PType         string `gorm:"column:ptype;type:varchar(100);not null;comment:'policy_definition'"`
	Role          Role   `gorm:"column:role;varchar(30);not null;index;comment:'用户角色'"`
	UrlRule       string `gorm:"column:urlRule;varchar(100);not null;comment:'url规则'"`
	RequestMethod string `gorm:"column:requestMethod;type:varchar(200);comment:'请求方式'"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
