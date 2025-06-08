package storage

// "encoding/binary" is used to convert data between byte slices and other types.
import (
	"encoding/binary"
)

// Node types
const (
	BNODE_NODE = iota + 1 // BNode is an internal node
	BNODE_LEAF            // BNode is a leaf node
)

// define sizes
// size of a node is taken as 1 OS page size for now i.e. 4KB
const HEADER = 4
const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VALUE_SIZE = 3000

func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
		// panic will stop the program and print the message
	}
}

func init() {
	node_max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VALUE_SIZE
	// 8B for a pointer, 2B for one KV offset, 4B for a (Key size, value size)
	assert(node_max <= BTREE_PAGE_SIZE, "Node size exceeds page size")
}

// BNode is a B+tree stored as a byte array
type BNode []byte

// Header contains:  Node type, no of keys
// == Header access methods ==

// get type of the node
func (node BNode) Type() uint16 {
	// read the first 2 bytes of the node
	// binary.BigEndian.Uint16 is used to convert the first 2 bytes of the node to a uint16
	// eg. 0x0001 is converted to 1
	return binary.BigEndian.Uint16(node[0:2])
}

// get number of keys in the node
func (node BNode) nKeys() uint16 {
	// read the next 2 bytes of the node
	return binary.BigEndian.Uint16(node[2:4])
}

// set the Header
func (node BNode) setHeader(nodeType uint16, nKeys uint16) {
	binary.BigEndian.PutUint16(node[0:2], nodeType) // set the type of the node
	binary.BigEndian.PutUint16(node[2:4], nKeys)    // set the number of keys in the node
}

// == Pointer access methods ==

// get the pointer to the given index
func (node BNode) getPtr(index uint16) uint64 {
	if index >= node.nKeys() {
		panic("getPtr : Index out of bounds")
	}
	pos := HEADER + 8*index
	return binary.BigEndian.Uint64(node[pos : pos+8])
}

// set the pointer to the given index
func (node BNode) setPtr(index uint16, ptr uint64) {
	if index >= node.nKeys() {
		panic("setPtr : Index out of bounds")
	}
	pos := HEADER + 8*index
	binary.BigEndian.PutUint64(node[pos:pos+8], ptr)
}

//  == Offset access methods ==

// get the offset to the given index
func (node BNode) getOffset(index uint16) uint16 {
	// if we have n pointers then n+1 offsets for the each boundary of the n KV pairs
	if index > node.nKeys() {
		panic("getOffset : Index out of bounds")
	}
	pos := HEADER + 8*node.nKeys() + 2*index
	return binary.BigEndian.Uint16(node[pos : pos+2])
}

// set the offset to the given index
func (node BNode) setOffset(index uint16, offset uint16) {
	if index > node.nKeys() {
		panic("setOffset : Index out of bounds")
	}
	pos := HEADER + 8*node.nKeys() + 2*index
	binary.BigEndian.PutUint16(node[pos:pos+2], offset)
}

// == KV access methods ==

// get the KV pais position at the given index
func (node BNode) kvPos(index uint16) uint16 {
	if index >= node.nKeys() {
		panic("kvPos : Index out of bounds")
	}
	return HEADER + 8*node.nKeys() + 2*(node.nKeys()+1) + node.getOffset(index)
}

// get the key at the given index
func (node BNode) getKey(index uint16) []byte {
	if index >= node.nKeys() {
		panic("getKey : Index out of bounds")
	}
	pos := node.kvPos(index)
	keySize := binary.BigEndian.Uint16(node[pos : pos+2])
	// next 2 bytes are the size of the value
	return node[pos+4 : pos+4+keySize]
}

// return the value at the given index
func (node BNode) getVal(index uint16) []byte {
	if index >= node.nKeys() {
		panic("getValue : Index out of bounds")
	}
	pos := node.kvPos(index)
	keySize := binary.BigEndian.Uint16(node[pos : pos+2])
	valueSize := binary.BigEndian.Uint16(node[pos+2 : pos+4])
	return node[pos+4+keySize : pos+4+keySize+valueSize]
}

