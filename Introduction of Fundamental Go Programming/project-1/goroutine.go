package main

import (
	"fmt"
	"sync"
	"time"
	"strings"
)

func randomize(data interface{}, index int) {
	fmt.Printf("%v %d\n", data, index)
}

func sequential(data interface{}, index int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	fmt.Printf("%v %d\n", data, index)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}

	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	for i := 0; i < 4; i++ {
		go randomize(data1, i+1)
		go randomize(data2, i+1)
	}
	
	time.Sleep(time.Second * 2)
	fmt.Println("\n", strings.Repeat("#", 30), "\n")
	
	for i := 0; i < 4; i++ {
		wg.Add(2)
		go sequential(data1, i+1, mutex, &wg)
		go sequential(data2, i+1, mutex, &wg)
		wg.Wait()
	}
	
}