// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package middleware

import (
	"github.com/Unknwon/macaron"
	// "github.com/macaron-contrib/csrf"
	"net/url"
)

type ToggleOptions struct {
	SignInRequire bool

	AdminRequire bool
	DisableCsrf  bool
}

func Toggle(options *ToggleOptions) macaron.Handler {
	return func(ctx *Context) {

		if options.SignInRequire {
			if !ctx.IsSigned {
				// Ignore watch repository operation.
				ctx.SetCookie("redirect_to", "/"+url.QueryEscape(ctx.Req.RequestURI), 0)
				ctx.Redirect("/user/login")
				return
			} else if !ctx.User.IsActive {
				ctx.Data["Title"] = ctx.Tr("auth.active_your_account")
				ctx.HTML(200, "user/auth/activate")
				return
			}
		}

		if options.AdminRequire {
			if !ctx.User.IsAdmin {
				ctx.Error(403, "request forbidden")
				return
			}
			ctx.Data["PageIsAdmin"] = true
		}
	}
}
