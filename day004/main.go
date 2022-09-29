package main

import "github.com/sirupsen/logrus"

// 建议全局统一格式，然后如果需要收集建议设置为 json 格式，方便收集到 ES 等日志系统
func main() {
	logrus.Info("------TextFormatter--------") // 默认格式
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 定义日期时间格式
		FullTimestamp:   true,
		DisableColors:   true,
	})
	logrus.Info("oscome info log")
	logrus.Debug("oscome debug log")

	logrus.WithFields(logrus.Fields{
		"name": "test",
	}).Infof("to do %v", "log")

	logrus.Info("------JSONFormatter--------")

	formatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 定义日期时间格式
		DataKey:         "test",                // key
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	}
	logrus.SetFormatter(formatter)
	logrus.SetReportCaller(true)      // 打印 log 产生的位置
	logrus.SetLevel(logrus.InfoLevel) // debug
	logrus.Info("oscome info log")
	logrus.Debug("oscome debug log")

	logrus.WithFields(logrus.Fields{
		"name": "test",
	}).Infof("to do %v", "log")
}
