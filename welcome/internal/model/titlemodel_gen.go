// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	titleFieldNames          = builder.RawFieldNames(&Title{})
	titleRows                = strings.Join(titleFieldNames, ",")
	titleRowsExpectAutoSet   = strings.Join(stringx.Remove(titleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	titleRowsWithPlaceHolder = strings.Join(stringx.Remove(titleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTitleIdPrefix = "cache:title:id:"
)

type (
	titleModel interface {
		Insert(ctx context.Context, data *Title) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Title, error)
		Update(ctx context.Context, newData *Title) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTitleModel struct {
		sqlc.CachedConn
		table string
	}

	Title struct {
		Id     int64          `db:"id"`
		Title  sql.NullString `db:"title"`
		Size   int64          `db:"size"`
		Family sql.NullString `db:"family"`
		Color  sql.NullString `db:"color"`
	}
)

func newTitleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTitleModel {
	return &defaultTitleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`title`",
	}
}

func (m *defaultTitleModel) Delete(ctx context.Context, id int64) error {
	titleIdKey := fmt.Sprintf("%s%v", cacheTitleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, titleIdKey)
	return err
}

func (m *defaultTitleModel) FindOne(ctx context.Context, id int64) (*Title, error) {
	titleIdKey := fmt.Sprintf("%s%v", cacheTitleIdPrefix, id)
	var resp Title
	err := m.QueryRowCtx(ctx, &resp, titleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", titleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTitleModel) Insert(ctx context.Context, data *Title) (sql.Result, error) {
	titleIdKey := fmt.Sprintf("%s%v", cacheTitleIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, titleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.Size, data.Family, data.Color)
	}, titleIdKey)
	return ret, err
}

func (m *defaultTitleModel) Update(ctx context.Context, data *Title) error {
	titleIdKey := fmt.Sprintf("%s%v", cacheTitleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, titleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Size, data.Family, data.Color, data.Id)
	}, titleIdKey)
	return err
}

func (m *defaultTitleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTitleIdPrefix, primary)
}

func (m *defaultTitleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", titleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTitleModel) tableName() string {
	return m.table
}
