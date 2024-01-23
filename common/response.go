package common

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	OK  = 200
	Err = 400
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResResult(w http.ResponseWriter, resp interface{}, err error) {
	body := Body{}
	if err != nil {
		body.Code = Err
		body.Msg = errors.Wrap(err, "Err").Error()
	} else {
		body.Code = OK
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
