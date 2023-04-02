package facade

import (
	"reflect"
	"testing"
)

func TestFacade(t *testing.T) {
	walletFacade := newWalletFacade("taro", 1234)
	expectedWalletFacade := &WalletDetail{
		name:    "taro",
		balance: 0,
		code:    1234,
	}

	if !reflect.DeepEqual(walletFacade.getWalletDetail(), expectedWalletFacade) {
		t.Fatalf("\nwant: %+v\ngot: %+v", walletFacade.getWalletDetail(), expectedWalletFacade)
	}
}
