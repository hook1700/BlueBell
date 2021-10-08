package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "jan"

var (
	ErrorExistUser       = errors.New("用户已存在")
	ErrorNotExistUser    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

//把每一步数据库封装成函数

// CheckUsernameExist 通过用户名检查用户是否存在
func CheckUsernameExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorExistUser
	}
	return
}

// InsertUser 添加用户
func InsertUser(user *models.User) (err error) {

	//加密插入密码
	user.Password = EncryptPassword(user.Password)
	sqlStr := `insert into user (user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return
}

func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func GetUserByUserName(user *models.User) (err error) {
	sqlStr := `select user_id,username,password from user where username =?`
	if err = db.Get(user, sqlStr, user.Username); err != nil {
		return ErrorNotExistUser
	}
	return
}

func GetUserById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id =?`
	if err = db.Get(user, sqlStr, id); err != nil {
		return nil, ErrorNotExistUser
	}
	return
}
