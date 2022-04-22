package handler

import (
	"net/http"

	"datac/internal/logic"
	"datac/internal/svc"
	"datac/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DatacHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDatacLogic(r.Context(), svcCtx)
		resp, err := l.Datac(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
