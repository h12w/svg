// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"fmt"
	"os"
	"svg/xml"
	"testing"
)

func Test_new(t *testing.T) {
	fmt.Printf("%v\n", NewSvg())
	fmt.Printf("%v\n", &Node{
		&xml.Document{
			xml.Nodes{
				&xml.ProcInst{"xml", []byte(`version="1.0" encoding="UTF-8"`)},
				&xml.Element{Name: xml.Name{Local: "svg"}},
			},
		}})
	return
	f, err := os.Open("xml/man.svg")
	c(err)
	defer f.Close()
	doc, err := Parse(f)
	c(err)
	doc.Svg().SetAttr("width", "987")
	//p(*doc.Svg().G().Attr("transform"))
	p(doc)
}

func c(err error) {
	if err != nil {
		panic(err)
	}
}

func p(v ...interface{}) {
	fmt.Println(v...)
}
