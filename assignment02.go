package assignment02

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

// global variable to inita
var initialize_rand = true

func GenerateNonce(blockData []Transaction) int {
	if initialize_rand {
		//seeding the rand function so will get different value every time
		rand.Seed(time.Now().UnixNano())
		initialize_rand = false // making it false bcz rand seed method already initialized
	}
	return rand.Intn(1000) + len(blockData)
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	var block Block
	block.BlockData = blockData
	block.PrevPointer = chainHead
	block.Nonce = GenerateNonce(blockData)
	block.CurrentHash = CalculateHash(block.BlockData, block.Nonce)
	if chainHead != nil {
		block.PrevHash = chainHead.CurrentHash
	} else {
		block.PrevHash = ""
	}
	return &block

}

func ListBlocks(chainHead *Block) {
	var counter = 1
	for chainHead != nil {
		fmt.Println("##############################################################################################")
		fmt.Println("				    Block ", counter)
		fmt.Println("----------------------------------------------------------------------------------------------")
		fmt.Println("Block Transactions : ")
		DisplayTransactions(chainHead.BlockData)
		fmt.Println("Block Nonce : ", chainHead.Nonce)
		fmt.Println("Block Hash                : " + chainHead.CurrentHash)
		fmt.Println("Block Previous Hash Value : " + chainHead.PrevHash)
		fmt.Println("##############################################################################################")
		chainHead = chainHead.PrevPointer
		counter = counter + 1
	}
}

func DisplayTransactions(blockData []Transaction) {
	for i := 0; i < len(blockData); i = i + 1 {
		fmt.Println("	****************************************************************")
		fmt.Println("			Transaction Sender   : " + blockData[i].Sender)
		fmt.Println("			Transaction Receiver : " + blockData[i].Receiver)
		fmt.Println("			Transaction Amount   : ", blockData[i].Amount)
		fmt.Println("	****************************************************************")
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	var transaction Transaction
	transaction.Sender = sender
	transaction.Receiver = receiver
	transaction.Amount = amount
	return transaction
}
