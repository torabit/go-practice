package facade

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

type WalletDetail struct {
	name    string
	balance int
	code    int
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	walletFacade := &WalletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
	}
	return walletFacade
}

func (w *WalletFacade) getWalletDetail() *WalletDetail {
	detail := &WalletDetail{
		name:    w.account.name,
		balance: w.wallet.balance,
		code:    w.securityCode.code,
	}
	return detail
}
