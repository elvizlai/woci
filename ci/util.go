/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/22 07:12.
 */

package ci

import (
	"fmt"
	"strings"
)

func getValue(x []interface{}, alias ...string) []string {
	data := []string{}
	if len(alias) == 0 {
		for _, v := range x {
			switch w := v.(type) {
			case map[interface{}]interface{}:
				for k, v := range w {
					data = append(data, strings.Replace(fmt.Sprint(v), "[KEY]", fmt.Sprint(k), -1))
				}
			default:
				data = append(data, fmt.Sprint(w))
			}
		}
	} else {
		for i := range alias {
			for _, v := range x {
				switch v := v.(type) {
				case map[interface{}]interface{}:
					if value, ok := v[alias[i]]; ok {
						data = append(data, strings.Replace(fmt.Sprint(value), "[KEY]", alias[i], -1))
						break
					}
				default:
					data = append(data, fmt.Sprint(v))
				}
			}
		}
	}
	return data
}
