// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"statedbl/service/sdb/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/createaccount",
				Handler: CreateAccountHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/subbalance",
				Handler: SubBalanceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/addbalance",
				Handler: AddBalanceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/getbalance",
				Handler: GetBalanceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/suicide",
				Handler: SuicideHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/getnonce",
				Handler: GetNonceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/setnonce",
				Handler: SetNonceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/getcodehash",
				Handler: GetCodeHashHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/getcode",
				Handler: GetCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/setcode",
				Handler: SetCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/getcodesize",
				Handler: GetCodeSizeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/addrefund",
				Handler: AddRefundHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/subrefund",
				Handler: SubRefundHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sdb/getrefund",
				Handler: GetRefundHandler(serverCtx),
			},
		},
	)
}
