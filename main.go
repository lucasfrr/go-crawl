package main

import (
	"fmt"
	"lucasfrr/go-crawl/db"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

func main() {
	visitLink("https://aprendagolang.com.br")
}

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

func visitLink(link string) {
	fmt.Printf("Visiting link: %s\n", link)
	response, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status not equal than 200: %d", response.StatusCode))
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

			if db.VisitedLink(link.String()) {
				fmt.Printf("Link was visited: %s\n", link)
				continue
			}

			visitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}

			db.Insert("links", visitedLink)

			visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}
