package blockchain

type Blockchain struct {
	Blocks []*Block
}

func InitBlockChain() *Blockchain {
	block := Genesis()
	return &Blockchain{[]*Block{block}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, block)
}
