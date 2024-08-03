package blockchain

import (
	"crypto/md5"
	"math/rand"
	"time"
)

type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

func CreateBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(10000)

	block := &Block{
		Hash:         "",
		Data:         data,
		PrevHash:     prevHash,
		Nonce:        initialNonce,
		Transactions: transactions,
	}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}

	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}

func (b *Block) ComputeHash() {
	hasher := md5.New()

	hasher.Write([]byte(b.Data))
	hasher.Write([]byte(b.PrevHash))

	b.Hash = string(hasher.Sum([]byte{}))
}
