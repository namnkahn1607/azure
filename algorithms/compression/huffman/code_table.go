/* API: Code Table */

package huffman

const R = 256

type BitCode struct {
	Bits uint64 // payload
	Len  uint8  // length
}

type CodeTable [R]BitCode
