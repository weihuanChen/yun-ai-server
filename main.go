package main

import (
	"yinglian.com/yun-ai-server/cmd"
	"yinglian.com/yun-ai-server/configs"
)

func main() {
	configs.InitCfg()
	cmd.Start()
}
