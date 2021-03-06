package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/FactomProject/factomd/common/adminBlock"
	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/entryCreditBlock"
	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
)

//var server string = "localhost:8088" //Localhost
//var server string = "52.17.183.121:8088" //TestNet
var server string = "52.18.72.212:8088" //MainNet

type DBlockHead struct {
	KeyMR string
}

func GetDBlock(keymr string) (interfaces.IDirectoryBlock, error) {
	raw, err := GetRaw(keymr)
	if err != nil {
		return nil, err
	}
	dblock, err := directoryBlock.UnmarshalDBlock(raw)
	if err != nil {
		return nil, err
	}
	return dblock, nil
}

func GetABlock(keymr string) (interfaces.IAdminBlock, error) {
	raw, err := GetRaw(keymr)
	if err != nil {
		return nil, err
	}
	block, err := adminBlock.UnmarshalABlock(raw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func GetECBlock(keymr string) (interfaces.IEntryCreditBlock, error) {
	raw, err := GetRaw(keymr)
	if err != nil {
		return nil, err
	}
	block, err := entryCreditBlock.UnmarshalECBlock(raw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func GetFBlock(keymr string) (interfaces.IFBlock, error) {
	raw, err := GetRaw(keymr)
	if err != nil {
		return nil, err
	}
	block, err := factoid.UnmarshalFBlock(raw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func GetEBlock(keymr string) (interfaces.IEntryBlock, error) {
	raw, err := GetRaw(keymr)
	if err != nil {
		return nil, err
	}
	block, err := entryBlock.UnmarshalEBlock(raw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func GetEntry(hash string) (interfaces.IEBEntry, error) {
	raw, err := GetRaw(hash)
	if err != nil {
		fmt.Printf("got error %s\n", err)
		fmt.Printf("called getraw with %s\n", hash)
		fmt.Printf("got result %s\n", raw)

		return nil, err
	}
	entry, err := entryBlock.UnmarshalEntry(raw)
	for err != nil { //just keep trying until it doesn't give an error
		fmt.Printf("got error %s\n", err)
		fmt.Printf("called entryBlock.UnmarshalEntry with %s\n", raw)
		fmt.Printf("got result %s\n", entry)
		//if we get an error like EOF, get the thing again after a short wait
		time.Sleep(20000 * time.Millisecond)
		raw, err = GetRaw(hash)
		if err != nil {
			return nil, err
		}
		entry, err = entryBlock.UnmarshalEntry(raw)
	}
	return entry, nil
}

func GetDBlockHead() (string, error) {
	//return "3a5ec711a1dc1c6e463b0c0344560f830eb0b56e42def141cb423b0d8487a1dc", nil //10
	//return "cde346e7ed87957edfd68c432c984f35596f29c7d23de6f279351cddecd5dc66", nil //100
	//return "d13472838f0156a8773d78af137ca507c91caf7bf3b73124d6b09ebb0a98e4d9", nil //200

	resp, err := http.Get(
		fmt.Sprintf("http://%s/v1/directory-block-head/", server))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(string(body))
	}

	d := new(DBlockHead)
	json.Unmarshal(body, d)

	return d.KeyMR, nil
}

type Data struct {
	Data string
}

func GetRaw(keymr string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/v1/get-raw-data/%s", server, keymr))
	for err != nil {
		//if the http code gave an error, give a little time and try again before panicking.
		fmt.Printf("got error %s, waiting 20 seconds\n", err)
		time.Sleep(20000 * time.Millisecond)
		resp, err = http.Get(fmt.Sprintf("http://%s/v1/get-raw-data/%s", server, keymr))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	for err != nil {
		//if the io reader code gave an error, give a little time and try again before panicking.
		fmt.Printf("got error %s, waiting 20 seconds\n", err)
		time.Sleep(20000 * time.Millisecond)
		body, err = ioutil.ReadAll(resp.Body)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}

	d := new(Data)
	if err := json.Unmarshal(body, d); err != nil {
		return nil, err
	}

	raw, err := hex.DecodeString(d.Data)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
