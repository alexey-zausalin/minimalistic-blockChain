package blockchain

import (
	"crypto/md5"
	"math/rand"
	"time"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
	Nonce    int
}

func CreateBlock(data string, prevHash string) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(10000)

	block := &Block{
		Hash:     "",
		Data:     data,
		PrevHash: prevHash,
		Nonce:    initialNonce,
	}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

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
