package state_test

import (
	"fmt"
	"github.com/FactomProject/factomd/common/primitives"
	. "github.com/FactomProject/factomd/state"
	"testing"
)

var _ = fmt.Print

func TestFedServer(t *testing.T) {
	state := new(State)
	pls := NewProcessLists(state)
	pl := pls.Get(0)
	pl.AddFedServer(primitives.NewHash([]byte("one")))
	pl.AddFedServer(primitives.NewHash([]byte("two")))
	pl.AddFedServer(primitives.NewHash([]byte("three")))
}
