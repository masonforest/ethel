package main

import (
    "fmt"
    "log"
    "strings"
    "time"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"c8a8b35742b10bdc00a724c6b76d5f85b6558b0c","crypto":{"cipher":"aes-128-ctr","ciphertext":"214a1ff15bbae62ca2dd90d2333cb23ee615df3f173d5158b1b89cf3329b18b6","cipherparams":{"iv":"5680a66256d7d09abaa7e95063a8c375"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"72760c070ebf2600a7200aa71158ff1a583d7a1676cbf8b1df3a3dc2e0bbf10f"},"mac":"82fbafdcdb69d78f6f3e9bd35d6a333d6518c425e0cf9dd4804778047e926bcc"},"id":"f972e6e1-ef3f-4bdb-b897-623797517c1b","version":3}`

func main() {
    // Create an IPC based RPC connection to a remote node and an authorized transactor
    conn, _ := ethclient.Dial("http://localhost:8545/")
    auth, err := bind.NewTransactor(strings.NewReader(key), "hNJTZzGTxXmz9KzdDTvfqnCS")
    if err != nil {
        log.Fatalf("Failed to create authorized transactor: %v", err)
    }
    // Deploy a new awesome contract for the binding demo
    address, tx, token, err := DeployToken(auth, conn)
    if err != nil {
        log.Fatalf("Failed to deploy new token contract: %v", err)
    }
    fmt.Printf("Contract pending deploy: 0x%x\n", address)
    fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

    // Don't even wait, check its presence in the local pending state
    time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

    name, err := token.Name(&bind.CallOpts{Pending: true})
    if err != nil {
        log.Fatalf("Failed to retrieve pending name: %v", err)
    }
    fmt.Println("Pending name:", name)
}
