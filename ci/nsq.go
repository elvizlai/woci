/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/15 11:44
 */

package ci

import (
	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func Nsq() {
	data, err := CMD(FMT("docker run -d --net=test --name %s-nsqd %s /nsqd", conf.Tracer, conf.NsqImage))
	if err != nil {
		log.Terrorf(conf.Tracer, data)
		log.Tfatal(conf.Tracer, err)
	}
	log.Tinfo(conf.Tracer, "nsq ok")
}
