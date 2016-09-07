package commands

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const CONRACT_GLOB_PATH = "*.sol"

func Deploy(c *cli.Context) error {
	contracts := compileContracts()
	conn, err := ethclient.Dial(c.String("host"))
	check(err)
	transactor := getTransactor(c)

	configData, _ := ioutil.ReadFile("./deploy.yml")
	var config map[string]map[string]string
	yaml.Unmarshal(configData, &config)
	for contractName, paramsYAML := range config {
		params := castContractParams(paramsYAML, contracts[contractName].Info.AbiDefinition)
		deployContract(contracts[contractName], params, conn, transactor)
		fmt.Printf("Deployed %s.\n", contractName)
	}

	printBalance(transactor, conn)

	return nil
}

func castContractParams(config map[string]string, abi interface{}) []interface{} {
	var params []interface{}
	d, _ := json.Marshal(abi)
	jsonParsed, _ := gabs.ParseJSON(d)
	children, _ := jsonParsed.Children()
	for _, child := range children {

		if child.Path("type").Data().(string) == "constructor" {
			inputs, _ := child.Path("inputs").Children()
			for _, input := range inputs {
				value := castValue(config[input.Path("name").Data().(string)], input.Path("type").Data().(string))
				params = append(params, value)
			}
		}
	}

	return params
}

func castValue(s string, t string) interface{} {
	switch t {
	case "uint8":
		i, _ := strconv.Atoi(s)
		return uint8(i)
	case "uint256":
		i, _ := strconv.Atoi(s)
		return big.NewInt(int64(i))
	case "string":
		return s
	default:
		panic("unrecognized type")
	}
}
func compileContracts() map[string]*compiler.Contract {
	os.Chdir("contracts")
	contractFiles, _ := filepath.Glob(CONRACT_GLOB_PATH)
	contracts, err := compiler.CompileSolidity("", contractFiles...)
	os.Chdir("..")
	check(err)
	return contracts
}

func deployContract(contract *compiler.Contract, params []interface{}, conn *ethclient.Client, transactor *bind.TransactOpts) {
	j, err := json.Marshal(contract.Info.AbiDefinition)
	check(err)
	abi, err := abi.JSON(strings.NewReader(string(j)))
	check(err)
	address, tx, _, err := bind.DeployContract(transactor, abi, common.FromHex(contract.Code), conn, params...)
	check(err)
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
}

func check(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
