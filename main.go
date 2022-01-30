package main

import (
	"bufio"
	"fmt"
	"math"
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

func rainbow(freq float64, i int) (int, int, int) {
	phase := freq * float64(i)
	var red int = int(math.Sin(phase+0)*127.0 + 128.0)
	var green int = int(math.Sin(phase+2.0*math.Pi/3.0)*127.0 + 128.0)
	var blue int = int(math.Sin(phase+4.0*math.Pi/3.0)*127.0 + 128.0)
	//"#%02X%02X%02X" % [ red, green, blue ]
	return red, green, blue
}

func roflize(scanner *bufio.Scanner) {
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
        lineCount += 1

		for index, runeValue := range line {
			red, green, blue := rainbow(0.1, index + lineCount)
			fmt.Printf("\033[38;2;%d;%d;%dm%s", red, green, blue, string(runeValue))
			//fmt.Printf("\\033[38;2;%d;%d;%dm%s\n", red, green, blue, string(runeValue))
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	roflize(scanner)
}
