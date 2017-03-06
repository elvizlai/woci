/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/23 16:14.
 */

package ci

import (
	"github.com/elvizlai/woci/cmd"
	"github.com/elvizlai/woci/config"
	"github.com/elvizlai/woci/logger"
)

func CLEAN() {
	for _, v := range config.GetConfig().CLEAN {
		out, err := cmd.CMD(v)
		if out != "" {
			logger.Debugf("%s", out)
		}
		if err != nil {
			logger.Errorf("CLEAN ERR: %v", err)
		}
	}
}
