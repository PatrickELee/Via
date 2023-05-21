package grpc

import (
	"context"
	"log"

	viapb "github.com/PatrickELee/Via/api/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	viapb.UnimplementedViaServer
}

func (*GRPCServer) GetIncident(ctx context.Context, req *viapb.GetIncidentRequest) (*viapb.GetIncidentResponse, error) {
	log.Println("GetIncident Request")

	return &viapb.GetIncidentResponse{}, nil
}

func (*GRPCServer) GetDangerProbability(ctx context.Context, req *viapb.GetDangerProbabilityRequest) (*viapb.GetDangerProbabilityResponse, error) {
	log.Println("GetDangerProbability Request")

	return &viapb.GetDangerProbabilityResponse{Danger: "high"}, nil
}

func MakeDangerProbRPCRequest(location, time string) string {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := viapb.NewViaClient(conn)

	response, err := c.GetDangerProbability(context.Background(), &viapb.GetDangerProbabilityRequest{Location: location, Time: time})
	if err != nil {
		log.Fatalf("Error when calling GetDangerProbability: %s", err)
	}
	log.Printf("Response from server: %s", response.Danger)

	return response.Danger
}
