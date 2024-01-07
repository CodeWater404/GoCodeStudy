package handler

import (
	"go-zero/feng/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/feng/api_study/api_v2/internal/logic"
	"go-zero/feng/api_study/api_v2/internal/svc"
	"go-zero/feng/api_study/api_v2/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)
	}
}
