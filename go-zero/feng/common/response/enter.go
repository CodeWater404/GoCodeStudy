package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2024/1/7
  @desc: 封装响应
**/

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// success to return
		r := &Body{
			Code: 0,
			Msg:  "success",
			Data: resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	// fail to return
	errCode := uint32(10086)
	errMsg := "server error"
	httpx.WriteJson(w, http.StatusBadGateway, &Body{
		Code: errCode,
		Msg:  errMsg,
		Data: nil,
	})
}
