package wallet

import (
	"fmt"
	"testing"
)

func TestMasterWallet(t *testing.T) {
	mw := NewMasterWallet()
	if mw == nil {
		t.Error("expected to have err but got nil")
	}
	fmt.Println(mw)
}
