// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

/*
TODO:
1. rename the package to xmltree
2. 
*/

import (
	"fmt"
)

type Node interface {
	fmt.Stringer
}

type Nodes []Node

type Document struct {
	Child Nodes
}

type Element struct {
	Name  Name
	Attr  Attrs
	Child Nodes
}

type Attr struct {
	Name  Name
	Value string
}

type Attrs []Attr

type Name struct {
	Space, Local string
}

type bytesNode struct {
	Bytes []byte
}

type (
	CharData   bytesNode
	Comment    bytesNode
	Dirivative bytesNode
)

type ProcInst struct {
	Target string
	Inst   []byte
}

