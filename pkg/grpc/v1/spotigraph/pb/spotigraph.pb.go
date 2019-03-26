// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/grpc/v1/spotigraph/pb/spotigraph.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	spotigraph "go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("pkg/grpc/v1/spotigraph/pb/spotigraph.proto", fileDescriptor_039c84dc8946bc09)
}

var fileDescriptor_039c84dc8946bc09 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0xc1, 0xf7, 0x5e, 0x1e, 0xcc, 0x72, 0x36, 0x0f, 0x7d, 0x8a, 0xa5, 0xcb, 0x16, 0x0c,
	0xda, 0x6d, 0xff, 0x51, 0xad, 0xa1, 0x42, 0xa1, 0x68, 0xdd, 0x74, 0x17, 0xf5, 0x12, 0x43, 0x53,
	0x33, 0x8e, 0xa3, 0xd0, 0x8f, 0xda, 0x75, 0x3f, 0x44, 0xb7, 0x65, 0x6e, 0xe2, 0x24, 0x0d, 0xf3,
	0x47, 0x5c, 0xb6, 0xe7, 0xdc, 0xc3, 0x99, 0x9f, 0x07, 0x42, 0xce, 0xd8, 0x6b, 0xe4, 0x47, 0x9c,
	0xcd, 0xfd, 0x5d, 0xd7, 0xdf, 0xb0, 0x54, 0xc4, 0x11, 0x0f, 0xd9, 0xd2, 0x67, 0xb3, 0xd2, 0x5f,
	0x1d, 0xc6, 0x53, 0x91, 0x52, 0x2a, 0x7d, 0x9d, 0xd2, 0xbf, 0x77, 0xdd, 0xc6, 0xb9, 0xbc, 0x47,
	0x79, 0x9e, 0x26, 0x95, 0x8c, 0x6a, 0x40, 0xef, 0xa3, 0x46, 0x7e, 0x4f, 0x37, 0xc0, 0xe9, 0x80,
	0x78, 0x7d, 0x0e, 0xa1, 0x00, 0xda, 0xfc, 0x99, 0xd7, 0x91, 0x72, 0x26, 0x8d, 0x61, 0xdd, 0xa8,
	0xaa, 0x93, 0x78, 0x15, 0x25, 0x20, 0x3d, 0x63, 0xd8, 0xd0, 0x6b, 0xf2, 0x2b, 0x00, 0x41, 0xeb,
	0x9a, 0x88, 0x00, 0x84, 0xfb, 0x7e, 0x40, 0xbc, 0x29, 0x5b, 0x98, 0x5a, 0x64, 0x92, 0x3b, 0xe5,
	0x92, 0x78, 0x03, 0x48, 0x40, 0x80, 0xad, 0xc8, 0xbf, 0x8a, 0x74, 0xff, 0xc6, 0xc4, 0xbb, 0xbc,
	0x0e, 0x88, 0x37, 0x81, 0x90, 0xcf, 0x97, 0xda, 0x0e, 0x99, 0x24, 0x03, 0xda, 0x15, 0xf5, 0x29,
	0x8c, 0xe2, 0x55, 0x28, 0x60, 0x91, 0xd7, 0xe8, 0x7d, 0xd6, 0xc8, 0x9f, 0xc9, 0x7a, 0x1b, 0x2e,
	0xe8, 0x50, 0xc1, 0x6d, 0x55, 0x8b, 0x4b, 0xbd, 0xa0, 0xdb, 0xd2, 0xbe, 0x0b, 0x4d, 0xb2, 0xda,
	0x6d, 0x86, 0xb7, 0xa1, 0x0b, 0xc9, 0x9f, 0xe5, 0x48, 0x18, 0x2a, 0xc0, 0xda, 0x26, 0x05, 0x61,
	0x47, 0xce, 0x95, 0x42, 0x6c, 0x2b, 0x63, 0x64, 0xfc, 0xa0, 0x18, 0x6b, 0x6b, 0x14, 0x90, 0x4f,
	0x4c, 0x90, 0xf7, 0x4d, 0x7a, 0x5f, 0x35, 0xf2, 0xb7, 0xbf, 0x0c, 0x99, 0x00, 0x4e, 0x47, 0x8a,
	0x73, 0xf5, 0xc7, 0xc9, 0x1d, 0x05, 0xe9, 0xb6, 0xf6, 0x7d, 0xb9, 0x2d, 0x9b, 0x22, 0xb2, 0x6e,
	0xea, 0x83, 0xf2, 0x07, 0x3a, 0x53, 0x46, 0x8a, 0xb7, 0xa1, 0x51, 0x41, 0xdc, 0x99, 0x75, 0xa3,
	0x98, 0xdb, 0x4b, 0x19, 0xa9, 0x3f, 0x2a, 0xea, 0x86, 0x32, 0x05, 0xf7, 0x53, 0x13, 0xf7, 0xa2,
	0x0f, 0xee, 0x3b, 0xd8, 0xc6, 0x89, 0x6d, 0xdf, 0xa8, 0xbb, 0xf6, 0x8d, 0x26, 0xcb, 0xbe, 0x51,
	0xb7, 0xee, 0x5b, 0x25, 0x98, 0xf7, 0x8d, 0x16, 0xd7, 0xbe, 0x55, 0x8e, 0x79, 0xdf, 0xe5, 0x32,
	0x47, 0xec, 0x1b, 0xcf, 0x0f, 0xd8, 0xf7, 0xbe, 0x09, 0x52, 0x7e, 0xe6, 0xf1, 0x0c, 0x2c, 0x94,
	0x51, 0x77, 0x51, 0x46, 0x93, 0x85, 0x32, 0xea, 0x56, 0xca, 0x2a, 0xc1, 0x4c, 0x19, 0x2d, 0x2e,
	0xca, 0x2a, 0xc7, 0x4c, 0xb9, 0x5c, 0xe6, 0x08, 0xca, 0x78, 0x7e, 0x00, 0xe5, 0x7d, 0x93, 0xbb,
	0xff, 0x2f, 0x75, 0xe3, 0x67, 0x77, 0xe6, 0xe1, 0xb7, 0xf2, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0x1c, 0x5b, 0x60, 0x9a, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Create(ctx context.Context, in *spotigraph.UserCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleUserRes, error)
	Get(ctx context.Context, in *spotigraph.UserGetReq, opts ...grpc.CallOption) (*spotigraph.SingleUserRes, error)
	Update(ctx context.Context, in *spotigraph.UserUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleUserRes, error)
	Delete(ctx context.Context, in *spotigraph.UserGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error)
	Search(ctx context.Context, in *spotigraph.UserSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedUserRes, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Create(ctx context.Context, in *spotigraph.UserCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleUserRes, error) {
	out := new(spotigraph.SingleUserRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.User/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Get(ctx context.Context, in *spotigraph.UserGetReq, opts ...grpc.CallOption) (*spotigraph.SingleUserRes, error) {
	out := new(spotigraph.SingleUserRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.User/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *spotigraph.UserUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleUserRes, error) {
	out := new(spotigraph.SingleUserRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.User/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *spotigraph.UserGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error) {
	out := new(spotigraph.EmptyRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.User/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Search(ctx context.Context, in *spotigraph.UserSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedUserRes, error) {
	out := new(spotigraph.PaginatedUserRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.User/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Create(context.Context, *spotigraph.UserCreateReq) (*spotigraph.SingleUserRes, error)
	Get(context.Context, *spotigraph.UserGetReq) (*spotigraph.SingleUserRes, error)
	Update(context.Context, *spotigraph.UserUpdateReq) (*spotigraph.SingleUserRes, error)
	Delete(context.Context, *spotigraph.UserGetReq) (*spotigraph.EmptyRes, error)
	Search(context.Context, *spotigraph.UserSearchReq) (*spotigraph.PaginatedUserRes, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.UserCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*spotigraph.UserCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.UserGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*spotigraph.UserGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.UserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*spotigraph.UserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.UserGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*spotigraph.UserGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.UserSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.User/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Search(ctx, req.(*spotigraph.UserSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.spotigraph.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _User_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
}

// SquadClient is the client API for Squad service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquadClient interface {
	Create(ctx context.Context, in *spotigraph.SquadCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleSquadRes, error)
	Get(ctx context.Context, in *spotigraph.SquadGetReq, opts ...grpc.CallOption) (*spotigraph.SingleSquadRes, error)
	Update(ctx context.Context, in *spotigraph.SquadUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleSquadRes, error)
	Delete(ctx context.Context, in *spotigraph.SquadGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error)
	Search(ctx context.Context, in *spotigraph.SquadSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedSquadRes, error)
}

type squadClient struct {
	cc *grpc.ClientConn
}

func NewSquadClient(cc *grpc.ClientConn) SquadClient {
	return &squadClient{cc}
}

func (c *squadClient) Create(ctx context.Context, in *spotigraph.SquadCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleSquadRes, error) {
	out := new(spotigraph.SingleSquadRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Squad/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *squadClient) Get(ctx context.Context, in *spotigraph.SquadGetReq, opts ...grpc.CallOption) (*spotigraph.SingleSquadRes, error) {
	out := new(spotigraph.SingleSquadRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Squad/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *squadClient) Update(ctx context.Context, in *spotigraph.SquadUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleSquadRes, error) {
	out := new(spotigraph.SingleSquadRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Squad/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *squadClient) Delete(ctx context.Context, in *spotigraph.SquadGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error) {
	out := new(spotigraph.EmptyRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Squad/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *squadClient) Search(ctx context.Context, in *spotigraph.SquadSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedSquadRes, error) {
	out := new(spotigraph.PaginatedSquadRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Squad/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquadServer is the server API for Squad service.
type SquadServer interface {
	Create(context.Context, *spotigraph.SquadCreateReq) (*spotigraph.SingleSquadRes, error)
	Get(context.Context, *spotigraph.SquadGetReq) (*spotigraph.SingleSquadRes, error)
	Update(context.Context, *spotigraph.SquadUpdateReq) (*spotigraph.SingleSquadRes, error)
	Delete(context.Context, *spotigraph.SquadGetReq) (*spotigraph.EmptyRes, error)
	Search(context.Context, *spotigraph.SquadSearchReq) (*spotigraph.PaginatedSquadRes, error)
}

func RegisterSquadServer(s *grpc.Server, srv SquadServer) {
	s.RegisterService(&_Squad_serviceDesc, srv)
}

func _Squad_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.SquadCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquadServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Squad/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquadServer).Create(ctx, req.(*spotigraph.SquadCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Squad_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.SquadGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquadServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Squad/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquadServer).Get(ctx, req.(*spotigraph.SquadGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Squad_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.SquadUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquadServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Squad/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquadServer).Update(ctx, req.(*spotigraph.SquadUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Squad_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.SquadGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquadServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Squad/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquadServer).Delete(ctx, req.(*spotigraph.SquadGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Squad_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.SquadSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquadServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Squad/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquadServer).Search(ctx, req.(*spotigraph.SquadSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Squad_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.spotigraph.v1.Squad",
	HandlerType: (*SquadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Squad_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Squad_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Squad_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Squad_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Squad_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
}

// ChapterClient is the client API for Chapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChapterClient interface {
	Create(ctx context.Context, in *spotigraph.ChapterCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleChapterRes, error)
	Get(ctx context.Context, in *spotigraph.ChapterGetReq, opts ...grpc.CallOption) (*spotigraph.SingleChapterRes, error)
	Update(ctx context.Context, in *spotigraph.ChapterUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleChapterRes, error)
	Delete(ctx context.Context, in *spotigraph.ChapterGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error)
	Search(ctx context.Context, in *spotigraph.ChapterSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedChapterRes, error)
}

type chapterClient struct {
	cc *grpc.ClientConn
}

func NewChapterClient(cc *grpc.ClientConn) ChapterClient {
	return &chapterClient{cc}
}

func (c *chapterClient) Create(ctx context.Context, in *spotigraph.ChapterCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleChapterRes, error) {
	out := new(spotigraph.SingleChapterRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Chapter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterClient) Get(ctx context.Context, in *spotigraph.ChapterGetReq, opts ...grpc.CallOption) (*spotigraph.SingleChapterRes, error) {
	out := new(spotigraph.SingleChapterRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Chapter/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterClient) Update(ctx context.Context, in *spotigraph.ChapterUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleChapterRes, error) {
	out := new(spotigraph.SingleChapterRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Chapter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterClient) Delete(ctx context.Context, in *spotigraph.ChapterGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error) {
	out := new(spotigraph.EmptyRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Chapter/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterClient) Search(ctx context.Context, in *spotigraph.ChapterSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedChapterRes, error) {
	out := new(spotigraph.PaginatedChapterRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Chapter/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChapterServer is the server API for Chapter service.
type ChapterServer interface {
	Create(context.Context, *spotigraph.ChapterCreateReq) (*spotigraph.SingleChapterRes, error)
	Get(context.Context, *spotigraph.ChapterGetReq) (*spotigraph.SingleChapterRes, error)
	Update(context.Context, *spotigraph.ChapterUpdateReq) (*spotigraph.SingleChapterRes, error)
	Delete(context.Context, *spotigraph.ChapterGetReq) (*spotigraph.EmptyRes, error)
	Search(context.Context, *spotigraph.ChapterSearchReq) (*spotigraph.PaginatedChapterRes, error)
}

func RegisterChapterServer(s *grpc.Server, srv ChapterServer) {
	s.RegisterService(&_Chapter_serviceDesc, srv)
}

func _Chapter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.ChapterCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Chapter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServer).Create(ctx, req.(*spotigraph.ChapterCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chapter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.ChapterGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Chapter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServer).Get(ctx, req.(*spotigraph.ChapterGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chapter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.ChapterUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Chapter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServer).Update(ctx, req.(*spotigraph.ChapterUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chapter_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.ChapterGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Chapter/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServer).Delete(ctx, req.(*spotigraph.ChapterGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chapter_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.ChapterSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Chapter/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServer).Search(ctx, req.(*spotigraph.ChapterSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.spotigraph.v1.Chapter",
	HandlerType: (*ChapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Chapter_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Chapter_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Chapter_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Chapter_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Chapter_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
}

// GuildClient is the client API for Guild service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuildClient interface {
	Create(ctx context.Context, in *spotigraph.GuildCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleGuildRes, error)
	Get(ctx context.Context, in *spotigraph.GuildGetReq, opts ...grpc.CallOption) (*spotigraph.SingleGuildRes, error)
	Update(ctx context.Context, in *spotigraph.GuildUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleGuildRes, error)
	Delete(ctx context.Context, in *spotigraph.GuildGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error)
	Search(ctx context.Context, in *spotigraph.GuildSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedGuildRes, error)
}

type guildClient struct {
	cc *grpc.ClientConn
}

func NewGuildClient(cc *grpc.ClientConn) GuildClient {
	return &guildClient{cc}
}

func (c *guildClient) Create(ctx context.Context, in *spotigraph.GuildCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleGuildRes, error) {
	out := new(spotigraph.SingleGuildRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Guild/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildClient) Get(ctx context.Context, in *spotigraph.GuildGetReq, opts ...grpc.CallOption) (*spotigraph.SingleGuildRes, error) {
	out := new(spotigraph.SingleGuildRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Guild/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildClient) Update(ctx context.Context, in *spotigraph.GuildUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleGuildRes, error) {
	out := new(spotigraph.SingleGuildRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Guild/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildClient) Delete(ctx context.Context, in *spotigraph.GuildGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error) {
	out := new(spotigraph.EmptyRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Guild/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildClient) Search(ctx context.Context, in *spotigraph.GuildSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedGuildRes, error) {
	out := new(spotigraph.PaginatedGuildRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Guild/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuildServer is the server API for Guild service.
type GuildServer interface {
	Create(context.Context, *spotigraph.GuildCreateReq) (*spotigraph.SingleGuildRes, error)
	Get(context.Context, *spotigraph.GuildGetReq) (*spotigraph.SingleGuildRes, error)
	Update(context.Context, *spotigraph.GuildUpdateReq) (*spotigraph.SingleGuildRes, error)
	Delete(context.Context, *spotigraph.GuildGetReq) (*spotigraph.EmptyRes, error)
	Search(context.Context, *spotigraph.GuildSearchReq) (*spotigraph.PaginatedGuildRes, error)
}

func RegisterGuildServer(s *grpc.Server, srv GuildServer) {
	s.RegisterService(&_Guild_serviceDesc, srv)
}

func _Guild_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.GuildCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Guild/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServer).Create(ctx, req.(*spotigraph.GuildCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guild_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.GuildGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Guild/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServer).Get(ctx, req.(*spotigraph.GuildGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guild_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.GuildUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Guild/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServer).Update(ctx, req.(*spotigraph.GuildUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guild_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.GuildGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Guild/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServer).Delete(ctx, req.(*spotigraph.GuildGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guild_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.GuildSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Guild/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServer).Search(ctx, req.(*spotigraph.GuildSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Guild_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.spotigraph.v1.Guild",
	HandlerType: (*GuildServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Guild_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Guild_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Guild_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Guild_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Guild_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
}

// TribeClient is the client API for Tribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TribeClient interface {
	Create(ctx context.Context, in *spotigraph.TribeCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleTribeRes, error)
	Get(ctx context.Context, in *spotigraph.TribeGetReq, opts ...grpc.CallOption) (*spotigraph.SingleTribeRes, error)
	Update(ctx context.Context, in *spotigraph.TribeUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleTribeRes, error)
	Delete(ctx context.Context, in *spotigraph.TribeGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error)
	Search(ctx context.Context, in *spotigraph.TribeSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedTribeRes, error)
}

type tribeClient struct {
	cc *grpc.ClientConn
}

func NewTribeClient(cc *grpc.ClientConn) TribeClient {
	return &tribeClient{cc}
}

func (c *tribeClient) Create(ctx context.Context, in *spotigraph.TribeCreateReq, opts ...grpc.CallOption) (*spotigraph.SingleTribeRes, error) {
	out := new(spotigraph.SingleTribeRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Tribe/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tribeClient) Get(ctx context.Context, in *spotigraph.TribeGetReq, opts ...grpc.CallOption) (*spotigraph.SingleTribeRes, error) {
	out := new(spotigraph.SingleTribeRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Tribe/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tribeClient) Update(ctx context.Context, in *spotigraph.TribeUpdateReq, opts ...grpc.CallOption) (*spotigraph.SingleTribeRes, error) {
	out := new(spotigraph.SingleTribeRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Tribe/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tribeClient) Delete(ctx context.Context, in *spotigraph.TribeGetReq, opts ...grpc.CallOption) (*spotigraph.EmptyRes, error) {
	out := new(spotigraph.EmptyRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Tribe/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tribeClient) Search(ctx context.Context, in *spotigraph.TribeSearchReq, opts ...grpc.CallOption) (*spotigraph.PaginatedTribeRes, error) {
	out := new(spotigraph.PaginatedTribeRes)
	err := c.cc.Invoke(ctx, "/grpc.spotigraph.v1.Tribe/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TribeServer is the server API for Tribe service.
type TribeServer interface {
	Create(context.Context, *spotigraph.TribeCreateReq) (*spotigraph.SingleTribeRes, error)
	Get(context.Context, *spotigraph.TribeGetReq) (*spotigraph.SingleTribeRes, error)
	Update(context.Context, *spotigraph.TribeUpdateReq) (*spotigraph.SingleTribeRes, error)
	Delete(context.Context, *spotigraph.TribeGetReq) (*spotigraph.EmptyRes, error)
	Search(context.Context, *spotigraph.TribeSearchReq) (*spotigraph.PaginatedTribeRes, error)
}

func RegisterTribeServer(s *grpc.Server, srv TribeServer) {
	s.RegisterService(&_Tribe_serviceDesc, srv)
}

func _Tribe_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.TribeCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TribeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Tribe/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TribeServer).Create(ctx, req.(*spotigraph.TribeCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tribe_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.TribeGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TribeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Tribe/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TribeServer).Get(ctx, req.(*spotigraph.TribeGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tribe_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.TribeUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TribeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Tribe/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TribeServer).Update(ctx, req.(*spotigraph.TribeUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tribe_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.TribeGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TribeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Tribe/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TribeServer).Delete(ctx, req.(*spotigraph.TribeGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tribe_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spotigraph.TribeSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TribeServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.spotigraph.v1.Tribe/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TribeServer).Search(ctx, req.(*spotigraph.TribeSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tribe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.spotigraph.v1.Tribe",
	HandlerType: (*TribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Tribe_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Tribe_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Tribe_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Tribe_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Tribe_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
}
