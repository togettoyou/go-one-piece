package util

import (
	"go-one-piece/util/conf"
	"go-one-piece/util/logging"
)

func Reset() {
	conf.Reset()
	logging.Reset()
}
