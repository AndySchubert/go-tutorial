package main

import "fmt"

// Arraysize is the size of the hashtable array
const ArraySize = 7

// Hashtable structure
type Hashtable struct {
	array [ArraySize]*bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

// bucketnode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert
func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete
func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("%v already exists!\n", k)
	}

}

// Search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// Hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	fmt.Println("STAN", hashTable.Search("STAN"))
	fmt.Println("KENNY", hashTable.Search("KENNY"))
}
