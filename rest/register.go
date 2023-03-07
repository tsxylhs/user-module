package rest

import (
	"user_module/rest/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAPIs(router gin.IRouter) {
	handler.NewUserApi().Register(router)
}
