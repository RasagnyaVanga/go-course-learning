package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "gRPC_activity/routeguide"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

// GetFeature returns the feature at the given point.
func (s *routeGuideServer) GetFeature(_ context.Context, point *pb.Point) (*pb.Feature, error) {
	//Normally we would check the feature of respective point in a slice of features.
	//To cutdown the logic complexity and focus on gRPC functionality
	//returning a dummy feature name for now.
	return &pb.Feature{Name: "Dummy Feature", Location: point}, nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream grpc.ServerStreamingServer[pb.Feature]) error {
	// Normally, we would check a slice of features and send to the stream
	// only those whose locations fall within the rectangle. Here, for simplicity,
	// we are just sending 3 features to the stream without any checks.
	for i := 0; i < 3; i++ {
		feature := &pb.Feature{
			Name: "dummy feature",
			Location: &pb.Point{
				Latitude:  int32(i),
				Longitude: int32(i),
			},
		}
		if err := stream.Send(feature); err != nil { //sending into stream
			return err
		}

	}
	return nil
}

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points,  number of known features visited, total distance traveled, and
// total time spent.
func (s *routeGuideServer) RecordRoute(stream grpc.ClientStreamingServer[pb.Point, pb.RouteSummary]) error {
	var pointCount, featureCount int32
	startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     100,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		log.Printf("RecordRoute received: %v,%v", point.Latitude, point.Longitude)
		pointCount++
		//check if the point has feature in features slice and add featurecount
		featureCount++
	}
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *routeGuideServer) RouteChat(stream grpc.BidiStreamingServer[pb.RouteNote, pb.RouteNote]) error {
	for {
		in, err := stream.Recv() //RouteNote received
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("RouteNote received from client: %d %d", in.Location.Latitude, in.Location.Longitude)
		reply := &pb.RouteNote{
			Message:  "Ack: " + in.Message,
			Location: in.Location,
		}

		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen ")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})

	log.Println("Server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
