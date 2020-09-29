package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-k8s-project/server/genproto"
	"log"
	"net"
	"strings"
)

var (
	port = "50051"
	orders = make(map[string]pb.Order)
)

type orderService struct {
	orders map[string]*pb.Order
}

func main()  {
	initSampleData()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	svc := new(orderService)

	log.Printf("Service config: %+v", svc)

	pb.RegisterOrderServiceServer(srv, svc)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initSampleData()  {
	orders["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Price: 1800.00}
	orders["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Price: 400.00}
	orders["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Price: 400.00}
	orders["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Price: 30.00}
	orders["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Price: 300.00}
}

func (s *orderService) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrappers.StringValue, error)  {
	log.Printf("Order Added, ID: %v", orderReq.Id)
	orders[orderReq.Id] = *orderReq

	if orderReq.Id == "-1" {
		log.Printf("order ID is invalid! -> received order ID %s", orderReq.Id)

		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		ds, err := errorStatus.WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field: "ID",
				Description: fmt.Sprintf("order ID received is not valid %s : %s", orderReq.Id),
			},
		)

		if err != nil {
			return nil, errorStatus.Err()
		}
		return nil, ds.Err()
	}

	return &wrappers.StringValue{Value: fmt.Sprintf("Order Added: %v", orderReq.Id)}, nil
}

func (s *orderService) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderService_SearchOrdersServer) error  {
	for key, order := range orders {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				err := stream.Send(&order)
				if err != nil {
					log.Fatalf("error sending message to stream: %v", err)
				}
				log.Printf("Matching order found: %v", key)
				break
			}
		}
	}
	return nil
}