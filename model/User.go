package model

import (
	"GinBlog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=0" label:"角色码"`
}

func CheckUser(name string)int{
	var data User
	db.Select("id").Where("username=?",name).First(&data)
	if data.ID>0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}
func CreateUser(user *User)int{

	err:=db.Create(&user).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetUsers(pageSize,pageNum int)[]User{
	var data []User
	err:=db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&data).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil
	}
	return data
}