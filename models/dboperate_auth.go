package models

import (
	"errors"
	"server/globals"
)

// 根据id获取用户信息
func GetUserInfoByID(userID string) (*Users, error) {
	var user Users
	result := globals.DB.Where("userid = ?", userID).Limit(1).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("无数据")
	}
	return &user, nil
}

// 用户登录
func LoginByUsername(username string, password string) (*Users, error) {
	var user Users
	result := globals.DB.Where("username = ? and password = ?", username, password).Limit(1).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("用户名密码错误")
	}
	return &user, nil
}

// 根据id更新用户信息
func UpdateUserInfoByID(user Users) error {
	result := globals.DB.Model(&user).Omit("userid", "password").Where("userid = ?", user.Userid).Updates(user)
	if result.RowsAffected < 1 {
		return result.Error
	}
	return nil
}

// 根据id修改用户密码
func UpdatePasswordByID(userid string, oldpw string, password string) error {
	var user Users
	r1 := globals.DB.Where("userid = ? and password = ?", userid, oldpw).First(&user)
	if r1.RowsAffected < 1 {
		return errors.New("旧密码错误")
	}
	if oldpw == password {
		return errors.New("新密码不得与旧密码相同")
	}
	r2 := globals.DB.Model(&user).Where("userid = ?", userid).Update("password", password)
	if r2.RowsAffected < 1 {
		return errors.New("密码修改失败，请稍后再试")
	}
	return nil
}
