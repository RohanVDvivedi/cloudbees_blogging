package main

import (
	"fmt"
	"sync"
)

import (
	"context"
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

func main() {
	//s := BloggingService{db: NewDB()}
	fmt.Println("Hello World");
}