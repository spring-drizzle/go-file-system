package models

type User struct {
	Id       string `gorm:"type:varchar(32);primary_key"`
	Name     string `gorm:"type:varchar(50)"`
	Surname  string `gorm:"type:varchar(50)"`
	Username string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`
}
