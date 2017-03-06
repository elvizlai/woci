/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/22 07:09.
 */

package ci

import (
	"github.com/elvizlai/woci/cmd"
	"github.com/elvizlai/woci/config"
	"github.com/elvizlai/woci/logger"
)

func BEFORE() {
	for _, v := range config.GetConfig().BEFORE {
		out, err := cmd.CMD(v)
		if out != "" {
			logger.Debugf("%s", out)
		}
		if err != nil {
			logger.Errorf("BEFORE ERR: %v", err)
		}
	}
}
