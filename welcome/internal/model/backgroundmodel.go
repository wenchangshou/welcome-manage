package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BackgroundModel = (*customBackgroundModel)(nil)

type (
	// BackgroundModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBackgroundModel.
	BackgroundModel interface {
		backgroundModel
	}

	customBackgroundModel struct {
		*defaultBackgroundModel
	}
)

// NewBackgroundModel returns a model for the database table.
func NewBackgroundModel(conn sqlx.SqlConn, c cache.CacheConf) BackgroundModel {
	return &customBackgroundModel{
		defaultBackgroundModel: newBackgroundModel(conn, c),
	}
}
