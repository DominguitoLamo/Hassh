package utils

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
)

func JsonSuccessfulWrite[T any](data T) *xhttp.BaseResponse[T] {
	res := xhttp.BaseResponse[T]{
		Code: 200,
		Data: data,
	}
	return &res
}

func JsonErrorWrite(w http.ResponseWriter, errorCode int, msg string) {
	// res := xhttp.BaseResponse[types.Nil]{
	// 	Code: errorCode,
	// 	Msg: msg,
	// }
	// return &res
	httpx.Error(w, errors.New(400, "参数错误"))	
}