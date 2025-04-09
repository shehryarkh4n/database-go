package datatypes

import "encoding/binary"

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

func (node BNode) Btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) Nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) SetHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}
