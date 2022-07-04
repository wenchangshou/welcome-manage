package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"welcome-manage/welcome/internal/logic"
	"welcome-manage/welcome/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadLogic(r, r.Context(), svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
