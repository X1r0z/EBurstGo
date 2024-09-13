package common

import (
	"flag"
	"github.com/fatih/color"
	"os"
)

func Parse(config *Config) {
	ParseArgs(config)
	ParseConfig(config)
}

func ParseArgs(config *Config) {
	flag.StringVar(&config.Url, "url", "", "Exchange url")
	flag.StringVar(&config.Mode, "mode", "", "Exchange endpoint")
	flag.BoolVar(&config.Check, "check", false, "check Exchange endpoints")
	flag.StringVar(&config.Domain, "domain", "", "AD domain")
	flag.StringVar(&config.Username, "user", "", "username")
	flag.StringVar(&config.Password, "pass", "", "password")
	flag.StringVar(&config.UserFile, "userf", "", "username file")
	flag.StringVar(&config.PassFile, "passf", "", "password file")
	flag.StringVar(&config.UserPassFile, "userpassf", "", "user-pass file (user:pass)")
	flag.BoolVar(&config.UserAsPass, "user-as-pass", false, "password is same as username")
	flag.BoolVar(&config.PthMode, "pth", false, "PTH mode (Pass The Hash)")
	flag.StringVar(&config.Proxy, "proxy", "", "socks/http(s) proxy")
	flag.StringVar(&config.OutputFile, "o", "result.txt", "output file")
	flag.IntVar(&config.Threads, "t", 2, "number of goroutines")
	flag.IntVar(&config.Delay, "delay", 0, "delay between requests")
	flag.BoolVar(&config.Verbose, "v", false, "show verbose info")
	flag.BoolVar(&config.NoColor, "nocolor", false, "disable output color")
	flag.Parse()
}

func ParseConfig(config *Config) {
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	color.NoColor = config.NoColor

	if config.UserPassFile != "" {
		UserPassDict = LoadAllDict(config.UserPassFile)
	}

	if config.UserFile != "" {
		UserDict = LoadDict(config.UserFile)
	}

	if config.PassFile != "" {
		PassDict = LoadDict(config.PassFile)
	}

	if config.UserAsPass {
		PassDict = UserDict
	}

	if config.Username != "" {
		UserDict = []string{config.Username}
	}

	if config.Password != "" {
		PassDict = []string{config.Password}
	}

	if len(UserPassDict) == 0 {
		for _, password := range PassDict {
			for _, username := range UserDict {
				UserPassDict[username] = password
			}
		}
	}

	Log = Logger{Verbose: config.Verbose}
}
