package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	sum := 0

	romanBytes := []byte(s)

	for idx := 0; idx < len(romanBytes); idx++ {

		if idx < len(romanBytes)-1 {

			if v, ok := romanMap[string(romanBytes[idx:idx+2])]; ok {
				sum += v
				idx += 2
				continue
			}
		}
		sum += romanMap[string(romanBytes[idx])]

	}
	return sum
}

func main() {
	number := "XI"
	newnumber := romanToInt(number)
	fmt.Println(newnumber)
}
