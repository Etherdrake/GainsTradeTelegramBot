package swaps

import (
	"fmt"
	"log"
	"testing"
)

func TestGetTokenInfo(t *testing.T) {
	tokenInfo, swapErr := GetTokenInfo(1, "0x582d872A1B094FC48F5DE31D3B73F2D9bE47def1")
	if swapErr != nil {
		log.Println("error getting token info")
		log.Println(swapErr)
		t.Fatal("error when getting token info")
	} else {
		log.Println("OK token info")
		fmt.Sprintln(tokenInfo)
	}
}
