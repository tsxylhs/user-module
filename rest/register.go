package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/tsxylhs/user_module/rest/handler"
)

func RegisterAPIs(router gin.IRouter) {
	handler.NewUserApi().Register(router)
}
