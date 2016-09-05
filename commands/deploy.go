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
	"log"
	"os"
	"strings"
)

func Deploy(c *cli.Context) error {
	contracts, err := compiler.CompileSolidity("", getFileName(c.Args().First()))
	check(err)
	conn, err := ethclient.Dial(c.String("host"))
	check(err)
	transactor := getTransactor(c)

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

func check(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
