package main

import (
	"fmt"
	"strings"
)

func main() {

	var content string = "As far as eye could reach he saw nothing but the stems of the great plants about him filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he receding in the violet shade, and far overhead the multiple transparency of huge leaves felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or  knowledge of man."
	loweredString := strings.ToLower(content)
	slicedContent := strings.Fields(loweredString)
	var slicedWOcomma []string
	var slicedWOsemicolon []string
	var slicedWOperiod []string
	var wordsCount = make(map[string]int)

	for _, item := range slicedContent {
		slicedWOcomma = append(slicedWOcomma, strings.Trim(item, ","))
	}
	for _, item := range slicedWOcomma {
		slicedWOsemicolon = append(slicedWOsemicolon, strings.Trim(item, ";"))
	}
	for _, item := range slicedWOsemicolon {
		slicedWOperiod = append(slicedWOperiod, strings.Trim(item, "."))
	}

	for _, item := range slicedWOperiod {
		count, ok := wordsCount[item]
		if ok {
			wordsCount[item] = count + 1
		} else {
			wordsCount[item] = count
		}
	}

	for k, v := range wordsCount {
		wordsCount[k] = v + 1
	}

	fmt.Println(wordsCount)
}
