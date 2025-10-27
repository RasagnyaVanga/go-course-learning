package main

import (
	"context"
	pb "grpc_crud/crud"
	dbpackage "grpc_crud/database"
	"log"
	"net"

	"google.golang.org/grpc"
)

type operationsImplementor struct {
	blogManager dbpackage.BlogManager //dependency
	pb.UnimplementedOperationServer
}

// Constructor function for DI
func NewOperationsImplementor(blogManager dbpackage.BlogManager) *operationsImplementor {
	return &operationsImplementor{blogManager: blogManager}
}

func (o *operationsImplementor) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	res, err := o.blogManager.CreatePost(req.Title, req.Author, req.Content)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: int32(res.ID)}, nil
}

func (o *operationsImplementor) ReadAll(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	var posts []dbpackage.Post
	posts, err := o.blogManager.ReadAllPosts()
	if err != nil {
		return nil, err
	}
	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, &pb.Post{
			Id:      int32(p.ID),
			Title:   p.Title,
			Content: p.Content,
			Author:  p.Author,
		})
	}
	return &pb.ReadResponse{Posts: pbPosts}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server running on :50051")

	grpcServer := grpc.NewServer()

	post, err := dbpackage.NewBlogDB()
	if err != nil {
		return
	}
	srv := NewOperationsImplementor(post)

	pb.RegisterOperationServer(grpcServer, srv)

	log.Println("Server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
