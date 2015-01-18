package project

import (
	// "github.com/astaxie/beego/orm"
	// "github.com/macaron-contrib/binding"
	// "tech_oa/form"
	"tech_oa/middleware"
	"tech_oa/models"
)

func List(ctx *middleware.Context) {
	var projects []*models.Project
	ctx.Orm.QueryTable("project").All(&projects)
	ctx.Data["projects"] = projects
	ctx.HTML(200, "admin/project/list.html")
}

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
