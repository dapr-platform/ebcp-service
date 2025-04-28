package config

import (
	"os"
)

var CLOSE_WEEK_DAY = "1"

func init() {
	if os.Getenv("CLOSE_WEEK_DAY") != "" {
		CLOSE_WEEK_DAY = os.Getenv("CLOSE_WEEK_DAY")
	}
}
