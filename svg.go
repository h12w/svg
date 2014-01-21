// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"container/list"
	"io"
	"strings"
	"github.com/hailiang/svg/xml"
)

func NewSvg() *Node {
	n, _ := Parse(strings.NewReader(
		`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"
 "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
</svg>`))
	return n
}

/*
&Node{
		&xml.Document{
			xml.Nodes{
				&xml.ProcInst{"xml", []byte(`version="1.0" encoding="UTF-8"`)},
				&xml.Element{Name: xml.Name{Local: "svg"}},
			},
		}}
*/

// -------------------------------------------

// Node represents a HTML node.
// Wrap html.Node so that chainable interface is possible
// Use pointer of it because we want to test with nil.
type Node struct {
	xml.Node
}

func NewNode(n xml.Node) *Node {
	if n == nil {
		return nil
	}
	return &Node{n}
}

func Parse(r io.Reader) (*Node, error) {
	n, err := xml.Parse(r)
	if err != nil {
		return nil, err
	}
	return NewNode(n), nil
}

func (n *Node) Find(cs ...Checker) *Node {
	return NewNode(Find(cs...)(n.Node))
}

func (n *Node) find(c Checker, cs []Checker) *Node {
	if n == nil {
		return nil
	}
	return n.Find(append([]Checker{c}, cs...)...)
}

func (n *Node) Svg(cs ...Checker) *Node {
	return n.find(Svg, cs)
}

func (n *Node) G(cs ...Checker) *Node {
	return n.find(G, cs)
}

func (n *Node) Attr(key string) *string {
	if n == nil {
		return nil
	}
	return GetAttr(n.Node, key)
}

func (n *Node) SetAttr(key, value string) {
	SetAttr(n.Node, key, value)
}

func (n *Node) AddChild(c *Node) {
	xml.AddChild(n.Node, c.Node)
}

// --------------------------------------------

type Checker func(xml.Node) xml.Node

func Not(c Checker) Checker {
	return func(n xml.Node) xml.Node {
		if c(n) == nil {
			return n
		}
		return nil
	}
}

func And(cs ...Checker) Checker {
	return func(n xml.Node) xml.Node {
		for _, c := range cs {
			if c(n) == nil {
				return nil
			}
		}
		return n
	}
}

func Pipe(cs ...Checker) Checker {
	return func(n xml.Node) xml.Node {
		for _, c := range cs {
			r := c(n)
			if r == nil {
				return nil
			} else {
				n = r
			}
		}
		return n
	}
}

func Or(cs ...Checker) Checker {
	return func(n xml.Node) xml.Node {
		for _, c := range cs {
			if c(n) != nil {
				return n
			}
		}
		return nil
	}
}

// --------------------------------------------

func ElementNode(n xml.Node) xml.Node {
	if _, ok := n.(*xml.Element); ok {
		return n
	}
	return nil
}

func AtomChecker(a string) Checker {
	return func(n xml.Node) xml.Node {
		if elem, ok := n.(*xml.Element); ok {
			if elem.Name.String() == a {
				return n
			}
		}
		return nil
	}
}

var (
	Svg = AtomChecker("svg")
	G   = AtomChecker("g")
)

// --------------------------------------------

func GetAttr(n xml.Node, key string) *string {
	if node, ok := n.(*xml.Element); ok {
		return node.GetAttr(key)
	}
	return nil
}

func SetAttr(n xml.Node, key, value string) {
	if node, ok := n.(*xml.Element); ok {
		node.SetAttr(key, value)
	}
}

// --------------------------------------------

func Children(n xml.Node) (nodes xml.Nodes) {
	switch node := n.(type) {
	case *xml.Element:
		nodes = node.Child
	case *xml.Document:
		nodes = node.Child
	default:
	}
	return
}

// Broad first search in all descendants
func Find(cs ...Checker) Checker {
	c := And(cs...)
	return func(n xml.Node) xml.Node {
		q := NewQueue()
		q.PushNodes(Children(n))
		for q.Len() > 0 {
			t := q.Pop()
			if c(t) != nil {
				return t
			} else {
				q.PushNodes(Children(t))
			}
		}
		return nil
	}
}

// FIFO queue.
type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Len() int {
	return q.l.Len()
}

func (q *Queue) Push(n xml.Node) {
	q.l.PushBack(n)
}

func (q *Queue) PushNodes(nodes xml.Nodes) {
	for _, node := range nodes {
		q.Push(node)
	}
}

func (q *Queue) Pop() xml.Node {
	if q.l.Front() == nil {
		return nil
	}
	return q.l.Remove(q.l.Front()).(xml.Node)
}
