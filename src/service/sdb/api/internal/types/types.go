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
	Amount string `json:"amount"`
}

type SubBalanceResponse struct {
	Empty bool `json:"empty"`
}

type AddBalanceRequest struct {
	Addr   string `json:"address"`
	Amount string `json:"amount"`
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

type EmptyRequest struct {
	Addr string `json:"address"`
}

type EmptyResponse struct {
	Is_empty bool `json:"is_empty"`
}

type ExistRequest struct {
	Addr string `json:"address"`
}

type ExistResponse struct {
	Is_exist bool `json:"is_exist"`
}

type AddressInAccessListRequest struct {
	Addr string `json:"address"`
}

type AddressInAccessListResponse struct {
	Is_in bool `json:"is_in"`
}

type SlotInAccessListRequest struct {
	Addr string `json:"address"`
	Hash string `json:"hash"`
}

type SlotInAccessListResponse struct {
	Addr_ok bool `json:"addr_ok"`
	Slot_ok bool `json:"slot_ok"`
}

type AddAddressToAccessListRequest struct {
	Addr string `json:"address"`
}

type AddAddressToAccessListResponse struct {
}

type AddSlotToAccessListRequest struct {
	Addr string `json:"address"`
	Slot string `json:"slot"`
}

type AddSlotToAccessListResponse struct {
}

type RevertToSnapshotRequest struct {
	Revid int32 `json:"revid"`
}

type RevertToSnapshotResponse struct {
}

type SnapshotRequest struct {
}

type SnapshotResponse struct {
	Revid int32 `json:"revid"`
}

type AddPreimageRequest struct {
	Hash     string `json:"hash"`
	Preimage string `json:"preimage"`
}

type AddPreimageResponse struct {
}

type AddLogRequest struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber uint64   `json:"blockNumber"`
	TxHash      string   `json:"transactionHash"`
	TxIndex     uint     `json:"transactionIndex"`
	BlockHash   string   `json:"blockHash"`
	Index       uint32   `json:"logIndex"`
	Removed     bool     `json:"removed"`
}

type AddLogResponse struct {
}

type Rules struct {
	ChainID          string `json:"chainid"`
	IsHomestead      bool   `json:"ishomestead"`
	IsEIP150         bool   `json:"iseip150"`
	IsEIP155         bool   `json:"iseip155"`
	IsEIP158         bool   `json:"iseip158"`
	IsByzantium      bool   `json:"isbyzantium"`
	IsConstantinople bool   `json:"isconstantinople"`
	IsPetersburg     bool   `json:"ispetersburg"`
	IsIstanbul       bool   `json:"isistanbul"`
	IsBerlin         bool   `json:"isberlin"`
	IsLondon         bool   `json:"islondon"`
	IsMerge          bool   `json:"ismerge"`
	IsShanghai       bool   `json:"isshanghai"`
	IsCancun         bool   `json:"iscancun"`
	IsPrague         bool   `json:"isprague"`
}

type AccessTuple struct {
	Addr        string   `json:"address"`
	StorageKeys []string `json:"storagekeys"`
}

type PrepareRequest struct {
	Rule         Rules         `json:"rules"`
	SenderAddr   string        `json:"senderaddr"`
	CoinbaseAddr string        `json:"coinbaseaddr"`
	DestAddr     string        `json:"destaddr"`
	PreCompiles  []string      `json:"precompiles"`
	List         []AccessTuple `json:"list"`
}

type PrepareResponse struct {
}
