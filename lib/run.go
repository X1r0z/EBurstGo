package lib

import (
	"net/url"
	"sync"
	"time"
)

func BruteRunner(targetUrl string, mode string, domain string, dict [][]string, t int, delay int, proxy string, o string, usePth bool, worker BruteWorker) {

	authPath := ExchangeUrls[mode]
	u, _ := url.JoinPath(targetUrl, authPath)
	Log.Info("[*] 使用 %v 接口爆破: %v", mode, targetUrl)

	t1 := time.Now()

	task := make(chan []string, len(dict))
	done := &DoneMap{done: make(map[string]string)}

	info := &TaskInfo{
		targetUrl: targetUrl,
		mode:      mode,
		u:         u,
		domain:    domain,
		task:      task,
		done:      done,
		delay:     delay,
		proxy:     proxy,
		o:         o,
		usePth:    usePth,
	}

	for _, data := range dict {
		task <- data
	}

	close(task)

	var wg sync.WaitGroup

	for i := 0; i < t; i++ {
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
