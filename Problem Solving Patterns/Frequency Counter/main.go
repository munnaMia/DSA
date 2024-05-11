package main

import "fmt"

func main() {
	fmt.Println(anagram("helol", "hello"))
}

func anagram(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	var track1 = make(map[string]int)
	var track2 = make(map[string]int)

	for _, v := range str1 {
		_, ok := track1[string(v)]
		if ok {
			track1[string(v)]++
		} else {
			track1[string(v)] = 1
		}
	}

	for _, v := range str2 {
		_, ok := track2[string(v)]
		if ok {
			track2[string(v)]++
		} else {
			track2[string(v)] = 1
		}
	}

	for i := range track1 {
		if track1[i] != track2[i] {
			return false
		}
	}

	return true
}
