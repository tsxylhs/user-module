package service

import (
	"context"

	"github.com/tsxylhs/user-module/module"
	"github.com/tsxylhs/user-module/module/model"

	code "github.com/tsxylhs/go-starter/domain"
	"github.com/tsxylhs/go-starter/starter"
)

type UserService struct {
	*starter.Module
}

var User = &UserService{
	Module: module.User,
}

func (service *UserService) Insert(ctx context.Context, form *model.User, result *code.Result) error {

	service.Db.Insert(&form)
	return nil

}
