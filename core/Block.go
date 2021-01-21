package core

import (
	"fmt"
	"time"
)

type Block struct{
	Timestamp int64 //the block`s createtime
	PrevBlockHash []byte // prev block`s hash 
	Data []byte //store data
	Hash []byte //the block`s hash
	Nonce int     //random number
}
//func (block *Block) SetHash(){
//	//1、timestamp  to []byte
//	//2、attributes other than hash are displayed as []byte
//	//3、hash the datas 256
//	timestamp:=[]byte(strconv.FormatInt(block.Timestamp,10))
//	headers:=bytes.Join([][]byte{block.PrevBlockHash,block.Data,timestamp},[]byte{})
//	hash:=sha256.Sum256(headers)
//	block.Hash=hash[:]
//
//
//}

func NewBlock(data string,PrevBlockHash []byte) *Block{
	//create block
	block :=&Block{
		Timestamp:time.Now().Unix(),
		PrevBlockHash:PrevBlockHash,
		Data:[]byte(data),
		Hash:[]byte{},
	}
	pow:=NewProofOfWork(block)
	nonce,hash:=pow.Run()
	block.Hash=hash[:]
	block.Nonce=nonce
	//set hash
	//block.SetHash();
	//fmt.Println(block.Timestamp);
	//return this block
	isValid:=pow.Validate()
	fmt.Print(isValid)
	return block;
}

func NewGenesisBlock() *Block{
	return NewBlock("Geneis Block",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}