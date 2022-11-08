//program multiply a number using function with pointers

package main

import (
	"fmt"
)

func doubleValue(number *int) {
	*number *= 2
}

func main() {
	number := 5
	doubleValue(&number)
	fmt.Println("Double value:", number)
}
