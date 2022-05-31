package main

import (
	"github.com/PGITAb/an-example-http-api-tests/mock-service/handler"
	pb "github.com/PGITAb/an-example-proto/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("an-example-service"),
	)

	// Register handler
	pb.RegisterAnExampleServiceHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
