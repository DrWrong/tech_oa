package user

import (
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/session"
	"tech_oa/form"
	"tech_oa/middleware"
)

func Login(ctx *middleware.Context) {
	if ctx.IsSigned {
		ctx.Redirect("/")
		return
	}
	ctx.HTML(200, "user/login")
}

func LoginPost(loginform form.LoginForm, ctx *middleware.Context, sess session.Store, errs binding.Errors) {
	if ctx.IsSigned {
		ctx.Redirect("/")
		return
	}
	if errs.Len() > 0 {
		ctx.Data["errors"] = errs
		ctx.HTML(200, "user/login")
		return
	}
	loginform.Login(sess)
	ctx.Redirect("/")
	return
}

func Logout(ctx *middleware.Context, sess session.Store) {
	if ctx.IsSigned {
		ctx.User.Logout(sess)
	}
	ctx.Redirect("/")
}
