package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	pb "grpc-k8s-project/client/genproto"
	"io"
	"log"
	"time"
)

var (
	k8sServiceAddr = "server:50051"
)

func main()  {
	conn, err := grpc.Dial(k8sServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewOrderServiceClient(conn)
	log.Print("client: %v", client)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	newOrder := pb.Order{Id: "110", Items: []string{"a", "b"}, Price: 450.20}
	res, err := client.AddOrder(ctx, &newOrder)

	if err != nil {
		got := status.Code(err)
		log.Printf("error occurred -> add order:, %v, err -> %v", got, err)
	} else {
		log.Print("AddOrder Response -> ", res.Value)
	}

	//Search Order: Server streaming scenario
	searchStream, _ := client.SearchOrders(ctx, &wrappers.StringValue{Value: "Google"})
	for {
		searchOrders, err := searchStream.Recv()
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			log.Print("Search result: ", searchOrders)
		}
	}

}
