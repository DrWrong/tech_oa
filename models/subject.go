package models

type Subject struct {
	Id   int
	Name string
	Desc string `orm:"default('')"`
}

type ProjectSubjectChosen struct {
	Id      int
	Project *Project `orm:"rel(fk)"`
	Subject *Subject `orm:"rel(fk)"`
	Group   *Group   `orm:"null;rel(one)"`
}

// func(ProjectSubjectChosen)TableName(){
//     return ""
// }
