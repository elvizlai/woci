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
	"time"

	_ "github.com/lib/pq"
	"github.com/wothing/woci/filewalk"
)

func Postgres(dsn string, dir string) {
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

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("extension pgcrypto error: %v", err))
		return
	}

	files := filewalk.WalkDir(dir, "sql").FileList()
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
}
