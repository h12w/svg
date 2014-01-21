// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"fmt"
	"strings"
)

// TODO: use writer instead of Stringer
func (d Document) String() string {
	return d.Child.String()
}

func (ns Nodes) String() string {
	ss := make([]string, len(ns))
	for i := range ss {
		ss[i] = ns[i].String()
	}
	return strings.Join(ss, "")
}

func (e Element) String() string {
	attr := ""
	if len(e.Attr) > 0 {
		attr = " " + e.Attr.String()
	}

	if len(e.Child) > 0 {
		return fmt.Sprintf(`<%v%s>%s</%v>`, e.Name, attr, e.Child, e.Name)
	}
	return fmt.Sprintf(`<%v%s/>`, e.Name, attr)
}

func (as Attrs) String() string {
	ss := make([]string, len(as))
	for i := range ss {
		ss[i] = as[i].String()
	}
	return strings.Join(ss, " ")
}

func (a Attr) String() string {
	return fmt.Sprintf(`%v="%s"`, a.Name, a.Value)
}

func (c Comment) String() string {
	return fmt.Sprintf(`<!--%s-->`, string(c.Bytes))
}

func (d Dirivative) String() string {
	return fmt.Sprintf(`<!%s>`, string(d.Bytes))
}

func (pi ProcInst) String() string {
	return fmt.Sprintf(`<?%s %s?>`, pi.Target, string(pi.Inst))
}

func (d CharData) String() string {
	return string(d.Bytes)
}

func (n Name) String() string {
	// TODO: consider name.Space
	return escape([]byte(n.Local))
}

