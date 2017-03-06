/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/21 15:24.
 */

package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"

	"github.com/elvizlai/woci/logger"
)

type conf struct {
	TRACER     string
	CONCURRENT int
	ENV        map[string]string
	INIT       []string
	BEFORE     []string
	BUILD      []interface{}
	START      []interface{}
	TEST       []interface{}
	SUCCEED    []string
	FAILED     []string
	CLEAN      []string
}

var c conf

func Initialize(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Fatalf("read config failed: %v", err)
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		logger.Fatalf("parse config failed: %v", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		logger.Fatalf("unmarshal config failed: %v", err)
	}

	dataStr := string(data)
	for i, j := 0, len(c.ENV); i < j; i++ {
		for k, v := range c.ENV {
			dataStr = strings.Replace(dataStr, "$"+strings.ToUpper(k), fmt.Sprint(v), -1)
		}
	}

	err = viper.ReadConfig(bytes.NewBufferString(dataStr))
	if err != nil {
		logger.Fatalf("parse config failed: %v", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		logger.Fatalf("unmarshal config failed: %v", err)
	}
}

func GetConfig() conf {
	return c
}
