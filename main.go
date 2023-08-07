package main

import (
	"EBurstGo/lib"
	"flag"
	"fmt"
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
	)
	flag.StringVar(&targetUrl, "url", "", "Exchange 服务器地址")
	flag.StringVar(&mode, "mode", "", "指定 Exchange 服务器接口")
	flag.BoolVar(&check, "check", false, "检查目标 Exchange 服务器可用接口")
	flag.StringVar(&domain, "domain", "", "AD 域名")
	flag.StringVar(&user, "user", "", "用户名字典")
	flag.StringVar(&pass, "pass", "", "密码字典")
	flag.IntVar(&n, "thread", 2, "协程数量")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if check {
		if targetUrl != "" {
			lib.Check(targetUrl)
		} else {
			fmt.Println("[-] Exchange 服务器地址为空")
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
			lib.NtlmBruteRun(targetUrl, mode, domain, userDict, passDict, n)
		case "activesync":
			lib.BasicBruteRun(targetUrl, mode, domain, userDict, passDict, n)
		case "owa", "ecp":
			lib.HttpBruteRun(targetUrl, mode, domain, userDict, passDict, n)
		case "powershell":
			lib.KerberosBruteRun(targetUrl, mode, domain, userDict, passDict, n)
		default:
			fmt.Println("[-] Exchange 接口无效")
		}
	}
}
