package form

import (
	"github.com/Unknwon/macaron"
	"github.com/astaxie/beego/orm"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/session"
	"tech_oa/models"
)

type LoginForm struct {
	Id       string `binding:"Required"`
	Password string `binding:"Required"`
	user     models.User
}

func (loginform *LoginForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	o := orm.NewOrm()
	loginform.user.Id = loginform.Id
	err := o.Read(&loginform.user)
	if err != nil || !loginform.user.CheckPassword(loginform.Password) {
		return append(errs, binding.Error{
			FieldNames:     []string{"id"},
			Classification: "autherror",
			Message:        "用户名或密码可能不正确"})
	}
	return errs
}

func (loginform *LoginForm) Login(sess session.Store) {
	loginform.user.Login(sess)
}
