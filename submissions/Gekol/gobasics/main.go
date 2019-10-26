package main

import (
	"fmt"
	"os"
	"strconv"
)

var digits = map[string][]string{}

func printInt(num int) {
	n_string := strconv.Itoa(num)
	var res = [7]string{}
	for pos, c := range n_string {
		for j := 0; j < 7; j++ {
			res[j] += digits[string(c)][j]
			if (pos < len(n_string) - 1) {
				res[j] += " "
			}
		}
	}
	for i := 0; i < 7; i++ {
		fmt.Println(res[i])
	}
}

func main() {
	digits["0"] = []string{" 000 ", "0   0", "0   0", "0   0", "0   0", "0   0", " 000 "}
	digits["1"] = []string{"  1  ", "1 1  ", "  1  ", "  1  ", "  1  ", "  1  ", "11111"}
	digits["2"] = []string{" 222 ", "2   2", "    2", "   2 ", "  2  ", " 2   ", "22222"}
	digits["3"] = []string{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "}
	digits["4"] = []string{"   4 ", "  44 ", " 4 4 ", "4  4 ", "44444", "   4 ", "   4 "}
	digits["5"] = []string{"55555", "5    ", "5    ", "5555 ", "    5", "5   5", " 555 "}
	digits["6"] = []string{" 666 ", "6   6", "6    ", "6666 ", "6   6", "6   6", " 666 "}
	digits["7"] = []string{"77777", "    7", "    7", "   7 ", "  7  ", " 7   ", " 7   "}
	digits["8"] = []string{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "}
	digits["9"] = []string{" 999 ", "9   9", "9   9", " 9999", "    9", "9   9", " 999 "}
	var num int;
	fmt.Print("Enter the number: ")
	fmt.Fscan(os.Stdin, &num)
	printInt(num)
}