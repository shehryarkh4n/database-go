package datatypes

import (
	"bytes"
	"database-go/constants"
	"encoding/binary"
)

// Offset code
func offsetPos(node BNode, idx uint16) uint16 {
	if !(1 <= idx && idx <= node.Nkeys()) {
		panic("offsetPos problem!")
	}
	return constants.HEADER + 8*node.Nkeys() + 2*(idx-1)
}

func (node BNode) GetOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}

	return binary.LittleEndian.Uint16(node[offsetPos(node, idx):])
}

func (node BNode) SetOffset(idx uint16, offset uint16)

// Key-Value Code
func (node BNode) KvPos(idx uint16) uint16 {
	if idx > node.Nkeys() {
		panic("kvPos problem!")
	}
	return constants.HEADER + 8*node.Nkeys() + 2*node.Nkeys() + node.GetOffset(idx)
}

func (node BNode) GetKey(idx uint16) []byte {
	if idx >= node.Nkeys() {
		panic("getKey KV problem!")
	}
	pos := node.KvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	// Very neat way of doing this in Go
	// This basically says go to the start of node's "pos+4"-th position
	// then take the first klen values from it
	// very different syntax from python, where this would be some 2D array index slicing
	// Had me so confused for a moment
	return node[pos+4:][:klen]
}

func (node BNode) getVal(idx uint16) []byte {
	if idx >= node.Nkeys() {
		panic("getVal KV problem!")
	}

	pos := node.KvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	vlen := binary.LittleEndian.Uint16(node[pos:])

	// This does make intuitive sense. You begin reading val bytes
	// after jumping head to the start of the val bytes by skipping
	// the klen/vlen bytes AND the key bytes
	return node[pos+4+klen:][:vlen]
}

// node size in bytes
// very smart lookup for node size.
// node.Nkeys() provides the total number of keys, which is automatically
// the idx of the last KV
func (node BNode) Nbytes() uint16 {
	return node.KvPos(node.Nkeys())
}

// returns the first child node whose range intersects the key so
// === (child[i] <= key)
func NodeLookupLessEqual(node BNode, key []byte) uint16 {
	found := uint16(0)

	// left := uint16(0)
	// right := ^uint16(0)
	// b_search_condition := left <= right

	// for b_search_condition {
	// 	if left > right {
	// 		b_search_condition = false
	// 	}

	// 	middle := (left + right) / 2 // do i need floor if its uint16?

	// 	cmp := bytes.Compare(node.getKey(middle), key)

	// 	if cmp <= -1 {
	// 		return found = middle
	// 	}

	// 	if cmp >= 0 {
	// 		break
	// 	}
	// }

	for i := uint16(1); i < node.Nkeys(); i++ {
		cmp := bytes.Compare(node.GetKey(i), key) // Compare returns -1 if a < b

		if cmp <= -1 { // why not just break here?
			found = i
		}
		if cmp >= 0 {
			break
		}
	}

	return found
}
