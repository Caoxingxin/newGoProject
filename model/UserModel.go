package model

//用户表
type Users struct {
	ID int64
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
	CreateTime int64 `gorm:"column:createtime"`
	//关联表
	Profile Profile `gorm:"foreignkey:UserId"`
}

//返回表名
func (user Users) TableName() string  {
	return "users"
}
