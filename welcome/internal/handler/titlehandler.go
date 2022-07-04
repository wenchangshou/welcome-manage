package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"welcome-manage/welcome/internal/logic"
	"welcome-manage/welcome/internal/svc"
	"welcome-manage/welcome/internal/types"
)

func TitleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetTitleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTitleLogic(r.Context(), svcCtx)
		resp, err := l.Title(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
