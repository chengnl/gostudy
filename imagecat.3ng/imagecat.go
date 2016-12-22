package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"imagecat.3ng/crawl"
)

var imageRoot = "/Users/chengnl/test/"

func getImage(image string) {
	fmt.Printf("image:%s =====================>start\n", image)
	resp, err := http.Get(image)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	imagename := image[strings.LastIndex(image, "/")+1:]
	if imagename == "" {
		return
	}
	if resp.StatusCode == http.StatusOK {
		imagefile, err := os.Create(fmt.Sprintf("%s%s", imageRoot, imagename))
		if err != nil {
			log.Fatal(err)
			return
		}
		io.Copy(imagefile, resp.Body)
		fmt.Printf("image:%s =====================>done\n", imagename)
	}
}

func main() {
	imageschan := make(chan string, 100)
	go func() {
		var baseUrl = "http://tu.duowan.com/m/meinv"
		var url string
		i := 0
		for {
			r := rand.New(rand.NewSource(1))
			url = fmt.Sprintf("%s?offset=%d&order=created&math=%.16f", baseUrl, i, r.Float64())
			fmt.Println(url)
			result := crawl.Crawl(url, imageschan)
			if result != nil && result.More && result.Enabled {
				i = result.Offset
			} else {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}()

	for image := range imageschan {
		go getImage(image)
	}
}
