package main

import "fmt"
import "./core"
// import "github.com/joho/godotenv"

func main(){
	// godotenv.Load()
	//block:=core.NewBlock("Geneis Block",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//block:=core.NewGenesisBlock()
	blockchain:=core.NewBlockchain()
	blockchain.AddBlock("send 20 OC to Allen")
	blockchain.AddBlock("send 20 OC to Allen2")

	//fmt.Printf("%x\n",block.PrevBlockHash);
	//fmt.Printf("%x\n",block.Hash);
	for _,block:=range blockchain.Blocks{
		fmt.Println(string(block.Data));
		fmt.Printf("PrevHash: %x\n",block.PrevBlockHash);
		fmt.Printf("Hash: %x\n",block.Hash);
	}
}