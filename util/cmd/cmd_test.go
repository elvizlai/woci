/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/26 07:10
 */

package cmd

import (
	"fmt"
	"testing"
)

func TestTCMD(t *testing.T) {
	fmt.Println(TCMD("STAGE", "tree -L 3"))
}
