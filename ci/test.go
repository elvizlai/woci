/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/06 09:46
 */

package ci

import (
	"strings"

	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/util/cmd"
	"github.com/wothing/woci/util/log"
)

func Test(names ...string) {
	if len(names) == 0 {
		for _, v := range conf.Config.Test {
			data, err := cmd.TCMD("TEEST", strings.Replace(v, "[PLACEHOLDER]", "", -1))
			if err != nil {
				log.TErrorORFatal(conf.Config.TRACER, "%v,%v", data, err)
			}
			log.Tinfof(conf.Config.TRACER, data)
		}
	} else {
		for _, name := range names {
			for _, v := range conf.Config.Test {
				data, err := cmd.TCMD("TEEST", strings.Replace(v, "[PLACEHOLDER]", name, -1))
				if err != nil {
					log.TErrorORFatal(conf.Config.TRACER, "%v,%v", data, err)
				}
				log.Tinfof(conf.Config.TRACER, data)
			}
		}
	}
}
