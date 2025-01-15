// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"strings"
	"testing"
	"text/template/parse"
)

type NoopBlock struct {
}

func (n *NoopBlock) Builtins() map[string]any {
	return nil
}

func (n *NoopBlock) Parse(tree *parse.Tree, pos parse.Pos, line int, pipe *parse.PipeNode, children *parse.ListNode) parse.Node {
	log.Println("Noop Template: ", children.String())
	return children
}

func TestNoopBlock(t *testing.T) {
	contents := `{{ noop . }} This is a basic noop block that basically returns the child subtree as is.  {{ if .Flag }} hello {{ else }} nohello {{ end }}{{ end }}`

	bm := parse.BlockMap{
		"noop": &NoopBlock{},
	}
	tmpl := Must(New("test").Blocks(bm).Parse(contents))
	buf := &strings.Builder{}
	want := " This is a basic noop block that basically returns the child subtree as is.   hello "
	if err := tmpl.Execute(buf, map[string]any{"Flag": true}); err != nil {
		t.Error("Expected: ", want, "Got Error: ", err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("expected\n%q\ngot\n%q", want, got)
	}
}
