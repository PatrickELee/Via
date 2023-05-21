package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/PatrickELee/Via/viapb"
	"github.com/soheilhy/cmux"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetDangerProbabilityJsonReq struct {
	Location string `json:"location"`
	Time     string `json:"time"`
}

type GetDangerProbabilityJsonResp struct {
	Danger string `json:"danger"`
}
type gRPCServer struct {
	viapb.UnimplementedViaServer
}

func (*gRPCServer) GetIncident(ctx context.Context, req *viapb.GetIncidentRequest) (*viapb.GetIncidentResponse, error) {
	log.Println("GetIncident Request")

	return &viapb.GetIncidentResponse{}, nil
}

func (*gRPCServer) GetDangerProbability(ctx context.Context, req *viapb.GetDangerProbabilityRequest) (*viapb.GetDangerProbabilityResponse, error) {
	log.Println("GetDangerProbability Request")

	return &viapb.GetDangerProbabilityResponse{Danger: "high"}, nil
}

func logError(err error) error {
	log.Println("Error with gRPC service:", err.Error())
	return status.Error(codes.Internal, err.Error())
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func dangerProbHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		enableCors(&w)
		log.Println("Get Danger Probabilty Request")
		d := json.NewDecoder((r.Body))
		dangerProbReq := &GetDangerProbabilityJsonReq{}
		err := d.Decode(dangerProbReq)
		if err != nil {
			fmt.Printf("Error occured, %v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Printf(dangerProbReq.Location)
		fmt.Printf(dangerProbReq.Time)

		dangerProbResp := &GetDangerProbabilityJsonResp{
			Danger: "high",
		}
		resp, _ := json.Marshal(dangerProbResp)
		w.Write(resp)
	default:
		fmt.Println("Error with request method")
	}
}

func main() {
	log.Println("Running Via web server")

	listener, err := net.Listen("tcp", "0.0.0.0:32314")
	if err != nil {
		log.Println("Error with hosting server:", err.Error())
	}

	m := cmux.New(listener)

	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	grpcServer := grpc.NewServer()
	viapb.RegisterViaServer(grpcServer, &gRPCServer{})

	serveMux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("website/build/"))

	serveMux.Handle("/", http.StripPrefix("/", fileServer))
	serveMux.HandleFunc("/api/dangerprobability", dangerProbHandler)
	httpServer := &http.Server{
		Handler: serveMux,
	}

	log.Printf("Server started at %v", listener.Addr().String())

	go grpcServer.Serve(grpcListener)
	go httpServer.Serve(httpListener)

	m.Serve()
}
