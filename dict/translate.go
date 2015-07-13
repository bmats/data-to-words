package dict

import (
	"bufio"
	"io"
	"math"
)

func (dict *Dictionary) Translate(in io.Reader, out io.Writer, delim string) error {
	bufWriter := bufio.NewWriter(out)
	defer bufWriter.Flush()

	maxBitCount := dict.Size()
	isFirstWord := true

	val, bitCount := 0, 0
	next, nextBitCount := 0, 0

	var buf []byte = make([]byte, 1)
	for {
		// Read a new byte?
		if nextBitCount == 0 {
			n, err := in.Read(buf)
			if err == nil {
				if n == 0 {
					continue
				}

				next = int(buf[0])
				nextBitCount = 8
			} else {
				if err == io.EOF {
					if bitCount != 0 && bitCount != maxBitCount {
						// Provide some 0's to fill up the rest
						next = 0
						nextBitCount = maxBitCount - bitCount
					} else {
						break
					}
				} else {
					return err
				}
			}
		}

		// Append the new bits onto val
		newBitCount := int(math.Min(float64(maxBitCount-bitCount), float64(nextBitCount)))
		val <<= uint(newBitCount)
		val |= next >> uint(nextBitCount-newBitCount)
		bitCount += newBitCount

		// Zero the used bits from next
		next &= (1 << uint(nextBitCount-newBitCount)) - 1
		nextBitCount -= newBitCount

		// Ready to print a word?
		if bitCount >= maxBitCount {
			if isFirstWord {
				isFirstWord = false
			} else {
				bufWriter.WriteString(delim)
			}

			bufWriter.WriteString(dict.Word(val))

			val, bitCount = 0, 0
		}
	}

	return nil
}
