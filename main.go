package main

import (
	"fmt"
	"sync"
	"net"
)

import (
	"context"
	"google.golang.org/grpc"
)

type Blog struct {
	PostID int32
    Title string
    Content string
    Author string
    PublicationDate string
    Tags []string
}

type DB struct {
	Lock sync.Mutex
	IDsUsed int32 // initialized to 0
	Blogs map[int32]Blog
}

// create a new DB
func NewDB() *DB {
	return &DB {
		IDsUsed: 0,
		Blogs: make(map[int32]Blog),
	}
}

// b.PostID is redundant
// returns id of the newly inserted blog
func (db *DB) Create(b Blog) int32 {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	b.PostID = db.IDsUsed
	db.Blogs[db.IDsUsed] = b
	db.IDsUsed++;
	return (int32)(len(db.Blogs)-1)
}

// returns blog if found and bool (set to true if found)
func (db *DB) Read(PostID int32) (Blog, bool) {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	b, ok := db.Blogs[PostID]
	return b, ok
}

// returns true, if update successfull
func (db *DB) Update(b Blog) bool {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if _, ok := db.Blogs[b.PostID]; !ok {
		return false
	}
	db.Blogs[b.PostID] = b
	return true
}

// returns true, if delete successfull
func (db *DB) Delete(PostID int32) bool {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if _, ok := db.Blogs[PostID]; !ok {
		return false
	}
	delete(db.Blogs, PostID)
	return true
}

type BloggingService struct {
	db *DB
}

func (s *BloggingService) Create(ctx context.Context, params *CreateParams) (*CreateResult, error) {
	b := Blog {
		Title: params.Title,
		Content: params.Content,
		Author: params.Author,
		PublicationDate: params.PublicationDate,
		Tags: params.Tags,
	}
	PostID := s.db.Create(b)
	return &CreateResult{PostID: PostID, Error: ""}, nil
}

func (s *BloggingService) Read(ctx context.Context, params *ReadParams) (*ReadResult, error) {
	b, ok := s.db.Read(params.PostID)
	if(ok) {
		return &ReadResult{
			PostID: b.PostID,
			Title: b.Title,
			Content: b.Content,
			Author: b.Author,
			PublicationDate: b.PublicationDate,
			Tags: b.Tags,
			Error: "",
		}, nil
	}
	return &ReadResult{Error: "BLOG NOT FOUND"}, nil
}

func (s *BloggingService) Update(ctx context.Context, params *UpdateParams) (*UpdateResult, error) {
	b := Blog {
		PostID: params.PostID,
		Title: params.Title,
		Content: params.Content,
		Author: params.Author,
		PublicationDate: params.PublicationDate,
		Tags: params.Tags,
	}
	ok := s.db.Update(b)
	if(ok) {
		return &UpdateResult{
			PostID: b.PostID,
			Title: b.Title,
			Content: b.Content,
			Author: b.Author,
			PublicationDate: b.PublicationDate,
			Tags: b.Tags,
			Error: "",
		}, nil
	}
	return &UpdateResult{Error: "BLOG NOT FOUND"}, nil
}

func (s *BloggingService) Delete(ctx context.Context, params *DeleteParams) (*DeleteResult, error) {
	ok := s.db.Delete(params.PostID)
	if(ok) {
		return &DeleteResult{Error: ""}, nil
	}
	return &DeleteResult{Error: "BLOG NOT FOUND"}, nil
}

func (s *BloggingService) mustEmbedUnimplementedBloggingServiceServer() {
}

var port int32 = 6959

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on localhost:%d", port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterBloggingServiceServer(grpcServer, &BloggingService{db: NewDB()})
	grpcServer.Serve(lis)
}