package lib

import (
	"crypto/tls"
	"github.com/fatih/color"
	"net/http"
	"net/url"
)

func Check(targetUrl string) {

	for k, v := range ExchangeUrls {
		u, _ := url.JoinPath(targetUrl, v)

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					Renegotiation:      tls.RenegotiateOnceAsClient,
				},
			},
		}
		res, err := client.Get(u)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 404 && res.StatusCode != 403 && res.StatusCode != 301 && res.StatusCode != 302 {
			color.Green("[+] 存在 %v 接口 (%v), 可以爆破", k, v)
		} else {
			color.Red("[-] 不存在 %v 接口 (%v)", k, v)
		}
	}
}
