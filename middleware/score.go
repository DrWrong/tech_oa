package middleware

import (
	"github.com/Unknwon/macaron"
	// "github.com/astaxie/beego/orm"
	"tech_oa/models"
)

func ScoreMiddleWare() macaron.Handler {
	return func(ctx *Context) {
		var t models.Task
		qs := models.GetAvaliableTasks()
		qs.Filter("Project", ctx.Project.Id).Filter("type", 3).One(&t)
		ctx.Task = &t
	}
}
