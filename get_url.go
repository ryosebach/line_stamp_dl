package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"flag"
)

func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find(".mdCMN09Image").Each(func(_ int, s *goquery.Selection) {
		elem, _ := s.Attr("style")
		imgUrl := strings.Split(strings.Split(elem, "url(")[1], ";compress")[0]
		fmt.Println(imgUrl)
	})
}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		fmt.Println("please input Stamp-URL")
		return
	}
	GetPage(url)
}
