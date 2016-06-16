package wsapi_test

import (
	"encoding/json"
	//"fmt"
	"strings"
	"testing"

	"github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/receipts"
	"github.com/FactomProject/factomd/testHelper"
	. "github.com/FactomProject/factomd/wsapi"
)

/*
func TestHandleFactoidSubmit(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleFactoidSubmit(context)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}
*/
/*
func TestHandleCommitChain(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleCommitChain(context)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}
*/
/*
func TestHandleRevealChain(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleRevealChain(context)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}
*/
/*
func TestHandleCommitEntry(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleCommitEntry(context)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}
*/
/*
func TestHandleRevealEntry(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleRevealEntry(context)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}
*/

func TestHandleDirectoryBlockHead(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleDirectoryBlockHead(context)

	if strings.Contains(testHelper.GetBody(context), "836ba9715fc4e83ae1e8755c40374e7e2265e4f312788710af2ff5478c2b495e") == false {
		t.Errorf("Context does not contain proper DBlock Head - %v", testHelper.GetBody(context))
	}
}

func TestHandleGetRaw(t *testing.T) {
	type RawData struct {
		Hash1 string
		Hash2 string
		Raw   string
	}

	toTest := []RawData{}
	var err error

	blockSet := testHelper.CreateTestBlockSet(nil)

	aBlock := blockSet.ABlock
	raw := RawData{}
	raw.Hash1 = aBlock.DatabasePrimaryIndex().String()
	raw.Hash2 = aBlock.DatabaseSecondaryIndex().String()
	hex, err := aBlock.MarshalBinary()
	if err != nil {
		panic(err)
	}
	raw.Raw = primitives.EncodeBinary(hex)
	toTest = append(toTest, raw) //1

	eBlock := blockSet.EBlock
	raw = RawData{}
	raw.Hash1 = eBlock.DatabasePrimaryIndex().String()
	raw.Hash2 = eBlock.DatabaseSecondaryIndex().String()
	hex, err = eBlock.MarshalBinary()
	if err != nil {
		panic(err)
	}
	raw.Raw = primitives.EncodeBinary(hex)
	toTest = append(toTest, raw) //2

	ecBlock := blockSet.ECBlock
	raw = RawData{}
	raw.Hash1 = ecBlock.(interfaces.DatabaseBatchable).DatabasePrimaryIndex().String()
	raw.Hash2 = ecBlock.(interfaces.DatabaseBatchable).DatabaseSecondaryIndex().String()
	hex, err = ecBlock.MarshalBinary()
	if err != nil {
		panic(err)
	}
	raw.Raw = primitives.EncodeBinary(hex)
	toTest = append(toTest, raw) //3

	fBlock := blockSet.FBlock
	raw = RawData{}
	raw.Hash1 = fBlock.(interfaces.DatabaseBatchable).DatabasePrimaryIndex().String()
	raw.Hash2 = fBlock.(interfaces.DatabaseBatchable).DatabaseSecondaryIndex().String()
	hex, err = fBlock.MarshalBinary()
	if err != nil {
		panic(err)
	}
	raw.Raw = primitives.EncodeBinary(hex)
	toTest = append(toTest, raw) //4

	dBlock := blockSet.DBlock
	raw = RawData{}
	raw.Hash1 = dBlock.DatabasePrimaryIndex().String()
	raw.Hash2 = dBlock.DatabaseSecondaryIndex().String()
	hex, err = dBlock.MarshalBinary()
	if err != nil {
		panic(err)
	}
	raw.Raw = primitives.EncodeBinary(hex)
	toTest = append(toTest, raw) //5

	context := testHelper.CreateWebContext()
	for i, v := range toTest {
		testHelper.ClearContextResponseWriter(context)
		HandleGetRaw(context, v.Hash1)

		if strings.Contains(testHelper.GetBody(context), v.Raw) == false {
			t.Errorf("Looking for %v", v.Hash1)
			t.Errorf("GetRaw %v/%v from Hash1 failed - %v", i, len(toTest), testHelper.GetBody(context))
		}

		testHelper.ClearContextResponseWriter(context)
		HandleGetRaw(context, v.Hash2)

		if strings.Contains(testHelper.GetBody(context), v.Raw) == false {
			t.Errorf("Looking for %v", v.Hash2)
			t.Errorf("GetRaw %v/%v from Hash2 failed - %v", i, len(toTest), testHelper.GetBody(context))
		}
	}
}

