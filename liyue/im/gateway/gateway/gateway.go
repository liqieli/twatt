// Code generated by goctl. DO NOT EDIT!
// Source: gateway.proto

package gateway

import (
	"context"

	"github.com/tangbo/twatt/liyue/im/gateway/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GreetReq  = pb.GreetReq
	GreetResp = pb.GreetResp

	Gateway interface {
		Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetResp, error)
	}

	defaultGateway struct {
		cli zrpc.Client
	}
)

func NewGateway(cli zrpc.Client) Gateway {
	return &defaultGateway{
		cli: cli,
	}
}

func (m *defaultGateway) Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetResp, error) {
	client := pb.NewGatewayClient(m.cli.Conn())
	return client.Greet(ctx, in, opts...)
}
