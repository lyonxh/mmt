package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mmt/mmt/internal/logic/user"
	"mmt/mmt/internal/svc"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
