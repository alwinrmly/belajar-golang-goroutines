package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Alwin")
	pool.Put("Ramli")
	pool.Put("Sasmita")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
