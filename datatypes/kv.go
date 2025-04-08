package datatypes

import (
	"database-go/constants"
	"encoding/binary"
)

func offsetPos(node BNode, idx uint16) uint16 {
	if !(1 <= idx && idx <= node.nkeys()) {
		panic("offsetPos problem!")
	}
	return constants.HEADER + 8*node.nkeys() + 2*(idx-1)
}

func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}

	return binary.LittleEndian.Uint16(node[offsetPos(node, idx):])
}

func (node BNode) setOffset(idx uint16, offset uint16)
