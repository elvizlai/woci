/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/29 22:53
 */

package ci

import (
	"strings"

	"github.com/wothing/log"
	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func Start() {
	log.Tinfof(conf.Config.TRACER, "START STAGE, COUNT:%d", len(conf.Config.Modules))

	for _, m := range conf.Config.Modules {
		data, err := TCMD("START", strings.Replace(m.Start, "[NAME]", m.Name, -1))
		if err != nil {
			log.Terrorf(conf.Config.TRACER, data)
			log.Tfatal(conf.Config.TRACER, err)
		}
	}
}
