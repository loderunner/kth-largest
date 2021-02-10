package main

type BinaryTree struct {
	data        int
	left, right *BinaryTree
}

func (b *BinaryTree) Push(i int) {
	// Item already in tree
	if i == b.data {
		return
	}

	// Lower than item, insert in left subtree
	if i < b.data {
		// Create left subtree
		if b.left == nil {
			b.left = &BinaryTree{data: i}
			return
		}

		// Recurse into left subtree
		b.left.Push(i)
		return
	}

	// Greater than item, insert in right subtree
	if i > b.data {
		// Create right subtree
		if b.right == nil {
			b.right = &BinaryTree{data: i}
			return
		}

		// Recurse into right subtree
		b.right.Push(i)
		return
	}
}

func (b *BinaryTree) FindKthLargest(k int) int {
	kthLargest, k := b.doFindKthLargest(k)
	if k != -1 {
		panic("k larger than tree size!")
	}

	return kthLargest.data
}

func (b *BinaryTree) doFindKthLargest(k int) (*BinaryTree, int) {
	// Check right subtree
	var current *BinaryTree
	if b.right != nil {
		current, k = b.right.doFindKthLargest(k)
	}

	// -1 marker means "found"
	if k == -1 {
		return current, -1
	}

	// Current item is largest in search
	if k == 0 {
		return b, -1
	}

	// Check left subtree
	if b.left != nil {
		return b.left.doFindKthLargest(k - 1)
	}

	// Backtrack by 1
	return b, k - 1
}
