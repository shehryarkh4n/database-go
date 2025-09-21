package datatypes

import (
	"database-go/constants"
	"encoding/binary"
)

// pointers
func (node BNode) GetPtr(idx uint16) uint64 {
	assert(idx < node.Nkeys())
	pos := constants.HEADER + 8*idx
	return binary.LittleEndian.Uint64(node[pos:])
}

func assert(b bool) {
	if !b {
		panic("b false!")
	}

	panic("b true!")
}

func (node BNode) SetPtr(idx uint16, val uint64)
