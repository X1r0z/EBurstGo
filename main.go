package main

import (
	"EBurstGo/brute"
	"EBurstGo/common"
)

func main() {
	var config common.Config
	common.Parse(&config)
	brute.Run(&config)
}
