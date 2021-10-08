package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//1.查询用户是否存在
	err = mysql.CheckUsernameExist(p.Username)
	if err != nil {
		// 数据查询出错
		return err
	}

	//2.生成UID
	uId := snowflake.GenID()

	//构造一个User实例

	u := &models.User{
		UserId:   uId,
		Username: p.Username,
		Password: p.Password,
	}
	//3.添加新的用户

	//redis xxx
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (user *models.User, token string, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//通过用户名获取密码
	if err = mysql.GetUserByUserName(user); err != nil {
		return nil, "", err
	}

	//查看秘密是否一致
	if mysql.EncryptPassword(p.Password) != user.Password {
		return nil, "", err
	}

	//通过jwt生成token
	token, _ = jwt.GenToken(user.UserId, user.Username)

	return user, token, nil

}
