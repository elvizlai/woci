/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/13 21:22
 */

package ci

import (
	"time"

	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

var prefix = "/17mei"

func Etcd() {
	CMD(FMT("docker run -d --net=test --name %s-etcd %s --listen-client-urls 'http://:2379,http://:4001' --advertise-client-urls 'http://:2379,http://:4001'", conf.Tracer, conf.EtcdImage))

	<-time.After(time.Second * 5)

	// initital pgsql
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/pgsql/host %s", conf.Tracer, prefix, conf.Tracer+"-pgsql"))
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/pgsql/port %s", conf.Tracer, prefix, "5432"))
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/pgsql/name %s", conf.Tracer, prefix, "meidb"))
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/pgsql/user %s", conf.Tracer, prefix, "postgres"))
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/pgsql/password %s", conf.Tracer, prefix, "wothing"))

	// initial redis
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/redis/host %s", conf.Tracer, prefix, conf.Tracer+"-redis"))
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/redis/port %s", conf.Tracer, prefix, "6379"))
	CMD(FMT("docker exec %s-etcd /etcdctl set %s/redis/password %s", conf.Tracer, prefix, `""`))

	// initial nsq
	log.Tinfo(conf.Tracer, "etcd ok")
}
