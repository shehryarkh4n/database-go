package actions

import (
	"bytes"
	"database-go/constants"
	"database-go/datatypes"
)

func treeInsert(tree *datatypes.BTree, node datatypes.BNode, key []byte, val []byte) datatypes.BNode {
	new := datatypes.BNode(make([]byte, 2*constants.BTREE_PAGE_SIZE))

	idx := datatypes.NodeLookupLessEqual(node, key) // insert key here

	switch node.Btype() {
	case datatypes.BNODE_LEAF:
		if bytes.Equal(key, node.GetKey(idx)) {
			leafUpdate(new, node, idx, key, val)
		} else {
			leafInsert(new, node, idx+1, key, val)
		}

	case datatypes.BNODE_NODE:
		nodeInsert(tree, new, node, idx, key, val)

	default:
		panic("TreeInsert bad!")
	}

	return new
}

func nodeInsert(tree *datatypes.BTree, new, node datatypes.BNode, idx uint16, key []byte, val []byte) {
	panic("unimplemented")
}
