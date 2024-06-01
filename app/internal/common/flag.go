package common

import "flag"

var (
	ConfigPathParam *string
)

func InitFlag() {
	ConfigPathParam = flag.String("conf", "conf.json", "config file path of elix")

	flag.Parse()
}
