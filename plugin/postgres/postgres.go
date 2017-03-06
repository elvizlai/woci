/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/21 23:50.
 */

package postgres

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/elvizlai/woci/plugin/filewalk"
)

func Postgres(dsn string, para ...string) {
	l := len(para)
	if l == 2 {
		exec(dsn, para[0], para[1])
	} else {
		os.Stderr.WriteString("plugin 'postgres' need three params\n")
		os.Exit(1)
	}
}

func exec(dsn string, cmd string, para string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error on connecting to '%s': %v\n", dsn, err))
		os.Exit(1)
		return
	}
	defer db.Close()

	for i := 0; ; i++ {
		if i > 30 {
			os.Stderr.WriteString(fmt.Sprintf("after 30s, still can't connection to '%s'\n", dsn))
			os.Exit(1)
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
			os.Stderr.WriteString(fmt.Sprintf("CMD '%s' error: %v\n", para, err))
			os.Exit(1)
		} else {
			os.Stdout.WriteString(fmt.Sprintf("'%s' ok\n", para))
		}
	case "file":
		files := filewalk.WalkDir(para, "sql").FileList()

		os.Stdout.WriteString("file list: \n")
		for _, f := range files {
			os.Stdout.WriteString(f + "\n")
		}

		for _, f := range files {
			_sql, err := ioutil.ReadFile(f)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("read '%s' error: %v\n", f, err))
				os.Exit(1)
				return
			}
			_, err = db.Exec(string(_sql))
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("file '%s' error: %v\n", f, err))
				os.Exit(1)
				return
			}
		}
	default:
		os.Stderr.WriteString(fmt.Sprintf("'%s' not acceptable, only accept CMD or FILE now!\n", cmd))
		os.Exit(1)
	}
}
