package model

import (
	"GinBlog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	// validate 用来做数据校验
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

func CheckUserNotExist(name string) int {
	var data User
	db.Select("id").Where("username=?", name).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}
func CreateUser(user *User) int {
	user.Password = GeneratePasswordHash(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetUsers(pageSize, pageNum int) ([]User, int64) {
	var data []User
	var count int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&data).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return data, count
}
func CheckPassword(user *User) (RoleCode, errCode int) {
	//if CheckUserNotExist(user.Username)!=errmsg.SUCCESS{
	//	return errmsg.ERROR_USER_NOT_EXIST
	//}
	var selectUser User
	err := db.Where("username=?", user.Username).First(&selectUser).Error
	if err != nil {
		log.Println("从数据库中查询密码出错", err, user)
		log.Println(selectUser)
		return 2, errmsg.ERROR
	}
	err = bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(user.Password))
	if err != nil {
		return 2, errmsg.ERROR_PASSWORD_WORON
	}
	return selectUser.Role, errmsg.SUCCESS
}
func GeneratePasswordHash(passwd string) string {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	if err != nil {
		log.Println("ERROR,generate password hash error" + passwd)
		return passwd
	}
	return string(hashPwd)
}

func EditUser(id int, user *User) int {
	data := map[string]interface{}{}
	data["username"] = user.Username
	data["role"] = user.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		log.Println("更新用户信息出错,", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func DeleteUser(id int) int {
	err := db.Model(&User{}).Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		log.Println("删除用户错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
