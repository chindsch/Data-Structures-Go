package main

import "fmt"

// Structure of the nodes
type Node struct {
	value int
	left *Node
	right *Node
}
	


// Put nodes in or insert nodes into the structure
func (n *Node) insertNode(num int) {
    if n.value > num {
	if n.left == nil {	
	  n.left = &Node{value: num}
 	 } else {
	  	n.left.insertNode(num)
	}
    } else if n.value <= num {
	if n.right == nil {
		n.right = &Node{value: num}
}	else {
		n.right.insertNode(num)
	}
    }
}



// Search values for nodes in trees to determine if they exist
func (n *Node) searchForNode(num int) bool {
	if n == nil {
	  return false
}
	if n.value > num {
		return n.left.searchForNode(num)
	} else if n.value < num {
		return n.right.searchForNode(num)
	}
	
	return true
}


//print main function
func main() {
	ourNode := &Node{value:17}
	fmt.Println(ourNode)
        
	ourNode.insertNode(5)
	fmt.Println(ourNode)
	ourNode.insertNode(300)
	fmt.Println(ourNode)
	ourNode.insertNode(9)
	fmt.Println(ourNode,ourNode.left)
        ourNode.insertNode(200)
        fmt.Println(ourNode,ourNode.right)
        ourNode.insertNode(1)
        fmt.Println(ourNode, ourNode.left)

	
	fmt.Println(ourNode.searchForNode(12))
	fmt.Println(ourNode.searchForNode(5))
	fmt.Println(ourNode.searchForNode(17))
	fmt.Println(ourNode.searchForNode(31245))
}

