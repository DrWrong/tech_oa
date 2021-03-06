package middleware

import (
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego/orm"
	"strconv"
	"tech_oa/models"
)

func ProjectMiddleware() macaron.Handler {
	return func(ctx *Context) {
		id, err := strconv.Atoi(ctx.Params(":id"))
		if err == nil {
			o := orm.NewOrm()
			project := models.Project{Id: id}
			err := o.Read(&project)
			if err == nil {
				ctx.Data["project"] = &project
				ctx.Project = &project
			} else {
				ctx.Error(404, "project not fuond")
				return
			}
		} else {
			ctx.Error(404, "project not found")
			return
		}

		if !ctx.User.CanAcessProject(ctx.Project) {
			ctx.Error(403, "don't hava right to access this project")
		}

	}
}