func TestHandleDirectoryBlock(t *testing.T) {
	context := testHelper.CreateWebContext()
	hash := "836ba9715fc4e83ae1e8755c40374e7e2265e4f312788710af2ff5478c2b495e"

	HandleDirectoryBlock(context, hash)

	if strings.Contains(testHelper.GetBody(context), "000000000000000000000000000000000000000000000000000000000000000a") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "956c41312070f58c628ca8027297e0af0aaaa7b8af7f84283fc5ad21a49cc00a") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "000000000000000000000000000000000000000000000000000000000000000c") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "37538b52f2336da8781caebeb471ca441f6c2a7374247e2645ef63816cc3ee87") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "000000000000000000000000000000000000000000000000000000000000000f") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "769c1a58507f8f62bf0dd7e7b8eb7f1edf61f19a6508a0fea632a492e7d93e00") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "6e7e64ac45ff57edbf8537a0c99fba2e9ee351ef3d3f4abd93af9f01107e592c") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "1127ed78303976572f25dfba2a058e475234c079ea0d0f645280d03caff08347") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}

	if strings.Contains(testHelper.GetBody(context), "ea0073c1378048ff620f5cf3d872366b791c32df0608b456c5859989c3b8d3ed") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}

func TestHandleEntryBlock(t *testing.T) {
	context := testHelper.CreateWebContext()
	chain, err := primitives.HexToHash("df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604")
	if err != nil {
		t.Error(err)
	}

	dbo := context.Server.Env["state"].(interfaces.IState).GetAndLockDB()
	defer context.Server.Env["state"].(interfaces.IState).UnlockDB()

	blocks, err := dbo.FetchAllEBlocksByChain(chain)
	if err != nil {
		t.Error(err)
	}
	fetched := 0
	for _, b := range blocks {
		hash := b.(*entryBlock.EBlock).DatabasePrimaryIndex().String()
		hash2 := b.(*entryBlock.EBlock).DatabaseSecondaryIndex().String()

		testHelper.ClearContextResponseWriter(context)
		HandleEntryBlock(context, hash)

		eBlock := new(EBlock)

		testHelper.UnmarshalRespDirectly(context, eBlock)

		if eBlock.Header.ChainID != "df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604" {
			t.Errorf("Wrong ChainID - %v", eBlock.Header.ChainID)
			t.Errorf("eBlock - %v", eBlock)
			t.Errorf("%v", testHelper.GetBody(context))
		}

		if eBlock.Header.DBHeight != int64(b.(*entryBlock.EBlock).GetHeader().GetDBHeight()) {
			t.Errorf("DBHeight is wrong - %v vs %v", eBlock.Header.DBHeight, b.(*entryBlock.EBlock).GetHeader().GetDBHeight())
		}

		testHelper.ClearContextResponseWriter(context)
		HandleEntryBlock(context, hash2)

		eBlock = new(EBlock)

		testHelper.UnmarshalRespDirectly(context, eBlock)

		if eBlock.Header.ChainID != "df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604" {
			t.Errorf("Wrong ChainID - %v", eBlock.Header.ChainID)
			t.Errorf("%v", testHelper.GetBody(context))
		}

		if eBlock.Header.DBHeight != int64(b.(*entryBlock.EBlock).GetHeader().GetDBHeight()) {
			t.Errorf("DBHeight is wrong - %v vs %v", eBlock.Header.DBHeight, b.(*entryBlock.EBlock).GetHeader().GetDBHeight())
		}

		fetched++
	}
	if fetched != testHelper.BlockCount {
		t.Errorf("Fetched %v blocks, expected %v", fetched, testHelper.BlockCount)
	}
}

