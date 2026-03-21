/* API: Binary Reader */

package binaryio

import (
	"bufio"
	"io"
)

type BinaryReader struct {
	in          *bufio.Reader
	accumulator byte
	left        uint8
}

/* Create a new Binary Reader based on I/O Reader. */
func NewBinaryReader(r io.Reader) *BinaryReader {
	return &BinaryReader{
		in: bufio.NewReader(r),
	}
}

/* Read 1 bit of data and return as a boolean value. */
func (br *BinaryReader) ReadBit() (bool, error) {
	if br.left == 0 {
		b, bufReadErr := br.in.ReadByte()
		if bufReadErr != nil {
			return false, bufReadErr
		}

		br.accumulator = b
		br.left = 8
	}

	br.left--
	bit := (br.accumulator>>br.left)&1 == 1

	return bit, nil
}

/* Read 8 bits of data and return as a char value. */
func (br *BinaryReader) ReadByte() (byte, error) {
	if br.left == 0 {
		return br.in.ReadByte()
	}

	var res byte
	for i := 7; i >= 0; i-- {
		bit, bitReadErr := br.ReadBit()
		if bitReadErr != nil {
			return 0, bitReadErr
		}

		if bit {
			res |= (1 << i)
		}
	}

	return res, nil
}

/* Read n bits of data and return as an unsigned integer value. */
func (br *BinaryReader) ReadMultiBits(n uint8) (uint64, error) {
	var res uint64

	for i := 0; i < int(n); i++ {
		res <<= 1
		
		bit, bitReadErr := br.ReadBit()
		if bitReadErr != nil {
			if bitReadErr == io.EOF {
				if i == 0 {
					return 0xFF, io.EOF
				}

				return 0xFF, io.ErrUnexpectedEOF
			}

			return 0xFF, bitReadErr
		}
		
		if bit {
			res |= 1
		}
	}
	
	return res, nil
}
