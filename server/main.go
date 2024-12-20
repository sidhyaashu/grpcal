package main

import (
	"context"
	"log"
	"net"

	"github.com/sidhyaashu/grpcal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context, in *pb.CalculationRequest,
)(*pb.CalculationResponse,error){
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	},nil
}

func (s *server) Subtract(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A - in.B,
	}, nil
}

func (s *server) Multiply(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A * in.B,
	}, nil
}

func (s *server) Divide(
	ctx context.Context, in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {

	if in.B == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument, "cannot divide by zero",
		)
	}

	return &pb.CalculationResponse{
		Result: in.A / in.B,
	}, nil
}

func (s *server) Sum(
	ctx context.Context, in *pb.NumbersRequest,
) (*pb.CalculationResponse, error) {
	var sum int64

	for _, num := range in.Numbers {
		sum += num
	}

	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}


func main() {
	listner , err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s,&server{})
	if err := s.Serve(listner); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}