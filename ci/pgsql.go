/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/04 21:03
 */

package ci

import (
	"database/sql"
	"io/ioutil"
	"time"

	_ "github.com/lib/pq"
	"github.com/wothing/log"

	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/filewalk"
)

// DBStage main entry point of this module
func Pgsql() {
	CMD(FMT("docker run -d --net=test --name %s-pgsql -e POSTGRES_DB=meidb -e POSTGRES_PASSWORD=wothing %s", conf.Tracer, conf.PGImage))
	pgInit(filewalk.WalkDir(conf.ProjectPath+"/"+conf.SQLDir, "sql").FileList()...)
}

func pgInit(files ...string) {
	dsn := FMT("postgres://postgres:wothing@%s-pgsql.test:5432/meidb?sslmode=disable", conf.Tracer)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Tfatalf(conf.Tracer, "error on connecting to db: %s", err)
	}
	defer db.Close()

	for i := 0; ; i++ {
		if i > 30 {
			log.Tfatal("After for a long time we can't connection to database")
		}
		if db.Ping() != nil {
			log.Tdebugf(conf.Tracer, "Try connection to database %d time(s)", i+1)
			time.Sleep(time.Second)
		} else {
			log.Tdebug(conf.Tracer, "connection to postgresql success")
			break
		}
	}

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`)
	if err != nil {
		log.Tfatal(conf.Tracer, "sql error:", err)
	}

	log.Tdebug(conf.Tracer, "sql list:", files)

	for _, f := range files {
		sql, err := ioutil.ReadFile(f)
		if err != nil {
			log.Tfatalf(conf.Tracer, "reading sql file error : %s", f)
		}
		_, err = db.Exec(string(sql))
		if err != nil {
			log.Tfatal(conf.Tracer, "sql error:", err)
		}
	}
	log.Tinfo(conf.Tracer, "postgres ok")
}
