// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package common

/*
import (
	"bytes"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"sync"
)

//var FactoidState state.IFactoidState

// factoid Chain
type FctChain struct {
	ChainID interfaces.IHash

	NextBlock       interfaces.IFBlock
	NextBlockHeight uint32
	BlockMutex      sync.Mutex
}

var _ interfaces.Printable = (*FctChain)(nil)

func (e *FctChain) JSONByte() ([]byte, error) {
	return primitives.EncodeJSON(e)
}

func (e *FctChain) JSONString() (string, error) {
	return primitives.EncodeJSONString(e)
}

func (e *FctChain) JSONBuffer(b *bytes.Buffer) error {
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *FctChain) String() string {
	str, _ := e.JSONString()
	return str
}

// factoid Block
*/
