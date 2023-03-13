package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string) int {
	re := regexp.MustCompile(`[\+\-\*/]`)
	operator := re.FindString(input)

	cleanInput := strings.Split(input, operator)
	val1 :=
		parser(cleanInput[0])
	val2 :=
		parser(cleanInput[1])

	switch operator {
	case "+":
		fmt.Println(val1 + val2)
		return val1 + val2
	case "-":
		fmt.Println(val1 - val2)
		return val1 - val2
	case "*":
		fmt.Println(val1 * val2)
		return val1 * val2
	case "/":
		fmt.Println(val1 / val2)
		return val1 / val2
	default:
		fmt.Println("Operador no soportado")
		return 0
	}
}

func parser(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("Ingresar operación. Ej 5+7: ")
	input := readInput()

	c := calc{}

	c.operate(input)
}
