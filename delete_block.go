package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/c313f9648c3742528bae0c24717d45aa")
	if err != nil {
		return
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("WALLET_PRIVATE_KEY"))
	if err != nil {
		return
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	// execute contract_deploy.go to obtain the contract address
	address := common.HexToAddress("0x43242b174f617489977742F6c988636517907436")
	instance, err := NewAccessControlIPFS(address, client)
	sender := common.HexToAddress(os.Args[1])
	block := os.Args[2]
	if len(block) > 32 {
		block_1 := [32]byte{}
		block_2 := [32]byte{}
		copy(block_1[:], []byte(block[:32]))
		copy(block_2[:], []byte(block[32:]))
		transaction, err := instance.DeleteBlockMultiple(
			auth,
			sender,
			[][32]byte{block_1, block_2},
		)
		if err != nil {
			fmt.Println("Delete Operation Failed!")
			return
		}
		fmt.Println("Delete Operation Successful.")
		fmt.Println("Transaction ID: ", transaction.Hash().Hex())
	} else {
		block_1 := [32]byte{}
		copy(block_1[:], []byte(block))
		transaction, err := instance.DeleteBlock(
			auth,
			sender,
			block_1,
		)
		if err != nil {
			fmt.Println("Delete Operation Failed!")
			return
		}
		fmt.Println("Delete Operation Successful.")
		fmt.Println("Transaction ID: ", transaction.Hash().Hex())
	}
}
