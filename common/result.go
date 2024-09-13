package common

import (
	"os"
	"sync"
)

type Result struct {
	Creds      map[string]string
	OutputFile string
	mutex      sync.RWMutex
}

func (r *Result) Get(username string) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, ok := r.Creds[username]
	return ok
}

func (r *Result) Put(username string, password string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.Creds[username] = password

	if r.OutputFile != "" {
		fp, _ := os.OpenFile(r.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_SYNC, 0644)
		defer fp.Close()

		fp.WriteString(username + ":" + password + "\n")
	}
}
