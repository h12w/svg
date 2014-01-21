// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"fmt"
)

func c(err error) {
	if err != nil {
		panic(err)
	}
}

func p(v ...interface{}) {
	fmt.Println(v...)
}
