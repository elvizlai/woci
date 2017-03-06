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

	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/util/cmd"
	"github.com/wothing/woci/util/log"
)

func beforeBuild() {
	for _, v := range conf.Config.Before {
		data, err := cmd.TCMD("BEFRE", v)
		if err != nil {
			log.TErrorORFatal(conf.Config.TRACER, "%v,%v", data, err)
		}
	}
}

func Build() {
	beforeBuild()

	log.Tinfof(conf.Config.TRACER, "[BUILD] COUNT: %d, Concurrent: %d", len(conf.Config.Modules), conf.Config.Concurrent)

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
		data, err := cmd.TCMD("BUILD", j)
		if err != nil {
			log.TErrorORFatal(conf.Config.TRACER, "%v,%v", data, err)
		}
		wg.Done()
	}
}
