// Code generated by goctl. DO NOT EDIT.
package types

type CreateAccountRequest struct {
	Addr string `json:"address"`
}

type CreateAccountResponse struct {
	Empty bool `json:"empty"`
}

type SubBalanceRequest struct {
	Addr   string `json:"address"`
	Amount int64  `json:"amount"`
}

type SubBalanceResponse struct {
	Empty bool `json:"empty"`
}

type AddBalanceRequest struct {
	Addr   string `json:"address"`
	Amount int64  `json:"amount"`
}

type AddBalanceResponse struct {
	Empty bool `json:"empty"`
}

type GetBalanceRequest struct {
	Addr string `json:"address"`
}

type GetBalanceResponse struct {
	Amount int64 `json:"amount"`
}

type SuicideRequest struct {
	Addr string `json:"address"`
}

type SuicideResponse struct {
	Is_suicide bool `json:"is_suicide"`
}

type HasSuicidedRequest struct {
	Addr string `json:"address"`
}

type HasSuicidedResponse struct {
	Is_suicide bool `json:"is_suicide"`
}

type GetNonceRequest struct {
	Addr string `json:"address"`
}

type GetNonceResponse struct {
	Amount uint64 `json:"amount"`
}

type SetNonceRequest struct {
	Addr   string `json:"address"`
	Amount uint64 `json:"amount"`
}

type SetNonceResponse struct {
	Empty bool `json:"empty"`
}

type GetCodeHashRequest struct {
	Addr string `json:"address"`
}

type GetCodeHashResponse struct {
	Hash string `json:"hash"`
}

type GetCodeRequest struct {
	Addr string `json:"address"`
}

type GetCodeResponse struct {
	Code string `json:"code"`
}

type SetCodeRequest struct {
	Addr string `json:"address"`
	Code string `json:"code"`
}

type SetCodeResponse struct {
}

type GetCodeSizeRequest struct {
	Addr string `json:"address"`
}

type GetCodeSizeResponse struct {
	Size int32 `json:"size"`
}

type AddRefundRequest struct {
	Amount uint64 `json:"amount"`
}

type AddRefundResponse struct {
}

type SubRefundRequest struct {
	Amount uint64 `json:"amount"`
}

type SubRefundResponse struct {
}

type GetRefundRequest struct {
}

type GetRefundResponse struct {
	Amount uint64 `json:"amount"`
}

type GetTransientStateRequest struct {
	Addr string `json:"address"`
	Key  string `json:"key"`
}

type GetTransientStateResponse struct {
	Value string `json:"value"`
}

type SetTransientStateRequest struct {
	Addr  string `json:"address"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetTransientStateResponse struct {
}
