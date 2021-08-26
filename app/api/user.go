package api

import (
	"gf-init/app/service"
	"gf-init/library/response"

	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*userApi) Index(r *ghttp.Request) {
	if err := service.User.GetAllUsers(); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}
