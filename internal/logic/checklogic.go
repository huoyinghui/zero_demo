package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	resp, err := l.svcCtx.Model.FindOne(in.Book)
	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
	// 手动代码结束
	return &types.CheckResp{
		Found: true,
		Price: 14,
	}, nil
}
