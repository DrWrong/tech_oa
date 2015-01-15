package form

import (
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego/orm"
	"github.com/macaron-contrib/binding"
	"tech_oa/models"
)

type ScoreForm struct {
	JudgeGroupId int `binding:"Requird"`
	Score        int `binding: "Range(1, 100)"`
	group        models.Group
}

func (s *ScoreForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	o := orm.NewOrm()
	s.group = models.Group{Id: s.JudgeGroupId}
	err := o.Read(&s.group)
	if err != nil {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"judge_group_id"},
			Classification: "GroupIdError",
			Message:        "选取的组不正确",
		})
	}
	return errs
}

func (s *ScoreForm) UpdateScore(fromgroup *models.Group) {
	o := orm.NewOrm()
	var groupscore models.GroupScore
	err := o.QueryTable(groupscore).Filter("FromGroup", fromgroup.Id).Filter("judgeGroup", s.JudgeGroupId).One(&groupscore)
	groupscore.Score = s.Score
	o.Update(&groupscore)
	if err != nil {
		groupscore.FromGroup = fromgroup
		groupscore.JudgeGroup = &s.group
		groupscore.Score = s.Score
		o.Insert(&groupscore)
	}
}
