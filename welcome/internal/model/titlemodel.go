package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TitleModel = (*customTitleModel)(nil)

type (
	// TitleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTitleModel.
	TitleModel interface {
		titleModel
	}

	customTitleModel struct {
		*defaultTitleModel
	}
)

// NewTitleModel returns a model for the database table.
func NewTitleModel(conn sqlx.SqlConn, c cache.CacheConf) TitleModel {
	return &customTitleModel{
		defaultTitleModel: newTitleModel(conn, c),
	}
}
