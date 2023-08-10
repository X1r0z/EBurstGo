package lib

import (
	"net/url"
	"sync"
	"time"
)

type TaskInfo struct {
	targetUrl string
	mode      string
	u         string
	domain    string
	task      chan []string
	done      *DoneMap
	delay     int
}

type DoneMap struct {
	mu      sync.RWMutex
	done    map[string]struct{}
	allDone bool
}

func (c *DoneMap) Get(user string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.done[user]
	return ok
}

func (c *DoneMap) Set(user string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.done[user] = struct{}{}
}

type BruteWorker func(info *TaskInfo)

func BruteRunner(targetUrl string, mode string, domain string, dict [][]string, n int, delay int, worker BruteWorker) {

	authPath := ExchangeUrls[mode]
	u, _ := url.JoinPath(targetUrl, authPath)
	Log.Info("[*] 使用 %v 接口爆破: %v", mode, targetUrl)

	t1 := time.Now()

	task := make(chan []string, len(dict))
	done := &DoneMap{done: make(map[string]struct{})}

	info := &TaskInfo{
		targetUrl: targetUrl,
		mode:      mode,
		u:         u,
		domain:    domain,
		task:      task,
		done:      done,
		delay:     delay,
	}

	for _, data := range dict {
		task <- data
	}

	close(task)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(info)
		}()
	}

	wg.Wait()

	t2 := time.Now()
	Log.Info("[*] 耗时: %v", t2.Sub(t1))
}
