package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/feng/api_study/api_swagger/internal/logic"
	"go-zero/feng/api_study/api_swagger/internal/svc"
	"go-zero/feng/api_study/api_swagger/internal/types"

	"go-zero/feng/common/response"
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
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		// custom template generate
		response.Response(r, w, resp, err)
	}
}
