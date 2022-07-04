package logic

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"welcome-manage/welcome/internal/svc"
	"welcome-manage/welcome/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 30 // 10 MB

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.r.ParseMultipartForm(maxFileSize)
	file, handler, err := l.r.FormFile("myfile")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)
	tempFile, err := os.Create(path.Join(l.svcCtx.Config.Path, handler.Filename))
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	return &types.Response{Code: 0, Msg: "成功"}, nil
}
