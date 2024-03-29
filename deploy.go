package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/sabbas/inbox/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"fmt"
	"os"
)

const key  = `paste the contents of your JSON key file here`

func main(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/fYe8qCnWi6TXZAXOVof9")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "passphrase associated with your JSON key file")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _:= contracts.DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
}