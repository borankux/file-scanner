package scan

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Scanner struct {
	startTime time.Time
	success   uint64
	failed    uint64
	dirs      uint64
	Callback  func(fullpath string, entry os.DirEntry)
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
	defer group.Done()
	dirs, err := os.ReadDir(path)
	if err != nil {
		atomic.AddUint64(&scanner.failed, 1)
		e := strings.Split(err.Error(), ":")[1]
		if e != " too many open files" {
			fmt.Println(color.RedString("[%s]", e))
		}
		return
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()
	for _, d := range dirs {
		atomic.AddUint64(&scanner.dirs, 1)
		wg.Add(1)
		newPath := path + "/" + d.Name()
		absPath, _ := filepath.Abs(newPath)
		go func(waitGroup *sync.WaitGroup) {
			if scanner.Callback != nil {
				scanner.Callback(absPath, d)
			}
			waitGroup.Done()
		}(&wg)
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
}

func (scanner *Scanner) Scan(path string, quite bool) {
	scanner.startTime = time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go scanner.scan(path, &wg, quite)

	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			scanner.prettyPrint()
		}
	}()
	wg.Wait()
	scanner.prettyPrint()
}
