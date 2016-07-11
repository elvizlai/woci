/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/30 08:10
 */

package ci

import (
	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/util/cmd"
	"github.com/wothing/woci/util/log"
)

func After() {
	for _, f := range conf.Config.After {
		data, err := cmd.TCMD("AFTER", f)
		if err != nil {
			log.TErrorORFatal(conf.Config.TRACER, data)
		}
	}
}
