package project

import (
	// "github.com/astaxie/beego/orm"
	"github.com/macaron-contrib/binding"
	"tech_oa/form"
	"tech_oa/middleware"
	// "tech_oa/models"
)

func Create(ctx *middleware.Context) {
	return
}

func Detail(ctx *middleware.Context) {
	// Id, err := strconv.Atoi(ctx.Params(":id"))
	// if err == nil {
	// 	o := orm.NewOrm()
	// 	project := models.Project{Id: Id}
	// 	err := o.Read(&project)
	// 	if err == nil {
	// 		ctx.Data["project"] = project
	// 		ctx.Data["groups"] = project.GetGroups()
	// 		ctx.HTML(200, "project/detail")
	// 	}
	// }
	ctx.Data["project"] = *ctx.Project
	ctx.Data["groups"] = ctx.Project.GetGroups()
	ctx.HTML(200, "project/detail")
	return
}

func Score(ctx *middleware.Context) {
	group, _ := ctx.User.GroupSpecifcProject(ctx.Project)
	ctx.Data["groupscores"] = group.GetJudegGroupScores()
	ctx.HTML(200, "project/score")
	return
}

func JudgeScore(ctx *middleware.Context, errs binding.Errors, scorform form.ScoreForm) {
	if errs.Len() > 0 {
		ctx.Data["errors"] = errs
		ctx.HTML(200, "project/detail")
		return
	}
	group, err := ctx.User.GroupSpecifcProject(ctx.Project)
	if err != nil {
		ctx.Data["error"] = "目前只支持小组长登陆"
		ctx.HTML(200, "project/detail")
		return
	}
	scorform.UpdateScore(group)
	ctx.Data["success"] = "评分成功"
	ctx.HTML(200, "project/detail")
}
