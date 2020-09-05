package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 12; i++ {
		id := rand.Intn(10) + 1
		if b, ok := queryCache(i); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(i); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found with id %v", id)
		time.Sleep(150 * time.Millisecond)

	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond) // a forced delay

	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false

}
