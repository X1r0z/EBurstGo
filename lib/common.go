package lib

import (
	"os"
	"sync"
)

var ExchangeUrls = map[string]string{
	"autodiscover": "/autodiscover",
	"ews":          "/ews",
	"mapi":         "/mapi",
	"activesync":   "/Microsoft-Server-ActiveSync",
	"oab":          "/oab/global.asax",
	"rpc":          "/rpc",
	"owa":          "/owa/auth.owa",
	"powershell":   "/powershell",
	"ecp":          "/owa/auth.owa",
}

var Log *Logging

type TaskInfo struct {
	targetUrl string
	mode      string
	u         string
	domain    string
	task      chan []string
	done      *DoneMap
	delay     int
	proxy     string
	o         string
}

type DoneMap struct {
	mu   sync.RWMutex
	done map[string]string
}

func (c *DoneMap) Get(user string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.done[user]
	return ok
}

func (c *DoneMap) Set(user string, pass string, o string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.done[user] = pass
	if o != "" {
		fp, _ := os.OpenFile(o, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_SYNC, 0644)
		defer fp.Close()
		fp.WriteString(user + ":" + pass)
		fp.WriteString("\n")
	}
}

type BruteWorker func(info *TaskInfo)
