package main

import (
	"fmt"
	"sync"
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
	Blogs []Blog
}

// PostID - redundant
func (db *DB) Create(b Blog) int32 {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	db.Blogs = append(db.Blogs, b)
	db.Blogs[len(db.Blogs)-1].PostID = (int32)(len(db.Blogs)-1)
	return (int32)(len(db.Blogs)-1)
}

// returns blog found and bool (set to true if found)
func (db *DB) Read(PostID int32) (Blog, bool) {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if(PostID < (int32)(len(db.Blogs))) {
		return db.Blogs[PostID], true
	}
	return Blog{}, false
}

func main() {
	fmt.Println("Hello World");
}