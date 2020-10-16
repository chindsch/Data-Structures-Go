package main

import "fmt"

type Stack struct {
	items []int
}

//Push data into stack
func (ourStack *Stack) pushIntoStack(num int){
	ourStack.items = append(ourStack.items,num)
}

//Pop data off the stack
func (ourStack *Stack) popOutOfStack() int {
	sizeofslice := len(ourStack.items)-1
	removedItem := ourStack.items[sizeofslice]
	ourStack.items := ourStack.items[:sizeofslice - 1]
	return removedItem
	
}


func main() {
	newStack := Stack{}

	newStack.pushIntoStack(5)
        newStack.pushIntoStack(25)
	newStack.pushIntoStack(340)
	fmt.Printf("%+v\n", newStack)
} 
