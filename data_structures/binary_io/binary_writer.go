/* API: Binary Writer */

package binaryio

import (
	"bufio"
	"io"
)

type BinaryWriter struct {
	out         *bufio.Writer
	accumulator byte
	count       uint8
}

/* Create a new Binary Writer based on I/O Writer. */
func NewBinaryWriter(w io.Writer) *BinaryWriter {
	return &BinaryWriter{
		out: bufio.NewWriter(w),
	}
}

/* Write the specified bit into bitstream. */
func (bw *BinaryWriter) WriteBit(bit bool) error {
	if bit {
		bw.accumulator |= (1 << (7 - bw.count))
	}

	bw.count++

	if bw.count == 8 {
		if bufWriErr := bw.out.WriteByte(bw.accumulator); bufWriErr != nil {
			return bufWriErr
		}

		bw.accumulator = 0
		bw.count = 0
	}

	return nil
}

/* Write the specified 8-bit char into bitstream. */
func (bw *BinaryWriter) WriteByte(b byte) error {
	if bw.count == 0 {
		if bufWriErr := bw.out.WriteByte(b); bufWriErr != nil {
			return bufWriErr
		}

		return nil
	}

	for i := 7; i >= 0; i-- {
		bit := (b>>i)&1 == 1

		if bitWriErr := bw.WriteBit(bit); bitWriErr != nil {
			return bitWriErr
		}
	}

	return nil
}

/* Write the n LSB of the specified unsigned integer value. */
func (bw *BinaryWriter) WriteLSBOf(val uint64, n uint8) error {
	for i := int(n) - 1; i >= 0; i-- {
		bit := (val>>i)&1 == 1

		if bitWriErr := bw.WriteBit(bit); bitWriErr != nil {
			return bitWriErr
		}
	}

	return nil
}

/* Close the writing bitstream. */
func (bw *BinaryWriter) Flush() error {
	if bw.count > 0 {
		if bufWriErr := bw.out.WriteByte(bw.accumulator); bufWriErr != nil {
			return bufWriErr
		}

		bw.accumulator = 0
		bw.count = 0
	}

	return bw.out.Flush()
}
