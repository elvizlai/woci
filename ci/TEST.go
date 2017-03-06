/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/22 09:28.
 */

package ci

import (
	"os"
	"strings"

	"github.com/elvizlai/woci/cmd"
	"github.com/elvizlai/woci/config"
	"github.com/elvizlai/woci/logger"
)

func TEST(placeHolder ...string) {
	if len(placeHolder) == 0 {
		for _, v := range getValue(config.GetConfig().TEST) {
			v = strings.Replace(v, "[PLACEHOLDER]", "", -1)
			out, err := cmd.CMD(v)
			os.Stdout.WriteString(out)
			if err != nil {
				logger.Errorf("TEST FAILED")
			}
		}
	} else {
		for _, x := range placeHolder {
			for _, y := range getValue(config.GetConfig().TEST) {
				y = strings.Replace(y, "[PLACEHOLDER]", x, -1)
				out, err := cmd.CMD(y)
				os.Stdout.WriteString(out)
				if err != nil {
					logger.Errorf("TEST FAILED")
				}
			}
		}
	}
}
