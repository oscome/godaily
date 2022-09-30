package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	cron := cron.New()
	cron.AddFunc("* * * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})

	cron.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05", // 定义日期时间格式
			FullTimestamp:   true,
			DisableColors:   true,
		})
		logrus.Info("game over")
		cron.Stop()
	}
}
