package handler

import (
	"context"

	"github.com/tsxylhs/user-module/module/model"
	"github.com/tsxylhs/user-module/service"

	"github.com/gin-gonic/gin"
	"github.com/tsxylhs/go-starter/log"
	"github.com/tsxylhs/go-starter/web"
	"go.uber.org/zap"
)

type UserApi struct {
	*web.RestHandler
}

func NewUserApi() *UserApi {
	return &UserApi{
		RestHandler: web.DefaultRestHandler,
	}
}

func (api *UserApi) Insert(c *gin.Context) {
	var form = model.User{}
	err := api.Bind(c, &form)
	if err != nil {
		log.Logger.Debug("fail to bind login params", zap.Error(err))
		api.BadRequestWithError(c, err)
		return
	}
	service.User.Insert(context.Background(), &form, nil)
	api.Success(c)
	return
}
func (api *UserApi) Register(router gin.IRouter) {
	v1 := router.Group("/v1")
	v1.POST("/insert", api.Insert)
}
