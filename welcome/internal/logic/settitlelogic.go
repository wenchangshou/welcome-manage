package logic

import (
	"context"
	"database/sql"

	"welcome-manage/welcome/internal/model"
	"welcome-manage/welcome/internal/svc"
	"welcome-manage/welcome/internal/types"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTitleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTitleLogic {
	return &SetTitleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetTitleLogic) SetTitle(req *types.SetTitleRequest) (resp *types.Response, err error) {
	t, err := l.svcCtx.TitleModel.FindOne(context.TODO(), 1)
	if t == nil || err == sqlx.ErrNotFound {
		t = &model.Title{
			Title:  sql.NullString{String: ""},
			Size:   300,
			Family: sql.NullString{String: "微软雅黑"},
			Color:  sql.NullString{String: "white"},
		}
		l.svcCtx.TitleModel.Insert(context.TODO(), t)
	}
	t.Size = int64(req.Size)
	t.Family = sql.NullString{String: req.Family}
	t.Title = sql.NullString{String: req.Title}
	t.Color = sql.NullString{String: req.Color}
	err = l.svcCtx.TitleModel.Update(context.TODO(), t)
	return
}
