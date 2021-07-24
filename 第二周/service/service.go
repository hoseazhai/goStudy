package service

import (
	"context"
	"err/dao"
	"err/global"
)

const (
	parentSpanGormKey = "opentracing:parent.span"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
