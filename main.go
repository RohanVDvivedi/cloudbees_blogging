package main

import (
	"fmt"
	"net"
)

import (
	"context"
	"google.golang.org/grpc"
)

import (
	"cloudbees_blogging/pb"
	"cloudbees_blogging/db"
)

type BloggingService struct {
	pb.UnimplementedBloggingServiceServer
	db *db.DB
}

func (s *BloggingService) Create(ctx context.Context, params *pb.CreateParams) (*pb.CreateResult, error) {
	b := db.Blog {
		Title: params.Title,
		Content: params.Content,
		Author: params.Author,
		PublicationDate: params.PublicationDate,
		Tags: params.Tags,
	}
	PostID := s.db.Create(b)
	return &pb.CreateResult{PostID: PostID, Error: ""}, nil
}

func (s *BloggingService) Read(ctx context.Context, params *pb.ReadParams) (*pb.ReadResult, error) {
	b, ok := s.db.Read(params.PostID)
	if(ok) {
		return &pb.ReadResult{
			PostID: b.PostID,
			Title: b.Title,
			Content: b.Content,
			Author: b.Author,
			PublicationDate: b.PublicationDate,
			Tags: b.Tags,
			Error: "",
		}, nil
	}
	return &pb.ReadResult{Error: "BLOG NOT FOUND"}, nil
}

func (s *BloggingService) Update(ctx context.Context, params *pb.UpdateParams) (*pb.UpdateResult, error) {
	b := db.Blog {
		PostID: params.PostID,
		Title: params.Title,
		Content: params.Content,
		Author: params.Author,
		PublicationDate: params.PublicationDate,
		Tags: params.Tags,
	}
	ok := s.db.Update(b)
	if(ok) {
		return &pb.UpdateResult{
			PostID: b.PostID,
			Title: b.Title,
			Content: b.Content,
			Author: b.Author,
			PublicationDate: b.PublicationDate,
			Tags: b.Tags,
			Error: "",
		}, nil
	}
	return &pb.UpdateResult{Error: "BLOG NOT FOUND"}, nil
}

func (s *BloggingService) Delete(ctx context.Context, params *pb.DeleteParams) (*pb.DeleteResult, error) {
	ok := s.db.Delete(params.PostID)
	if(ok) {
		return &pb.DeleteResult{Error: ""}, nil
	}
	return &pb.DeleteResult{Error: "BLOG NOT FOUND"}, nil
}

var port int32 = 6959

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on localhost:%d\n", port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBloggingServiceServer(grpcServer, &BloggingService{db: db.NewDB()})
	grpcServer.Serve(lis)
}