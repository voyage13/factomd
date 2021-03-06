// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages_test

import (
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/directoryBlock"
	//"github.com/FactomProject/factomd/common/interfaces"
	. "github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/primitives"
	"testing"
)

func TestMarshalUnmarshalDirectoryBlockSignature(t *testing.T) {
	msg := newDirectoryBlockSignature()

	hex, err := msg.MarshalBinary()
	if err != nil {
		t.Error("#1 ", err)
	}
	t.Logf("Marshalled - %x", hex)

	msg2, err := UnmarshalMessage(hex)
	if err != nil {
		t.Error("#2 ", err)
	}
	str := msg2.String()
	t.Logf("str - %v", str)

	if msg2.Type() != constants.DIRECTORY_BLOCK_SIGNATURE_MSG {
		t.Error("Invalid message type unmarshalled")
	}

	hex2, err := msg2.(*DirectoryBlockSignature).MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(hex) != len(hex2) {
		t.Error("Hexes aren't of identical length")
	}
	for i := range hex {
		if hex[i] != hex2[i] {
			t.Error("Hexes do not match")
		}
	}

	if msg.IsSameAs(msg2.(*DirectoryBlockSignature)) != true {
		t.Errorf("DirectoryBlockSignature messages are not identical")
	}
}

func TestSignAndVerifyDirectoryBlockSignature(t *testing.T) {
	dbs := newSignedDirectoryBlockSignature()

	hex, err := dbs.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Marshalled - %x", hex)

	t.Logf("Sig - %x", *dbs.Signature.GetSignature())
	if len(*dbs.Signature.GetSignature()) == 0 {
		t.Error("Signature not present")
	}

	valid, err := dbs.VerifySignature()
	if err != nil {
		t.Error(err)
	}
	if valid == false {
		t.Error("Signature is not valid")
	}

	dbs2, err := UnmarshalMessage(hex)
	if err != nil {
		t.Error(err)
	}

	if dbs2.Type() != constants.DIRECTORY_BLOCK_SIGNATURE_MSG {
		t.Error("Invalid message type unmarshalled")
	}
	dbsProper := dbs2.(*DirectoryBlockSignature)

	valid, err = dbsProper.VerifySignature()
	if err != nil {
		t.Error(err)
	}
	if valid == false {
		t.Error("Signature 2 is not valid")
	}
}

func newDirectoryBlockSignature() *DirectoryBlockSignature {
	dbs := new(DirectoryBlockSignature)
	dbs.DBHeight = 123456
	//hash, _ := primitives.NewShaHashFromStr("cbd3d09db6defdc25dfc7d57f3479b339a077183cd67022e6d1ef6c041522b40")
	//dbs.DirectoryBlockKeyMR = hash
	hash, _ := primitives.NewShaHashFromStr("a077183cd67022e6d1ef6c041522b40cbd3d09db6defdc25dfc7d57f3479b339")
	dbs.ServerIdentityChainID = hash
	tmp := directoryBlock.NewDBlockHeader()
	dbs.DirectoryBlockHeader = tmp
	return dbs
}

func newSignedDirectoryBlockSignature() *DirectoryBlockSignature {
	dbs := newDirectoryBlockSignature()
	key, err := primitives.NewPrivateKeyFromHex("07c0d52cb74f4ca3106d80c4a70488426886bccc6ebc10c6bafb37bf8a65f4c38cee85c62a9e48039d4ac294da97943c2001be1539809ea5f54721f0c5477a0a")
	if err != nil {
		panic(err)
	}
	err = dbs.Sign(key)
	if err != nil {
		panic(err)
	}
	return dbs
}
