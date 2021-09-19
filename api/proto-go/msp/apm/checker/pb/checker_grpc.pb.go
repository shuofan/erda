// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: checker.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// CheckerServiceClient is the client API for CheckerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckerServiceClient interface {
	CreateChecker(ctx context.Context, in *CreateCheckerRequest, opts ...grpc.CallOption) (*CreateCheckerResponse, error)
	UpdateChecker(ctx context.Context, in *UpdateCheckerRequest, opts ...grpc.CallOption) (*UpdateCheckerResponse, error)
	DeleteChecker(ctx context.Context, in *UpdateCheckerRequest, opts ...grpc.CallOption) (*UpdateCheckerResponse, error)
	ListCheckers(ctx context.Context, in *ListCheckersRequest, opts ...grpc.CallOption) (*ListCheckersResponse, error)
	DescribeCheckers(ctx context.Context, in *DescribeCheckersRequest, opts ...grpc.CallOption) (*DescribeCheckersResponse, error)
	DescribeChecker(ctx context.Context, in *DescribeCheckerRequest, opts ...grpc.CallOption) (*DescribeCheckerResponse, error)
}

type checkerServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewCheckerServiceClient(cc grpc1.ClientConnInterface) CheckerServiceClient {
	return &checkerServiceClient{cc}
}

