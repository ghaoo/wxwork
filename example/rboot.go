package main

import (
	"github.com/ghaoo/rboot"
	_ "github.com/ghaoo/rboot/adapter"
	"github.com/sirupsen/logrus"
)

func main() {
	bot := rboot.New()

	pay := New()
	pay.registerPlugin()

	bot.Go()
}

func init() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}
