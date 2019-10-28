package main

import (
	"fmt"
	"os"

	//"os"
	"strconv"
)

const digits =
	`
 000 
0   0
0   0
0   0
0   0
0   0
 000 

  1  
1 1  
  1  
  1  
  1  
  1  
11111

 222 
2   2
    2
   2 
  2  
 2   
22222

 333 
3   3
    3
  33 
    3
3   3
 333 

   4 
  44 
 4 4 
4  4 
44444
   4 
   4 

55555
5    
5    
5555 
    5
5   5
 555 

 666 
6   6
6    
6666 
6   6
6   6
 666 

77777
    7
    7
   7 
  7  
 7   
 7   

 888 
8   8
8   8
 888 
8   8
8   8
 888 

 999 
9   9
9   9
 9999
    9
9   9
 999 
`

func write_num(num string, digits []rune) {
	current_digits := []rune(num)
	for j := 0; j < 7; j++ {
		for i := 0; i < len(current_digits) ; i++ {
			digit, _ := strconv.Atoi(string(current_digits[i]))
			index := 43 * digit + 1
			for k := 0; k < 5; k++ {
				fmt.Printf("%c", digits[index + k + j * 6])
			}
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}

func main() {
	symbols := []rune(digits)
	var num string = os.Args[1]
	write_num(num, symbols)
}