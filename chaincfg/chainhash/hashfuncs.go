// Copyright (c) 2015 The Decred developers
// Copyright (c) 2016-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash


import (
	"github.com/ProjectVisor/xvrd/keccak"
)

// HashB calculates hash(b) and returns the resulting bytes.
func HashB(b []byte) []byte {




	h := keccak.New256()
	h.Write(b)
	hash := h.Sum(nil)

	return hash[:]
}

// HashH calculates hash(b) and returns the resulting bytes as a Hash.
func HashH(b []byte) Hash {
	h := keccak.New256()
	h.Write(b)
	return Hash(h.Sum(nil))
}

// DoubleHashB calculates hash(hash(b)) and returns the resulting bytes.
func DoubleHashB(b []byte) []byte {

	h := keccak.New256()
	h.Write(b)
	first := h.Sum(nil)

	h := keccak.New256()
	h.Write(first[:])
	second := h.Sum(nil)
	return second[:]
}

// DoubleHashH calculates hash(hash(b)) and returns the resulting bytes as a
// Hash.
func DoubleHashH(b []byte) Hash {

	h := keccak.New256()
	h.Write(b)
	first := h.Sum(nil)
	h := keccak.New256()
	h.Write(first[:])
	return Hash(h.Sum(nil))
}
