// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"bytes"
	"encoding/xml"
	"io"
)

func Parse(r io.Reader) (*Document, error) {
	decoder := xml.NewDecoder(r)
	doc := &Document{}
	return doc, doc.Child.parse(decoder, nil)
}

func (nodes *Nodes) parse(decoder *xml.Decoder, parent *Element) error {
TokenLoop:
	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			elem := &Element{
				Name: parseName(t.Name),
				Attr: parseAttrs(t.Attr),
			}
			nodes.Add(elem)
			elem.Child.parse(decoder, elem)
		case xml.EndElement:
			break TokenLoop
		case xml.CharData:
			nodes.Add(&CharData{t.Copy()})
		case xml.Comment:
			nodes.Add(&Comment{t.Copy()})
		case xml.Directive:
			nodes.Add(&Dirivative{t.Copy()})
		case xml.ProcInst:
			nodes.Add(&ProcInst{t.Target, t.Copy().Inst})
		default:
			panic("unknown type")
		}
	}
	return nil

}

func parseName(n xml.Name) Name {
	return Name{
		Space: n.Space,
		Local: n.Local}
}

func parseAttr(a xml.Attr) Attr {
	return Attr{
		Name:  parseName(a.Name),
		Value: a.Value,
	}
}

func parseAttrs(as []xml.Attr) Attrs {
	attrs := make([]Attr, len(as))
	for i := range attrs {
		attrs[i] = parseAttr(as[i])
	}
	return attrs
}

func escape(s []byte) string {
	var buf bytes.Buffer
	_ = xml.EscapeText(&buf, s)
	return buf.String()
}
