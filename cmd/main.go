package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// to connect any account to make transaction
func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		fmt.Println("HERE?")
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) //Type assertion := used to convert the PublicKey to an *ecdsa.PublicKey
	if !ok {
		panic("invalid key")
	}
	//derive the ethereum address from the public key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetching the last used nonce of the account
	//this is the nonce to be used in the next transaction
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	//getting the chain ID of the ethereum network
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	//creating a new transaction authorization object using the private key and the chian id
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(1000000)

	return auth

}

func main() {
	//address of ethereum env
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		panic(err)
	}
	fmt.Println("client", client)
	auth := getAccountAuth(client, "878d3c676514439710c33b49d8246715933b75ccd6a8847a0fa5c71ed0704957")

	fmt.Println("auth", auth)
}
