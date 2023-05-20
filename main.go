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

type server struct {
	viapb.UnimplementedViaServer
}

// var (
// 	offer_client   client.OfferClient
// 	catalog_client client.CatalogClient
// )

func (*server) GetIncident(ctx context.Context, req *viapb.GetIncidentRequest) (*viapb.GetIncidentResponse, error) {
	log.Println("Order Service - Called CreateOrder")

	return &viapb.GetIncidentResponse{}, nil
}

func error_response(err error) error {
	log.Println("Order Service - ERROR:", err.Error())
	return status.Error(codes.Internal, err.Error())
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("Rendering home\n")
// 	http.ServeFile(w, r)
// }

func main() {
	log.Println("Running Order Service")

	lis, err := net.Listen("tcp", "0.0.0.0:55052")
	if err != nil {
		log.Println("Order Service - ERROR:", err.Error())
	}

	m := cmux.New(lis)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := grpc.NewServer()
	viapb.RegisterViaServer(grpcS, &server{})

	serveMux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	serveMux.Handle("/", http.StripPrefix("/static", fileServer))
	httpS := &http.Server{
		Handler: serveMux,
	}

	log.Printf("Server started at %v", lis.Addr().String())

	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	m.Serve()
}
