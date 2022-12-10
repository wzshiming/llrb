package llrb

type node[K Ordered, V Any] struct {
	key         K
	value       V
	red         bool
	left, right *node[K, V]
}

func (n *node[K, V]) isRed() bool {
	return n != nil && n.red
}

func (n *node[K, V]) rotateLeft() *node[K, V] {
	child := n.right
	n.right = child.left
	child.left = n
	child.red = n.red
	n.red = true
	return child
}

func (n *node[K, V]) rotateRight() *node[K, V] {
	child := n.left
	n.left = child.right
	child.right = n
	child.red = n.red
	n.red = true
	return child
}

func (n *node[K, V]) moveRedLeft() *node[K, V] {
	n = n.colorFlip()
	if n.right.left.isRed() {
		n.right = n.right.rotateRight()
		n = n.rotateLeft()
		n = n.colorFlip()
	}
	return n
}

func (n *node[K, V]) moveRedRight() *node[K, V] {
	n = n.colorFlip()
	if n.left != nil && n.left.left.isRed() {
		n = n.rotateRight()
		n = n.colorFlip()
	}
	return n
}

func (n *node[K, V]) deleteMax() (*node[K, V], *node[K, V]) {
	if n == nil {
		return nil, nil
	}
	if n.left.isRed() {
		n = n.rotateRight()
	}
	if n.right == nil {
		return nil, n
	}
	if !n.right.isRed() && !n.right.left.isRed() {
		n = n.moveRedRight()
	}

	right, deleted := n.right.deleteMax()
	n.right = right

	return n.fixUp(), deleted
}

// deleteMin code for LLRB 2-3 trees
func (n *node[K, V]) deleteMin() (*node[K, V], *node[K, V]) {
	if n == nil {
		return nil, nil
	}
	if n.left == nil {
		return nil, n
	}

	if !n.left.isRed() && !n.left.left.isRed() {
		n = n.moveRedLeft()
	}

	left, deleted := n.left.deleteMin()
	n.left = left

	return n.fixUp(), deleted
}

func (n *node[K, V]) colorFlip() *node[K, V] {
	n.red = !n.red
	if n.left != nil {
		n.left.red = !n.left.red
	}
	if n.right != nil {
		n.right.red = !n.right.red
	}
	return n
}

func (n *node[K, V]) fixUp() *node[K, V] {
	if n.right.isRed() {
		n = n.rotateLeft()
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = n.rotateRight()
	}
	if n.left.isRed() && n.right.isRed() {
		n = n.colorFlip()
	}
	return n
}

func (n *node[K, V]) Range(fn func(K, V) bool) {
	if n == nil {
		return
	}
	n.left.Range(fn)
	if fn(n.key, n.value) {
		n.right.Range(fn)
	}
}

func (n *node[K, V]) Min() *node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node[K, V]) Max() *node[K, V] {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node[K, V]) Put(k K, v V) (*node[K, V], bool) {
	ok := false
	if n == nil {
		return &node[K, V]{
			key:   k,
			value: v,
			red:   true,
		}, true
	}

	switch {
	case k < n.key:
		n.left, ok = n.left.Put(k, v)
	case n.key < k:
		n.right, ok = n.right.Put(k, v)
	default:
		n.value = v
	}
	return n.fixUp(), ok
}

func (n *node[K, V]) delete(k K) (*node[K, V], *node[K, V]) {
	var deleted *node[K, V]
	if k < n.key {
		if n.left != nil {
			if !n.left.isRed() && !n.left.left.isRed() {
				n = n.moveRedLeft()
			}
			n.left, deleted = n.left.delete(k)
		}
	} else {
		if n.left.isRed() {
			n = n.rotateRight()
		}
		if !(k < n.key) && !(n.key < k) && n.right == nil {
			return nil, n
		}
		if n.right != nil {
			if !n.right.isRed() && !n.right.left.isRed() {
				n = n.moveRedRight()
			}
			if !(k < n.key) && !(n.key < k) {
				smallest := n.right.Min()
				n.key = smallest.key
				n.value = smallest.value
				n.right, deleted = n.right.deleteMin()
			} else {
				n.right, deleted = n.right.delete(k)
			}
		}
	}
	return n.fixUp(), deleted
}

func (n *node[K, V]) Get(k K) (v V, ok bool) {
	if n == nil {
		return v, false
	}
	switch {
	case k < n.key:
		return n.left.Get(k)
	case n.key < k:
		return n.right.Get(k)
	default:
		return n.value, true
	}
}
