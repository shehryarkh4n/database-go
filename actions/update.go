package actions

import (
	"database-go/constants"
	"database-go/datatypes"
	"encoding/binary"
)

func leafInsert(new datatypes.BNode, old datatypes.BNode, idx uint16, key []byte, val []byte) {

	// defining a new node, so we go through the whole struct again
	new.SetHeader(datatypes.BNODE_LEAF, old.Nkeys()+1)
	nodeAppendRange(new, old, 0, 0, idx)
	nodeAppendKV(new, idx, 0, key, val)
	nodeAppendRange(new, old, idx+1, idx, old.Nkeys()-idx)
}

func leafUpdate(new datatypes.BNode, old datatypes.BNode, idx uint16, key []byte, val []byte)

func nodeAppendKV(new datatypes.BNode, idx uint16, ptr uint64, key []byte, val []byte) {
	// set a ptr
	new.SetPtr(idx, ptr)
	//KVs
	pos := new.KvPos(idx)
	binary.LittleEndian.PutUint16(new[pos+0:], uint16(len(key)))
	binary.LittleEndian.PutUint16(new[pos+2:], uint16(len(val)))
	copy(new[pos+4:], key)
	copy(new[pos+4+uint16(len(key)):], val)

	new.SetOffset(idx+1, new.GetOffset(idx)+4+uint16(len(key)+len(val)))
}

// copy multiple KVs into the position from the old node
func nodeAppendRange(
	new datatypes.BNode, old datatypes.BNode,
	dstNew uint16, srcOld uint16, n uint16,
)

func nodeReplaceChildN(
	tree *datatypes.BTree,
	new datatypes.BNode,
	old datatypes.BNode,
	idx uint16,
	children ...datatypes.BNode,
) {
	inc := uint16(len(children))
	new.SetHeader(datatypes.BNODE_NODE, old.Nkeys()+inc-1)
	nodeAppendRange(new, old, 0, 0, idx)
	for i, node := range children {
		nodeAppendKV(new, idx+uint16(i), tree.New(node), node.GetKey(0), nil)
	}
	nodeAppendRange(new, old, idx+inc, idx+1, old.Nkeys()-(idx+1))
}

func nodeSplit2(left datatypes.BNode, right datatypes.BNode, old datatypes.BNode) {
	panic("unimplemented")
}

func nodeSplit3(old datatypes.BNode) (uint16, [3]datatypes.BNode) {
	if old.Nbytes() <= constants.BTREE_PAGE_SIZE {
		old = old[:constants.BTREE_PAGE_SIZE]
		return 1, [3]datatypes.BNode{old} // not split
	}
	left := datatypes.BNode(make([]byte, 2*constants.BTREE_PAGE_SIZE)) // might be split later
	right := datatypes.BNode(make([]byte, constants.BTREE_PAGE_SIZE))
	nodeSplit2(left, right, old)
	if left.Nbytes() <= constants.BTREE_PAGE_SIZE {
		left = left[:constants.BTREE_PAGE_SIZE]
		return 2, [3]datatypes.BNode{left, right} // 2 nodes
	}
	leftleft := datatypes.BNode(make([]byte, constants.BTREE_PAGE_SIZE))
	middle := datatypes.BNode(make([]byte, constants.BTREE_PAGE_SIZE))
	nodeSplit2(leftleft, middle, left)
	// assert(leftleft.Nbytes() <= constants.BTREE_PAGE_SIZE)
	return 3, [3]datatypes.BNode{leftleft, middle, right} // 3 nodes
}
