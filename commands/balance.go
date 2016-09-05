package commands

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"io/ioutil"
)

func Balance(c *cli.Context) error {
	conn, err := ethclient.Dial(c.String("host"))
	check(err)
	t := getTransactor(c)

	fmt.Printf("Address: %s\n", t.From.Hex())
	printBalance(t, conn)

	return nil
}

func getTransactor(c *cli.Context) *bind.TransactOpts {
	var config Config
	var keyJSON string

	configJSON, err := ioutil.ReadFile(getConfigFile())
	check(err)
	err = json.Unmarshal(configJSON, &config)
	check(err)
	if c.Bool("testnet") {
		keyJSON = config.Testnet
	} else {
		keyJSON = config.Live
	}

	return bind.NewKeyedTransactor(unmarshalKey([]byte(string(keyJSON))).PrivateKey)
}

func unmarshalKey(s []byte) accounts.Key {
	key := &accounts.Key{}
	if isKeyEncypted(s) {
		key = decyptKey(s)
	} else {
		key.UnmarshalJSON(s)
	}

	return *key
}

func decyptKey(s []byte) *accounts.Key {
	p := askString("Password")
	key, err := accounts.DecryptKey(s, p)
	if (err) != nil {
		fmt.Println("Invalid password. Please try again.")
		return decyptKey(s)
	}

	return key
}

func isKeyEncypted(s []byte) bool {
	var key map[string]interface{}
	err := json.Unmarshal(s, &key)
	check(err)
	return key["crypto"] != nil
}

func printBalance(transactor *bind.TransactOpts, conn *ethclient.Client) {
	bal, err := conn.BalanceAt(context.TODO(), transactor.From, nil)
	check(err)
	fmt.Printf("Balance: %s\n", common.CurrencyToString(bal))
}
