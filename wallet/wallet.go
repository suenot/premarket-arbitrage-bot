package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// DeriveKey derives the first Ethereum private key from a BIP-39 mnemonic.
// It uses the standard path m/44'/60'/0'/0/0.
func DeriveKey(mnemonic string) (*ecdsa.PrivateKey, common.Address, error) {
	mnemonic = strings.TrimSpace(mnemonic)

	w, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("invalid mnemonic: %w", err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := w.Derive(path, false)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("derivation failed: %w", err)
	}

	privKey, err := w.PrivateKey(account)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("private key extraction failed: %w", err)
	}

	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	return privKey, addr, nil
}

// Verify checks that the derived address matches the expected wallet address.
func Verify(mnemonic, expectedAddr string) error {
	_, addr, err := DeriveKey(mnemonic)
	if err != nil {
		return err
	}

	if !strings.EqualFold(addr.Hex(), expectedAddr) {
		return fmt.Errorf("address mismatch: derived %s, expected %s", addr.Hex(), expectedAddr)
	}

	return nil
}
