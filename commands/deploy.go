package commands

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"log"
	"os"
	"strings"
)

const key = `{"address":"c8a8b35742b10bdc00a724c6b76d5f85b6558b0c","crypto":{"cipher":"aes-128-ctr","ciphertext":"214a1ff15bbae62ca2dd90d2333cb23ee615df3f173d5158b1b89cf3329b18b6","cipherparams":{"iv":"5680a66256d7d09abaa7e95063a8c375"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"72760c070ebf2600a7200aa71158ff1a583d7a1676cbf8b1df3a3dc2e0bbf10f"},"mac":"82fbafdcdb69d78f6f3e9bd35d6a333d6518c425e0cf9dd4804778047e926bcc"},"id":"f972e6e1-ef3f-4bdb-b897-623797517c1b","version":3}`
const password = "hNJTZzGTxXmz9KzdDTvfqnCS"

func Deploy(c *cli.Context) error {
	contracts, err := compiler.CompileSolidity("", getFileName(c.Args().First()))
	check(err)
	conn, err := ethclient.Dial("http://localhost:8545/")
	check(err)
	transactor, err := bind.NewTransactor(strings.NewReader(key), password)
	check(err)

	for _, contract := range contracts {
		deployContract(contract, conn, transactor)
	}

	fmt.Println("Deployed", getFileName(c.Args().First()))

	printBalance(transactor, conn)

	return nil
}

func getFileName(s string) string {
	if _, err := os.Stat("contracts/" + s + ".sol"); err == nil {
		return "contracts/" + s + ".sol"
	} else if _, err := os.Stat(s); err == nil {
		return s
	} else {
		log.Fatalf("%s does not exist", s)
		return ""
	}
}
func deployContract(contract *compiler.Contract, conn *ethclient.Client, transactor *bind.TransactOpts) {
	j, err := json.Marshal(contract.Info.AbiDefinition)
	check(err)
	abi, err := abi.JSON(strings.NewReader(string(j)))
	check(err)
	address, tx, _, err := bind.DeployContract(transactor, abi, common.FromHex(contract.Code), conn)
	check(err)
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
}

func printBalance(transactor *bind.TransactOpts, conn *ethclient.Client) {
	bal, err := conn.BalanceAt(context.TODO(), transactor.From, nil)
	check(err)
	fmt.Printf("Remaining Balance: %s\n", common.CurrencyToString(bal))
}

func check(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
