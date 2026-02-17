package main

import (
	"strings"
	"slices"
)

func filterWords(s string) string {
	
	badWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	sSplit := strings.Split(s, " ")
	sLowSplit := strings.Split(strings.ToLower(s), " ")

	for i, word := range(sLowSplit) {
		if slices.Contains(badWords, word) {
			sSplit[i] = "****"
		}
	}


	return strings.Join(sSplit, " ")
}
