package handler

import (
	"datac/common/result"
	"net/http"

	"datac/internal/logic"
	"datac/internal/svc"
	"datac/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DatacHandler2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDatacHandler2Logic(r.Context(), svcCtx)
		resp, err := l.DatacHandler2(&req)
		result.Response(w, resp, err)
	}
}
