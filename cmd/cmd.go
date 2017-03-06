/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/21 17:15.
 */

package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/elvizlai/woci/logger"
)

func CMD(order string) (string, error) {
	logger.Infof("run cmd-> %s", order)

	c := exec.Command("sh", "-e", "-c", order)

	var stdout bytes.Buffer
	c.Stdout = &stdout

	var stderr bytes.Buffer
	c.Stderr = &stderr

	err := c.Run()

	// TODO remove
	//fmt.Println("err",err)
	//fmt.Println("stdout",stdout.String())
	//fmt.Println("stderr",stderr.String())

	if err != nil {
		if e := stderr.String(); e != "" {
			return stdout.String(), fmt.Errorf("stderr: %s", e)
		}

		if e := stdout.String(); e != "" {
			return stdout.String(), fmt.Errorf("stdout: %s", e)
		}

		return "", err
	}

	return stdout.String(), nil
}
