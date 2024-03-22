package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"mmt/utils"
	"mmt/mmt/internal/logic/user"
	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		common.ResResult(w, resp, err)
	}
}
