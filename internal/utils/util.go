package utils

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func logError(err error) error {
	log.Println("Error with gRPC service:", err.Error())
	return status.Error(codes.Internal, err.Error())
}
