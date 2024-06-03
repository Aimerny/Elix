package common

import "flag"

var (
	ConfigPathParam     *string
	UpgradeOngeDatabase *bool
)

func InitFlag() {
	ConfigPathParam = flag.String("conf", "conf.json", "config file path of elix")
	UpgradeOngeDatabase = flag.Bool("db-update", false, "update database of onge service")
	flag.Parse()
}
