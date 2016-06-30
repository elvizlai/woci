/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/07 07:27
 */

package ci

import (
	"strings"

	"github.com/wothing/log"
	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

// do not care error
func Clean() {
	log.Tinfof(conf.Config.TRACER, "CLEAN STAGE, COUNT:%d", len(conf.Config.Modules))

	for _, m := range conf.Config.Modules {
		data, err := TCMD("CLEAN", strings.Replace(m.Clean, "[NAME]", m.Name, -1))
		if err != nil {
			log.Twarn(conf.Config.TRACER, data)
		}
	}
}
