package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

var counter uint64

func scan(path string, group *sync.WaitGroup) {
	dirs, _ := os.ReadDir(path)
	for _, d := range dirs {
		newPath := path + "/" + d.Name()
		if d.IsDir() {
			group.Add(1)
			go scan(newPath, group)
		} else {
			atomic.AddUint64(&counter, 1)
			fmt.Println(newPath)
		}
	}
	group.Done()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("provide path")
		return
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	now := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go scan(absPath, &wg)
	wg.Wait()
	diff := time.Since(now)
	fmt.Printf("%d files, scanned at:%v\n", counter, diff)
}
