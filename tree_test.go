package llrb

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewTree[int, int]()
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

	if k, v := tree.Min(); k != 0 && v != 0 {
		t.Errorf("failed to delete min, got key %v, value %v", k, v)
	}
	if k, v := tree.DeleteMin(); k != 0 && v != 0 {
		t.Errorf("failed to delete min, got key %v, value %v", k, v)
	}

	if k, v := tree.Max(); k != 9 && v != 81 {
		t.Errorf("failed to delete max, got key %v, value %v", k, v)
	}
	if k, v := tree.DeleteMax(); k != 9 && v != 81 {
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
