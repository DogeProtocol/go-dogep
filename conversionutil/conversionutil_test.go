package conversionutil

import (
	"fmt"
	"github.com/DogeProtocol/dp/common"
	"strings"
	"testing"
)

var data = []byte{25, 71, 244, 125, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 48, 120, 100, 97, 48, 50, 53, 53, 51, 67, 48, 68, 54, 56, 65, 50, 53, 49, 70, 53, 56, 48, 50, 52, 99, 50, 51, 69, 55, 54, 99, 57, 57, 101, 52, 56, 51, 49, 53, 70, 99, 67, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 48, 120, 52, 100, 52, 51, 55, 53, 98, 100, 52, 102, 55, 98, 54, 101, 97, 48, 53, 51, 51, 55, 98, 53, 100, 56, 48, 54, 101, 53, 50, 51, 99, 57, 98, 98, 51, 99, 49, 101, 102, 54, 97, 50, 51, 57, 99, 48, 48, 102, 101, 51, 50, 49, 102, 55, 54, 57, 100, 98, 49, 57, 49, 99, 48, 101, 52, 55, 53, 100, 53, 98, 49, 98, 50, 99, 57, 97, 57, 50, 50, 50, 51, 99, 49, 102, 52, 50, 55, 48, 97, 99, 98, 55, 48, 49, 57, 48, 55, 55, 51, 52, 51, 102, 99, 55, 48, 101, 57, 97, 57, 97, 102, 52, 56, 100, 97, 97, 98, 56, 100, 57, 98, 57, 97, 99, 56, 53, 53, 55, 49, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func TestData(t *testing.T) {
	quantumAddr := common.HexToAddress("0x0ee725628df715c17975fbec62f564b8c6456af3defc02c04c1837d4e28faa65")
	ethAddr, err := VerifyDataAndGetEthereumAddress(quantumAddr, data)
	if err != nil {
		t.Fatalf("failed a")
	}

	fmt.Println(ethAddr)
	if strings.Compare(ethAddr, "0xda02553C0D68A251F58024c23E76c99e48315FcC") != 0 {
		t.Fatalf("failed b")
	}
}

func TestSnapshot(t *testing.T) {
	a1 := strings.ToLower("0xB1eC2A23d60E6D903b1Bc6F09094557ba0a65E9c")
	_, ok := SnapshotMap[a1]

	// If the key exists
	if ok == false {
		t.Fatalf("failed")
	}
}
