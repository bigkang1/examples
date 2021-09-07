// Code generated by proto-gen-vine. DO NOT EDIT.
// source: pb/hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	client "github.com/vine-io/vine/core/client"
	server "github.com/vine-io/vine/core/server"
	api "github.com/vine-io/vine/lib/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// API Endpoints for Hello service
func NewHelloEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Hello service
type HelloService interface {
	Echo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) Echo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Hello.Echo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service
type HelloHandler interface {
	Echo(context.Context, *Request, *Response) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) error {
	type helloImpl interface {
		Echo(ctx context.Context, in *Request, out *Response) error
	}
	type Hello struct {
		helloImpl
	}
	h := &helloHandler{hdlr}
	return s.Handle(s.NewHandler(&Hello{h}, opts...))
}

type helloHandler struct {
	HelloHandler
}

func (h *helloHandler) Echo(ctx context.Context, in *Request, out *Response) error {
	return h.HelloHandler.Echo(ctx, in, out)
}
