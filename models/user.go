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

func (u *User) GroupSpecifcProject(project *Project) (*Group, error) {
	o := orm.NewOrm()
	// group := Group{Project: project, GroupLeader: u}
	// err := o.Read(&group)
	var group Group
	err := o.QueryTable("group").Filter("Project", project.Id).Filter("GroupLeader", u.Id).One(&group)
	fmt.Println(err)
	// if err == nil {}
	return &group, err
}

func (u *User) CanAcessProject(project *Project) bool {
	// err := o.QueryTable("rel_user_project").Filter()
	if u.IsAdmin {
		return true
	}
	for _, uproject := range u.GetProjects() {
		if uproject == project {
			return true
		}
	}
	return false
}

type RScore struct {
	Desc  string
	Score int
}

type ScoreResponse struct {
	Group  *Group
	Scores []RScore
}

func (u *User) GetJudegGroupScores(project *Project) []ScoreResponse {
	var groupscores []*GroupScore
	groups := project.GetGroups()
	groupIds := make([]int, len(groups))
	for i, group := range groups {
		groupIds[i] = group.Id
	}
	o := orm.NewOrm()
	o.QueryTable("group_score").Filter(
		"FromUser", u.Id).Filter(
		"JudgeGroup__Id__in", groupIds).RelatedSel(
		"JudgeGroup").RelatedSel("Task").OrderBy(
		"JudgeGroup__Id", "type").All(&groupscores)

	scoreResponses := []ScoreResponse{}
	scoreResponse := ScoreResponse{
		Group:  groupscores[0].JudgeGroup,
		Scores: []RScore{},
	}
	group := groupscores[0].JudgeGroup

	for _, groupscore := range groupscores {
		if groupscore.JudgeGroup.Id != group.Id {
			// fmt.Println("new scoreResponse")
			scoreResponses = append(scoreResponses, scoreResponse)
			scoreResponse = ScoreResponse{
				Group:  groupscore.JudgeGroup,
				Scores: []RScore{},
			}
			group = groupscore.JudgeGroup
		}

		scoreResponse.Scores = append(scoreResponse.Scores, RScore{
			Score: groupscore.Score,
			Desc:  groupscore.GetTypeDesc(),
		})
		// fmt.Printf("desc: %s\n", groupscore.GetTypeDesc())
		// fmt.Println(len(scoreResponse.Scores))

		// fmt.Println(scoreResponse.Scores)
	}
	scoreResponses = append(scoreResponses, scoreResponse)
	fmt.Println(scoreResponses)
	return scoreResponses
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

// func (u *User) IsAdmin() bool {
// 	return u.Role
// }

// func (u *User) IsActive()
