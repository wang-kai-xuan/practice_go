package main

import (
	"fmt"
	"testing"
)

type OutInfo struct {
	index int
	addr  string
}

func TestArrType(t *testing.T) {
	var outs []OutInfo
	outs = append(outs, OutInfo{1, "item1"})
	outs = append(outs, OutInfo{2, "item2"})

	for _, item := range outs {
		fmt.Println(item)
	}
}

func TestArr(t *testing.T) {
	var name []string
	name = append(name, "wkx")
	name = append(name, "wkx")
	name = append(name, "wkx")
	name = append(name, "wkx")

	nameLen := len(name)
	fmt.Println(nameLen)
}

func TestMapArr(t *testing.T) {
	txs := make(map[string]string)
	txs["txid1"] = "txinf1"
	txs["txid2"] = "txinf12"
	txs["txid3"] = "txinf123"
	txs["txid4"] = "txinf1234"
	txs["txid5"] = "txinf12345"

	len1 := len(txs["txid1"])
	len1 = len(txs["txid2"])
	len1 = len(txs["txid3"])
	len1 = len(txs["txid4"])
	len1 = len(txs["txid5"])
	t1 := txs["txid5"] == ""
	t1 = txs["txid6"] == ""
	fmt.Println(len1)
	fmt.Println(t1)
}
