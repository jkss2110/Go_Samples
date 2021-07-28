package controller

import (
	"dummy_golang/models"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]models.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func ConnectingBooks() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	cachech := make(chan models.Book)
	dbch := make(chan models.Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(1)
		// First go
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex, cachech chan<- models.Book) {
			b, ok := queryCache(id, m)
			if ok {
				cachech <- b
			}
			wg.Done()
		}(id, wg, m, cachech)
		wg.Add(1)
		// Second go
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex, dbch chan<- models.Book) {
			if b, ok := queryDatabase(id, m); ok {
				dbch <- b
			}
			wg.Done()
		}(id, wg, m, dbch)

		go func(cachech <-chan models.Book, dbch <-chan models.Book) {
			select {
			case b := <-cachech:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbch // to come out of error when go routine slept
			case b := <-dbch:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cachech, dbch)
		wg.Wait()
	}
}

func queryCache(id int, m *sync.Mutex) (models.Book, bool) {
	m.Lock()
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}

func queryDatabase(id int, m *sync.Mutex) (models.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range models.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return models.Book{}, false
}
