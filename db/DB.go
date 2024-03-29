package db

import (
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
	IDsUsed int32 // initialized to 1
	Blogs map[int32]Blog
}

// create a new DB
func NewDB() *DB {
	return &DB {
		IDsUsed: 1,
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
	return b.PostID
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