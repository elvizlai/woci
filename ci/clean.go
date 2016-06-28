/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/07 07:27
 */

package ci

import (
	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

// all blew do not care error

func DataClean() {
	log.Tinfo(conf.Tracer, "data clean")
	CMD(FMT("docker stop %s-etcd", conf.Tracer))
	CMD(FMT("docker stop %s-redis", conf.Tracer))
	CMD(FMT("docker stop %s-pgsql", conf.Tracer))
	CMD(FMT("docker stop %s-nsqd", conf.Tracer))

	CMD(FMT("docker rm %s-etcd", conf.Tracer))
	CMD(FMT("docker rm %s-redis", conf.Tracer))
	CMD(FMT("docker rm %s-pgsql", conf.Tracer))
	CMD(FMT("docker rm %s-nsqd", conf.Tracer))
}

func AppClean() {
	log.Tinfof(conf.Tracer, "app clean, count:%d", len(conf.Services))
	for _, s := range conf.Services {
		CMD(FMT("docker stop %s-%s", conf.Tracer, s.Name))
		CMD(FMT("docker rm %s-%s", conf.Tracer, s.Name))
	}
}
