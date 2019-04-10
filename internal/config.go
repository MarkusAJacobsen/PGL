package internal

import (
	"flag"
	"strconv"
)

const Root = "/var/pgl/logs"
const RootDebug = "./logs"

func getLogPath() (p string) {
	d := getDebug()
	if d {
		p = RootDebug
	} else {
		p = Root
	}
	return p
}

func getDebug() (d bool) {
	d, err := strconv.ParseBool(flag.Lookup("d").Value.String())
	if err != nil {
		panic(err)
	}
	return
}
