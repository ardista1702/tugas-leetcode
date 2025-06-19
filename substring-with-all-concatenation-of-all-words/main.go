package main

import "fmt"

func getFrequencyOfWords(words []string) map[string]int {
	result := map[string]int{}
	for _, word := range words {
		result[word]++
	}
	return result
}

func findSubstring(s string, words []string) []int {
	result := []int{}
	wordLen := len(words[0])
	totalLen := wordLen * len(words)
	frequencyOfWords := getFrequencyOfWords(words)

	memo := map[string]bool{}

	for i := 0; i <= len(s)-totalLen; i++ {
		substring := s[i : i+totalLen]

		if exist, foundSubstring := memo[substring]; foundSubstring {
			if exist {
				result = append(result, i)
			}
			continue
		}

		seen := map[string]int{}
		valid := true
		for j := 0; j < totalLen; j += wordLen {
			word := substring[j : j+wordLen]
			seen[word]++

			if seen[word] > frequencyOfWords[word] {
				valid = false
				break
			}

		}
		memo[substring] = valid

		if valid && len(frequencyOfWords) == len(seen){
			result = append(result, i)
		}
	}

	return result
}

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar","foo","the"}
	result := findSubstring(s,words)
	fmt.Println(result)
}