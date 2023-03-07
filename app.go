package app

import (
	starter "github.com/tsxylhs/go-starter/starter"
)

var Service = starter.NewService("user-module-service", true, false)
var Api = starter.NewWeb("user-nodule-api")
