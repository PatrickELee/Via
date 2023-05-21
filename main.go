package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/PatrickELee/Via/viapb"
	"github.com/soheilhy/cmux"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCServer struct {
	viapb.UnimplementedViaServer
}

func (*gRPCServer) GetIncident(ctx context.Context, req *viapb.GetIncidentRequest) (*viapb.GetIncidentResponse, error) {
	log.Println("GetIncident Request")

	return &viapb.GetIncidentResponse{}, nil
}

func logError(err error) error {
	log.Println("Error with gRPC service:", err.Error())
	return status.Error(codes.Internal, err.Error())
}

func main() {
	log.Println("Running Via web server")

	listener, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		log.Println("Error with hosting server:", err.Error())
	}

	m := cmux.New(listener)

	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	grpcServer := grpc.NewServer()
	viapb.RegisterViaServer(grpcServer, &gRPCServer{})

	serveMux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("website/build/"))

	serveMux.Handle("/", http.StripPrefix("/", fileServer))
	httpServer := &http.Server{
		Handler: serveMux,
	}

	log.Printf("Server started at %v", listener.Addr().String())

	go grpcServer.Serve(grpcListener)
	go httpServer.Serve(httpListener)

	m.Serve()
}
