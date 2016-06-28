/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/04/23 10:44
 */

package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pborman/uuid"
	"github.com/wothing/log"
)

var (
	Tracer      string
	Concurrent  int
	REPO        string
	ProjectPath string // This is a absolute PATH
	SQLDir      string
	PGImage     string
	RedisImage  string
	EtcdImage   string
	NsqImage    string
	Services    []Service
)

type Service struct {
	Name string
	Path string
	Para string
}

func ParseConfig(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Tfatalf(Tracer, "read %s error: %v", configFile, err)
	}

	cm := make(map[string]interface{})
	err = json.Unmarshal(data, &cm)
	if err != nil {
		log.Tfatalf(Tracer, "unmarshal %s error: %v", configFile, err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Tfatalf(Tracer, "%s panic--> %v", configFile, r)
		}
	}()

	Concurrent = int(cm["Concurrent"].(float64))
	REPO = cm["REPO"].(string)
	ProjectPath = cm["ProjectPath"].(string) // This is a absolute PATH
	SQLDir = cm["SQLDir"].(string)

	PGImage = cm["PGImage"].(string)
	RedisImage = cm["RedisImage"].(string)
	EtcdImage = cm["EtcdImage"].(string)
	NsqImage = cm["NsqImage"].(string)

	services := cm["Services"].([]interface{})
	for _, v := range services {
		s := Service{
			Name: v.(map[string]interface{})["Name"].(string),
			Path: v.(map[string]interface{})["Path"].(string),
			Para: v.(map[string]interface{})["Para"].(string),
		}
		Services = append(Services, s)
	}
}

func GenUUID() {
	Tracer = uuid.New()[:8]
	err := ioutil.WriteFile("/tracer", []byte(Tracer), os.ModeTemporary)
	if err != nil {
		log.Fatal(err)
	}
}

func RestoreUUID() {
	data, err := ioutil.ReadFile("/tracer")
	if err != nil {
		log.Fatal("can not do this, read file '/tracer' failed")
	}
	Tracer = string(data)
}
