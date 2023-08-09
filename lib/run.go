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
	done      *DoneFlag
	delay     int
}

type DoneFlag struct {
	mu   sync.RWMutex
	done bool
}

func (c *DoneFlag) GetDone() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.done
}

func (c *DoneFlag) SetDone() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.done = true
}

type BruteWorker func(info *TaskInfo)

func BruteRunner(targetUrl string, mode string, domain string, userDict []string, passDict []string, n int, delay int, worker BruteWorker) {

	authPath := ExchangeUrls[mode]
	u, _ := url.JoinPath(targetUrl, authPath)
	Log.Info("[*] 使用 %v 接口爆破: %v", mode, targetUrl)
	Log.Info("[*] 用户名:%v 密码:%v 共计:%v", len(userDict), len(passDict), len(userDict)*len(passDict))

	t1 := time.Now()

	for _, username := range userDict {

		task := make(chan []string, len(passDict))
		done := &DoneFlag{done: false}

		info := &TaskInfo{
			targetUrl: targetUrl,
			mode:      mode,
			u:         u,
			domain:    domain,
			task:      task,
			done:      done,
			delay:     delay,
		}

		for _, password := range passDict {
			data := []string{username, password}
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
	}

	t2 := time.Now()
	Log.Info("[*] 耗时: %v", t2.Sub(t1))
}
