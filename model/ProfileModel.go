package model

//档案
type Profile struct {
	ID int64
	Name string `gorm:"column:name"`
	//User Users `gorm:"foreignkey:id"`
	UserId int64 `gorm:"column:userId"`
}

func (profile Profile) TableName() string {
	return "profile"
}
