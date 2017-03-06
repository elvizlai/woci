/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/29 21:40
 */

package plugin

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/wothing/woci/util/filewalk"
)

func Postgres(dsn string, para ...string) {
	l := len(para)
	if l == 2 {
		exec(dsn, para[0], para[1])
	} else {
		os.Stderr.WriteString("plugin 'postgres' need three params, check")
	}
}

func exec(dsn string, cmd string, para string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error on connecting to '%s': %v", dsn, err))
		return
	}
	defer db.Close()

	for i := 0; ; i++ {
		if i > 30 {
			os.Stderr.WriteString(fmt.Sprintf("after 30s we can't connection to '%s'", dsn))
			return
		}
		if db.Ping() != nil {
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	switch strings.ToLower(cmd) {
	case "cmd":
		_, err = db.Exec(para)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%s '%s' error: %v", cmd, para, err))
		} else {
			os.Stdout.WriteString(fmt.Sprintf("'%s' done", para))
		}
	case "file":
		files := filewalk.WalkDir(para, "sql").FileList()
		os.Stdout.WriteString(fmt.Sprintf("sql list: %v\n", files))

		for _, f := range files {
			sql, err := ioutil.ReadFile(f)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("read '%s' error: %v", f, err))
				return
			}
			_, err = db.Exec(string(sql))
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("exec '%s' error: %v", f, err))
				return
			}
		}
	default:
		os.Stderr.WriteString(fmt.Sprintf("'%s' not acceptable, only accept CMD and FILE now!", cmd))
	}
}
