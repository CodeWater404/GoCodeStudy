package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/feng/rpc_study/user_api_rpc/api/internal/logic"
	"go-zero/feng/rpc_study/user_api_rpc/api/internal/svc"
	"go-zero/feng/rpc_study/user_api_rpc/api/internal/types"

	"go-zero/feng/common/response"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		// custom template generate
		response.Response(r, w, resp, err)
	}
}
