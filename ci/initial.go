/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/29 22:45
 */

package ci

import (
	"github.com/wothing/log"
	. "github.com/wothing/woci/base"
	"github.com/wothing/woci/conf"
)

func Initial() {
	for _, v := range conf.Config.Initial {
		data, err := TCMD("INITIAL", v)
		if err != nil {
			log.Fatal(data, "\n", err)
		}
		log.Debug(data)
	}
}
