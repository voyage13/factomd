// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import (
	"bytes"
	"encoding/hex"
	"fmt"
	. "github.com/FactomProject/factomd/common/constants"
	. "github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// The default FactoidSignature doesn't care about indexing.  We will extend this
// FactoidSignature for multisig
type FactoidSignature struct {
	FactoidSignature [SIGNATURE_LENGTH]byte // The FactoidSignature
}

var _ ISignature = (*FactoidSignature)(nil)

func (t *FactoidSignature) GetHash() IHash {
	return nil
}

func (b FactoidSignature) String() string {
	txt, err := b.CustomMarshalText()
	if err != nil {
		return "<error>"
	}
	return string(txt)
}

func (FactoidSignature) GetDBHash() IHash {
	return primitives.Sha([]byte("Signature"))
}

func (w1 FactoidSignature) GetNewInstance() IBlock {
	return new(FactoidSignature)
}

// Checks that the FactoidSignatures are the same.
func (s1 *FactoidSignature) IsEqual(sig IBlock) []IBlock {
	s2, ok := sig.(*FactoidSignature)
	if !ok || // Not the right kind of IBlock
		s1.FactoidSignature != s2.FactoidSignature { // Not the right rcd
		r := make([]IBlock, 0, 5)
		return append(r, s1)
	}
	return nil
}

// Index is ignored.  We only have one FactoidSignature
func (s *FactoidSignature) SetSignature(sig []byte) error {
	if len(sig) != SIGNATURE_LENGTH {
		return fmt.Errorf("Bad FactoidSignature.  Should not happen")
	}
	copy(s.FactoidSignature[:], sig)
	return nil
}

func (s *FactoidSignature) GetSignature() *[SIGNATURE_LENGTH]byte {
	return &s.FactoidSignature
}

func (s FactoidSignature) MarshalBinary() ([]byte, error) {
	var out bytes.Buffer

	out.Write(s.FactoidSignature[:])

	return out.Bytes(), nil
}

func (s FactoidSignature) CustomMarshalText() ([]byte, error) {
	var out bytes.Buffer

	out.WriteString(" FactoidSignature: ")
	out.WriteString(hex.EncodeToString(s.FactoidSignature[:]))
	out.WriteString("\n")

	return out.Bytes(), nil
}

func (s *FactoidSignature) UnmarshalBinaryData(data []byte) ([]byte, error) {
	copy(s.FactoidSignature[:], data[:SIGNATURE_LENGTH])
	return data[SIGNATURE_LENGTH:], nil
}

func (s *FactoidSignature) UnmarshalBinary(data []byte) error {
	_, err := s.UnmarshalBinaryData(data)
	return err
}