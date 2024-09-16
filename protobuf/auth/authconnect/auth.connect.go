// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protobuf/auth/auth.proto

package authconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	auth "github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceLoginProcedure is the fully-qualified name of the AuthService's Login RPC.
	AuthServiceLoginProcedure = "/AuthService/Login"
	// AuthServiceSubmitOtpProcedure is the fully-qualified name of the AuthService's SubmitOtp RPC.
	AuthServiceSubmitOtpProcedure = "/AuthService/SubmitOtp"
	// AuthServiceAuthenticateProcedure is the fully-qualified name of the AuthService's Authenticate
	// RPC.
	AuthServiceAuthenticateProcedure = "/AuthService/Authenticate"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authServiceServiceDescriptor            = auth.File_protobuf_auth_auth_proto.Services().ByName("AuthService")
	authServiceLoginMethodDescriptor        = authServiceServiceDescriptor.Methods().ByName("Login")
	authServiceSubmitOtpMethodDescriptor    = authServiceServiceDescriptor.Methods().ByName("SubmitOtp")
	authServiceAuthenticateMethodDescriptor = authServiceServiceDescriptor.Methods().ByName("Authenticate")
)

// AuthServiceClient is a client for the AuthService service.
type AuthServiceClient interface {
	Login(context.Context, *connect.Request[auth.LoginRequest]) (*connect.Response[auth.LoginResponse], error)
	SubmitOtp(context.Context, *connect.Request[auth.SubmitOtpRequest]) (*connect.Response[auth.SubmitOtpResponse], error)
	Authenticate(context.Context, *connect.Request[auth.AuthenticateRequest]) (*connect.Response[auth.AuthenticateResponse], error)
}

// NewAuthServiceClient constructs a client for the AuthService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		login: connect.NewClient[auth.LoginRequest, auth.LoginResponse](
			httpClient,
			baseURL+AuthServiceLoginProcedure,
			connect.WithSchema(authServiceLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		submitOtp: connect.NewClient[auth.SubmitOtpRequest, auth.SubmitOtpResponse](
			httpClient,
			baseURL+AuthServiceSubmitOtpProcedure,
			connect.WithSchema(authServiceSubmitOtpMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		authenticate: connect.NewClient[auth.AuthenticateRequest, auth.AuthenticateResponse](
			httpClient,
			baseURL+AuthServiceAuthenticateProcedure,
			connect.WithSchema(authServiceAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	login        *connect.Client[auth.LoginRequest, auth.LoginResponse]
	submitOtp    *connect.Client[auth.SubmitOtpRequest, auth.SubmitOtpResponse]
	authenticate *connect.Client[auth.AuthenticateRequest, auth.AuthenticateResponse]
}

// Login calls AuthService.Login.
func (c *authServiceClient) Login(ctx context.Context, req *connect.Request[auth.LoginRequest]) (*connect.Response[auth.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// SubmitOtp calls AuthService.SubmitOtp.
func (c *authServiceClient) SubmitOtp(ctx context.Context, req *connect.Request[auth.SubmitOtpRequest]) (*connect.Response[auth.SubmitOtpResponse], error) {
	return c.submitOtp.CallUnary(ctx, req)
}

// Authenticate calls AuthService.Authenticate.
func (c *authServiceClient) Authenticate(ctx context.Context, req *connect.Request[auth.AuthenticateRequest]) (*connect.Response[auth.AuthenticateResponse], error) {
	return c.authenticate.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the AuthService service.
type AuthServiceHandler interface {
	Login(context.Context, *connect.Request[auth.LoginRequest]) (*connect.Response[auth.LoginResponse], error)
	SubmitOtp(context.Context, *connect.Request[auth.SubmitOtpRequest]) (*connect.Response[auth.SubmitOtpResponse], error)
	Authenticate(context.Context, *connect.Request[auth.AuthenticateRequest]) (*connect.Response[auth.AuthenticateResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceLoginHandler := connect.NewUnaryHandler(
		AuthServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(authServiceLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceSubmitOtpHandler := connect.NewUnaryHandler(
		AuthServiceSubmitOtpProcedure,
		svc.SubmitOtp,
		connect.WithSchema(authServiceSubmitOtpMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceAuthenticateHandler := connect.NewUnaryHandler(
		AuthServiceAuthenticateProcedure,
		svc.Authenticate,
		connect.WithSchema(authServiceAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceLoginProcedure:
			authServiceLoginHandler.ServeHTTP(w, r)
		case AuthServiceSubmitOtpProcedure:
			authServiceSubmitOtpHandler.ServeHTTP(w, r)
		case AuthServiceAuthenticateProcedure:
			authServiceAuthenticateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) Login(context.Context, *connect.Request[auth.LoginRequest]) (*connect.Response[auth.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("AuthService.Login is not implemented"))
}

func (UnimplementedAuthServiceHandler) SubmitOtp(context.Context, *connect.Request[auth.SubmitOtpRequest]) (*connect.Response[auth.SubmitOtpResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("AuthService.SubmitOtp is not implemented"))
}

func (UnimplementedAuthServiceHandler) Authenticate(context.Context, *connect.Request[auth.AuthenticateRequest]) (*connect.Response[auth.AuthenticateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("AuthService.Authenticate is not implemented"))
}
