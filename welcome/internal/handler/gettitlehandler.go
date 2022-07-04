package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"welcome-manage/welcome/internal/logic"
	"welcome-manage/welcome/internal/svc"
	"welcome-manage/welcome/internal/types"
)

func GetTitleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NullRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetTitleLogic(r.Context(), svcCtx)
		resp, err := l.GetTitle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
