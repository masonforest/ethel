package commands

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/pborman/uuid"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var configFile string

type Config struct {
	Testnet string `json:"testnet"`
	Live    string `json:"live"`
}

func Init(c *cli.Context) error {
	testnetKey, testnetJSON := createKey("testnet")
	liveKey, liveJSON := createKey("live")

	config := &Config{
		Testnet: string(testnetJSON),
		Live:    string(liveJSON),
	}

	configJSON, err := json.Marshal(config)
	check(err)
	writeKeyFile(getConfigFile(), configJSON)

	fmt.Println("Created .ethel")
	fmt.Printf("Address: %s\n", liveKey.Address.Hex())
	fmt.Printf("Testnet Address: %s\n", testnetKey.Address.Hex())
	return nil
}

func writeKeyFile(file string, content []byte) error {
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()
	return os.Rename(f.Name(), file)
}

func askString(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + ": ")
	s, _ := reader.ReadString('\n')
	return strings.Trim(s, "\n")
}

func askBool(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + " (y/n): ")
	s, _ := reader.ReadString('\n')

	if strings.Trim(s, "\n") == "y" {
		return true
	} else if strings.Trim(s, "\n") == "n" {
		return false
	} else {
		return askBool(question)
	}
}

func createKey(network string) (*accounts.Key, []byte) {
	if askBool("Encrypt " + network + " key?") {
		return createEncryptedKey(rand.Reader)
	} else {
		return createPlainTextKey(rand.Reader)
	}
}

func getConfigFile() string {
	usr, err := user.Current()

	check(err)

	return usr.HomeDir + "/.ethel.json"
}

func createEncryptedKey(rand io.Reader) (*accounts.Key, []byte) {
	privateKeyECDSA, err := ecdsa.GenerateKey(secp256k1.S256(), rand)
	check(err)
	key := newKeyFromECDSA(privateKeyECDSA)
	password := askString("Password:")
	json, err := accounts.EncryptKey(
		key,
		password,
		accounts.StandardScryptN,
		accounts.StandardScryptP,
	)
	check(err)
	return key, json
}

func createPlainTextKey(rand io.Reader) (*accounts.Key, []byte) {
	privateKeyECDSA, err := ecdsa.GenerateKey(secp256k1.S256(), rand)
	check(err)
	key := newKeyFromECDSA(privateKeyECDSA)
	json, err := key.MarshalJSON()
	check(err)
	return key, json
}

func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *accounts.Key {
	id := uuid.NewRandom()
	key := &accounts.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}
