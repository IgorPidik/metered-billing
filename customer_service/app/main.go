package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "customer_service_proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, dbErr := initDB()
	if dbErr != nil {
		log.Fatalf("failed to init db: %v", dbErr)
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, NewCustomerServiceServer(db))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
