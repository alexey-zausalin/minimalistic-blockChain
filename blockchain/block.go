package blockchain

import (
	"crypto/md5"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
}

func (b *Block) ComputeHash() {
	hasher := md5.New()

	hasher.Write([]byte(b.Data))
	hasher.Write([]byte(b.PrevHash))

	b.Hash = string(hasher.Sum([]byte{}))
}
