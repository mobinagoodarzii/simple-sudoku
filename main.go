package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomList(input int) [][]int {
	rand.Seed(time.Now().Unix())
	var allRows [][]int
	for len(allRows) < input {
		isDupVal := false
		var rowValues []int
		for len(rowValues) < input {
			isDuplicate := false
			num := rand.Intn(input) + 1
			for _, eachVal := range rowValues {
				if num == eachVal {
					isDuplicate = true
					break
				}
			}
			if isDuplicate == false {
				rowValues = append(rowValues, num)
			}
		}
		for _, eachSlice := range allRows {
			for i := 0; i < len(eachSlice); i++ {
				if rowValues[i] == eachSlice[i] {
					isDupVal = true
					break
				}
			}
			if isDupVal == true {
				break
			}
		}
		if isDupVal == false {
			allRows = append(allRows, rowValues)
			//fmt.Println("rowValues:", rowValues) //(to test)
			//fmt.Println("allRows:", allRows) //(to test)
			//fmt.Println(len(allRows)) //(to test)
			if len(allRows) > 1 {
				y := len(allRows)
				for x := 1; x < 10; x++ {
					startOfY := y - ((y - 3*((y-1)/3)) - 1)
					startOfX := x - ((x - 3*((x-1)/3)) - 1)
					for j := startOfY; j < startOfY+((y-3*((y-1)/3))-1); j++ {
						for i := startOfX; i < startOfX+3; i++ {
							//fmt.Println("y is:", y, "x is:", x) //(to test)
							//fmt.Println(allRows[y-1][x-1], allRows[j-1][i-1], j, i) //(to test)
							if allRows[y-1][x-1] == allRows[j-1][i-1] {
								isDupVal = true
								allRows = allRows[:len(allRows)-1]
								break
							}
						}
						if isDupVal == true {
							break
						}
					}
					if isDupVal == true {
						break
					}
				}
			}
		}
	}
	fmt.Println(allRows)
	return allRows
}

func makeTable(input int, tableValues [][]int) {
	fmt.Print("┌")
	for i := 0; i < input-1; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("-")
		}
		fmt.Print("┬")
	}
	for j := 0; j < 3; j++ {
		fmt.Print("-")
	}
	fmt.Print("┐\n")
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			fmt.Print("|")
			fmt.Print(" ")
			fmt.Print(tableValues[i][j])
			fmt.Print(" ")
		}
		fmt.Print("|\n")
		fmt.Print("└")
		for j := 0; j < (input)-1; j++ {
			for k := 0; k < 3; k++ {
				fmt.Print("-")
			}
			fmt.Print("┴")
		}
		for k := 0; k < 3; k++ {
			fmt.Print("-")
		}
		fmt.Print("┘")
		fmt.Println()
	}
}
func main() {
	fmt.Println("Let's make a sudoku table!!")
	fmt.Println("Enter the number of rows and columns:")
	//The program is coded for number 9.
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	makeTable(input, randomList(input))
}

//Note

//برای y = 3 و x = 3
//x = 2 , y = 3 لازم نیس
//x = 1 , y = 3 لازم نیس
//x = 3 , y = 2 چک میکنه
//x = 2 , y = 2 چک میکنه
//x = 1 , y = 2 چک میکنه
//x = 3 , y = 1 چک میکنه
//x = 2 , y = 1 چک میکنه
//x = 1 , y = 1 چک میکنه

//برای y = 9 و x = 9
//x = 8 , y = 9 لازم نیس
//x = 7 , y = 9 لازم نیس
//x = 9 , y = 8 چک میکنه
//x = 8 , y = 8 چک میکنه
//x = 7 , y = 8 چک میکنه
//x = 9 , y = 7 چک میکنه
//x = 8 , y = 7 چک میکنه
//x = 7 , y = 7 چک میکنه

//برای y = 2 و x = 8
//x = 7 , y = 2 لازم نیس
//x = 9 , y = 1 چک میکنه
//x = 8 , y = 1 چک میکنه
//x = 7 , y = 1 چک میکنه

//برای y = 4 و x = 5
//x = 4 , y = 4 لازم نیس

//برای y = 8 و x = 3
//x = 2 , y = 8 لازم نیس
//x = 1 , y = 8 لازم نیس
//x = 3 , y = 7 چک میکنه
//x = 2 , y = 7 چک میکنه
//x = 1 , y = 7 چک میکنه

//1 - 3 * 0
//2 - 3 * 0
//3 - 3 * 0
//4 - 3 * 1
//5 - 3 * 1
//6 - 3 * 1
//7 - 3 * 2
//8 - 3 * 2
//9 - 3 * 2
//y - 3 *(y-1)/3
// len(y) = y - 3 * (y-1)/3

//1 + 2 -> x +(1*3 - x)
//2 + 1 -> x +(1*3 - x)
//3 + 0 -> x +(1*3 - x)
//4 + 2 -> x +(2*3 - x)
//5 + 1 -> x +(2*3 - x)
//6 + 0 -> x +(2*3 - x)
//7 + 2 -> x +(3*3 - x)
//8 + 1 -> x +(3*3 - x)
//9 + 0 -> x +(3*3 - x)
//x+((x+2)/3)*3 - x)

//1 - (1-1)0 = 1
//2 - (2-1)1 = 1
//3 - (3-1)2 = 1
//4 - ((x - 3 * (x-1)/3)-1)0 = 4
//5 - ((x - 3 * (x-1)/3)-1)1 = 4
//6 - 2 = 4
//7 - 0 = 7
//8 - 1 = 7
//9 - 2 = 7
