package main

import (
	"context"
	pb "customer_service_proto"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CustomerServiceServer struct {
	pb.UnimplementedCustomerServiceServer
	CustomerHandler *CustomerHandler
}

func NewCustomerServiceServer(db *gorm.DB) *CustomerServiceServer {
	return &CustomerServiceServer{
		CustomerHandler: &CustomerHandler{DB: db},
	}
}

func (s *CustomerServiceServer) CheckCustomerDetailsValidity(context context.Context, request *pb.CustomerDetails) (*pb.CustomerDetailsValid, error) {
	found, err := s.CustomerHandler.ServiceBelongsToCustomer(uint(request.ServiceID), uint(request.CustomerID))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.CustomerDetailsValid{Valid: found}, nil
}

func (s *CustomerServiceServer) CreateCustomer(context context.Context, request *pb.CreateCustomerRequest) (*pb.CreateCustomerReply, error) {
	customer, err := s.CustomerHandler.CreateCustomer(request.CustomerName)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.CreateCustomerReply{
		CustomerID: uint32(customer.Model.ID),
		// CustomerName: customer.Name,
	}, nil
}

func (s *CustomerServiceServer) CreateCustomerService(context context.Context, request *pb.CreateCustomerServiceRequest) (*pb.CreateCustomerServiceReply, error) {
	service, err := s.CustomerHandler.CreateService(uint(request.CustomerID), request.ServiceName)
	if err != nil {
		if errors.Is(err, InvalidCustomerIdErr) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.CreateCustomerServiceReply{ServiceID: uint32(service.Model.ID)}, nil
}
