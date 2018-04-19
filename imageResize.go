package main

import (
	"fmt"
	"github.com/disintegration/imaging"
)


func main() {
	src, err := imaging.Open("コウペンちゃん３/16.png")

	if err != nil {
		fmt.Println(err)
	}

	dstImageFill := imaging.Fill(src, 128, 128, imaging.Center, imaging.Lanczos)

	err = imaging.Save(dstImageFill, "test/out_example.png")
	if err != nil {
		fmt.Println(err)
	}
}
