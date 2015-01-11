package models

type Question struct {
	Id             int
	ChoiceQuestion bool `orm:"default(1)"`
	Desc           string
	Choices        string `orm:"default('')"` //若是选择题，choices为json描述的字符串
}
