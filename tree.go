package llrb

// Tree is a left-leaning red-black tree.
type Tree[K Ordered, V Any] struct {
	root   *node[K, V]
	length int
}

// NewTree returns a new Tree.
func NewTree[K Ordered, V Any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Put inserts a key-value pair into the tree.
func (t *Tree[K, V]) Put(k K, v V) (ok bool) {
	t.root, ok = t.root.Put(k, v)
	t.root.red = false
	if ok {
		t.length++
	}
	return ok
}

// Get returns the value associated with the given key.
func (t *Tree[K, V]) Get(k K) (v V, ok bool) {
	return t.root.Get(k)
}

// Range calls fn for each key-value pair in the tree in ascending order.
func (t *Tree[K, V]) Range(fn func(K, V) bool) {
	t.root.Range(fn)
}

// Keys returns a slice of all keys in the tree in ascending order.
func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, 0, t.length)
	t.Range(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

// Values returns a slice of all values in the tree in ascending order.
func (t *Tree[K, V]) Values() []V {
	values := make([]V, 0, t.length)
	t.root.Range(func(k K, v V) bool {
		values = append(values, v)
		return true
	})
	return values
}

// Min returns the minimum key-value pair in the tree.
func (t *Tree[K, V]) Min() (k K, v V, ok bool) {
	if t.root == nil {
		return k, v, false
	}
	n := t.root.Min()
	if n != nil {
		k = n.key
		v = n.value
	}
	return k, v, true
}

// Max returns the maximum key-value pair in the tree.
func (t *Tree[K, V]) Max() (k K, v V, ok bool) {
	if t.root == nil {
		return k, v, false
	}
	n := t.root.Max()
	if n != nil {
		k = n.key
		v = n.value
	}
	return k, v, true
}

// Len returns the number of key-value pairs in the tree.
func (t *Tree[K, V]) Len() int {
	return t.length
}

// Delete removes the key-value pair associated with the given key.
func (t *Tree[K, V]) Delete(k K) (v V, ok bool) {
	if t.root == nil {
		return v, false
	}
	root, deleted := t.root.delete(k)
	t.root = root
	if t.root != nil {
		t.root.red = false
	}
	if deleted == nil {
		return v, false
	}
	t.length--
	return deleted.value, true
}

// DeleteMin removes the minimum key-value pair from the tree.
func (t *Tree[K, V]) DeleteMin() (k K, v V, ok bool) {
	root, deleted := t.root.deleteMin()
	t.root = root
	if deleted != nil {
		k = deleted.key
		v = deleted.value
	}
	if deleted == nil {
		return k, v, false
	}
	t.length--
	return k, v, true
}

// DeleteMax removes the maximum key-value pair from the tree.
func (t *Tree[K, V]) DeleteMax() (k K, v V, ok bool) {
	root, deleted := t.root.deleteMax()
	t.root = root
	if deleted != nil {
		k = deleted.key
		v = deleted.value
	}
	if deleted == nil {
		return k, v, false
	}
	t.length--
	return k, v, true
}
