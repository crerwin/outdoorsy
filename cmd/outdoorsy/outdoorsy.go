package main

import (
	"github.com/crerwin/outdoorsy/pkg/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	odsFormatter := new(log.TextFormatter)
	odsFormatter.TimestampFormat = "2006-01-02 15:04:05"
	odsFormatter.FullTimestamp = true
	log.SetFormatter(odsFormatter)

	log.SetLevel(log.InfoLevel)

	cmd.Execute()
}
