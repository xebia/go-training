// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: appointments.proto

/*
Package appointmentapi is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package appointmentapi

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_AppointmentInternal_GetAppointmentsOnStatus_0(ctx context.Context, marshaler runtime.Marshaler, client AppointmentInternalClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppointmentsOnStatusRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["status"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "status")
	}

	e, err = runtime.Enum(val, AppointmentStatus_value)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "status", err)
	}

	protoReq.Status = AppointmentStatus(e)

	msg, err := client.GetAppointmentsOnStatus(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AppointmentInternal_GetAppointmentsOnStatus_0(ctx context.Context, marshaler runtime.Marshaler, server AppointmentInternalServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppointmentsOnStatusRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["status"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "status")
	}

	e, err = runtime.Enum(val, AppointmentStatus_value)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "status", err)
	}

	protoReq.Status = AppointmentStatus(e)

	msg, err := server.GetAppointmentsOnStatus(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_AppointmentInternal_ModifyAppointmentStatus_0 = &utilities.DoubleArray{Encoding: map[string]int{"appointmentUid": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_AppointmentInternal_ModifyAppointmentStatus_0(ctx context.Context, marshaler runtime.Marshaler, client AppointmentInternalClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ModifyAppointmentStatusRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["appointmentUid"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "appointmentUid")
	}

	protoReq.AppointmentUid, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "appointmentUid", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_AppointmentInternal_ModifyAppointmentStatus_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ModifyAppointmentStatus(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AppointmentInternal_ModifyAppointmentStatus_0(ctx context.Context, marshaler runtime.Marshaler, server AppointmentInternalServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ModifyAppointmentStatusRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["appointmentUid"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "appointmentUid")
	}

	protoReq.AppointmentUid, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "appointmentUid", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_AppointmentInternal_ModifyAppointmentStatus_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ModifyAppointmentStatus(ctx, &protoReq)
	return msg, metadata, err

}

func request_AppointmentExternal_GetAppointmentsOnUser_0(ctx context.Context, marshaler runtime.Marshaler, client AppointmentExternalClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppointmentsOnUserRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["userUid"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "userUid")
	}

	protoReq.UserUid, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "userUid", err)
	}

	msg, err := client.GetAppointmentsOnUser(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AppointmentExternal_GetAppointmentsOnUser_0(ctx context.Context, marshaler runtime.Marshaler, server AppointmentExternalServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppointmentsOnUserRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["userUid"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "userUid")
	}

	protoReq.UserUid, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "userUid", err)
	}

	msg, err := server.GetAppointmentsOnUser(ctx, &protoReq)
	return msg, metadata, err

}

func request_AppointmentExternal_RequestAppointment_0(ctx context.Context, marshaler runtime.Marshaler, client AppointmentExternalClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RequestAppointmentRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.Appointment); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.RequestAppointment(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AppointmentExternal_RequestAppointment_0(ctx context.Context, marshaler runtime.Marshaler, server AppointmentExternalServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RequestAppointmentRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.Appointment); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.RequestAppointment(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterAppointmentInternalHandlerServer registers the http handlers for service AppointmentInternal to "mux".
// UnaryRPC     :call AppointmentInternalServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterAppointmentInternalHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AppointmentInternalServer) error {

	mux.Handle("GET", pattern_AppointmentInternal_GetAppointmentsOnStatus_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AppointmentInternal_GetAppointmentsOnStatus_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentInternal_GetAppointmentsOnStatus_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_AppointmentInternal_ModifyAppointmentStatus_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AppointmentInternal_ModifyAppointmentStatus_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentInternal_ModifyAppointmentStatus_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterAppointmentExternalHandlerServer registers the http handlers for service AppointmentExternal to "mux".
// UnaryRPC     :call AppointmentExternalServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterAppointmentExternalHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AppointmentExternalServer) error {

	mux.Handle("GET", pattern_AppointmentExternal_GetAppointmentsOnUser_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AppointmentExternal_GetAppointmentsOnUser_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentExternal_GetAppointmentsOnUser_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_AppointmentExternal_RequestAppointment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AppointmentExternal_RequestAppointment_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentExternal_RequestAppointment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterAppointmentInternalHandlerFromEndpoint is same as RegisterAppointmentInternalHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAppointmentInternalHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAppointmentInternalHandler(ctx, mux, conn)
}

// RegisterAppointmentInternalHandler registers the http handlers for service AppointmentInternal to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAppointmentInternalHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAppointmentInternalHandlerClient(ctx, mux, NewAppointmentInternalClient(conn))
}

// RegisterAppointmentInternalHandlerClient registers the http handlers for service AppointmentInternal
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AppointmentInternalClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AppointmentInternalClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AppointmentInternalClient" to call the correct interceptors.
func RegisterAppointmentInternalHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AppointmentInternalClient) error {

	mux.Handle("GET", pattern_AppointmentInternal_GetAppointmentsOnStatus_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AppointmentInternal_GetAppointmentsOnStatus_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentInternal_GetAppointmentsOnStatus_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_AppointmentInternal_ModifyAppointmentStatus_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AppointmentInternal_ModifyAppointmentStatus_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentInternal_ModifyAppointmentStatus_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_AppointmentInternal_GetAppointmentsOnStatus_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 2}, []string{"api", "appointment", "status"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_AppointmentInternal_ModifyAppointmentStatus_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3}, []string{"api", "appointment", "appointmentUid", "status"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_AppointmentInternal_GetAppointmentsOnStatus_0 = runtime.ForwardResponseMessage

	forward_AppointmentInternal_ModifyAppointmentStatus_0 = runtime.ForwardResponseMessage
)

// RegisterAppointmentExternalHandlerFromEndpoint is same as RegisterAppointmentExternalHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAppointmentExternalHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAppointmentExternalHandler(ctx, mux, conn)
}

// RegisterAppointmentExternalHandler registers the http handlers for service AppointmentExternal to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAppointmentExternalHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAppointmentExternalHandlerClient(ctx, mux, NewAppointmentExternalClient(conn))
}

// RegisterAppointmentExternalHandlerClient registers the http handlers for service AppointmentExternal
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AppointmentExternalClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AppointmentExternalClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AppointmentExternalClient" to call the correct interceptors.
func RegisterAppointmentExternalHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AppointmentExternalClient) error {

	mux.Handle("GET", pattern_AppointmentExternal_GetAppointmentsOnUser_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AppointmentExternal_GetAppointmentsOnUser_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentExternal_GetAppointmentsOnUser_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_AppointmentExternal_RequestAppointment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AppointmentExternal_RequestAppointment_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AppointmentExternal_RequestAppointment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_AppointmentExternal_GetAppointmentsOnUser_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "appointment", "user", "userUid"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_AppointmentExternal_RequestAppointment_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api", "appointment"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_AppointmentExternal_GetAppointmentsOnUser_0 = runtime.ForwardResponseMessage

	forward_AppointmentExternal_RequestAppointment_0 = runtime.ForwardResponseMessage
)