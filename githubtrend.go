package githubtrend

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	Href string
	Star int
}

var numberRegexp = regexp.MustCompile("\\d+")

func ScrapeGitHubTrend() ([]ScrapeResult, error) {
	var ret []ScrapeResult
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	document.Find(".Box-row").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Find("a").First().Attr("href")
		if !ok {
			return
		}
		href = strings.Replace(href, "/login?return_to=", "", 1)
		href, err = url.PathUnescape(href)
		if err != nil {
			return
		}
		star := numberRegexp.FindString(s.Find("span.d-inline-block.float-sm-right").Text())
		starNum, err := strconv.Atoi(star)
		if err != nil {
			return
		}
		res := ScrapeResult{
			Href: fmt.Sprintf("https://github.com%s", href),
			Star: starNum,
		}
		ret = append(ret, res)
	})

	sort.Slice(ret, func(i, j int) bool { return ret[i].Star > ret[j].Star })
	return ret, nil
}
