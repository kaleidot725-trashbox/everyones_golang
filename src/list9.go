package list9

import (
	logrus "github.com/sirupsen/logrus"
)

fun main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.Info("Succeeded")
	logrus.Warn("not correct")
	logrus.Error("something error")
	logrus.Fatal("panic")	
}