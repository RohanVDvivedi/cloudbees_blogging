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

// b.PostID is redundant
// returns id of the newly inserted blog
func (db *DB) Create(b Blog) int32 {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	db.Blogs = append(db.Blogs, b)
	db.Blogs[len(db.Blogs)-1].PostID = (int32)(len(db.Blogs)-1)
	return (int32)(len(db.Blogs)-1)
}

// returns blog if found and bool (set to true if found)
func (db *DB) Read(PostID int32) (Blog, bool) {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if(PostID >= (int32)(len(db.Blogs))) {
		return Blog{}, false
	}
	return db.Blogs[PostID], true
}

// returns true, if update successfull
func (db *DB) Update(b Blog) bool {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if(b.PostID >= (int32)(len(db.Blogs))) {
		return false
	}
	db.Blogs[b.PostID] = b
	return true
}

func main() {
	fmt.Println("Hello World");
}