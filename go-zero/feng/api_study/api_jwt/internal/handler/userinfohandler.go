package handler

import (
	"net/http"

	"go-zero/feng/api_study/api_jwt/internal/logic"
	"go-zero/feng/api_study/api_jwt/internal/svc"
	"go-zero/feng/common/response"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		// custom template generate
		response.Response(r, w, resp, err)
	}
}
