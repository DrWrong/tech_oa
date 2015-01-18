package form

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego/orm"
	"github.com/macaron-contrib/binding"
	"tech_oa/middleware"
	"tech_oa/models"
)

type ScoreForm struct {
	GroupId int   `binding:"Requird"`
	Score   []int `binding: "Range(1, 100)"`
	group   models.Group
}

func (s *ScoreForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	o := orm.NewOrm()
	s.group = models.Group{Id: s.GroupId}
	err := o.Read(&s.group)
	if err != nil {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"judge_group_id"},
			Classification: "GroupIdError",
			Message:        "选取的组不正确",
		})
	}
	// fmt.Println(s.Score)
	return errs
}

func (s *ScoreForm) UpdateScore(ctx *middleware.Context) {
	o := orm.NewOrm()
	fmt.Println("Start update score ....")
	configs := ctx.Task.GetScoreConfig()
	for i, config := range configs {
		if s.Score[i] != 0 {
			// fmt.Println(s.Score[i])
			var groupscore models.GroupScore
			err := o.QueryTable("group_score").Filter(
				"FromUser", ctx.User.Id).Filter(
				"JudgeGroup", s.group.Id).Filter(
				"Type", config.Type).One(&groupscore)
			if err == nil && groupscore.Score != s.Score[i] {
				groupscore.Score = s.Score[i]
				groupscore.ScoreWeight = config.ScoreWeight
				groupscore.UserWeight = config.StuWeight
				o.Update(&groupscore)
			} else {
				fmt.Println(err)
				groupscore.FromUser = ctx.User
				groupscore.JudgeGroup = &s.group
				groupscore.Type = config.Type
				groupscore.Score = s.Score[i]
				groupscore.ScoreWeight = config.ScoreWeight
				groupscore.UserWeight = config.StuWeight
				groupscore.Project = ctx.Project
				groupscore.Task = ctx.Task
				_, err := o.Insert(&groupscore)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("insert row ok")
			}
		}

	}
}
