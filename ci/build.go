/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/29 22:53
 */

package ci

import (
	"strings"
	"sync"

	"github.com/wothing/log"
	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func beforeBuild() {
	for _, v := range conf.Config.Before {
		data, err := TCMD("BEFORE", v)
		if err != nil {
			log.Terrorf(conf.Config.TRACER, data)
			log.Tfatal(conf.Config.TRACER, err)
		}
	}
}

func Build() {
	beforeBuild()

	log.Tinfof(conf.Config.TRACER, "BUILD STAGE, COUNT:%d", len(conf.Config.Modules))

	jobCount := len(conf.Config.Modules)
	jobs := make(chan string, jobCount)

	wg := &sync.WaitGroup{}
	wg.Add(jobCount)

	for i, j := 0, conf.Config.Concurrent; i < j; i++ {
		go builder(wg, jobs)
	}

	//add jobs
	for _, m := range conf.Config.Modules {
		jobs <- strings.Replace(m.Build, "[NAME]", m.Name, -1)
	}

	wg.Wait()
}

func builder(wg *sync.WaitGroup, jobs <-chan string) {
	for j := range jobs {
		data, err := TCMD("BUILD", j)
		if err != nil {
			log.Terrorf(conf.Config.TRACER, data)
			log.Tfatal(conf.Config.TRACER, err)
		}
		wg.Done()
	}
}
