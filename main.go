package main

import (
	"github.com/schidstorm/sshd-puller/cli"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := cli.Run()
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}
