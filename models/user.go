package models

import (
	"crypto/md5"
	"fmt"
	"github.com/macaron-contrib/session"
)

type User struct {
	Id          string `orm:"pk"` //学生的学号或者教师的工号
	Name        string
	Major       string
	Grade       int
	ClassNumber string
	Email       string     `orm:"default('')"`
	Password    string     `orm:"default('')"`
	Projects    []*Project `orm:"reverse(many)"`
	IsAdmin     bool       `orm:"default(0)"` //身份0表示学生 1 表示教师
	IsActive    bool       `orm: "default(1)`
}

func (u *User) CheckPassword(password string) bool {
	if !u.IsActive {
		return false
	}
	data := []byte(password)
	// md5byte := md5.Sum(data)
	return u.Password == fmt.Sprintf("%x", md5.Sum(data))
}

func (u *User) Login(sess session.Store) {
	sess.Set("uid", u.Id)
}

func (u *User) Logout(sess session.Store) {
	sess.Delete("uid")
}

// func (u *User) IsAdmin() bool {
// 	return u.Role
// }

// func (u *User) IsActive()
