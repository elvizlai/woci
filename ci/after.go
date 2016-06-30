/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/30 08:10
 */

package ci

import (
	"github.com/wothing/log"
	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func After() {
	for _, f := range conf.Config.After {
		data, err := TCMD("AFTER", f)
		if err != nil {
			log.Twarn(conf.Config.TRACER, data)
		}
	}
}
