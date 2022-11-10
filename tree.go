package llrb

type Tree[K Ordered, V Any] struct {
	root   *node[K, V]
	length int
}

func NewTree[K Ordered, V Any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

func (t *Tree[K, V]) Put(k K, v V) (ok bool) {
	t.root, ok = t.root.Put(k, v)
	t.root.red = false
	if ok {
		t.length++
	}
	return ok
}

func (t *Tree[K, V]) Get(k K) (v V, ok bool) {
	return t.root.Get(k)
}

func (t *Tree[K, V]) Range(fn func(K, V) bool) {
	t.root.Range(fn)
}

func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, 0, t.length)
	t.Range(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

func (t *Tree[K, V]) Values() []V {
	values := make([]V, 0, t.length)
	t.root.Range(func(k K, v V) bool {
		values = append(values, v)
		return true
	})
	return values
}

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

func (t *Tree[K, V]) Len() int {
	return t.length
}

func (t *Tree[K, V]) Delete(k K) (v V, ok bool) {
	if t.root == nil {
		return v, false
	}
	var deleted *node[K, V]
	t.root, deleted = t.root.delete(k)
	if t.root != nil {
		t.root.red = false
	}
	if deleted == nil {
		return v, false
	}
	t.length--
	return deleted.value, true
}

func (t *Tree[K, V]) DeleteMin() (k K, v V, ok bool) {
	if t.root == nil {
		return k, v, false
	}
	var deleted *node[K, V]
	t.root, deleted = t.root.deleteMin()
	if deleted != nil {
		k = deleted.key
		v = deleted.value
	}
	t.length--
	return k, v, true
}

func (t *Tree[K, V]) DeleteMax() (k K, v V, ok bool) {
	if t.root == nil {
		return k, v, false
	}
	var deleted *node[K, V]
	t.root, deleted = t.root.deleteMax()
	if deleted != nil {
		k = deleted.key
		v = deleted.value
	}
	t.length--
	return k, v, true
}