/*
func TestHandleEntry(t *testing.T) {
	context := testHelper.CreateWebContext()
	hash := ""

	HandleEntry(context, hash)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}
*/
func TestHandleChainHead(t *testing.T) {
	context := testHelper.CreateWebContext()
	hash := "000000000000000000000000000000000000000000000000000000000000000d"

	HandleChainHead(context, hash)

	if strings.Contains(testHelper.GetBody(context), "836ba9715fc4e83ae1e8755c40374e7e2265e4f312788710af2ff5478c2b495e") == false {
		t.Errorf("Invalid directory block head: %v", testHelper.GetBody(context))
	}

	hash = "000000000000000000000000000000000000000000000000000000000000000a"

	testHelper.ClearContextResponseWriter(context)
	HandleChainHead(context, hash)

	if strings.Contains(testHelper.GetBody(context), "956c41312070f58c628ca8027297e0af0aaaa7b8af7f84283fc5ad21a49cc00a") == false {
		t.Errorf("Invalid admin block head: %v", testHelper.GetBody(context))
	}

	hash = "000000000000000000000000000000000000000000000000000000000000000c"

	testHelper.ClearContextResponseWriter(context)
	HandleChainHead(context, hash)

	if strings.Contains(testHelper.GetBody(context), "37538b52f2336da8781caebeb471ca441f6c2a7374247e2645ef63816cc3ee87") == false {
		t.Errorf("Invalid entry credit block head: %v", testHelper.GetBody(context))
	}

	hash = "000000000000000000000000000000000000000000000000000000000000000f"

	testHelper.ClearContextResponseWriter(context)
	HandleChainHead(context, hash)

	if strings.Contains(testHelper.GetBody(context), "769c1a58507f8f62bf0dd7e7b8eb7f1edf61f19a6508a0fea632a492e7d93e00") == false {
		t.Errorf("Invalid factoid block head: %v", testHelper.GetBody(context))
	}

	hash = "6e7e64ac45ff57edbf8537a0c99fba2e9ee351ef3d3f4abd93af9f01107e592c"

	testHelper.ClearContextResponseWriter(context)
	HandleChainHead(context, hash)

	if strings.Contains(testHelper.GetBody(context), "1127ed78303976572f25dfba2a058e475234c079ea0d0f645280d03caff08347") == false {
		t.Errorf("Invalid entry block head: %v", testHelper.GetBody(context))
	}

	hash = "df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604"

	testHelper.ClearContextResponseWriter(context)
	HandleChainHead(context, hash)

	if strings.Contains(testHelper.GetBody(context), "ea0073c1378048ff620f5cf3d872366b791c32df0608b456c5859989c3b8d3ed") == false {
		t.Errorf("Invalid entry block head: %v", testHelper.GetBody(context))
	}
}

//func TestHandleEntryCreditBalance(t *testing.T) {
//	context := testHelper.CreateWebContext()
//
//	HandleEntryCreditBalance(context)
//
//	if strings.Contains(testHelper.GetBody(context), "") == false {
//		t.Errorf("%v", testHelper.GetBody(context))
//	}
//}

func TestHandleFactoidBalance(t *testing.T) {
	context := testHelper.CreateWebContext()
	eckey := testHelper.NewFactoidRCDAddressString(0)

	t.Logf("%v\n", eckey)

	HandleFactoidBalance(context, eckey)

	//expectedAmount := fmt.Sprintf("%v", uint64(testHelper.BlockCount)*testHelper.DefaultCoinbaseAmount)
	expectedAmount := "999889000"
	if strings.Contains(testHelper.GetBody(context), expectedAmount) == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}

func TestHandleGetFee(t *testing.T) {
	context := testHelper.CreateWebContext()

	HandleGetFee(context)

	if strings.Contains(testHelper.GetBody(context), "") == false {
		t.Errorf("%v", testHelper.GetBody(context))
	}
}

func TestBlockIteration(t *testing.T) {
	context := testHelper.CreateWebContext()
	hash := "000000000000000000000000000000000000000000000000000000000000000d"

	HandleChainHead(context, hash)

	j := testHelper.GetRespText(context)
	head := new(CHead)
	err := primitives.DecodeJSONString(j, head)
	if err != nil {
		panic(err)
	}

	prev := head.ChainHead
	fetched := 0
	for {
		if prev == "0000000000000000000000000000000000000000000000000000000000000000" || prev == "" {
			break
		}
		testHelper.ClearContextResponseWriter(context)
		HandleDirectoryBlock(context, prev)

		j = testHelper.GetRespText(context)
		block := new(DBlock)
		err = primitives.DecodeJSONString(j, block)
		if err != nil {
			panic(err)
		}
		prev = block.Header.PrevBlockKeyMR
		fetched++
	}
	if fetched != testHelper.BlockCount {
		t.Errorf("DBlock only found %v blocks, was expecting %v", fetched, testHelper.BlockCount)
	}
}

func TestHandleGetReceipt(t *testing.T) {
	context := testHelper.CreateWebContext()
	hash := "be5fb8c3ba92c0436269fab394ff7277c67e9b2de4431b723ce5d89799c0b93a"

	HandleGetReceipt(context, hash)

	j := testHelper.GetRespMap(context)

	if j == nil {
		t.Error("Receipt not found!")
		return
	}

	dbo := context.Server.Env["state"].(interfaces.IState).GetAndLockDB()
	defer context.Server.Env["state"].(interfaces.IState).UnlockDB()

	receipt := j["receipt"].(map[string]interface{})
	marshalled, err := json.Marshal(receipt)
	if err != nil {
		t.Error(err)
	}

	err = receipts.VerifyFullReceipt(dbo, string(marshalled))
	if err != nil {
		t.Logf("receipt - %v", j)
		t.Error(err)
	}
}
