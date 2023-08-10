package main

import (
	"EBurstGo/lib"
	"flag"
	"github.com/fatih/color"
	"io"
	"os"
	"strings"
)

func main() {

	var (
		targetUrl  string
		mode       string
		check      bool
		domain     string
		user       string
		pass       string
		userf      string
		passf      string
		userpassf  string
		userAsPass bool
		proxy      string
		t          int
		v          bool
		delay      int
		debug      bool
		nocolor    bool
		o          string
		nosave     bool
	)
	flag.StringVar(&targetUrl, "url", "", "Exchange 服务器地址")
	flag.StringVar(&mode, "mode", "", "指定 Exchange Web 接口")
	flag.BoolVar(&check, "check", false, "检测目标 Exchange 可用接口")
	flag.StringVar(&domain, "domain", "", "AD 域名")
	flag.StringVar(&user, "user", "", "指定用户名")
	flag.StringVar(&pass, "pass", "", "指定密码")
	flag.StringVar(&userf, "userf", "", "用户名字典")
	flag.StringVar(&passf, "passf", "", "密码字典")
	flag.StringVar(&userpassf, "userpassf", "", "指定用户名密码字典 (user:pass)")
	flag.BoolVar(&userAsPass, "user-as-pass", false, "指定密码与用户名相同")
	flag.StringVar(&proxy, "proxy", "", "指定 socks/http(s) 代理")
	flag.StringVar(&o, "o", "result.txt", "指定结果输出文件")
	flag.BoolVar(&nosave, "nosave", false, "不将结果输出至文件")
	flag.IntVar(&t, "t", 2, "协程数量")
	flag.IntVar(&delay, "delay", 0, "请求延时")
	flag.BoolVar(&v, "v", false, "显示详细信息")
	flag.BoolVar(&debug, "debug", false, "显示 Debug 信息")
	flag.BoolVar(&nocolor, "nocolor", false, "关闭输出颜色")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	if nocolor {
		color.NoColor = true
	}

	if nosave {
		o = ""
	}

	lib.Log = &lib.Logging{Verbose: v, IsDebug: debug}

	var dict [][]string

	if check {
		lib.Check(targetUrl)
	} else {
		if userpassf != "" {
			fp, _ := os.Open(userpassf)
			defer fp.Close()
			b, _ := io.ReadAll(fp)
			for _, v := range strings.Split(string(b), "\n") {
				if v != "" {
					u, p, _ := strings.Cut(v, ":")
					dict = append(dict, []string{u, p})
				}
			}
			lib.Log.Info("[*] 用户名:密码共计:%v", len(dict))
		} else {
			var userDict []string
			var passDict []string

			if user != "" {
				userDict = []string{user}
			}

			if userf != "" {
				fp, _ := os.Open(userf)
				defer fp.Close()
				b, _ := io.ReadAll(fp)
				for _, v := range strings.Split(string(b), "\n") {
					if v != "" {
						userDict = append(userDict, v)
					}
				}
			}

			if pass != "" {
				passDict = []string{pass}
			}

			if passf != "" {
				fp, _ := os.Open(passf)
				defer fp.Close()
				b, _ := io.ReadAll(fp)
				for _, v := range strings.Split(string(b), "\n") {
					if v != "" {
						passDict = append(passDict, v)
					}
				}
			}

			for _, u := range userDict {
				if userAsPass {
					dict = append(dict, []string{u, u})
				} else {
					for _, p := range passDict {
						dict = append(dict, []string{u, p})
					}
				}
			}

			if userAsPass {
				lib.Log.Info("[*] 用户名:%v 密码:%v 共计:%v", len(userDict), len(userDict), len(dict))
			} else {
				lib.Log.Info("[*] 用户名:%v 密码:%v 共计:%v", len(userDict), len(passDict), len(dict))
			}
		}

		var worker lib.BruteWorker

		switch mode {
		case "autodiscover", "ews", "mapi", "oab", "rpc":
			worker = lib.NtlmBruteWorker
		case "activesync":
			worker = lib.BasicBruteWorker
		case "owa", "ecp":
			worker = lib.HttpBruteWorker
		case "powershell":
			worker = lib.KerberosBruteWorker
		default:
			lib.Log.Failed("[-] Exchange 接口无效")
			return
		}

		lib.BruteRunner(targetUrl, mode, domain, dict, t, delay, proxy, o, worker)
	}
}
