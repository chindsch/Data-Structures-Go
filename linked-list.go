package main

import "fmt"

// Implement the structure of the node data type
type node struct {
	data string
	nextElement *node
}


// Implement the structure of a list
type linkedList struct {
	head *node
	length int
}


// Add items to list
func (myList *linkedList) addNodeToList(n *node) {
	//Create pointer from the current new head node to the next item in list
	n.nextElement = myList.head
	//set up new head in the list
	myList.head =  n
	//Update the length of the list
	myList.length++
}


// Delete items from list
func (myList *linkedList) deleteNodeFromList(node int) {
	currentHead := myList.head
	for i:= 1; i < node-1; i++ {
		currentHead = currentHead.nextElement
	}
	//reassigining the next element
	currentHead.nextElement = currentHead.nextElement.nextElement
	myList.length--
}

//See whats in the list
func (myList linkedList) listNodes() {
	printMe := myList.head
	for myList.length > 0 {
		fmt.Printf("%v\n", printMe.data)
		printMe = printMe.nextElement
		myList.length--
	}
}




func main() {
	ourList := linkedList{}
	// dataList := []string{"one","two","three"}
	ourList.addNodeToList(&node{data: "one"})
	ourList.addNodeToList(&node{data: "two"})
	ourList.addNodeToList(&node{data: "three"})
	ourList.addNodeToList(&node{data: "four"})


	ourList.deleteNodeFromList(2)

	fmt.Printf("%+v\n", ourList)

	ourList.listNodes()

}















