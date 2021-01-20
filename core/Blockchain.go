package core

type Blockchain struct {
	Blocks []*Block //store block
}

//add a new block
func (blockchain *Blockchain) AddBlock(data string){
	//1、create a new block
	//2、add blockchain
	preBlock:=blockchain.Blocks[len(blockchain.Blocks)-1]
	newblock:=NewBlock(data,preBlock.Hash);
	blockchain.Blocks=append(blockchain.Blocks,newblock)
}

//create the geneis blockchain
func NewBlockchain() *Blockchain{
	return  &Blockchain{[]*Block{NewGenesisBlock()}}
}