/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/11 09:41
 */

package log

import (
	"os"

	"github.com/wothing/log"
)

var DebugMode bool
var ForceMode bool

func Initial() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Llevel)
	if DebugMode {
		log.SetOutputLevel(log.Ldebug)
	}
}

func Debugf(tracer string, format string, args ...interface{}) {
	log.Tdebugf(tracer, format, args...)
}

func Tinfof(tracer string, format string, args ...interface{}) {
	log.Tinfof(tracer, format, args...)
}

func TErrorORFatal(tracer string, format string, args ...interface{}) {
	if ForceMode {
		log.Tfatalf(tracer, format, args...)
	} else {
		log.Terrorf(tracer, format, args...)
	}
}

func Tfatalf(tracer string, format string, args ...interface{}) {
	log.Tfatalf(tracer, format, args...)
}