// get total no of bytes used by the node
func (node BNode) nbytes() uint16 {
	if node.nKeys() == 0 {
		return 0
	}
	lastIndex := node.nKeys() - 1
	pos := node.kvPos(lastIndex)
	keySize := binary.BigEndian.Uint16(node[pos : pos+2])
	valueSize := binary.BigEndian.Uint16(node[pos+2 : pos+4])
	return pos + 4 + keySize + valueSize
}

// == lookup methods ==

// find largest key less than or equal to the given key
func nodeLookupLE(node BNode, key []byte) uint16 {
	nkeys := node.nKeys()
	idxFound := uint16(0)

	// Binary search for the key
	lowIdx := uint16(0)
	highIdx := nkeys - 1
	for lowIdx <= highIdx {
		midIdx := (lowIdx + highIdx) / 2
		midKey := node.getKey(midIdx)
		if string(midKey) == string(key) {
			idxFound = midIdx
			break
		}
		if string(midKey) < string(key) {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}

	return idxFound
}

// == Node update methods ==
// Node insertion methods

// append a new key-value pair to the node
func nodeAppendKV(new BNode, idx uint16, ptr uint64, key []byte, value []byte) {
	// Set the pointer
	new.setPtr(idx, ptr)

	// get position of the KV pair
	pos := new.kvPos(idx)

	// set the KV lengths
	binary.BigEndian.PutUint16(new[pos:pos+2], uint16(len(key)))     // key size
	binary.BigEndian.PutUint16(new[pos+2:pos+4], uint16(len(value))) // value size

	// copy the key and value data
	copy(new[pos+4:], key)                    // copy the key
	copy(new[pos+4+uint16(len(key)):], value) // copy the value

	// if this is first KV in the new node, set the lefmost offset of the first KV
	if idx == 0 {
		new.setOffset(0, HEADER+8*new.nKeys()+2*(new.nKeys()+1))
	}

	// set the offset for the end of current KV pair
	new.setOffset(idx+1, new.getOffset(idx)+4+uint16(len(key))+uint16(len(value)))
}

// append a range of key-value pairs from one node to another
func nodeAppendRange(new BNode, oldBNode BNode, dstIdx uint16, srcIdx uint16, n uint16) {
	// @dstIdx: index in the new node, @srcIdx: index in the old node
	// @n : number of keys to copy
	if n == 0 {
		return
	}

	// copy the pointers
	for i := uint16(0); i < n; i++ {
		new.setPtr(dstIdx+i, oldBNode.getPtr(srcIdx+i))
	}

	// get the positions from indexes
	dstPos := new.kvPos(dstIdx)
	srcPos := oldBNode.kvPos(srcIdx)

	// find the numner of bytes to copy
	var copySize uint16 = 0
	if srcIdx+n >= oldBNode.nKeys() {
		// copy till the end of the node
		copySize = oldBNode.nbytes() - srcPos
	} else {
		// copy till the specified index
		copySize = oldBNode.kvPos(srcIdx+n) - srcPos
	}

	// copy the key-value pairs
	copy(new[dstPos:], oldBNode[srcPos:srcPos+copySize])

	// if this is first KV in the new node, set the lefmost offset of the first KV
	if dstIdx == 0 {
		new.setOffset(0, HEADER+8*new.nKeys()+2*(new.nKeys()+1))
	}

	// update the offsets
	for i := uint16(1); i <= n; i++ {
		// rel offset for this pair in old node
		relOffset := oldBNode.getOffset(srcIdx+i) - srcPos
		// set the offset for the new node
		new.setOffset(dstIdx+i, new.getOffset(dstIdx)+relOffset)
	}

}

// add a new key to leaf node using copy on write
func leafInsert(
	newNode BNode, oldBNode BNode, idx uint16, key []byte, value []byte) {

	// set the header
	newNode.setHeader(BNODE_LEAF, oldBNode.nKeys()+1)

	// copy before insertion index
	nodeAppendRange(newNode, oldBNode, 0, 0, idx)
	// append the new key-value pair
	nodeAppendKV(newNode, idx, 0, key, value)
	// copy after insertion index
	nodeAppendRange(newNode, oldBNode, idx+1, idx, oldBNode.nKeys()-idx)
}
