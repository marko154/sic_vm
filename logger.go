package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	logfile, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	Log = logrus.New()
	Log.SetOutput(logfile)
}
