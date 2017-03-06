/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/29 22:45
 */

package ci

import (
	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/util/cmd"
	"github.com/wothing/woci/util/log"
)

func Initial() {
	for _, v := range conf.Config.Initial {
		data, err := cmd.TCMD("INITL", v)
		if err != nil {
			log.TErrorORFatal(conf.Config.TRACER, "%v,%v", data, err)
		}
		log.Debugf(conf.Config.TRACER, data)
	}
}
