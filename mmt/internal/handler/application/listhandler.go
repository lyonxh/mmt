package application

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mmt/mmt/internal/logic/application"
	"mmt/mmt/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := application.NewListLogic(r.Context(), svcCtx)
		err := l.List()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
