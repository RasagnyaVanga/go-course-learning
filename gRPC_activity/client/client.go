package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "gRPC_activity/routeguide"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//calling GetFeature
	feature, err := client.GetFeature(ctx, &pb.Point{Latitude: 0, Longitude: 0})
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(feature)

	//calling ListFeatures  - serverStreaming RPC
	stream, err := client.ListFeatures(ctx, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	})
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}

	//calling RecordRoute - clientStreaming RPC
	stream2, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	for i := 0; i < 3; i++ {
		if err := stream2.Send(&pb.Point{Latitude: int32(i), Longitude: int32(i)}); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send failed: %v", err)
		}
	}
	reply, err := stream2.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	log.Printf("Route summary: %v", reply)

	//calling RouteChat - bidirectional RPC  -- sending and receiving - used channels to send and receive from stream.
	stream3, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("client.RouteChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream3.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.RouteChat failed: %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for i := 0; i < 3; i++ {
		if err := stream3.Send(&pb.RouteNote{
			Location: &pb.Point{Latitude: int32(i), Longitude: int32(i)},
			Message:  "Message " + string('A'+i),
		}); err != nil {
			log.Fatalf("client.RouteChat: stream.Send failed: %v", err)
		}
	}
	stream3.CloseSend()
	<-waitc

}
