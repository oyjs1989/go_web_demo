package bptree

const maxlen = 16

type BPItem struct {
	// Key is the key of the item.
	Key int64
	Val interface{} // 一般是一个指针
}

type BPNode struct {
	// Items is the items of the node.
	Nodes []*BPNode
	Items []BPItem
	Next  *BPNode
}

type BPTree struct {
	// Root is the root of the tree.
	root *BPNode
}

func NewBPTree() *BPTree {
	return &BPTree{}
}

func (t *BPTree) Insert(key int64, val interface{}) {
	if t.root == nil {
		t.root = &BPNode{Items: []BPItem{{Key: key, Val: val}}}
		return
	}
	t.root.Insert(key, val)
}

func (n *BPNode) Insert(key int64, val interface{}) {
	if n.Nodes == nil {
		n.insertItem(key, val)
		return
	}
	for i, item := range n.Items {
		if key < item.Key {
			n.Nodes[i].Insert(key, val)
			return
		}
	}
	n.Nodes[len(n.Items)].Insert(key, val)
}

func (n *BPNode) insertItem(key int64, val interface{}) {

}
