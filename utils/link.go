package utils

import (
	"fmt"
	"lucasfrr/go-crawl/db"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"golang.org/x/net/html"
)

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

func VisitLink(link string) {
	fmt.Printf("Visiting link: %s\n", link)
	response, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("status not equal than 200: %d\n", response.StatusCode)
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		panic(err)
	}

	extractLinks(doc)
}

func extractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}

			match, err := regexp.MatchString("(http)", link.String())
			if err != nil {
				fmt.Println(err)
				continue
			}

			if !match {
				fmt.Printf("Link %s not valid. Go to next link\n", link.String())
				continue
			}

			if db.VisitedLink(link.String()) {
				fmt.Printf("Link ->%s<- already visited\n", link)
				continue
			}

			visitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}

			db.Insert("links", visitedLink)

			go VisitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}
