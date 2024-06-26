package main

import (
	"fmt"

	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
)

func main() {

	compareWordSet := []struct {
		str1 string
		str2 string
	}{
		{"22.04", "22.04.1"},
		{"22.04", "20.04"},
		{"20.04", "18.04"},
		{"x86_64", "amd64"},
		{"hvm-ssd", "ssd"},
		{"hvm-ssd", "hdd"},
	}

	for _, set := range compareWordSet {
		fmt.Printf("Comparing '%s' with '%s':\n", set.str1, set.str2)
		fmt.Printf(" - LevenshteinDistance, Similarity ratio: %.2f\n", recommendation.CalculateSimilarityByLevenshteinDistance(set.str1, set.str2))
		fmt.Printf(" - SequenceMatcher, Similarity ratio: %.2f\n", recommendation.CalculateSimilarityBySequenceMatcher(set.str1, set.str2))
		fmt.Println("--------------------------------------------------------")
	}

	keywords := "Ubuntu 22.04.4 LTS (Jammy Jellyfish) x86_64 SSD"
	vmImages := []string{
		"ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20220609",
		"ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-20191002",
	}

	// Select VM OS image via LevenshteinDistance-based text similarity
	delimiters1 := []string{" ", "-", "_", ",", "(", ")", "[", "]", "/"}
	delimiters2 := delimiters1

	for _, image := range vmImages {
		fmt.Printf("Comparing keywords with VM Image:\n")
		fmt.Printf("Keywords: '%s'\n", keywords)
		fmt.Printf("VM Image: '%s'\n", image)
		score := recommendation.CalculateSimilarity(keywords, delimiters1, image, delimiters2)
		fmt.Printf(" - Similarity Score: %.2f\n", score)
		fmt.Println("--------------------------------------------------------")
	}
}
