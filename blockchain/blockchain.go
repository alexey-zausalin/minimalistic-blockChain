package blockchain

type Blockchain struct {
	Blocks []*Block
}

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

func InitBlockChain() *Blockchain {
	block := Genesis()
	return &Blockchain{[]*Block{block}}
}

func (bc *Blockchain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseRcpt,
		Amount:   10.0,
		Coinbase: true,
	}

	newBlock := CreateBlock(data, prevBlock.Hash, append([]*Transaction{coinbaseTransaction}, transactions...))

	bc.Blocks = append(bc.Blocks, newBlock)
}
