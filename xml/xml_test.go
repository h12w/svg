// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"os"
	"testing"
)

const svg = `<?xml version="1.0" encoding="UTF-8"?><svg width="400" height="300"></svg>`

func Test_new(t *testing.T) {
	f, err := os.Open("man.svg")
	c(err)
	defer f.Close()
	tree, err := Parse(f)
	c(err)
	p(tree)
}