func (c *checkerServiceClient) CreateChecker(ctx context.Context, in *CreateCheckerRequest, opts ...grpc.CallOption) (*CreateCheckerResponse, error) {
	out := new(CreateCheckerResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.checker.CheckerService/CreateChecker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerServiceClient) UpdateChecker(ctx context.Context, in *UpdateCheckerRequest, opts ...grpc.CallOption) (*UpdateCheckerResponse, error) {
	out := new(UpdateCheckerResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.checker.CheckerService/UpdateChecker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerServiceClient) DeleteChecker(ctx context.Context, in *UpdateCheckerRequest, opts ...grpc.CallOption) (*UpdateCheckerResponse, error) {
	out := new(UpdateCheckerResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.checker.CheckerService/DeleteChecker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerServiceClient) ListCheckers(ctx context.Context, in *ListCheckersRequest, opts ...grpc.CallOption) (*ListCheckersResponse, error) {
	out := new(ListCheckersResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.checker.CheckerService/ListCheckers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerServiceClient) DescribeCheckers(ctx context.Context, in *DescribeCheckersRequest, opts ...grpc.CallOption) (*DescribeCheckersResponse, error) {
	out := new(DescribeCheckersResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.checker.CheckerService/DescribeCheckers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerServiceClient) DescribeChecker(ctx context.Context, in *DescribeCheckerRequest, opts ...grpc.CallOption) (*DescribeCheckerResponse, error) {
	out := new(DescribeCheckerResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.checker.CheckerService/DescribeChecker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServiceServer is the server API for CheckerService service.
// All implementations should embed UnimplementedCheckerServiceServer
// for forward compatibility
type CheckerServiceServer interface {
	CreateChecker(context.Context, *CreateCheckerRequest) (*CreateCheckerResponse, error)
	UpdateChecker(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error)
	DeleteChecker(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error)
	ListCheckers(context.Context, *ListCheckersRequest) (*ListCheckersResponse, error)
	DescribeCheckers(context.Context, *DescribeCheckersRequest) (*DescribeCheckersResponse, error)
	DescribeChecker(context.Context, *DescribeCheckerRequest) (*DescribeCheckerResponse, error)
}

// UnimplementedCheckerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCheckerServiceServer struct {
}

func (*UnimplementedCheckerServiceServer) CreateChecker(context.Context, *CreateCheckerRequest) (*CreateCheckerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChecker not implemented")
}
func (*UnimplementedCheckerServiceServer) UpdateChecker(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChecker not implemented")
}
func (*UnimplementedCheckerServiceServer) DeleteChecker(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChecker not implemented")
}
func (*UnimplementedCheckerServiceServer) ListCheckers(context.Context, *ListCheckersRequest) (*ListCheckersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCheckers not implemented")
}
func (*UnimplementedCheckerServiceServer) DescribeCheckers(context.Context, *DescribeCheckersRequest) (*DescribeCheckersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCheckers not implemented")
}
func (*UnimplementedCheckerServiceServer) DescribeChecker(context.Context, *DescribeCheckerRequest) (*DescribeCheckerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeChecker not implemented")
}

func RegisterCheckerServiceServer(s grpc1.ServiceRegistrar, srv CheckerServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_CheckerService_serviceDesc(srv, opts...), srv)
}

var _CheckerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.msp.apm.checker.CheckerService",
	HandlerType: (*CheckerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "checker.proto",
}

func _get_CheckerService_serviceDesc(srv CheckerServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_CheckerService_CreateChecker_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateChecker(ctx, req.(*CreateCheckerRequest))
	}
	var _CheckerService_CreateChecker_info transport.ServiceInfo
	if h.Interceptor != nil {
		_CheckerService_CreateChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "CreateChecker", srv)
		_CheckerService_CreateChecker_Handler = h.Interceptor(_CheckerService_CreateChecker_Handler)
	}

	_CheckerService_UpdateChecker_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateChecker(ctx, req.(*UpdateCheckerRequest))
	}
	var _CheckerService_UpdateChecker_info transport.ServiceInfo
	if h.Interceptor != nil {
		_CheckerService_UpdateChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "UpdateChecker", srv)
		_CheckerService_UpdateChecker_Handler = h.Interceptor(_CheckerService_UpdateChecker_Handler)
	}

	_CheckerService_DeleteChecker_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteChecker(ctx, req.(*UpdateCheckerRequest))
	}
	var _CheckerService_DeleteChecker_info transport.ServiceInfo
	if h.Interceptor != nil {
		_CheckerService_DeleteChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "DeleteChecker", srv)
		_CheckerService_DeleteChecker_Handler = h.Interceptor(_CheckerService_DeleteChecker_Handler)
	}

	_CheckerService_ListCheckers_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListCheckers(ctx, req.(*ListCheckersRequest))
	}
	var _CheckerService_ListCheckers_info transport.ServiceInfo
	if h.Interceptor != nil {
		_CheckerService_ListCheckers_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "ListCheckers", srv)
		_CheckerService_ListCheckers_Handler = h.Interceptor(_CheckerService_ListCheckers_Handler)
	}

	_CheckerService_DescribeCheckers_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DescribeCheckers(ctx, req.(*DescribeCheckersRequest))
	}
	var _CheckerService_DescribeCheckers_info transport.ServiceInfo
	if h.Interceptor != nil {
		_CheckerService_DescribeCheckers_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "DescribeCheckers", srv)
		_CheckerService_DescribeCheckers_Handler = h.Interceptor(_CheckerService_DescribeCheckers_Handler)
	}

	_CheckerService_DescribeChecker_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DescribeChecker(ctx, req.(*DescribeCheckerRequest))
	}
	var _CheckerService_DescribeChecker_info transport.ServiceInfo
	if h.Interceptor != nil {
		_CheckerService_DescribeChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "DescribeChecker", srv)
		_CheckerService_DescribeChecker_Handler = h.Interceptor(_CheckerService_DescribeChecker_Handler)
	}

	var serviceDesc = _CheckerService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CreateChecker",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateCheckerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(CheckerServiceServer).CreateChecker(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _CheckerService_CreateChecker_info)
				}
				if interceptor == nil {
					return _CheckerService_CreateChecker_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.checker.CheckerService/CreateChecker",
				}
				return interceptor(ctx, in, info, _CheckerService_CreateChecker_Handler)
			},
		},
		{
			MethodName: "UpdateChecker",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateCheckerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(CheckerServiceServer).UpdateChecker(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _CheckerService_UpdateChecker_info)
				}
				if interceptor == nil {
					return _CheckerService_UpdateChecker_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.checker.CheckerService/UpdateChecker",
				}
				return interceptor(ctx, in, info, _CheckerService_UpdateChecker_Handler)
			},
		},
		{
			MethodName: "DeleteChecker",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateCheckerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(CheckerServiceServer).DeleteChecker(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _CheckerService_DeleteChecker_info)
				}
				if interceptor == nil {
					return _CheckerService_DeleteChecker_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.checker.CheckerService/DeleteChecker",
				}
				return interceptor(ctx, in, info, _CheckerService_DeleteChecker_Handler)
			},
		},
		{
			MethodName: "ListCheckers",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListCheckersRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(CheckerServiceServer).ListCheckers(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _CheckerService_ListCheckers_info)
				}
				if interceptor == nil {
					return _CheckerService_ListCheckers_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.checker.CheckerService/ListCheckers",
				}
				return interceptor(ctx, in, info, _CheckerService_ListCheckers_Handler)
			},
		},
		{
			MethodName: "DescribeCheckers",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DescribeCheckersRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(CheckerServiceServer).DescribeCheckers(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _CheckerService_DescribeCheckers_info)
				}
				if interceptor == nil {
					return _CheckerService_DescribeCheckers_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.checker.CheckerService/DescribeCheckers",
				}
				return interceptor(ctx, in, info, _CheckerService_DescribeCheckers_Handler)
			},
		},
		{
			MethodName: "DescribeChecker",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DescribeCheckerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(CheckerServiceServer).DescribeChecker(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _CheckerService_DescribeChecker_info)
				}
				if interceptor == nil {
					return _CheckerService_DescribeChecker_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.checker.CheckerService/DescribeChecker",
				}
				return interceptor(ctx, in, info, _CheckerService_DescribeChecker_Handler)
			},
		},
	}
	return &serviceDesc
}