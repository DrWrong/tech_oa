package middleware

import (
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego/orm"
	"github.com/macaron-contrib/session"
	"tech_oa/models"
)

type Context struct {
	*macaron.Context
	User     *models.User
	IsSigned bool
	Project  *models.Project
}

func (ctx *Context) GetUserBySession(sess session.Store) {
	// sess.Set("uid", "123")

	if uid, ok := sess.Get("uid").(string); ok {
		o := orm.NewOrm()
		user := models.User{Id: uid}
		err := o.Read(&user)
		ctx.User = &user
		if err == nil {
			ctx.IsSigned = true
		}
	}
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context, sess session.Store) {
		ctx := &Context{
			Context:  c,
			IsSigned: false,
		}

		ctx.GetUserBySession(sess)
		ctx.Data["ctx"] = ctx
		c.Map(ctx)
	}
}
