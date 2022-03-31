package scan

import (
	"fmt"
	_ "github.com/spf13/cobra"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Scanner struct {
	success     uint64
	failed      uint64
	elapsed     time.Duration
	Callback    func(entry os.DirEntry)
	PrettyPrint func(success uint64, failed uint64, elapse time.Duration)
}

func (scanner *Scanner) scan(path string, group *sync.WaitGroup, quite bool) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		atomic.AddUint64(&scanner.failed, 1)
	}

	wg := sync.WaitGroup{}
	for _, d := range dirs {
		wg.Add(1)
		go func() {
			if scanner.Callback != nil {
				scanner.Callback(d)
			}
			wg.Done()
		}()
		newPath := path + "/" + d.Name()
		if d.IsDir() {
			group.Add(1)
			go scanner.scan(newPath, group, quite)
		} else {
			atomic.AddUint64(&scanner.success, 1)
			if !quite {
				fmt.Println(newPath)
			}
		}
	}
	wg.Wait()
	group.Done()
}

func (scanner *Scanner) Scan(path string, quite bool) {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go scanner.scan(path, &wg, quite)
	wg.Wait()
	if scanner.PrettyPrint != nil {
		scanner.PrettyPrint(scanner.success, scanner.failed, time.Since(start))
	}
}
