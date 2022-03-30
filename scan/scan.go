package main

import (
	"fmt"
	"sync"
	"time"
)

func scan(path string, queue chan string, wg *sync.WaitGroup) {
	fmt.Printf("job started for：%s\n", path)
	wg2 := sync.WaitGroup{}
	if path != "fuck" {
		for _, d := range []string{"1", "2", "3", "4", "fuck"} {
			wg2.Add(1)
			go func(text string) {
				time.Sleep(time.Second * 5)
				queue <- text
				wg2.Done()
			}(d)
		}
	}

	fmt.Println("job ended")
	wg2.Wait()
	wg.Done()
}

func worker(queue chan string, wg *sync.WaitGroup) {
	fmt.Println("worker started")

	wg2 := sync.WaitGroup{}
	for e := range queue {
		wg2.Add(1)
		go scan(e, queue, wg)
	}
	fmt.Println("worker ended")
	wg2.Wait()
	wg.Done()
}

func dispatch(path string) {
	fmt.Println("")
	queue := make(chan string, 2)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go scan(path, queue, &wg)

	wg.Add(1)
	go worker(queue, &wg)

	wg.Wait()
}

func main() {
	dispatch("./starter")
}
