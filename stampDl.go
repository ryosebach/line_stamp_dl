package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/disintegration/imaging"
)

func GetImageAndResizeAndSave(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Not valid url! Please confirm url")
		return
	}

	if _, err := os.Stat("images"); err != nil {
		os.Mkdir("images", 0755)
	}

	dirName := "images/" + strings.Split(doc.Find("title").Text(), " -")[0]
	fmt.Println(dirName)
	if err := os.Mkdir(dirName, 0755); err != nil {
		fmt.Println(err)
	}
	if err := os.Mkdir(dirName+"/stamp", 0755); err != nil {
		fmt.Println(err)
	}

	var indexNum int = 0
	doc.Find(".mdCMN09Image.FnPreview").Each(func(_ int, s *goquery.Selection) {
		elem, _ := s.Attr("style")
		imgUrl := strings.Split(strings.Split(elem, "url(")[1], ";compress")[0]

		response, err := http.Get(imgUrl)
		body, err := ioutil.ReadAll(response.Body)
		filename := dirName + "/" + strconv.Itoa(indexNum) + ".png"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println(err)
		}

		defer func() {
			file.Close()
		}()

		file.Write(body)

		src, err := imaging.Open(filename)

		if err != nil {
			fmt.Println(err)
		}

		fillImg := imaging.Fill(src, 128, 128, imaging.Center, imaging.Lanczos)

		if err := imaging.Save(fillImg, dirName+"/stamp/"+strconv.Itoa(indexNum)+".png"); err != nil {
			fmt.Println(err)
		}
		indexNum++
	})
}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		fmt.Println("please input Stamp-URL")
		return
	}
	GetImageAndResizeAndSave(url)
}
