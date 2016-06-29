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

	"github.com/wothing/log"

	"github.com/wothing/woci/conf"
)

var FMT = fmt.Sprintf

func CMD(order string) (string, error) {
	log.Tdebugf(conf.Tracer, "CMD --> %s", order)

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
		log.Tdebugf(conf.Tracer, "%s --> %s\n, STDERR --> %s\n, STDOUT -- > %s\n", order, err.Error(), stderr.String(), stdout.String())
		if stderr.String() == "" {
			return stdout.String(), err
		}
		return stderr.String(), err
	}
	return stdout.String(), nil
}
