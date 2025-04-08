package datatypes

// BNode Type
type BNode []byte

type BTree struct {
	// pointer (a nonzero page number)
	root uint64

	// callbacks for managing on-disk pages
	get func(uint64) []byte // derefence a pointer
	new func([]byte) uint64 // allocate a new page
	del func(uint64)        // deallocate a page
}
