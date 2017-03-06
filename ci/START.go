/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/22 07:58.
 */

package ci

import (
	"github.com/elvizlai/woci/cmd"
	"github.com/elvizlai/woci/config"
	"github.com/elvizlai/woci/logger"
)

func START(alias ...string) {
	cmds := getValue(config.GetConfig().START, alias...)
	logger.Infof("start count: %d", len(cmds))
	for _, v := range cmds {
		out, err := cmd.CMD(v)
		if out != "" {
			logger.Debugf("%s", out)
		}
		if err != nil {
			logger.Errorf("BUILD ERR: %v", err)
		}
	}
}
