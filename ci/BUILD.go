/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/22 07:08.
 */

package ci

import (
	"github.com/elvizlai/woci/cmd"
	"github.com/elvizlai/woci/config"
	"github.com/elvizlai/woci/logger"
)

func BUILD(alias ...string) {
	cmds := getValue(config.GetConfig().BUILD, alias...)
	logger.Infof("build count: %d", len(cmds))
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
