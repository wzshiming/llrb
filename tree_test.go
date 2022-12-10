package llrb

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

func TestTree(t *testing.T) {
	tree := NewTree[int, int]()

	_, _, ok := tree.Min()
	if ok {
		t.Errorf("the tree is not empty")
	}

	_, _, ok = tree.Max()
	if ok {
		t.Errorf("the tree is not empty")
	}

	_, _, ok = tree.DeleteMin()
	if ok {
		t.Errorf("the tree is not empty")
	}

	_, _, ok = tree.DeleteMax()
	if ok {
		t.Errorf("the tree is not empty")
	}

	_, ok = tree.Delete(0)
	if ok {
		t.Errorf("the tree is not empty")
	}

	for i := 0; i != 5; i++ {
		tree.Put(i, i*i)
	}
	for _, number := range []int{0, 1, 2, 3, 4, 5, 9, 8, 7, 6} {
		tree.Put(number, number*number)
	}

	for _, number := range []int{0, 1, 5, 8, 9} {
		got, found := tree.Get(number)
		if !found {
			t.Errorf("failed to find %d", number)
		}
		want := number * number
		if got != want {
			t.Errorf("value is %d should be %d", got, want)
		}
	}
	for _, number := range []int{-1, -21, 10, 11, 148} {
		_, found := tree.Get(number)
		if found {
			t.Errorf("should not have found %d", number)
		}
	}

	if k, v, ok := tree.Min(); !ok || k != 0 || v != 0 {
		t.Errorf("failed to delete min, got key %v, value %v", k, v)
	}
	if k, v, ok := tree.DeleteMin(); !ok || k != 0 || v != 0 {
		t.Errorf("failed to delete min, got key %v, value %v", k, v)
	}

	if k, v, ok := tree.Max(); !ok || k != 9 || v != 81 {
		t.Errorf("failed to delete max, got key %v, value %v", k, v)
	}
	if k, v, ok := tree.DeleteMax(); !ok || k != 9 || v != 81 {
		t.Errorf("failed to delete max, got key %v, value %v", k, v)
	}
	length := 7
	for i, number := range []int{1, 5, 8} {
		if v, deleted := tree.Delete(number); !deleted && v != i*i {
			t.Errorf("failed to delete %d", number)
		}
		if tree.Len() != length-i {
			t.Errorf("map len %d should be %d", tree.Len(), length-i)
		}
	}
	for _, number := range []int{-1, -21, 10, 11, 148} {
		if _, deleted := tree.Delete(number); deleted {
			t.Errorf("should not have deleted nonexistent %d", number)
		}
	}
	if tree.Len() != 5 {
		t.Errorf("map len %d should be 5", tree.Len())
	}

	wantKeys := []int{2, 3, 4, 6, 7}
	gotKeys := tree.Keys()
	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("keys %v should be %v", gotKeys, wantKeys)
	}

	wantValues := []int{4, 9, 16, 36, 49}
	gotValues := tree.Values()
	if !reflect.DeepEqual(gotValues, wantValues) {
		t.Errorf("values %v should be %v", gotValues, wantValues)
	}
}

func BenchmarkFind(b *testing.B) {
	const base = 1000000
	intRBTree := NewTree[int, int]()
	for i := 0; i < base; i++ {
		intRBTree.Put(i, i)
	}

	b.Run("success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			intRBTree.Get(i % base)
		}
	})
	b.Run("failure", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			intRBTree.Get(i % base / 2)
		}
	})
}

func FuzzTree(f *testing.F) {
	for i := 0; i < 1000; i++ {
		k := strconv.FormatInt(int64(rand.Int()), 10)
		f.Add(k)
	}

	tree := NewTree[string, string]()
	f.Fuzz(func(t *testing.T, key string) {
		value := key

		got, ok := tree.Get(key)
		if ok {
			_, ok := tree.Delete(key)
			if !ok {
				t.Errorf("failed to delete %s", key)
			}
			if got != value {
				t.Errorf("value is %s should be %s", got, value)
			}
		} else {
			ok := tree.Put(key, value)
			if !ok {
				t.Errorf("failed to put %s", key)
			}

			got, found := tree.Get(key)
			if !found {
				t.Errorf("failed to find %s", key)
			}
			if got != value {
				t.Errorf("value is %s should be %s", got, value)
			}

			if rand.Int()%2 == 0 {
				_, ok := tree.Delete(key)
				if !ok {
					t.Errorf("failed to delete %s", key)
				}
			}
		}

		keys := tree.Keys()
		if !slices.IsSorted(keys) {
			t.Errorf("keys are not sorted")
		}

		if !tree.root.check() {
			t.Errorf("tree is not valid")
		}
	})
}
