/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/9/21 17:18.
 */

package cmd

import "testing"

func TestCMD(t *testing.T) {
	_, err := CMD("xxoo")
	if err == nil {
		t.Fatal("err should not be nil")
	} else {
		t.Log(err)
	}

	_, err = CMD("docker ls")
	if err == nil {
		t.Fatal("err should not be nil")
	} else {
		t.Log(err)
	}

	_, err = CMD("docker ls | true && docker ps")
	if err != nil {
		t.Fatal(err)
	}

	_, err = CMD("docker ps")
	if err != nil {
		t.Fatal(err)
	}

	_, err = CMD("docker ps -a")
	if err != nil {
		t.Fatal(err)
	}

}
