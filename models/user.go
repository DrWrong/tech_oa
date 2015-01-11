package models

type User struct {
	Id          string `orm:"pk"` //学生的学号或者教师的工号
	Name        string
	Major       string
	Grade       int
	ClassNumber string
	Email       string     `orm:"default('')"`
	Password    string     `orm:"default('')"`
	Projects    []*Project `orm:"reverse(many)"`
	Role        int        `orm:"default(0)"` //身份0表示学生 1 表示教师
	Avaliable   bool       `orm: "default(1)`
}
