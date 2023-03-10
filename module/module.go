package module

import (
	app "github.com/tsxylhs/user-module"

	starter "github.com/tsxylhs/go-starter/starter"
)

var Role = starter.NewModule("role", "role", "/v1/role")
var User = starter.NewModule("user", "user", "v1/user")

func init() {
	app.Service.Register(Role, User)
}
