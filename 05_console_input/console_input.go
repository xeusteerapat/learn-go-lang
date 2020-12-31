package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the your birth year: ")
	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("You will be %d year old at the end of 2020", 2020-input)
}
