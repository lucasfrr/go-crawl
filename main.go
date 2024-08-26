package main

import (
	"flag"
	"lucasfrr/go-crawl/utils"
)

var link string

func init() {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br", "url to init requests")
}

func main() {
	flag.Parse()

	done := make(chan bool)
	go utils.VisitLink(link)

	<-done
}
