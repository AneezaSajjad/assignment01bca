package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	head *node
	tail *node
}
type node struct {
	prev         *node
	next         *node
	transaction  string
	nonce        int
	previousHash string
	change       string
}

func (b *block) NewBlock(transaction string, nonce int, previousHash string) *block {
	no := &node{
		transaction:  transaction,
		nonce:        nonce,
		previousHash: previousHash,
		change:       CalculateHash(transaction),
	}
	if b.head == nil {
		b.head = no
		b.tail = no
	} else {
		currentNode := b.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = no
		b.tail = no

	}

	//b.DisplayBlocks()

	return b
}

func (b *block) DisplayBlocks() {
	currentNode := b.head

	if currentNode == nil {
		fmt.Println("node is empty.")

	}
	for currentNode != nil {
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.next
	}

}
func (b *block) ChangeBlock() {
	fmt.Println("the new block chain is")
	currentNode := b.head

	if currentNode == nil {
		fmt.Println("node is empty.")

	}
	for currentNode != nil {
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.next
	}
}
func (b *block) VerifyChain() bool {
	currentNode := b.head
	nextnode := b.head.next
	var temp bool
	temp = true
	for currentNode.next != nil {
		if nextnode.previousHash != currentNode.change {
			temp = false
			break
		}
		currentNode = currentNode.next
		nextnode = nextnode.next
	}
	return temp
}
func CalculateHash(stringToHash string) string {

	fmt.Printf("String Received: %s\n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func main() {
	//bl block
	bl := new(block)

	bl.NewBlock("GHTOIJ", 21, "NULL")

	bl.NewBlock("DETOFG", 11, bl.tail.change)
	bl.NewBlock("ABTOCD", 7, bl.tail.change)
	bl.DisplayBlocks()
	bl.tail.transaction = "knsfdfs"

	bl.ChangeBlock()
	bl.VerifyChain()
	if bl.VerifyChain() == true {
		fmt.Println("there was a change in a blockchain")
	} else {
		fmt.Println("there was no change in a blockchain")
	}
}
