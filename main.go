package main

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/macaron-contrib/session"
	_ "github.com/macaron-contrib/session/mysql"
	"middleware"
	"os"
	_ "tech_oa/models"
	_ "tech_oa/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	fmt.Println(beego.AppConfig.String("mysqlstring"))
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
		Provider: "mysql",
		Config: session.Config{
			Gclifetime:     3600,
			ProviderConfig: beego.AppConfig.String("mysqlstring"),
		},
	}))
	m.Use(middleware.Contexter())
	return m

}

func runWeb() {
	// m := newMacaron()
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		if os.Args[1] == "run" {
			runWeb()
		} else if os.Args[1] == "orm" {
			orm.RunCommand()
		}
	}

}
