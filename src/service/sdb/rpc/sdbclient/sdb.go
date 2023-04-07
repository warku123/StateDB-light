// Code generated by goctl. DO NOT EDIT.
// Source: sdb.proto

package sdbclient

import (
	"context"

	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddBalanceRequest         = sdb.AddBalanceRequest
	AddBalanceResponse        = sdb.AddBalanceResponse
	AddRefundRequest          = sdb.AddRefundRequest
	AddRefundResponse         = sdb.AddRefundResponse
	CreateAccountRequest      = sdb.CreateAccountRequest
	CreateAccountResponse     = sdb.CreateAccountResponse
	GetBalanceRequest         = sdb.GetBalanceRequest
	GetBalanceResponse        = sdb.GetBalanceResponse
	GetCodeHashRequest        = sdb.GetCodeHashRequest
	GetCodeHashResponse       = sdb.GetCodeHashResponse
	GetCodeRequest            = sdb.GetCodeRequest
	GetCodeResponse           = sdb.GetCodeResponse
	GetCodeSizeRequest        = sdb.GetCodeSizeRequest
	GetCodeSizeResponse       = sdb.GetCodeSizeResponse
	GetNonceRequest           = sdb.GetNonceRequest
	GetNonceResponse          = sdb.GetNonceResponse
	GetRefundRequest          = sdb.GetRefundRequest
	GetRefundResponse         = sdb.GetRefundResponse
	GetTransientStateRequest  = sdb.GetTransientStateRequest
	GetTransientStateResponse = sdb.GetTransientStateResponse
	HasSuicidedRequest        = sdb.HasSuicidedRequest
	HasSuicidedResponse       = sdb.HasSuicidedResponse
	SetCodeRequest            = sdb.SetCodeRequest
	SetCodeResponse           = sdb.SetCodeResponse
	SetNonceRequest           = sdb.SetNonceRequest
	SetNonceResponse          = sdb.SetNonceResponse
	SetTransientStateRequest  = sdb.SetTransientStateRequest
	SetTransientStateResponse = sdb.SetTransientStateResponse
	SubBalanceRequest         = sdb.SubBalanceRequest
	SubBalanceResponse        = sdb.SubBalanceResponse
	SubRefundRequest          = sdb.SubRefundRequest
	SubRefundResponse         = sdb.SubRefundResponse
	SuicideRequest            = sdb.SuicideRequest
	SuicideResponse           = sdb.SuicideResponse

	Sdb interface {
		CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
		SubBalance(ctx context.Context, in *SubBalanceRequest, opts ...grpc.CallOption) (*SubBalanceResponse, error)
		AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error)
		GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
		Suicide(ctx context.Context, in *SuicideRequest, opts ...grpc.CallOption) (*SuicideResponse, error)
		HasSuicided(ctx context.Context, in *HasSuicidedRequest, opts ...grpc.CallOption) (*HasSuicidedResponse, error)
		GetNonce(ctx context.Context, in *GetNonceRequest, opts ...grpc.CallOption) (*GetNonceResponse, error)
		SetNonce(ctx context.Context, in *SetNonceRequest, opts ...grpc.CallOption) (*SetNonceResponse, error)
		GetCodeHash(ctx context.Context, in *GetCodeHashRequest, opts ...grpc.CallOption) (*GetCodeHashResponse, error)
		GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*GetCodeResponse, error)
		SetCode(ctx context.Context, in *SetCodeRequest, opts ...grpc.CallOption) (*SetCodeResponse, error)
		GetCodeSize(ctx context.Context, in *GetCodeSizeRequest, opts ...grpc.CallOption) (*GetCodeSizeResponse, error)
		AddRefund(ctx context.Context, in *AddRefundRequest, opts ...grpc.CallOption) (*AddRefundResponse, error)
		SubRefund(ctx context.Context, in *SubRefundRequest, opts ...grpc.CallOption) (*SubRefundResponse, error)
		GetRefund(ctx context.Context, in *GetRefundRequest, opts ...grpc.CallOption) (*GetRefundResponse, error)
		GetTransientState(ctx context.Context, in *GetTransientStateRequest, opts ...grpc.CallOption) (*GetTransientStateResponse, error)
		SetTransientState(ctx context.Context, in *SetTransientStateRequest, opts ...grpc.CallOption) (*SetTransientStateResponse, error)
	}

	defaultSdb struct {
		cli zrpc.Client
	}
)

func NewSdb(cli zrpc.Client) Sdb {
	return &defaultSdb{
		cli: cli,
	}
}

func (m *defaultSdb) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.CreateAccount(ctx, in, opts...)
}

func (m *defaultSdb) SubBalance(ctx context.Context, in *SubBalanceRequest, opts ...grpc.CallOption) (*SubBalanceResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.SubBalance(ctx, in, opts...)
}

func (m *defaultSdb) AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.AddBalance(ctx, in, opts...)
}

func (m *defaultSdb) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetBalance(ctx, in, opts...)
}

func (m *defaultSdb) Suicide(ctx context.Context, in *SuicideRequest, opts ...grpc.CallOption) (*SuicideResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.Suicide(ctx, in, opts...)
}

func (m *defaultSdb) HasSuicided(ctx context.Context, in *HasSuicidedRequest, opts ...grpc.CallOption) (*HasSuicidedResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.HasSuicided(ctx, in, opts...)
}

func (m *defaultSdb) GetNonce(ctx context.Context, in *GetNonceRequest, opts ...grpc.CallOption) (*GetNonceResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetNonce(ctx, in, opts...)
}

func (m *defaultSdb) SetNonce(ctx context.Context, in *SetNonceRequest, opts ...grpc.CallOption) (*SetNonceResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.SetNonce(ctx, in, opts...)
}

func (m *defaultSdb) GetCodeHash(ctx context.Context, in *GetCodeHashRequest, opts ...grpc.CallOption) (*GetCodeHashResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetCodeHash(ctx, in, opts...)
}

func (m *defaultSdb) GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*GetCodeResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetCode(ctx, in, opts...)
}

func (m *defaultSdb) SetCode(ctx context.Context, in *SetCodeRequest, opts ...grpc.CallOption) (*SetCodeResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.SetCode(ctx, in, opts...)
}

func (m *defaultSdb) GetCodeSize(ctx context.Context, in *GetCodeSizeRequest, opts ...grpc.CallOption) (*GetCodeSizeResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetCodeSize(ctx, in, opts...)
}

func (m *defaultSdb) AddRefund(ctx context.Context, in *AddRefundRequest, opts ...grpc.CallOption) (*AddRefundResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.AddRefund(ctx, in, opts...)
}

func (m *defaultSdb) SubRefund(ctx context.Context, in *SubRefundRequest, opts ...grpc.CallOption) (*SubRefundResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.SubRefund(ctx, in, opts...)
}

func (m *defaultSdb) GetRefund(ctx context.Context, in *GetRefundRequest, opts ...grpc.CallOption) (*GetRefundResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetRefund(ctx, in, opts...)
}

func (m *defaultSdb) GetTransientState(ctx context.Context, in *GetTransientStateRequest, opts ...grpc.CallOption) (*GetTransientStateResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.GetTransientState(ctx, in, opts...)
}

func (m *defaultSdb) SetTransientState(ctx context.Context, in *SetTransientStateRequest, opts ...grpc.CallOption) (*SetTransientStateResponse, error) {
	client := sdb.NewSdbClient(m.cli.Conn())
	return client.SetTransientState(ctx, in, opts...)
}
