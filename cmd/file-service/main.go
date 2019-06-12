package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bitstored/file-service/pb"
	"github.com/bitstored/file-service/pkg/server"
	"github.com/bitstored/file-service/pkg/service"

	"github.com/cenkalti/backoff"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	ServiceName = "file-management"
)

var (
	grpcAddr = flag.String("grpc", ":4005", "gRPC API address")
	httpAddr = flag.String("http", ":5005", "HTTP API address")
	cert     = flag.String("cert", "scripts/server.crt", "certificate pathname")
	certKey  = flag.String("certkey", "scripts/server.key", "private key pathname")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	fmt.Println(os.Args)

	service := service.NewService()
	gRPCListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %s", *grpcAddr, err)
	}

	devServer := server.NewServer(service)

	// Register standard server metrics and customized metrics to registry.
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	gRPCServer := grpc.NewServer()

	pb.RegisterFileManagementServer(gRPCServer, devServer)
	reflection.Register(gRPCServer)
	grpc_prometheus.Register(gRPCServer)
	grpcMetrics.InitializeMetrics(gRPCServer)

	reg := prometheus.NewRegistry()
	reg.MustRegister(grpcMetrics)

	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	go func() {
		if err := gRPCServer.Serve(gRPCListener); err != nil {
			log.Fatalf("Failed to serve gRPC: %s", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, *grpcAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	gw, err := server.NewGateway(ctx, conn)
	if err != nil {
		log.Fatalf("Unable to create gateway- %v", err)
	}
	mux.Handle("/", gw)

	httpServer := &http.Server{
		Handler:      mux,
		Addr:         *httpAddr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		if err := httpServer.ListenAndServeTLS(*cert, *certKey); err != http.ErrServerClosed {
			log.Fatalf("Unable to start a http server - %s", err)
		}
	}()
	fmt.Printf("File server listening on  %s for gRPC\nFile server listening on on %s for http\n", *grpcAddr, *httpAddr)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	// Wait for signal
	<-done

	// Create wait deadline

	// Doesn't block if no connections, will wait until the timeout deadline otherwise.
	log.Println("shutting down")
	err = httpServer.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

func retry(ctx context.Context, f func() error) error {
	return backoff.Retry(f, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
}
