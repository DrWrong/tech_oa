package score

import (
	"github.com/macaron-contrib/binding"
	"tech_oa/form"
	"tech_oa/middleware"
)

func Home(ctx *middleware.Context) {
	ctx.HTML(200, "project/score")
}

func JudgeScore(ctx *middleware.Context, errs binding.Errors, scorform form.ScoreForm) {
	if errs.Len() > 0 {
		ctx.Data["errors"] = errs
		ctx.HTML(200, "project/score")
		return
	}

	// group, err := ctx.User.GroupSpecifcProject(ctx.Project)
	// if err != nil {
	// 	ctx.Data["error"] = "目前只支持小组长登陆"
	// 	ctx.HTML(200, "project/detail")
	// 	return
	// }
	scorform.UpdateScore(ctx)
	ctx.Data["success"] = "评分成功"
	ctx.HTML(200, "project/score")
	// ctx.Redirect("/")
}
