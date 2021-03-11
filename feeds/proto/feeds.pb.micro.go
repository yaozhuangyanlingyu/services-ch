// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/feeds.proto

package feeds

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Feeds service

func NewFeedsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Feeds service

type FeedsService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error)
	Entries(ctx context.Context, in *EntriesRequest, opts ...client.CallOption) (*EntriesResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type feedsService struct {
	c    client.Client
	name string
}

func NewFeedsService(name string, c client.Client) FeedsService {
	return &feedsService{
		c:    c,
		name: name,
	}
}

func (c *feedsService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "Feeds.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsService) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error) {
	req := c.c.NewRequest(c.name, "Feeds.Remove", in)
	out := new(RemoveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsService) Entries(ctx context.Context, in *EntriesRequest, opts ...client.CallOption) (*EntriesResponse, error) {
	req := c.c.NewRequest(c.name, "Feeds.Entries", in)
	out := new(EntriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Feeds.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Feeds service

type FeedsHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Remove(context.Context, *RemoveRequest, *RemoveResponse) error
	Entries(context.Context, *EntriesRequest, *EntriesResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterFeedsHandler(s server.Server, hdlr FeedsHandler, opts ...server.HandlerOption) error {
	type feeds interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error
		Entries(ctx context.Context, in *EntriesRequest, out *EntriesResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Feeds struct {
		feeds
	}
	h := &feedsHandler{hdlr}
	return s.Handle(s.NewHandler(&Feeds{h}, opts...))
}

type feedsHandler struct {
	FeedsHandler
}

func (h *feedsHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.FeedsHandler.Add(ctx, in, out)
}

func (h *feedsHandler) Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error {
	return h.FeedsHandler.Remove(ctx, in, out)
}

func (h *feedsHandler) Entries(ctx context.Context, in *EntriesRequest, out *EntriesResponse) error {
	return h.FeedsHandler.Entries(ctx, in, out)
}

func (h *feedsHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.FeedsHandler.List(ctx, in, out)
}
