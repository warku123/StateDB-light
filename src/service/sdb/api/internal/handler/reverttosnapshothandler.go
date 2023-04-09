package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"statedbl/service/sdb/api/internal/logic"
	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
)

func RevertToSnapshotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RevertToSnapshotRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRevertToSnapshotLogic(r.Context(), svcCtx)
		resp, err := l.RevertToSnapshot(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
