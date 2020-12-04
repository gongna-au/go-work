package main

import "fmt"

func Repalce(s string, target, replace byte) string {

	slice := []byte(s)

	if len(slice) == 0 {
		return ""
	}
	for i, value := range slice {

		if value == target {
			slice[i] = replace

		}
	}

	result := string(slice)
	return result
}

func main() {
	s := "You are interesting."
	s = Repalce(s, 'i', 'e')
	s1 := ""
	s1 = Repalce(s1, 'i', 'e')
	fmt.Println(s, s1)

}
