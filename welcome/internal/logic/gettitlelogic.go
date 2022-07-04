package logic

import (
	"context"
	"errors"

	"welcome-manage/welcome/internal/svc"
	"welcome-manage/welcome/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTitleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTitleLogic {
	return &GetTitleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTitleLogic) GetTitle(req *types.NullRequest) (resp *types.GetTitleResponse, err error) {
	t, err := l.svcCtx.TitleModel.FindOne(context.TODO(), 1)
	if err != nil {
		return nil, errors.New("find one err:" + err.Error())
	}
	resp = &types.GetTitleResponse{
		Color:  t.Color.String,
		Family: t.Family.String,
		Title:  t.Title.String,
		Size:   int(t.Size),
	}
	return
}
