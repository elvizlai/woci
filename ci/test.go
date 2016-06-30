/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/06 09:46
 */

package ci

import (
	"strings"

	"github.com/wothing/log"
	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func Test(names ...string) {
	if len(names) == 0 {
		for _, v := range conf.Config.Test {
			data, err := TCMD("TEST", strings.Replace(v, "[PLACEHOLDER]", "", -1))
			if err != nil {
				log.Terror(conf.Config.TRACER, data)
				log.Tfatal(conf.Config.TRACER, err)
			}
			log.Tinfo(conf.Config.TRACER, data)
		}
	} else {
		for _, name := range names {
			for _, v := range conf.Config.Test {
				data, err := TCMD("TEST", strings.Replace(v, "[PLACEHOLDER]", name, -1))
				if err != nil {
					log.Terror(conf.Config.TRACER, data)
					log.Tfatal(conf.Config.TRACER, err)
				}
				log.Tinfo(conf.Config.TRACER, data)
			}
		}
	}
}
