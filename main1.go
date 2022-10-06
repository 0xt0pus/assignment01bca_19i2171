package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash        []byte
	PrevHash    []byte
	Nonce       []byte
	Transaction []byte
}

type BlockChain struct {
	blocks []*Block
}

func CalculateHash(stringToHash []byte) []byte {
	sum := sha256.Sum256((stringToHash))
	return sum[:] //convert int[32] to full byte and returned
}

func createBlock(prevhash []byte, nonce string, transaction string) *Block {

	hash := bytes.Join([][]byte{prevhash, []byte(nonce), []byte(transaction)}, []byte{})
	block := &Block{(CalculateHash(hash)), prevhash, []byte(nonce), []byte(transaction)}
	return block

}

func (chain *BlockChain) NewBlock(trans string, nonc string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	// fmt.Printf("BOOM %x", prevBlock.PrevHash)

	new := createBlock(prevBlock.Hash, nonc, trans)
	chain.blocks = append(chain.blocks, new)
}

func InitBlockChain() *BlockChain {

	block := createBlock([]byte{}, "1234", "Genisis")
	bchain := BlockChain{[]*Block{block}}
	return &bchain
	// return &BlockChain{[]*Block{(createBlock([]byte{}, "1234", "Genisis"))}}
}

func ListBlocks(chain *BlockChain){


	for _, block := range chain.blocks {

		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %s\n", block.Nonce)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Print("\n")

	}

}

func VerifyChain(chain *BlockChain) {

	
	i:=0;

	for _, block := range chain.blocks {

		if(i==len(chain.blocks)-1){
			break;
		}
		
		if(bytes.Compare(chain.blocks[i+1].PrevHash, block.Hash)!=0){

			fmt.Println("verification problems")
		}
		// else
		// {
		// 	fmt.Printf("%x==%x\n",chain.blocks[i+1].PrevHash, block.Hash);

		// }		
		 	i++;
	}

	fmt.Println("BlockChain Verified")

}


func main() {
	// x := CalculateHash("Hello")
	// str := string(x[:])
	// fmt.Printf("%x", str)

	chain := InitBlockChain()
	chain.NewBlock("Alice To Bob", "87")
	chain.NewBlock("Bob To Trudy", "433432")
	chain.NewBlock("Trudy to Ahmed", "7362323")
	chain.NewBlock("Trudy to James", "23231")
	chain.NewBlock("James to Alice", "6575")
	chain.NewBlock("Bob to Ahmed", "09344")
	ListBlocks(chain)
	VerifyChain(chain)
	
}

