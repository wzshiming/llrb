package llrb

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tree := NewTree[int, int]()
	for i := 0; i != 10; i++ {
		tree.Put(i, i*i)
	}

	want := `
     ┌─ 0: 0
   ┌─ 1: 1
   │ └─ 2: 4
─── 3: 9
   │   ┌─ 4: 16
   │ ┌─ 5: 25
   │ │ └─ 6: 36
   └─ 7: 49
     │ ┌─ 8: 64
     └─ 9: 81
`
	want = want[1:]
	got := tree.String()
	if got != want {
		t.Errorf("the output is not expected")
		fmt.Printf("got:\n%s\nwant:\n%s\n", got, want)
	}
}
