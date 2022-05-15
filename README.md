# githubtrend
Crawl [GitHub Trending](https://github.com/trending) to use something like daily chat bot

# Usage

write the code like this.

```
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
```

```
go run cmd/main.go
100Star https://github.com/example/example
99Star https://github.com/example/example2
...
```