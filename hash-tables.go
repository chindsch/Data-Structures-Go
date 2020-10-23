package main

import (
	"fmt"
)

const sizeOfArray = 10


func hash(name string) int {
 //interate over a string and cast to integer
 letterSum := 0
  for _, letter := range name {
	letterSum += int(letter)
 }

  return letterSum % sizeOfArray

}

func main() {
	fmt.Println(hash("hello"))
}
