package llrb

// check that the tree is a valid red-black tree
func (n *node[K, V]) check() bool {
	if n == nil {
		return true
	}
	if n.isRed() {
		if n.left.isRed() || n.right.isRed() {
			return false
		}
	}
	return n.left.check() && n.right.check()
}
