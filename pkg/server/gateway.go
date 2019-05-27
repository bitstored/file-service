package server

import (
	"context"
	"github.com/bitstored/file-service/pb"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// NewGateway returns a new gateway server which translates HTTP into gRPC.
func NewGateway(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {

	mux := gwruntime.NewServeMux()

	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		pb.RegisterFileManagementHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}
