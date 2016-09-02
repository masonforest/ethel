package main

//go:generate abigen --sol token.sol --pkg main --out token.go

import (
	"crypto/ecdsa"
	"flag"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var (
	key0      *ecdsa.PrivateKey
	key1      *ecdsa.PrivateKey
	addr0     common.Address
	addr1     common.Address
	auth      *bind.TransactOpts
	backend       *backends.SimulatedBackend
	GAS_LIMIT = big.NewInt(500000)
	NONCE     = big.NewInt(2)
)

func TestMain(m *testing.M) {
	key0, _ = crypto.GenerateKey()
	key1, _ = crypto.GenerateKey()
	addr0 = crypto.PubkeyToAddress(key0.PublicKey)
	addr1 = crypto.PubkeyToAddress(key1.PublicKey)
	auth = bind.NewKeyedTransactor(key0)
	backend = backends.NewSimulatedBackend(
		core.GenesisAccount{
			Address: addr0,
			Balance: big.NewInt(9223372036854775807),
		},
		core.GenesisAccount{
			Address: addr1,
			Balance: big.NewInt(9223372036854775807),
		})
	flag.Parse()
	os.Exit(m.Run())
}

func Signer(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signature, err := crypto.Sign(tx.SigHash().Bytes(), key0)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signature)
	}
}

func deploy(initialSupply *big.Int, name string) *TokenSession {
	_, _, token, err := DeployToken(
		bind.NewKeyedTransactor(key0),
		backend,
		initialSupply,
		name,
	)

	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	backend.Commit()

	return &TokenSession{
		Contract:     token,
		TransactOpts: *bind.NewKeyedTransactor(key0),
	}

}

func TestInitializer(t *testing.T) {
	token := deploy(big.NewInt(100), "Token")

	name, _ := token.Name()
	assert.Equal(t, "Token", name)
	balance, _ := token.Balances(addr0)
	assert.Equal(t, big.NewInt(100), balance)
}

func TestTransfer(t *testing.T) {
	token := deploy(big.NewInt(3), "Token")

	_, err := token.Transfer(addr1, big.NewInt(1))

	if err != nil {
		panic(err)
	}

	backend.Commit()

	senderBalance, _ := token.Balances(addr0)
	recipientBalance, _ := token.Balances(addr1)
	assert.Equal(t, big.NewInt(2), senderBalance)
	assert.Equal(t, big.NewInt(1), recipientBalance)
}
