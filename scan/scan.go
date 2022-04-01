package scan

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/spf13/cobra"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Scanner struct {
	startTime time.Time
	success   uint64
	failed    uint64
	Callback  func(entry os.DirEntry)
}

func (scanner *Scanner) prettyPrint() {
	fmt.Print("\033[H\033[2J")
	strSuccess := color.GreenString(strconv.Itoa(int(scanner.success)))
	strFailed := color.RedString(strconv.Itoa(int(scanner.failed)))
	strElapsed := color.YellowString(scanner.getElapsed().String())
	msg := fmt.Sprintf("%s:%s\n%s :%s\n%s:%s", color.WhiteString("success"), strSuccess, color.WhiteString("failed"), strFailed, color.WhiteString("elapsed"), strElapsed)
	fmt.Println(msg)
}

func (scanner *Scanner) getElapsed() time.Duration {
	return time.Since(scanner.startTime)
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
	scanner.startTime = time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go scanner.scan(path, &wg, quite)
	wg.Wait()
	scanner.prettyPrint()
}
