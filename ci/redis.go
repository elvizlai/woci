/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/04/23 14:05
 */

package ci

import (
	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func Redis() {
	data, err := CMD(FMT("docker run -d --net=test --name %s-redis %s", conf.Tracer, conf.RedisImage))
	if err != nil {
		log.Terrorf(conf.Tracer, data)
		log.Tfatal(conf.Tracer, err)
	}
	log.Tinfo(conf.Tracer, "redis ok")
}
