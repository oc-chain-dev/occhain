package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

type ProofOfwork struct{
	block *Block
	target *big.Int  //store difficulty
}
const targetBits=20
var (
	maxNonce=math.MaxInt64
)

func  NewProofOfWork(block *Block) *ProofOfwork {
	target:=big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
	pow :=&ProofOfwork{block,target}
	return  pow
}

func (pow *ProofOfwork) prepareData(nonce int) []byte{
	data :=bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
		)
	return  data
}


func (pow *ProofOfwork) Run() (int,[]byte){
	nonce:=0
	var hashInt big.Int
	var hash [32]byte

	fmt.Printf("Minig the block containing \"%s\"\n",pow.block.Data)
	//var hashInt
	for nonce<maxNonce{
		data:=pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target)==-1{
			break
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}

//Verify that the pow is valid
func (pow *ProofOfwork) Validate() bool{
	var hashInt big.Int
	data :=pow.prepareData(pow.block.Nonce)
	hash:=sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid:=hashInt.Cmp(pow.target)==-1
	return isValid
}