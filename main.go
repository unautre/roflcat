package main

import (
	"bufio"
	"fmt"
	"os"
)

const RESET = 0
const BOLD = 1
const DIM = 2
const ITALIC = 3
const UNDERLINE = 4

const BG = 10

const BLACK = 30
const RED = 31
const GREEN = 32
const YELLOW = 33
const BLUE = 34

func roflize(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()

		for index, runeValue := range line {
			fmt.Printf("\033[%dm%s", (index%10)+30, string(runeValue))
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	roflize(scanner)
}
