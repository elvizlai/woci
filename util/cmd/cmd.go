/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/06 09:48
 */

package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/util/log"
)

func TCMD(stage, order string, args ...interface{}) (string, error) {
	order = strings.Replace(fmt.Sprintf(order, args...), "[TRACER]", conf.Config.TRACER, -1)
	log.Tinfof(conf.Config.TRACER, "[%-5s] %s", stage, order)

	c := exec.Command("bash")

	var stdout bytes.Buffer
	c.Stdout = &stdout

	var stderr bytes.Buffer
	c.Stderr = &stderr

	in := bytes.NewBuffer(nil)
	c.Stdin = in
	in.WriteString(order)

	err := c.Run()
	if err != nil {
		if stderr.String() == "" {
			return stdout.String(), err
		}
		return stderr.String(), err
	}
	return stdout.String(), nil
}
