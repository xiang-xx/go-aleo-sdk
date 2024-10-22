package transaction

import (
	"github.com/xiang-xx/go-aleo-sdk/account"
)

// NewCoinbaseTransaction crafts a transaction that can be used for coinbase rewards.
func NewCoinbaseTransaction(address *account.Address, value int64, random []byte) string {
	return newCoinbaseTransaction(address, value, random)
}

// NewTransferTransaction consumes a single record and crafts a transaction that sends an amount to a recipient.
func NewTransferTransaction(privateKey *account.PrivateKey, to *account.Address, in string, ledgerProofs []string, amount, fee int64) (string, error) {
	return newTransferTransaction(privateKey, to, in, ledgerProofs, amount, fee)
}
