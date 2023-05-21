package server

import (
	"log"
	"net"
	"net/http"

	viapb "github.com/PatrickELee/Via/api/proto"
	igrpc "github.com/PatrickELee/Via/internal/grpc"
	ihttp "github.com/PatrickELee/Via/internal/http"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const Protocol = "tcp"
const Ip = "0.0.0.0"
const Port = "32314"

type ViaServerMux struct {
	m cmux.CMux
}

type ViaServer struct {
	gRPCServer   *grpc.Server
	gRPCListener net.Listener
	httpServer   *http.Server
	httpListener net.Listener
	vsm          *ViaServerMux
}

func CreateViaServerMux() *ViaServerMux {
	listener, err := net.Listen(Protocol, Ip+":"+Port)
	if err != nil {
		log.Println("Error with hosting server:", err.Error())
	}

	log.Printf("Server started at %v", listener.Addr().String())

	vsm := &ViaServerMux{
		m: cmux.New(listener),
	}

	return vsm
}

func (vsm *ViaServerMux) CreateGRPCServerAndListener() (*grpc.Server, net.Listener) {
	grpcListener := vsm.m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	grpcServer := grpc.NewServer()
	viapb.RegisterViaServer(grpcServer, &igrpc.GRPCServer{})

	return grpcServer, grpcListener
}

func (vsm *ViaServerMux) CreateHTTPServerAndListener() (*http.Server, net.Listener) {
	httpListener := vsm.m.Match(cmux.HTTP1Fast())
	serveMux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("website/build/"))

	serveMux.Handle("/", http.StripPrefix("/", fileServer))
	serveMux.HandleFunc("/api/dangerprobability", ihttp.DangerProbHandler)
	httpServer := &http.Server{
		Handler: serveMux,
	}

	return httpServer, httpListener
}

func CreateViaWebServers() *ViaServer {
	log.Println("Launching Via web servers")

	vs := &ViaServer{}
	vs.vsm = CreateViaServerMux()
	vs.gRPCServer, vs.gRPCListener = vs.vsm.CreateGRPCServerAndListener()
	vs.httpServer, vs.httpListener = vs.vsm.CreateHTTPServerAndListener()

	return vs
}

func (vs *ViaServer) ListenAndServe() {
	go vs.gRPCServer.Serve(vs.gRPCListener)
	go vs.httpServer.Serve(vs.httpListener)

	vs.vsm.m.Serve()
}
