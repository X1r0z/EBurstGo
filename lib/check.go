package lib

import (
	"github.com/fatih/color"
	"net/url"
)

func Check(targetUrl string) {

	for _, v := range ExchangeUrls {
		u, _ := url.JoinPath(targetUrl, v)
		res, err := Client.Get(u)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 404 && res.StatusCode != 403 && res.StatusCode != 301 && res.StatusCode != 302 {
			color.Green("[+] 存在 %v 接口, 可以爆破", u)
		} else {
			color.Red("[-] 不存在 %v 接口", u)
		}
	}
}
