/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/29 22:53
 */

package ci

import (
	"strings"

	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/util/cmd"
	"github.com/wothing/woci/util/log"
)

func Start() {
	log.Tinfof(conf.Config.TRACER, "START STAGE, COUNT:%d", len(conf.Config.Modules))

	for _, m := range conf.Config.Modules {
		data, err := cmd.TCMD("START", strings.Replace(m.Start, "[NAME]", m.Name, -1))
		if err != nil {
			log.TErrorORFatal(conf.Config.TRACER, "%v,%v", data, err)
		}
	}
}
