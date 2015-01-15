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
}

func (ctx *Context) GetUserBySession(sess session.Store) {
	// sess.Set("uid", "123")

	if uid, ok := sess.Get("uid").(string); ok {
		o := orm.NewOrm()
		user := models.User{Id: uid}
		err := o.Read(&user)
		if err == nil {
			ctx.User = &user
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
		c.Map(ctx)
	}
}
