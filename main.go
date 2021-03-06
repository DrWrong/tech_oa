package main

import (
	// "fmt"
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/pongo2"
	"github.com/macaron-contrib/session"
	_ "github.com/macaron-contrib/session/mysql"
	"os"
	// "tech_oa/middleware/binding"
	"tech_oa/controllers/project"
	"tech_oa/controllers/score"
	"tech_oa/controllers/user"
	"tech_oa/form"
	"tech_oa/middleware"
	_ "tech_oa/models"
	_ "tech_oa/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	// fmt.Println(beego.AppConfig.String("mysqlstring"))
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlstring"))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}

func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("static"))
	m.Use(session.Sessioner(session.Options{
		Provider:       "mysql",
		ProviderConfig: beego.AppConfig.String("mysqlstring"),
	}))
	m.Use(middleware.Contexter())
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory: "views",
	}))
	return m

}

func runWeb() {
	m := newMacaron()
	// bindIgnErr := binding.BindIgnErr
	loginReq := middleware.Toggle(&middleware.ToggleOptions{SignInRequire: true})
	m.Get("/", loginReq, func(ctx *middleware.Context) {
		// ctx.Data["ctx"] = ctx
		ctx.HTML(200, "dashbord")
	})
	m.Group("/project/:id", func() {
		m.Get("", project.Detail)
		m.Group("/score", func() {
			m.Get("", score.Home)
			m.Post("", binding.BindIgnErr(form.ScoreForm{}), score.JudgeScore)
		}, middleware.ScoreMiddleWare())
	}, loginReq, middleware.ProjectMiddleware())

	m.Group("/user", func() {
		m.Get("/login", user.Login)
		m.Post("/login", binding.BindIgnErr(form.LoginForm{}), user.LoginPost)
		m.Get("/logout", user.Logout)
		// r.Post("/login", user.LoginPost)
	})
	adminReq := middleware.Toggle(&middleware.ToggleOptions{SignInRequire: true, AdminRequire: true})
	m.Group("/admin", func() {
		m.Get("", func(ctx *middleware.Context) {
			// ctx.Data["ctx"] = ctx
			ctx.HTML(200, "admin/home")
		})
		m.Group("/project", func() {
			m.Get("", project.List)
			m.Get("/create", project.Create)
			// m.Post("/")
		})
	}, adminReq)

	m.Run()
}

func main() {
	if len(os.Args) > 1 {
		// fmt.Println(os.Args[1])
		if os.Args[1] == "run" {
			runWeb()
		} else if os.Args[1] == "orm" {
			orm.RunCommand()
		}
	}

}
