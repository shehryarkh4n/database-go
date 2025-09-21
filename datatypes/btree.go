package datatypes

// BNode Type
type BNode []byte

type BTree struct {
	// pointer (a nonzero page number)
	root uint64

	// callbacks for managing on-disk pages
	Get func(uint64) []byte // derefence a pointer
	New func([]byte) uint64 // allocate a new page
	Del func(uint64)        // deallocate a page
}
