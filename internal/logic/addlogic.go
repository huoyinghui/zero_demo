package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	_, err := l.svcCtx.Model.Insert(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &add.AddResp{
		Ok: true,
	}, nil
	// 手动代码结束
	// return &types.AddResp{}, nil
}
