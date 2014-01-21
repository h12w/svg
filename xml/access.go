// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

func (e *Element) GetAttr(key string) *string {
	for _, a := range e.Attr {
		if a.Name.String() == key {
			return &a.Value
		}
	}
	return nil
}

func (e *Element) SetAttr(key, value string) {
	set := false
	for i, a := range e.Attr {
		if a.Name.String() == key {
			e.Attr[i].Value = value
			set = true
		}
	}
	if !set {
		e.Attr.Add(key, value)
	}
}

func (nodes *Nodes) Add(node Node) {
	*nodes = append(*nodes, node)
}

func (attrs *Attrs) Add(key, value string) {
	*attrs = append(*attrs, Attr{Name{Local: key}, value})
}

func AddChild(parent, child Node) {
	switch n := parent.(type) {
		case *Document:
			n.Child.Add(child)
		case *Element:
			n.Child.Add(child)
	}
}
