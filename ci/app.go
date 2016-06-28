/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/05 14:15
 */

package ci

import (
	"strings"
	"sync"

	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func AppBuild() {
	log.Tinfof(conf.Tracer, "app build, count:%d", len(conf.Services))

	data, err := CMD("make -C " + conf.ProjectPath + " idl ve")
	if err != nil {
		log.Terrorf(conf.Tracer, data)
		log.Tfatal(conf.Tracer, err)
	}

	jobCount := len(conf.Services)
	jobs := make(chan string, jobCount)

	wg := &sync.WaitGroup{}
	wg.Add(jobCount)

	for i, j := 0, conf.Concurrent; i < j; i++ {
		go builder(wg, jobs)
	}

	//add jobs
	for _, s := range conf.Services {
		jobs <- FMT("cd %s/%s && CGO_ENABLED=0 GOBIN=/app go install -ldflags '-X github.com/wothing/17mei/db._DEBUG=TRUE'", conf.ProjectPath, s.Path)
	}

	wg.Wait()
}

func builder(wg *sync.WaitGroup, jobs <-chan string) {
	for j := range jobs {
		data, err := CMD(j)
		if err != nil {
			log.Terrorf(conf.Tracer, data)
			log.Tfatal(conf.Tracer, err)
		}
		wg.Done()
	}
}

func AppStart() {
	log.Tinfof(conf.Tracer, "app start, count:%d", len(conf.Services))
	for _, s := range conf.Services {
		v := FMT("docker run -it -d --net=test -v app:/app --name %s-%s daocloud.io/sdrzlyz/alpine-ca /app/%s %s", conf.Tracer, s.Name, s.Name, s.Para)
		v = strings.Replace(v, "[TRACER]", conf.Tracer, -1)
		data, err := CMD(v)
		if err != nil {
			log.Terrorf(conf.Tracer, data)
			log.Tfatal(conf.Tracer, err)
		}
	}
}
