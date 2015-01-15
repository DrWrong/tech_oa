package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
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
	Projects    []*Project `orm:"rel(m2m); rel_through(tech_oa/models.UserProject)"`
	IsAdmin     bool       `orm:"default(0)"` //身份0表示学生 1 表示教师
	IsActive    bool       `orm: "default(1)`
}

type UserProject struct {
	Id      int
	Project *Project `orm:"rel(fk)"`
	User    *User    `orm:"rel(fk)"`
	// Goup   *Group   `orm:"null;rel(one)"`
}

func (u *UserProject) TableName() string {
	return "rel_user_project"
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

func (u *User) GetProjects() []*Project {
	if u == nil {
		return []*Project{}
	}
	if len(u.Projects) > 0 {
		return u.Projects
	}
	o := orm.NewOrm()
	o.LoadRelated(u, "Projects")
	// var projects []*Project
	// o.QueryTable("project").Filter("Users__User__Id", u.Id).All(&projects)
	// u.Projects = projects
	return u.Projects

}

// 现在只允许小组长来操作
func (u *User) GroupSpecifcProject(project *Project) (*Group, error) {
	o := orm.NewOrm()
	// group := Group{Project: project, GroupLeader: u}
	// err := o.Read(&group)
	var group Group
	err := o.QueryTable("group").Filter("Project", project.Id).Filter("GroupLeader", u.Id).One(&group)
	fmt.Println(err)
	return &group, err
}

// func (u *User) IsAdmin() bool {
// 	return u.Role
// }

// func (u *User) IsActive()
