/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/06 09:48
 */

package base

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/wothing/log"
	"github.com/wothing/woci/conf"
)

var FMT = fmt.Sprintf

func TCMD(stage, order string) (string, error) {
	order = strings.Replace(order, "[TRACER]", conf.Config.TRACER, -1)
	log.Tinfof(conf.Config.TRACER, "%s: %s", stage, order)

	cmd := exec.Command("bash")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	in.WriteString(order)

	err := cmd.Run()
	if err != nil {
		if stderr.String() == "" {
			return stdout.String(), err
		}
		return stderr.String(), err
	}
	return stdout.String(), nil
}
