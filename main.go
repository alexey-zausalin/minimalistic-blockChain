package main

import (
	"fmt"
	"minimalistic-blockChain/blockchain"
)

func main() {
	bc := blockchain.InitBlockChain()
	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")
	bc.AddBlock("Block 3")

	fmt.Println("BlockChain:")
	for _, block := range bc.Blocks {
		fmt.Printf("\tPrevHash: %x\n", block.PrevHash)
		fmt.Printf("\tData: %s\n", block.Data)
		fmt.Printf("\tHash: %x\n", block.Hash)
		fmt.Println("----------------------------------------")
	}
}
