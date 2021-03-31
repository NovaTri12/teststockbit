package main

import (
	"fmt"
	"strings"
	"sort"
)

var (
	anagramList = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func sortString(keyString string) string {
	stringAnagram := strings.Split(keyString, "")
	sort.Strings(stringAnagram)

	return strings.Join(stringAnagram, "")
}

func main() {
	list := make(map[string][]string)

	for _, word := range anagramList {
		key := sortString(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}