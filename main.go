// -*- compile-command: "tinygo test ./..."; -*-

package main

import (
	"strings"

	"github.com/extism/go-pdk"
)

const (
	defaultVowels = "aeiouAEIOU"
)

// VowelReport represents the JSON struct returned to the host.
type VowelReport struct {
	Count  int
	Total  int
	Vowels string
}

//go:export count_vowels
func CountVowels() int32 {
	input := pdk.InputString()

	vowels := getVowels()
	var count int
	for _, r := range input {
		if strings.ContainsRune(vowels, r) {
			count++
		}
	}

	total := getTotal() + count
	storeTotal(total)

	result := VowelReport{
		Count:  count,
		Total:  total,
		Vowels: vowels,
	}

	pdk.OutputJSON(result)
	return 0
}

func getTotal() int {
	return pdk.GetVarInt("total")
}

func storeTotal(total int) {
	pdk.SetVarInt("total", total)
}

func getVowels() string {
	vowels, ok := pdk.GetConfig("vowels")
	if !ok || vowels == "" {
		return defaultVowels
	}
	return vowels
}

func main() {}
