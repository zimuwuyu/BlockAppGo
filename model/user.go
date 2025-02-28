package model

type Role string

const (
	AdminRole      Role = "ADMIN"
	UserRole       Role = "USER"
	ViewerRole     Role = "VIEWER"
	SuperAdminRole Role = "SUPERADMIN"
	SchedulerRole  Role = "SCHEDULER"
)

type User struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:id;varchar(30);not null;comment:'用户名称'"`
	_PassWord   string `gorm:"column:password;string;not null;comment:'用户密码'"`
	_Role       Role   `gorm:"column:role;varchar(30);not null;index;comment:'用户角色'"`
	PhoneNumber string `gorm:"column:phone_number;varchar(30);comment:'用户电话号码'"`
	Email       string `gorm:"column:email;varchar(30);comment:'用户邮箱'"`
	CreateTime  string `gorm:"column:create_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime  string `gorm:"column:update_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}

type UserFeedback struct {
	ID         int    `gorm:"column:id;primaryKey"`
	UserId     int    `gorm:"column:user_id;int;not null;comment:'用户id'"`
	Message    string `gorm:"column:message;string;not null;comment:'反馈信息'"`
	CreateTime string `gorm:"column:create_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime string `gorm:"column:update_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}
