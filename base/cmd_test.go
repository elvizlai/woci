/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/26 07:10
 */

package base

import (
	"fmt"
	"testing"
)

func TestCMD(t *testing.T) {
	fmt.Println(CMD("tree -L 3"))
}
