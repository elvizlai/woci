/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/06 09:46
 */

package ci

import (
	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func AppTest() {
	data, err := CMD(FMT("CGO_ENABLED=0 TestEnv=CI CiTracer=%s go test -v %s/gateway/tests/*.go", conf.Tracer, conf.ProjectPath))
	if err != nil {
		log.Terror(conf.Tracer, data)
		log.Tfatal(conf.Tracer, err)
	}
	log.Tinfo(conf.Tracer, data)
}
