// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/proto/user.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for UserSrv service

func NewUserSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserSrv service

type UserSrvService interface {
	GetUserById(ctx context.Context, in *SearchId, opts ...client.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *SearchParams, opts ...client.CallOption) (*Users, error)
	CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *SearchId, opts ...client.CallOption) (*AffectedCount, error)
	BeforeCreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ValidationErr, error)
	BeforeUpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ValidationErr, error)
	BeforeDeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*ValidationErr, error)
	AfterCreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*AfterFuncErr, error)
	AfterUpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*AfterFuncErr, error)
	AfterDeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*AfterFuncErr, error)
}

type userSrvService struct {
	c    client.Client
	name string
}

func NewUserSrvService(name string, c client.Client) UserSrvService {
	return &userSrvService{
		c:    c,
		name: name,
	}
}

func (c *userSrvService) GetUserById(ctx context.Context, in *SearchId, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserSrv.GetUserById", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) GetUsers(ctx context.Context, in *SearchParams, opts ...client.CallOption) (*Users, error) {
	req := c.c.NewRequest(c.name, "UserSrv.GetUsers", in)
	out := new(Users)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserSrv.CreateUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) UpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserSrv.UpdateUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) DeleteUser(ctx context.Context, in *SearchId, opts ...client.CallOption) (*AffectedCount, error) {
	req := c.c.NewRequest(c.name, "UserSrv.DeleteUser", in)
	out := new(AffectedCount)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) BeforeCreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ValidationErr, error) {
	req := c.c.NewRequest(c.name, "UserSrv.BeforeCreateUser", in)
	out := new(ValidationErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) BeforeUpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ValidationErr, error) {
	req := c.c.NewRequest(c.name, "UserSrv.BeforeUpdateUser", in)
	out := new(ValidationErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) BeforeDeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*ValidationErr, error) {
	req := c.c.NewRequest(c.name, "UserSrv.BeforeDeleteUser", in)
	out := new(ValidationErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) AfterCreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*AfterFuncErr, error) {
	req := c.c.NewRequest(c.name, "UserSrv.AfterCreateUser", in)
	out := new(AfterFuncErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) AfterUpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*AfterFuncErr, error) {
	req := c.c.NewRequest(c.name, "UserSrv.AfterUpdateUser", in)
	out := new(AfterFuncErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) AfterDeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*AfterFuncErr, error) {
	req := c.c.NewRequest(c.name, "UserSrv.AfterDeleteUser", in)
	out := new(AfterFuncErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserSrv service

type UserSrvHandler interface {
	GetUserById(context.Context, *SearchId, *User) error
	GetUsers(context.Context, *SearchParams, *Users) error
	CreateUser(context.Context, *User, *User) error
	UpdateUser(context.Context, *User, *User) error
	DeleteUser(context.Context, *SearchId, *AffectedCount) error
	BeforeCreateUser(context.Context, *User, *ValidationErr) error
	BeforeUpdateUser(context.Context, *User, *ValidationErr) error
	BeforeDeleteUser(context.Context, *User, *ValidationErr) error
	AfterCreateUser(context.Context, *User, *AfterFuncErr) error
	AfterUpdateUser(context.Context, *User, *AfterFuncErr) error
	AfterDeleteUser(context.Context, *User, *AfterFuncErr) error
}

func RegisterUserSrvHandler(s server.Server, hdlr UserSrvHandler, opts ...server.HandlerOption) error {
	type userSrv interface {
		GetUserById(ctx context.Context, in *SearchId, out *User) error
		GetUsers(ctx context.Context, in *SearchParams, out *Users) error
		CreateUser(ctx context.Context, in *User, out *User) error
		UpdateUser(ctx context.Context, in *User, out *User) error
		DeleteUser(ctx context.Context, in *SearchId, out *AffectedCount) error
		BeforeCreateUser(ctx context.Context, in *User, out *ValidationErr) error
		BeforeUpdateUser(ctx context.Context, in *User, out *ValidationErr) error
		BeforeDeleteUser(ctx context.Context, in *User, out *ValidationErr) error
		AfterCreateUser(ctx context.Context, in *User, out *AfterFuncErr) error
		AfterUpdateUser(ctx context.Context, in *User, out *AfterFuncErr) error
		AfterDeleteUser(ctx context.Context, in *User, out *AfterFuncErr) error
	}
	type UserSrv struct {
		userSrv
	}
	h := &userSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&UserSrv{h}, opts...))
}

type userSrvHandler struct {
	UserSrvHandler
}

func (h *userSrvHandler) GetUserById(ctx context.Context, in *SearchId, out *User) error {
	return h.UserSrvHandler.GetUserById(ctx, in, out)
}

func (h *userSrvHandler) GetUsers(ctx context.Context, in *SearchParams, out *Users) error {
	return h.UserSrvHandler.GetUsers(ctx, in, out)
}

func (h *userSrvHandler) CreateUser(ctx context.Context, in *User, out *User) error {
	return h.UserSrvHandler.CreateUser(ctx, in, out)
}

func (h *userSrvHandler) UpdateUser(ctx context.Context, in *User, out *User) error {
	return h.UserSrvHandler.UpdateUser(ctx, in, out)
}

func (h *userSrvHandler) DeleteUser(ctx context.Context, in *SearchId, out *AffectedCount) error {
	return h.UserSrvHandler.DeleteUser(ctx, in, out)
}

func (h *userSrvHandler) BeforeCreateUser(ctx context.Context, in *User, out *ValidationErr) error {
	return h.UserSrvHandler.BeforeCreateUser(ctx, in, out)
}

func (h *userSrvHandler) BeforeUpdateUser(ctx context.Context, in *User, out *ValidationErr) error {
	return h.UserSrvHandler.BeforeUpdateUser(ctx, in, out)
}

func (h *userSrvHandler) BeforeDeleteUser(ctx context.Context, in *User, out *ValidationErr) error {
	return h.UserSrvHandler.BeforeDeleteUser(ctx, in, out)
}

func (h *userSrvHandler) AfterCreateUser(ctx context.Context, in *User, out *AfterFuncErr) error {
	return h.UserSrvHandler.AfterCreateUser(ctx, in, out)
}

func (h *userSrvHandler) AfterUpdateUser(ctx context.Context, in *User, out *AfterFuncErr) error {
	return h.UserSrvHandler.AfterUpdateUser(ctx, in, out)
}

func (h *userSrvHandler) AfterDeleteUser(ctx context.Context, in *User, out *AfterFuncErr) error {
	return h.UserSrvHandler.AfterDeleteUser(ctx, in, out)
}