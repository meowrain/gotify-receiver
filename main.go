package main

import (
	_ "gotify-client/pkg/logger"

	"gotify-client/cmd/client"
	"time"
)

func main() {
	_, _ = time.LoadLocation("Asia/Shanghai")
	client.Main()
}
