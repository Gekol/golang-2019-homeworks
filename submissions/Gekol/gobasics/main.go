package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	digitLength   = 43
	digitHeight   = 7
	digitWidth    = 5
	digitInterval = 6
)

const digits = `
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

func writeAsterisk(num string) {
	for i := 0; i < len(num)*digitInterval - 1; i++ {
		fmt.Printf("*")
	}
	fmt.Println("")
}

func writeNum(num string, digits []rune) {
	writeAsterisk(num)
	for j := 0; j < digitHeight; j++ {
		for _, i := range num {
			digit, _ := strconv.Atoi(string(i))
			index := digitLength*digit + 1
			for k := 0; k < digitWidth; k++ {
				fmt.Printf("%c", digits[index+k+j*digitInterval])
			}
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
	writeAsterisk(num)
}

func main() {
	symbols := []rune(digits)
	var num string = os.Args[1]
	if _, err := strconv.Atoi(num); err == nil {
		writeNum(num, symbols)
	} else {
		fmt.Println("Invalid input!!!")
	}
}
