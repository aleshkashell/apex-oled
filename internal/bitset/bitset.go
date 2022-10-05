package bitset

import "golang.org/x/exp/constraints"

type BitSet interface {
	byte | constraints.Integer
}

func SetBit[B BitSet](n B, pos uint) B {
	n |= 1 << pos
	return n
}

// Clears the bit at pos in n.
func ClearBit[B BitSet](n B, pos uint) B {
	var mask B
	mask = ^(1 << pos)
	n &= mask
	return n
}

func HasBit[B BitSet](n B, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}
