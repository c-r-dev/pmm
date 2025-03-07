// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: managementpb/alerting/alerting.proto

package alertingv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Alerting_ListTemplates_FullMethodName  = "/alerting.v1.Alerting/ListTemplates"
	Alerting_CreateTemplate_FullMethodName = "/alerting.v1.Alerting/CreateTemplate"
	Alerting_UpdateTemplate_FullMethodName = "/alerting.v1.Alerting/UpdateTemplate"
	Alerting_DeleteTemplate_FullMethodName = "/alerting.v1.Alerting/DeleteTemplate"
	Alerting_CreateRule_FullMethodName     = "/alerting.v1.Alerting/CreateRule"
)

// AlertingClient is the client API for Alerting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Alerting service lets to manage alerting templates and create alerting rules from them.
type AlertingClient interface {
	// ListTemplates returns a list of all collected alert rule templates.
	ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error)
	// CreateTemplate creates a new template.
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	// UpdateTemplate updates existing template, previously created via API.
	UpdateTemplate(ctx context.Context, in *UpdateTemplateRequest, opts ...grpc.CallOption) (*UpdateTemplateResponse, error)
	// DeleteTemplate deletes existing, previously created via API.
	DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*DeleteTemplateResponse, error)
	// CreateRule creates alerting rule from the given template.
	CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleResponse, error)
}

type alertingClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertingClient(cc grpc.ClientConnInterface) AlertingClient {
	return &alertingClient{cc}
}

func (c *alertingClient) ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTemplatesResponse)
	err := c.cc.Invoke(ctx, Alerting_ListTemplates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, Alerting_CreateTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) UpdateTemplate(ctx context.Context, in *UpdateTemplateRequest, opts ...grpc.CallOption) (*UpdateTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTemplateResponse)
	err := c.cc.Invoke(ctx, Alerting_UpdateTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*DeleteTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTemplateResponse)
	err := c.cc.Invoke(ctx, Alerting_DeleteTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRuleResponse)
	err := c.cc.Invoke(ctx, Alerting_CreateRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertingServer is the server API for Alerting service.
// All implementations must embed UnimplementedAlertingServer
// for forward compatibility.
//
// Alerting service lets to manage alerting templates and create alerting rules from them.
type AlertingServer interface {
	// ListTemplates returns a list of all collected alert rule templates.
	ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error)
	// CreateTemplate creates a new template.
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	// UpdateTemplate updates existing template, previously created via API.
	UpdateTemplate(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error)
	// DeleteTemplate deletes existing, previously created via API.
	DeleteTemplate(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error)
	// CreateRule creates alerting rule from the given template.
	CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleResponse, error)
	mustEmbedUnimplementedAlertingServer()
}

// UnimplementedAlertingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlertingServer struct{}

func (UnimplementedAlertingServer) ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplates not implemented")
}

func (UnimplementedAlertingServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}

func (UnimplementedAlertingServer) UpdateTemplate(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}

func (UnimplementedAlertingServer) DeleteTemplate(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}

func (UnimplementedAlertingServer) CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRule not implemented")
}
func (UnimplementedAlertingServer) mustEmbedUnimplementedAlertingServer() {}
func (UnimplementedAlertingServer) testEmbeddedByValue()                  {}

// UnsafeAlertingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertingServer will
// result in compilation errors.
type UnsafeAlertingServer interface {
	mustEmbedUnimplementedAlertingServer()
}

func RegisterAlertingServer(s grpc.ServiceRegistrar, srv AlertingServer) {
	// If the following call pancis, it indicates UnimplementedAlertingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Alerting_ServiceDesc, srv)
}

func _Alerting_ListTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).ListTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alerting_ListTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).ListTemplates(ctx, req.(*ListTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alerting_CreateTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alerting_UpdateTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).UpdateTemplate(ctx, req.(*UpdateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alerting_DeleteTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).DeleteTemplate(ctx, req.(*DeleteTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alerting_CreateRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).CreateRule(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Alerting_ServiceDesc is the grpc.ServiceDesc for Alerting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Alerting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerting.v1.Alerting",
	HandlerType: (*AlertingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTemplates",
			Handler:    _Alerting_ListTemplates_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _Alerting_CreateTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _Alerting_UpdateTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _Alerting_DeleteTemplate_Handler,
		},
		{
			MethodName: "CreateRule",
			Handler:    _Alerting_CreateRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/alerting/alerting.proto",
}
