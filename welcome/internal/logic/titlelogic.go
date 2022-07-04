package logic

import (
	"context"
	"database/sql"

	"welcome-manage/welcome/internal/svc"
	"welcome-manage/welcome/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TitleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TitleLogic {
	return &TitleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TitleLogic) Title(req *types.SetTitleRequest) (resp *types.Response, err error) {
	d, err := l.svcCtx.TitleModel.FindOne(context.TODO(), 1)
	if err != nil {
		return nil, err
	}

	(*d).Family = sql.NullString{String: req.Title}
	return
}
