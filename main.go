package main

import (
	"EBurstGo/lib"
	"flag"
	"io"
	"os"
	"strings"
)

func main() {

	var (
		targetUrl string
		mode      string
		check     bool
		domain    string
		user      string
		pass      string
		n         int
		v         bool
		delay     int
		debug     bool
	)
	flag.StringVar(&targetUrl, "url", "", "Exchange 服务器地址")
	flag.StringVar(&mode, "mode", "", "指定 Exchange Web 接口")
	flag.BoolVar(&check, "check", false, "检测目标 Exchange 可用接口")
	flag.StringVar(&domain, "domain", "", "AD 域名")
	flag.StringVar(&user, "user", "", "用户名字典")
	flag.StringVar(&pass, "pass", "", "密码字典")
	flag.IntVar(&n, "thread", 2, "协程数量")
	flag.IntVar(&delay, "delay", 0, "请求延时")
	flag.BoolVar(&v, "verbose", false, "显示详细信息")
	flag.BoolVar(&debug, "debug", false, "显示 Debug 信息")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	lib.Log = &lib.Logging{Verbose: v, IsDebug: debug}

	if check {
		if targetUrl != "" {
			lib.Check(targetUrl)
		} else {
			lib.Log.Failed("[-] Exchange 服务器地址为空")
		}
	} else {

		userFp, _ := os.Open(user)
		passFp, _ := os.Open(pass)

		defer userFp.Close()
		defer passFp.Close()

		userBytes, _ := io.ReadAll(userFp)
		passBytes, _ := io.ReadAll(passFp)

		userDict := strings.Split(string(userBytes), "\n")
		passDict := strings.Split(string(passBytes), "\n")
		userDict = userDict[:len(userDict)-1]
		passDict = passDict[:len(passDict)-1]

		switch mode {
		case "autodiscover", "ews", "mapi", "oab", "rpc":
			lib.NtlmBruteRun(targetUrl, mode, domain, userDict, passDict, n, delay)
		case "activesync":
			lib.BasicBruteRun(targetUrl, mode, domain, userDict, passDict, n, delay)
		case "owa", "ecp":
			lib.HttpBruteRun(targetUrl, mode, domain, userDict, passDict, n, delay)
		case "powershell":
			lib.KerberosBruteRun(targetUrl, mode, domain, userDict, passDict, n, delay)
		default:
			lib.Log.Failed("[-] Exchange 接口无效")
		}
	}
}
