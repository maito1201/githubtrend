package main

import (
	"fmt"

	"github.com/maito1201/githubtrend"
)

func main() {
	result, err := githubtrend.ScrapeGitHubTrend()
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		fmt.Printf("%dStar %s\n", v.Star, v.Href)
	}
}
