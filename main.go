package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter math expression and press enter to calculate: ")
	scanner.Scan()
	exp := scanner.Text()
	splitOfExp := strings.Split(exp,"+")
	elOne, _ := strconv.Atoi(splitOfExp[0])
	elTwo, _ := strconv.Atoi(splitOfExp[1])

	fmt.Println(elOne + elTwo)
}