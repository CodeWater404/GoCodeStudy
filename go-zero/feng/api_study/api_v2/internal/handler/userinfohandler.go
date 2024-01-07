package handler

import (
	"go-zero/feng/common/response"
	"net/http"

	"go-zero/feng/api_study/api_v2/internal/logic"
	"go-zero/feng/api_study/api_v2/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)
	}
}
