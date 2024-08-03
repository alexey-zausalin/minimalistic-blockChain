package blockchain

import (
	"crypto/md5"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
}

func CreateBlock(data string, prevHash string) *Block {
	block := &Block{Hash: "", Data: data, PrevHash: prevHash}
	block.ComputeHash()

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}

func (b *Block) ComputeHash() {
	hasher := md5.New()

	hasher.Write([]byte(b.Data))
	hasher.Write([]byte(b.PrevHash))

	b.Hash = string(hasher.Sum([]byte{}))
}
