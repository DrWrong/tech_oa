package middleware

import (
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego/orm"
	"github.com/macaron-contrib/session"
	"models"
)

type Context struct {
	*macaron.Context
	User *models.User
}

func (ctx *Context) GetUserBySession(sess session.Store) {
	uid = sess.Get("uid").(string)
	if uid != nil {
		o := orm.NewOrm()
		user := User{Id: uid}
		err = o.Read(&user)
		if err != nil {
			User = &user
		}
	}
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context, sess session.Store) {
		ctx := &Context{
			Context: c,
		}
		ctx.GetUserBySession(sess)
		c.Map(ctx)
	}
}
