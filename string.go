package llrb

import (
	"bytes"
	"fmt"
	"io"
)

func (t *Tree[K, V]) String() string {
	return t.root.String()
}

func (n *node[K, V]) String() string {
	buf := bytes.NewBuffer(nil)
	n.format(buf, edges{}, edgeRoot)
	return buf.String()
}

func (n *node[K, V]) format(wr io.Writer, pre edges, edge edge) {
	if n == nil {
		return
	}

	{
		es := edgeSpace
		if edge == edgeRight {
			es = edgeLink
		}
		n.left.format(wr, append(pre, es), edgeLeft)
	}

	_, _ = fmt.Fprintf(wr, "%s %v: %v\n", append(pre, edge), n.key, n.value)

	{
		es := edgeSpace
		if edge == edgeLeft {
			es = edgeLink
		}
		n.right.format(wr, append(pre, es), edgeRight)
	}
}

type edges []edge

func (e edges) String() string {
	buf := make([]rune, 0, len(e))
	for _, v := range e {
		buf = append(buf, []rune(v.String())...)
	}
	return string(buf)
}

type edge uint

const (
	_ = edge(iota)
	edgeSpace
	edgeLink
	edgeRoot
	edgeLeft
	edgeRight
)

var edgeMap = map[edge]string{
	edgeSpace: `  `,
	edgeLink:  ` │`,
	edgeRoot:  `───`,
	edgeLeft:  ` ┌─`,
	edgeRight: ` └─`,
}

func (e edge) String() string {
	return edgeMap[e]
}
